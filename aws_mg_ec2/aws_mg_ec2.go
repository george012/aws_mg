package aws_mg_ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/george012/aws_mg/aws_mg_common"
	"github.com/george012/aws_mg/aws_mg_model"
	"github.com/george012/gtbox/gtbox_log"
	"log"
	"time"
)

// 等待EC2实例状态变为"running"
func waitForInstancesRunning(ctx context.Context, svc *ec2.Client, instanceIds []string) error {
	for {
		describeInstancesOutput, err := svc.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			InstanceIds: instanceIds,
		})
		if err != nil {
			return err
		}

		allRunning := true
		for _, reservation := range describeInstancesOutput.Reservations {
			for _, instance := range reservation.Instances {
				if *&instance.State.Name != "running" {
					allRunning = false
					break
				}
			}
		}

		if allRunning {
			fmt.Println("所有EC2实例的状态已变为'running'")
			return nil
		}

		fmt.Println("等待中...")
		time.Sleep(10 * time.Second)
	}
}

func CreateInstanceFromAWSManager(instance_pre_config *aws_mg_model.AWSInstancePreConfig, aws_config *aws.Config, ec2_client *ec2.Client, end_func func(result_info interface{}, err error)) {

	// 创建EC2服务客户端
	ec2_client = ec2.NewFromConfig(*aws_config)

	// 创建EC2实例
	runInput := &ec2.RunInstancesInput{
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
		InstanceType: instance_pre_config.InstanceType,
		BlockDeviceMappings: []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &types.EbsBlockDevice{
					VolumeSize: aws.Int32(instance_pre_config.VolumeSize),
				},
			},
		},
		SubnetId:         aws.String(instance_pre_config.SubnetID),
		SecurityGroupIds: instance_pre_config.SafeGroupIDs,
		ImageId:          aws.String(instance_pre_config.AMIId),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(instance_pre_config.NameTag),
					},
				},
			},
		},
	}

	runOutput, err := ec2_client.RunInstances(context.TODO(), runInput)
	if err != nil {
		end_func(nil, errors.New(fmt.Sprintf("无法创建EC2实例: %s", err.Error())))
		return
	}

	// 打印实例ID
	gtbox_log.LogInfof("创建的实例ID:")
	for _, instance := range runOutput.Instances {
		gtbox_log.LogInfof("%v", *instance.InstanceId)
	}

	// 等待EC2实例状态变为"running"
	gtbox_log.LogInfof("等待EC2实例状态变为'running'...")
	instanceIds := []string{}
	for _, instance := range runOutput.Instances {
		instanceIds = append(instanceIds, *instance.InstanceId)
	}
	if err := waitForInstancesRunning(context.TODO(), ec2_client, instanceIds); err != nil {
		end_func(nil, errors.New(fmt.Sprintf("等待EC2实例状态变为'running'时出错: %s", err.Error())))
		return
	}

	// 为每个实例申请并绑定Elastic IP
	for _, instance := range runOutput.Instances {
		// 申请Elastic IP
		allocateOutput, err := ec2_client.AllocateAddress(context.TODO(), &ec2.AllocateAddressInput{})
		if err != nil {
			end_func(nil, errors.New(fmt.Sprintf("无法分配Elastic IP: %s", err.Error())))
			return
		}

		// 绑定Elastic IP到实例
		_, err = ec2_client.AssociateAddress(context.TODO(), &ec2.AssociateAddressInput{
			InstanceId: instance.InstanceId,
			PublicIp:   allocateOutput.PublicIp,
		})

		if err != nil {
			end_func(nil, errors.New(fmt.Sprintf("无法绑定Elastic IP到实例: %s", err.Error())))
			return
		}
	}
	end_func(runOutput.Instances, nil)
}

func ListInstanceFromAWSManager(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client) map[string]*types.Instance {
	// 创建EC2服务客户端
	ec2_client = ec2.NewFromConfig(*aws_config)
	result, err := ec2_client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})

	if err != nil {
		gtbox_log.LogErrorf("Got an error retrieving information about your Amazon EC2 instances: %s", err)
		return nil
	}

	ret_map := map[string]*types.Instance{}

	for _, r := range result.Reservations {

		for _, ins := range r.Instances {
			ret_map[*r.ReservationId] = &ins
		}
	}
	return ret_map
}

func StopInstance(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client, nstanceIds []string) error {
	// 创建EC2服务客户端
	ec2_client = ec2.NewFromConfig(*aws_config)

	//停止aws
	result, err := ec2_client.StopInstances(context.TODO(), &ec2.StopInstancesInput{InstanceIds: nstanceIds})
	if err != nil {
		fmt.Println("Error stopping instance:", err)
		return err
	}
	fmt.Println(result)
	return nil
}

func DeleteInstance(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client, instanceIds []string) error {
	// 创建EC2服务客户端
	ec2_client = ec2.NewFromConfig(*aws_config)
	//停止aws
	_, err := ec2_client.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{InstanceIds: instanceIds})
	if err != nil {
		fmt.Println("Error stopping instance:", err)
		return err
	}
	return nil
}

func RebootInstance(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client, instanceIds []string) error {
	// 创建EC2服务客户端
	ec2_client = ec2.NewFromConfig(*aws_config)
	_, err := ec2_client.RebootInstances(context.TODO(), &ec2.RebootInstancesInput{InstanceIds: instanceIds})
	if err != nil {
		gtbox_log.LogErrorf("重启失败")
		return err
	}
	return nil
}

func StartInstance(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client, instanceIds []string) error {
	// 创建EC2服务客户端
	ec2_client = ec2.NewFromConfig(*aws_config)
	_, err := ec2_client.StartInstances(context.TODO(), &ec2.StartInstancesInput{InstanceIds: instanceIds})
	if err != nil {
		gtbox_log.LogErrorf("启动失败")
		return err
	}
	return nil
}

func ImportCertificate(region aws_mg_common.AWSRegion, aws_config *aws.Config, acm_client *acm.Client, certificate aws_mg_model.Certificate) error {
	acm_client = acm.NewFromConfig(*aws_config)

	// 创建上传请求
	input := &acm.ImportCertificateInput{
		//证书
		Certificate: []byte(certificate.Certificate),
		//私钥
		PrivateKey: []byte(certificate.PrivateKey),
		//证书链
		CertificateChain: []byte(certificate.CertificateChain),
	}

	_, err := acm_client.ImportCertificate(context.TODO(), input)

	if err != nil {
		gtbox_log.LogErrorf("上传证书失败")
		return err
	}
	return nil
}

func DeleteCertificate(region aws_mg_common.AWSRegion, aws_config *aws.Config, acm_client *acm.Client, certificateArn string) error {
	acm_client = acm.NewFromConfig(*aws_config)

	// 创建上传请求
	input := &acm.DeleteCertificateInput{
		CertificateArn: &certificateArn,
	}

	_, err := acm_client.DeleteCertificate(context.TODO(), input)

	if err != nil {
		gtbox_log.LogErrorf("删除证书失败")
		return err
	}
	return nil
}

func GetIpList(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client, prefixListIDs []string) []types.ManagedPrefixList {
	ec2_client = ec2.NewFromConfig(*aws_config)

	// 查询 AWS 全球 IP 地址范围
	input := &ec2.DescribeManagedPrefixListsInput{
		PrefixListIds: prefixListIDs,
	}

	result, err := ec2_client.DescribeManagedPrefixLists(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}
	return result.PrefixLists
}

func DeleteIp(region aws_mg_common.AWSRegion, aws_config *aws.Config, ec2_client *ec2.Client, prefixListID string) {
	ec2_client = ec2.NewFromConfig(*aws_config)

	// 准备删除 IP 地址列表的输入参数
	input := &ec2.DeleteManagedPrefixListInput{
		PrefixListId: aws.String(prefixListID),
	}

	// 删除 IP 地址列表
	_, err := ec2_client.DeleteManagedPrefixList(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}
}

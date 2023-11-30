package aws_mg

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/george012/aws_mg/aws_mg_common"
	"github.com/george012/aws_mg/aws_mg_ec2"
	"github.com/george012/aws_mg/aws_mg_model"
	"sync"
)

type AWSManager struct {
	region               aws_mg_common.AWSRegion
	awsAk                string
	awsSk                string
	aWSConfig            *aws.Config
	ec2instancePreConfig *aws_mg_model.AWSInstancePreConfig
	ec2Client            *ec2.Client
	acmClient            *acm.Client
}

var (
	aws_mg_once    sync.Once
	currentManager *AWSManager
)

func instanceOnce() *AWSManager {
	aws_mg_once.Do(func() {
		currentManager = &AWSManager{}
	})

	return currentManager
}

func (aws_mg *AWSManager) ListInstance() map[string]*types.Instance {
	return aws_mg_ec2.ListInstanceFromAWSManager(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client)
}

func (aws_mg *AWSManager) DeleteInstance(instanceIds []string) {
	aws_mg_ec2.DeleteInstance(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, instanceIds)
}

func (aws_mg *AWSManager) StopInstance(instanceIds []string) {
	aws_mg_ec2.StopInstance(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, instanceIds)
}

func (aws_mg *AWSManager) RebootInstance(instanceIds []string) {
	aws_mg_ec2.RebootInstance(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, instanceIds)
}

func (aws_mg *AWSManager) StartInstance(instanceIds []string) {
	aws_mg_ec2.StartInstance(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, instanceIds)
}

func (aws_mg *AWSManager) ImportCertificate(certificate aws_mg_model.Certificate) {
	aws_mg_ec2.ImportCertificate(aws_mg.region, aws_mg.aWSConfig, aws_mg.acmClient, certificate)
}

func (aws_mg *AWSManager) DeleteCertificate(certificateArn string) {
	aws_mg_ec2.DeleteCertificate(aws_mg.region, aws_mg.aWSConfig, aws_mg.acmClient, certificateArn)
}
func (aws_mg *AWSManager) GetIpList(prefixListIDs []string) []types.ManagedPrefixList {
	return aws_mg_ec2.GetIpList(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, prefixListIDs)
}
func (aws_mg *AWSManager) DeleteIp(prefixListID string) {
	aws_mg_ec2.DeleteIp(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, prefixListID)
}

func (aws_mg *AWSManager) AllocateEIP() *ec2.AllocateAddressOutput {
	return aws_mg_ec2.AllocateEIP(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client)
}

func (aws_mg *AWSManager) CreateNetworkInterface(subnetID string) *ec2.CreateNetworkInterfaceOutput {
	return aws_mg_ec2.CreateNetworkInterface(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, subnetID)
}

func (aws_mg *AWSManager) DeleteNetworkInterface(networkInterfaceId string) {
	aws_mg_ec2.DeleteNetworkInterface(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, networkInterfaceId)
}

func (aws_mg *AWSManager) AssociateEIP(allocationId, networkInterfaceId string) *ec2.AssociateAddressOutput {
	return aws_mg_ec2.AssociateEIP(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, allocationId, networkInterfaceId)
}

func (aws_mg *AWSManager) DisassociateAddress(associationId string) {
	aws_mg_ec2.DisassociateAddress(aws_mg.region, aws_mg.aWSConfig, aws_mg.ec2Client, associationId)
}

func (aws_mg *AWSManager) CreateEC2Instance(instance_pre_config *aws_mg_model.AWSInstancePreConfig, end_func func(result_info interface{}, err error)) {
	aws_mg_ec2.CreateInstanceFromAWSManager(instance_pre_config, aws_mg.aWSConfig, aws_mg.ec2Client, end_func)
}

// NewAWSManager 初始化工具，仅运行一次。
func NewAWSManager(ak string, sk string, default_region aws_mg_common.AWSRegion) (*AWSManager, error) {

	instanceOnce().awsAk = ak
	instanceOnce().awsSk = sk

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(currentManager.awsAk, currentManager.awsSk, "")),
		config.WithRegion(default_region.String()),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("无法配置AWS SDK: %s", err.Error()))
	}

	instanceOnce().aWSConfig = &cfg
	instanceOnce().region = default_region

	return instanceOnce(), nil
}

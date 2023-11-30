package aws_mg_test

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/george012/aws_mg"
	"github.com/george012/aws_mg/aws_mg_common"
	"github.com/george012/aws_mg/aws_mg_model"
	"github.com/george012/gtbox/gtbox_log"
	"testing"
)

func TestCreateEC2(t *testing.T) {
	ami_ubuntu2004 := "ami-0f8e81a3da6e2510a"

	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_pre_conf := aws_mg_model.NewAWSInstancePreConfig(
		aws_mg_common.AWSRegion_US_West_1_California_North,
		ami_ubuntu2004, 150,
		types.InstanceTypeT3Large, []string{},
		"",
		"aws_mg_test",
		"",
		true,
	)

	a_mg.CreateEC2Instance(a_pre_conf, func(result_info interface{}, err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
}

// 开通网络接口，在开通弹性IP，弹性IP上绑定网络接口，再开实例，实例上“自动分配公有 IP”选择禁用
func TestCreateEc2AndAssociateEIP(t *testing.T) {
	ami_ubuntu2004 := "ami-0f8e81a3da6e2510a"

	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}
	//分配弹性Ip
	allocateAddressOutput := a_mg.AllocateEIP()
	//创建网络接口
	createNetworkInterfaceOutput := a_mg.CreateNetworkInterface("")
	//将弹性 IP 绑定到网络接口
	associateAddressOutput := a_mg.AssociateEIP(*allocateAddressOutput.AllocationId, *createNetworkInterfaceOutput.NetworkInterface.NetworkInterfaceId)

	a_pre_conf := aws_mg_model.NewAWSInstancePreConfig(
		aws_mg_common.AWSRegion_US_West_1_California_North,
		ami_ubuntu2004, 150,
		types.InstanceTypeT3Large, []string{},
		"",
		"aws_mg_test",
		*associateAddressOutput.AssociationId,
		false,
	)

	a_mg.CreateEC2Instance(a_pre_conf, func(result_info interface{}, err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
}

func TestDeleteNetworkInterface(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.DeleteNetworkInterface("")
}

func TestDisassociateAddress(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.DisassociateAddress("")
}

func TestStopInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.StopInstance([]string{""})
}

func TestStartInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.StartInstance([]string{""})
}

func TestRebootInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.RebootInstance([]string{""})
}

func TestDeleteInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.DeleteInstance([]string{""})
}

func TestImportCertificate(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.ImportCertificate(aws_mg_model.Certificate{Certificate: "", PrivateKey: "", CertificateChain: ""})
}

func TestDeleteCertificate(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.DeleteCertificate("certificateArn")
}

func TestGetIpList(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk/", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}
	var s []string
	s = append(s, "")
	a_mg.GetIpList(s)
}

func TestDeleteIp(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk/", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}
	a_mg.DeleteIp("")
}

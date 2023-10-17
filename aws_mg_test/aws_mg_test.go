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
	)

	a_mg.CreateEC2Instance(a_pre_conf, func(result_info interface{}, err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
}

func TestStopInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.StopInstance([]string{"i-009cc2e0f3098a3da"})
}

func TestStartInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.StartInstance([]string{"i-009cc2e0f3098a3da"})
}

func TestRebootInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.RebootInstance([]string{"i-009cc2e0f3098a3da"})
}

func TestDeleteInstanceEC2(t *testing.T) {
	a_mg, err := aws_mg.NewAWSManager("test_ak", "test_sk", aws_mg_common.AWSRegion_US_West_2_Oregon)
	if err != nil {
		gtbox_log.LogErrorf("%s", err)
	}

	a_mg.DeleteInstance([]string{"i-009cc2e0f3098a3da"})
}

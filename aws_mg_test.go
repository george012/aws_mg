package aws_mg

import (
	"github.com/george012/aws_mg/aws_mg_common"
	"github.com/george012/gtbox/gtbox_log"
	"testing"
)

func TestCreateEC2(t *testing.T) {
	ami_ubuntu2004 := "ami-0f8e81a3da6e2510a"

	SetupAWSManager("test_ak", "test_sk", func(err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})

	NewEC2WithRegion(aws_mg_common.AWSRegion_US_West_1_California_North, ami_ubuntu2004, func(result_info interface{}, err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
}

func TestStopInstanceEC2(t *testing.T) {
	currentRegion = aws_mg_common.AWSRegion_US_West_2_Oregon
	SetupAWSManager("test_ak", "test_ak/", func(err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
	instanceOnce().StopInstance([]string{"i-009cc2e0f3098a3da"})
}

func TestStartInstanceEC2(t *testing.T) {
	currentRegion = aws_mg_common.AWSRegion_US_West_2_Oregon
	SetupAWSManager("test_ak", "test_ak/", func(err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
	instanceOnce().StartInstance([]string{"i-009cc2e0f3098a3da"})
}

func TestRebootInstanceEC2(t *testing.T) {
	currentRegion = aws_mg_common.AWSRegion_US_West_2_Oregon
	SetupAWSManager("test_ak", "test_ak/", func(err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
	instanceOnce().RebootInstance([]string{"i-009cc2e0f3098a3da"})
}

func TestDeleteInstanceEC2(t *testing.T) {
	currentRegion = aws_mg_common.AWSRegion_US_West_2_Oregon
	SetupAWSManager("test_ak", "test_ak/", func(err error) {
		if err != nil {
			gtbox_log.LogErrorf("%s", err)
		}
	})
	instanceOnce().DeleteInstance([]string{"i-009cc2e0f3098a3da"})
}

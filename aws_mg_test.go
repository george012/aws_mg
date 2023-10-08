package aws_mg

import (
	"github.com/s-c-f-d/aws_mg/aws_mg_common"
	"testing"
)

func TestCreateEC2(t *testing.T) {
	ami_ubuntu2004 := "ami-0f8e81a3da6e2510a"
	SetupAWSManager("test_ak", "test_sk")
	NewEC2WithRegion(aws_mg_common.AWSRegion_US_West_1_California_North, ami_ubuntu2004)
}

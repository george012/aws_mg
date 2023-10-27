package aws_mg

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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

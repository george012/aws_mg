package aws_mg_model

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/george012/aws_mg/aws_mg_common"
)

type AWSInstancePreConfig struct {
	Region                     aws_mg_common.AWSRegion
	VolumeSize                 int32              // 硬盘大小 GB
	InstanceType               types.InstanceType // 实例类型
	SafeGroupIDs               []string           // 安全组ID
	SubnetID                   string             // 子网ID
	NameTag                    string             // tag名字
	AMIId                      string             // 镜像id
	NetworkInterfaceId         string             // 网络接口id
	IsAssociatePublicIpAddress bool               //是否禁用公网ip
}

// NewAWSInstancePreConfig 创建EC2
// region 区域
// ami_id 系统镜像ID
// volume_size 硬盘大小
// aws_instanceType 实例类型
// safe_groups 安全组
// subnet_id 子网ID
// name_tag 名字标签
func NewAWSInstancePreConfig(region aws_mg_common.AWSRegion, ami_id string, volume_size int32, aws_instanceType types.InstanceType, safe_groups []string, subnet_id string, name_tag string, networkInterfaceId string, isAssociatePublicIpAddress bool) *AWSInstancePreConfig {
	instancePreConfig := &AWSInstancePreConfig{
		Region:                     region,
		VolumeSize:                 volume_size,
		InstanceType:               aws_instanceType,
		SafeGroupIDs:               safe_groups,
		SubnetID:                   subnet_id,
		NameTag:                    name_tag,
		AMIId:                      ami_id,
		NetworkInterfaceId:         networkInterfaceId,
		IsAssociatePublicIpAddress: isAssociatePublicIpAddress,
	}
	return instancePreConfig
}

type Certificate struct {
	Certificate      string //证书
	PrivateKey       string //私钥
	CertificateChain string //证书链
}

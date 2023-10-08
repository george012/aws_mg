package aws_mg_model

type AWSInstancePreConfig struct {
	VolumeSize   int32    // 硬盘大小 GB
	InstanceType string   // 实例类型
	SafeGroupIDs []string // 安全组ID
	SubnetID     string   // 子网ID
	NameTag      string   // tag名字
	AMIId        string   // 镜像id
}

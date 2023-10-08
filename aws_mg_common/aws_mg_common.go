package aws_mg_common

type AWSSubnetInfo struct {
	SubnetID  string
	SubnetTag string
}

type AWSVPCInfo struct {
	Subnets []*AWSSubnetInfo
}

type AWSRegionInfo struct {
	RegionTag string
	VpcInfo   []*AWSVPCInfo
}

type AWSRegion int64

const (
	AWSRegion_US_East_2_Ohio AWSRegion = iota
	AWSRegion_US_East_1_Virginia_North
	AWSRegion_US_West_1_California_North
	AWSRegion_US_West_2_Oregon
	AWSRegion_AF_South_1_Cape_Town
	AWSRegion_AP_East_1_Hong_Kong
	AWSRegion_AP_South_2_Hyderabad
	AWSRegion_AP_Southeast_3_Jakarta
	AWSRegion_AP_Southeast_4_Melbourne
	AWSRegion_AP_South_1_Mumbai
	AWSRegion_AP_Northeast_3_Osaka
	AWSRegion_AP_Northeast_2_Seoul
	AWSRegion_AP_Southeast_1_Singapore
	AWSRegion_AP_Southeast_2_Sydney
	AWSRegion_AP_Northeast_1_Tokyo
	AWSRegion_CA_Central_1_Central
	AWSRegion_EU_Central_1_Frankfurt
	AWSRegion_EU_West_1_Ireland
	AWSRegion_EU_West_2_London
	AWSRegion_EU_South_1_Milan
	AWSRegion_EU_West_3_Paris
	AWSRegion_EU_South_2_Spain
	AWSRegion_EU_North_1_Stockholm
	AWSRegion_EU_Central_2_Zurich
	AWSRegion_IL_Central_1_Tel_Aviv
	AWSRegion_ME_South_1_Bahrain
	AWSRegion_ME_Central_1_UAE
	AWSRegion_SA_East_1_Sao_Paulo
)

// String AWS可用区 字符串标识
func (region AWSRegion) String() string {
	switch region {
	case AWSRegion_US_East_2_Ohio:
		return "us-east-2"
	case AWSRegion_US_East_1_Virginia_North:
		return "us-east-1"
	case AWSRegion_US_West_1_California_North:
		return "us-west-1"
	case AWSRegion_US_West_2_Oregon:
		return "us-west-2"
	case AWSRegion_AF_South_1_Cape_Town:
		return "af-south-1"
	case AWSRegion_AP_East_1_Hong_Kong:
		return "ap-east-1"
	case AWSRegion_AP_South_2_Hyderabad:
		return "ap-south-2"
	case AWSRegion_AP_Southeast_3_Jakarta:
		return "ap-southeast-3"
	case AWSRegion_AP_Southeast_4_Melbourne:
		return "ap-southeast-4"
	case AWSRegion_AP_South_1_Mumbai:
		return "ap-south-1"
	case AWSRegion_AP_Northeast_3_Osaka:
		return "ap-northeast-3"
	case AWSRegion_AP_Northeast_2_Seoul:
		return "ap-northeast-2"
	case AWSRegion_AP_Southeast_1_Singapore:
		return "ap-southeast-1"
	case AWSRegion_AP_Southeast_2_Sydney:
		return "ap-southeast-2"
	case AWSRegion_AP_Northeast_1_Tokyo:
		return "ap-northeast-1"
	case AWSRegion_CA_Central_1_Central:
		return "ca-central-1"
	case AWSRegion_EU_Central_1_Frankfurt:
		return "eu-central-1"
	case AWSRegion_EU_West_1_Ireland:
		return "eu-west-1"
	case AWSRegion_EU_West_2_London:
		return "eu-west-2"
	case AWSRegion_EU_South_1_Milan:
		return "eu-south-1"
	case AWSRegion_EU_West_3_Paris:
		return "eu-west-3"
	case AWSRegion_EU_South_2_Spain:
		return "eu-south-2"
	case AWSRegion_EU_North_1_Stockholm:
		return "eu-north-1"
	case AWSRegion_EU_Central_2_Zurich:
		return "eu-central-2"
	case AWSRegion_IL_Central_1_Tel_Aviv:
		return "il-central-1"
	case AWSRegion_ME_South_1_Bahrain:
		return "me-south-1"
	case AWSRegion_ME_Central_1_UAE:
		return "me-central-1"
	case AWSRegion_SA_East_1_Sao_Paulo:
		return "sa-east-1"
	default:
		return ""
	}
}

// DescriptionCN AWS可用区 中文描述
func (region AWSRegion) DescriptionCN() string {
	switch region {
	case AWSRegion_US_East_2_Ohio:
		return "美国东部（俄亥俄）"
	case AWSRegion_US_East_1_Virginia_North:
		return "美国东部（弗吉尼亚北部）"
	case AWSRegion_US_West_1_California_North:
		return "美国西部（加利福尼亚北部）"
	case AWSRegion_US_West_2_Oregon:
		return "美国西部（俄勒冈）"
	case AWSRegion_AF_South_1_Cape_Town:
		return "非洲（开普敦）"
	case AWSRegion_AP_East_1_Hong_Kong:
		return "亚太地区（香港）"
	case AWSRegion_AP_South_2_Hyderabad:
		return "亚太地区（海得拉巴）"
	case AWSRegion_AP_Southeast_3_Jakarta:
		return "亚太地区（雅加达）"
	case AWSRegion_AP_Southeast_4_Melbourne:
		return "亚太地区（墨尔本）"
	case AWSRegion_AP_South_1_Mumbai:
		return "亚太地区（孟买）"
	case AWSRegion_AP_Northeast_3_Osaka:
		return "亚太地区（大阪）"
	case AWSRegion_AP_Northeast_2_Seoul:
		return "亚太地区（首尔）"
	case AWSRegion_AP_Southeast_1_Singapore:
		return "亚太地区（新加坡）"
	case AWSRegion_AP_Southeast_2_Sydney:
		return "亚太地区（悉尼）"
	case AWSRegion_AP_Northeast_1_Tokyo:
		return "亚太地区（东京）"
	case AWSRegion_CA_Central_1_Central:
		return "加拿大（中部）"
	case AWSRegion_EU_Central_1_Frankfurt:
		return "欧洲（法兰克福）"
	case AWSRegion_EU_West_1_Ireland:
		return "欧洲（爱尔兰）"
	case AWSRegion_EU_West_2_London:
		return "欧洲（伦敦）"
	case AWSRegion_EU_South_1_Milan:
		return "欧洲（米兰）"
	case AWSRegion_EU_West_3_Paris:
		return "欧洲（巴黎）"
	case AWSRegion_EU_South_2_Spain:
		return "欧洲（西班牙）"
	case AWSRegion_EU_North_1_Stockholm:
		return "欧洲地区（斯德哥尔摩）"
	case AWSRegion_EU_Central_2_Zurich:
		return "欧洲（苏黎世）"
	case AWSRegion_IL_Central_1_Tel_Aviv:
		return "以色列（特拉维夫）"
	case AWSRegion_ME_South_1_Bahrain:
		return "中东（巴林）"
	case AWSRegion_ME_Central_1_UAE:
		return "中东（阿联酋）"
	case AWSRegion_SA_East_1_Sao_Paulo:
		return "南美洲（圣保罗）"
	default:
		return ""
	}
}

func (region AWSRegion) GetRegionInfos() *AWSRegionInfo {
	region_info := &AWSRegionInfo{}

	return region_info
}

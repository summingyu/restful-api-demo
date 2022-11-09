package host

import "context"

// define host app service interface
type Service interface {
	CreateHost(context.Context, *Host) (*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	DescribeHost(context.Context, *QueryHostRequest) (*Host, error)
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

func NewHost() *Host {
	return &Host{
		Base:     &Base{},
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

// define the host struct
type Host struct {
	*Base
	*Resource
	*Describe
}

type Vendor int

type Base struct {
	Id           string `json:"id"`            // 全局唯一Id
	SyncAt       int64  `json:"sync_at"`       // 同步时间
	Vendor       Vendor `json:"vendor"`        // 厂商
	Region       string `json:"region"`        // 地域
	Zone         string `json:"zone"`          // 区域
	CreateAt     int64  `json:"create_at"`     // 创建时间
	ResourceHash string `json:"resource_hash"` // 基础数据Hash
	DescribeHash string `json:"describe_hash"` // 描述数据Hash
}

type Resource struct {
	ExpireAt    int64             `json:"expire_at"`   // 过期时间
	Category    string            `json:"category"`    // 种类
	Type        string            `json:"type"`        // 规格
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	Account     string            `json:"account"`     // 资源所有者
	PublicIP    string            `json:"public_ip"`   // 公网IP
	PrivateIP   string            `json:"private_ip"`  // 内网IP
	PayType     string            `json:"pay_type"`    // 实例付费方式
}

type Describe struct {
	CPU          int    `json:"cpu"`           // 核数
	Memory       int    `json:"memory"`        // 内存
	GPUAmount    int    `json:"gpu_amount"`    // GPU数量
	GPUSpec      string `json:"gpu_spec"`      // GPU类型
	OSType       string `json:"os_type"`       // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`       // 操作系统名称
	SerialNumber string `json:"serial_number"` // 序列号
}

type HostSet struct {
	Items []*Host
	Total int
}

type QueryHostRequest struct {
}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct {
	Id string `json:"id"`
}

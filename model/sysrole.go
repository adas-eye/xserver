package model

// SysRole
type SysRoleOpt struct {
	ID      uint   `json:"roleId" gorm:"primary_key"`
	Name    string `json:"roleName" gorm:"not null;comment:角色名称;"`
	Code    string `json:"roleCode" gorm:"unique_index;size:100;not null;comment:角色权限字符串;"`
	Type    string `json:"roleType" gorm:"size:1;default:'1';"`
	Enable  string `json:"enable" gorm:"size:1,default:'1';comment:0禁用1正常;"`
	MenuIds JUints `json:"menuIds" gorm:""`
	Details string `json:"details"`
}

type SysRole struct {
	SysRoleOpt
	CreatedAt jtime  `json:"createdAt"`
	CreatedBy string `json:"createdBy" gorm:"comment:创建者;"`
}

func (o *SysRole) TableName() string {
	return "t_sysrole"
}

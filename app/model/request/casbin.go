package request

import (
	"gva-lbx/app/model"
	"strconv"
)

type Rule struct {
	Path   string `json:"path" example:"路由路径"`
	Method string `json:"method" example:"请求方法"`
}

type CasbinAddPolicies struct {
	RoleId uint `json:"roleId" swaggertype:"string" example:"uint 角色id"`
	Rules  []Rule
}

func (r *CasbinAddPolicies) Update() [][]string {
	length := len(r.Rules)
	rules := make([][]string, 0, length)
	roleId := strconv.Itoa(int(r.RoleId))
	for i := 0; i < length; i++ {
		rules = append(rules, []string{roleId, r.Rules[i].Path, r.Rules[i].Method})
	}
	return rules
}

type CasbinUpdate struct {
	OldPath   string `json:"oldPath" example:"旧路径"`
	NewPath   string `json:"newPath" example:"新路径"`
	OldMethod string `json:"oldMethod" example:"旧请求方法"`
	NewMethod string `json:"newMethod" example:"新请求方法"`
}

func (r *CasbinUpdate) Update() model.Casbin {
	return model.Casbin{
		Path:   r.NewPath,
		Method: r.NewMethod,
	}
}

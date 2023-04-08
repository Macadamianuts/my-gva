package request

import (
	"go/token"
	"gva-lbx/plugin/coder/model"
	"gva-lbx/utils"
	"sort"
	"strings"
	"text/template"
)

// AutoCodeCreate 代码生成器
type AutoCodeCreate struct {
	Type          string           `json:"type" example:"模版类型"`
	Plugin        string           `json:"plugin" example:"插件名"`
	Struct        string           `json:"struct" example:"结构体名称"`
	Filename      string           `json:"filename" example:"文件夹名称"`
	TableName     string           `json:"table" example:"表名"`
	Description   string           `json:"description" example:"结构体中文名称"`
	Abbreviation  string           `json:"abbreviation" example:"结构体简称"`
	UnderlineName string           `json:"humpPackageName" example:"go文件名(下划线)"`
	AutoMoveFile  bool             `json:"autoMoveFile" swaggertype:"string" example:"bool 是否自动移动文件"`
	AutoCreateApi bool             `json:"autoCreateApi" swaggertype:"string" example:"bool 是否自动创建api"`
	Fields        []*AutoCodeField `json:"fields,omitempty"`
}

func (r *AutoCodeCreate) AutoCodeHistoryRepeat() AutoCodeHistoryRepeat {
	return AutoCodeHistoryRepeat{
		Struct: r.Struct,
		Plugin: r.Plugin,
	}
}

func (r *AutoCodeCreate) Create() model.AutoCodeHistory {
	length := len(r.Fields)
	fields := make([]model.AutoCodeHistoryField, 0, length)
	for i := 0; i < length; i++ {
		fields = append(fields, model.AutoCodeHistoryField{
			Name:        r.Fields[i].Name,
			Type:        r.Fields[i].Type,
			Json:        r.Fields[i].Json,
			Description: r.Fields[i].Description,
			Size:        r.Fields[i].Size,
			Where:       r.Fields[i].Where,
			Column:      r.Fields[i].Column,
			Comment:     r.Fields[i].Comment,
			Sort:        r.Fields[i].Sort,
			ErrorText:   r.Fields[i].ErrorText,
			Dictionary:  r.Fields[i].Dictionary,
			Require:     r.Fields[i].Require,
			Clearable:   r.Fields[i].Clearable,
		})
	}
	return model.AutoCodeHistory{
		Type:          r.Type,
		Plugin:        r.Plugin,
		Struct:        r.Struct,
		Filename:      r.Filename,
		TableName:     r.TableName,
		Description:   r.Description,
		Abbreviation:  r.Abbreviation,
		UnderlineName: r.UnderlineName,
		AutoMoveFile:  r.AutoMoveFile,
		AutoCreateApi: r.AutoCreateApi,
		Fields:        fields,
	}
}

func (r *AutoCodeCreate) Functions() template.FuncMap {
	return map[string]any{
		"HasTime": r.HasTime,
		"HasSort": r.HasSort,
	}
}

func (r *AutoCodeCreate) HasTime() bool {
	for i := 0; i < len(r.Fields); i++ {
		if r.Fields[i].Type == "time.Time" {
			return true
		}
	}
	return false
}

func (r *AutoCodeCreate) HasSort() bool {
	for i := 0; i < len(r.Fields); i++ {
		if r.Fields[i].Sort {
			return true
		}
	}
	return false
}

func (r *AutoCodeCreate) Pretreatment() {
	r.KeyWord()
	r.TrimSpace()
	r.AutoCodeFieldStruct()
}

// TrimSpace 去空格
func (r *AutoCodeCreate) TrimSpace() {
	utils.TrimSpace(r)
	for i := 0; i < len(r.Fields); i++ {
		utils.TrimSpace(r.Fields[i])
	}
}

func (r *AutoCodeCreate) AutoCodeFieldStruct() {
	for i := 0; i < len(r.Fields); i++ {
		r.Fields[i].Struct = r.Struct
	}
}

func (r *AutoCodeCreate) KeyWord() {
	if token.IsKeyword(r.Abbreviation) {
		r.Abbreviation = r.Abbreviation + "_"
	}
	if strings.HasSuffix(r.UnderlineName, "test") {
		r.UnderlineName = r.UnderlineName + "_"
	}
}

func (r *AutoCodeCreate) Sort() {
	length := len(r.Fields)
	fieldMap := make(map[string]*AutoCodeField, length)
	fieldTypeMap := make(map[string][]*AutoCodeField, length)
	slices := make(sort.StringSlice, 0, length)
	for i := 0; i < length; i++ {
		{ // 字段类型
			value, ok := fieldTypeMap[r.Fields[i].Type]
			if !ok {
				value = make([]*AutoCodeField, 0, 5)
			}
			value = append(value, r.Fields[i])
			fieldTypeMap[r.Fields[i].Type] = value
		}
		{ // 字段名
			_, ok := fieldMap[r.Fields[i].Name]
			if !ok {
				fieldMap[r.Fields[i].Name] = r.Fields[i]
			}
			slices = append(slices, r.Fields[i].Name)
		}
	}
	sort.Strings(slices)
}

// AutoCodeField 代码生成器字段
type AutoCodeField struct {
	// 后端相关
	Name        string `json:"name" example:"字段名"`
	Type        string `json:"type" example:"字段数据类型"`
	Json        string `json:"json" example:"字段 json tag"`
	Struct      string `json:"-" example:"结构体名称"`
	Description string `json:"fieldDescription" example:"字段中文名"`
	// 数据库相关
	Size    string `json:"size" example:"数据库字段长度"`
	Where   string `json:"where" example:"数据库字段搜索条件"`
	Column  string `json:"column" example:"数据库字段列名"`
	Comment string `json:"comment" example:"数据库字段描述"`
	Sort    bool   `json:"sort" swaggertype:"string" example:"bool 是否增加排序"`
	// 前端相关
	ErrorText  string `json:"errorText" example:"校验失败文字"`
	Dictionary string `json:"dictionary" example:"关联字典"`
	Require    bool   `json:"require" swaggertype:"string" example:"bool 是否必填"`
	Clearable  bool   `json:"clearable" swaggertype:"string" example:"bool 是否可清空"`
}

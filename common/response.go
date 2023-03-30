package common

type PageResult struct {
	PageInfo
	List  interface{} `json:"list" swaggertype:"string" example:"interface 数据"`
	Count int64       `json:"count" swaggertype:"string" example:"int64 总数"`
}

func NewPageResult(list interface{}, count int64, pageInfo PageInfo) PageResult {
	if list == nil {
		return PageResult{
			PageInfo: PageInfo{
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			},
			List:  []struct{}{},
			Count: count,
		}
	}
	return PageResult{
		PageInfo: PageInfo{
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		},
		List:  list,
		Count: count,
	}
}

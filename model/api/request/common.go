package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`          // 页码
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,min=10"` // 每页大小
}

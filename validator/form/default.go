package form

// ListRequest 接收数据
type ListRequest struct {
	Path string `form:"path" json:"path"`
}

// ListResponse 响应数据
type ListResponse struct {
	Path  string `form:"path" json:"path" validate:"required"`
	IsDir bool   `form:"isDir" json:"isDir" validate:"required"`
	Time  string `form:"Time" json:"time" validate:"required"`
	Name  string `form:"Name" json:"name" validate:"required"`
	Size  int64  `form:"Size" json:"size" validate:"required"`
}

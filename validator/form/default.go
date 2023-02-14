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

// GetRequest 接收数据
type GetRequest struct {
	Path string `form:"path" json:"path"`
}

// GetResponse 响应数据
type GetResponse struct {
	Path  string `form:"path" json:"path" validate:"required"`
	IsDir bool   `form:"isDir" json:"isDir" validate:"required"`
	Time  string `form:"Time" json:"time" validate:"required"`
	Name  string `form:"Name" json:"name" validate:"required"`
	Size  int64  `form:"Size" json:"size" validate:"required"`
}

// DownloadRequest 接收数据
type DownloadRequest struct {
	Path string `form:"path" json:"path"`
}

// LoginRequest 登录
type LoginRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

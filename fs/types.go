package fs

// 通用 API 响应结构
type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// ListRequest 请求结构体
type ListRequest struct {
	Path    string `json:"path"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
	Refresh bool   `json:"refresh,omitempty"`
}

// ListResponse 响应结构体
type ListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content  []FileInfo `json:"content"`
		Total    int        `json:"total"`
		Readme   string     `json:"readme"`
		Header   string     `json:"header"`
		Write    bool       `json:"write"`
		Provider string     `json:"provider"`
	} `json:"data"`
}

// GetRequest 请求结构体
type GetRequest struct {
	Path     string `json:"path"`
	Password string `json:"password,omitempty"`
}

// GetResponse 响应结构体
type GetResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    FileInfo `json:"data"`
}

// FileInfo 结构体
type FileInfo struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	IsDir    bool   `json:"is_dir"`
	Modified string `json:"modified"`
	Created  string `json:"created"`
	Sign     string `json:"sign"`
	Thumb    string `json:"thumb"`
	Type     int    `json:"type"`
	RawURL   string `json:"raw_url"`
}

// CreateDirRequest 创建文件夹请求
type CreateDirRequest struct {
	Path string `json:"path"`
}

// RenameRequest 重命名文件或目录
type RenameRequest struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

// DeleteRequest 删除文件或目录
type DeleteRequest struct {
	Dir   string   `json:"dir"`
	Names []string `json:"names"`
}

package auth

// 设置登录请求体
type loginReqest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Optcode  string `json:"opt_code,omitempty"`
}

// 设置登录响应体
type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

// 2FA 生成响应体
type TwoFAResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		QR     string `json:"qr"`     // 2FA QR码的base64数据
		Sercet string `json:"secret"` // 2FA 密钥
	} `json:"data"`
}

// 2FA 验证请求体
type VerifyTwoFAReqest struct {
	Code   string `json:"code"`
	Sercet string `json:"secret"`
}

// 通用响应结构
type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type UserResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    UserInfo `json:"data"`
}

// 用户信息结构体
type UserInfo struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	BasePath   string `json:"base_path"`
	Role       int    `json:"role"`
	Disabled   bool   `json:"disabled"`
	Permission int    `json:"permission"`
	SSOID      string `json:"sso_id"`
	Opt        bool   `json:"opt"`
}

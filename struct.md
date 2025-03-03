初始化 Go SDK 项目可以分为几个步骤，确保项目的结构正确，并且能够顺利地进行开发和版本管理。以下是详细的步骤说明：

### 1. **创建项目目录**

首先，创建项目的根目录，并进入该目录：

```bash
mkdir go-alist-sdk
cd go-alist-sdk
```

### 2. **初始化 Go 模块**

在项目根目录下初始化 Go 模块，这将创建 `go.mod` 和 `go.sum` 文件，用于管理项目的依赖关系。

```bash
go mod init github.com/yourusername/go-alist-sdk
```

替换 `yourusername` 为你在 GitHub 或其他代码托管平台上的用户名。

### 3. **创建目录结构**

在项目根目录下，创建你之前设计的目录结构：

```bash
mkdir -p client auth fs public admin utils
```

你可以使用 `mkdir -p` 一次性创建多层目录。

### 4. **创建 `go.mod` 和 `go.sum` 文件**

`go.mod` 文件已经在初始化 Go 模块时创建。如果你有额外的第三方依赖（例如用于单元测试或日志记录的库），可以在 `go.mod` 中添加并通过 `go get` 下载它们。

例如，安装 `gorilla/mux`（如果你需要用它来处理路由）：

```bash
go get github.com/gorilla/mux
```

这会自动更新 `go.mod` 和 `go.sum` 文件，添加该依赖。

### 5. **添加必要的依赖**

如果需要其他依赖（如日志库、HTTP 客户端、JSON 解析库等），你可以通过 `go get` 命令安装它们。例如，Go 的标准库已经包含了 JSON 处理和 HTTP 请求功能，但你可能还需要第三方日志库：

```bash
go get github.com/sirupsen/logrus
```

这将安装 `logrus` 库并更新 `go.mod`。

### 6. **编写基础代码**

根据之前的结构说明，你可以从 `client`、`auth`、`fs`、`public` 等模块开始编写代码。以下是每个模块的简单实现示例。

#### 创建 `client/client.go`

```go
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	AuthToken  string
	HTTPClient *http.Client
}

// NewClient creates a new Alist API client
func NewClient(baseURL, authToken string) *Client {
	return &Client{
		BaseURL:    baseURL,
		AuthToken:  authToken,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// SendRequest sends a request to the Alist API and returns the response body
func (c *Client) SendRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var requestBody []byte
	if body != nil {
		var err error
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
```

#### 创建 `auth/auth.go`

```go
package auth

import (
	"fmt"
	"github.com/yourusername/go-alist-sdk/client"
)

type AuthResponse struct {
	Token string `json:"token"`
}

// Login handles user login and returns an AuthResponse with a Token
func Login(client *client.Client, username, password string) (*AuthResponse, error) {
	endpoint := "/auth/login"
	body := map[string]string{
		"username": username,
		"password": password,
	}

	data, err := client.SendRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}

	var authResp AuthResponse
	if err := json.Unmarshal(data, &authResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &authResp, nil
}
```

#### 创建 `fs/file.go`

```go
package fs

import (
	"fmt"
	"github.com/yourusername/go-alist-sdk/client"
)

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// UploadFile handles file upload to Alist
func UploadFile(c *client.Client, filePath string, fileContent []byte) (*File, error) {
	endpoint := "/fs/upload"
	body := map[string]interface{}{
		"file_path":   filePath,
		"file_content": fileContent,
	}

	data, err := c.SendRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}

	var fileResp File
	if err := json.Unmarshal(data, &fileResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &fileResp, nil
}
```

### 7. **编写测试代码**

在 `auth`、`fs` 和其他模块中，编写测试代码以确保 API 调用按预期工作。你可以使用 Go 的 `testing` 包进行单元测试。

#### 例如，创建 `auth/auth_test.go`

```go
package auth

import (
	"testing"
	"github.com/yourusername/go-alist-sdk/client"
)

func TestLogin(t *testing.T) {
	client := client.NewClient("https://api.alist.com", "")
	authResp, err := Login(client, "testuser", "password123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if authResp.Token == "" {
		t.Fatalf("expected token, got empty string")
	}
}
```

### 8. **运行项目**

确保你已经完成了所有的实现和依赖安装，接下来就可以通过以下命令来运行或测试你的代码：

- **运行测试**

```bash
go test ./...
```

- **构建项目**

```bash
go build -o go-alist-sdk
```

- **运行项目**

```bash
go run main.go
```

### 9. **创建 Git 仓库**

在项目根目录下初始化 Git 仓库（如果你还没有的话）：

```bash
git init
git add .
git commit -m "Initial commit"
```

你可以将项目推送到 GitHub 或其他代码托管平台：

```bash
git remote add origin https://github.com/yourusername/go-alist-sdk.git
git push -u origin master
```

### 10. **版本控制和发布**

当你完成了开发，可以使用 Git 标签来发布版本：

```bash
git tag v1.0.0
git push origin v1.0.0
```

这样，SDK 就可以被其他开发者使用，并且你可以在需要时进行版本更新。

---

通过以上步骤，你就成功地初始化了一个 Go SDK 项目，并且准备好开发和维护 Alist 的 API 集成。
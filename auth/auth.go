package auth

import (
	"encoding/json"
	"fmt"

	"github.com/Singert/go-alist-sdk/client"
)

// Login 进行用户验证,并返回 JWT TOKEN
func Login(c *client.Client, username, password, optcode string) (string, error) {

	//设置请求endpoint
	endpoint := "/api/auth/login"

	//构建请求body
	body := loginReqest{
		Username: username,
		Password: password,
		Optcode:  optcode,
	}

	//发送请求
	resp, respBody, err := c.SendRequest("POST", endpoint, body)
	if err != nil {
		return "", fmt.Errorf("failed to send request in login func: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to login, status: %d, response: %s", resp.StatusCode, string(respBody))
	}

	//处理response body
	var loginResp LoginResponse
	err = json.Unmarshal(respBody, &loginResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body in login func: %w", err)
	}
	if loginResp.Code != 200 {
		return "", fmt.Errorf("login failed, code: %d, message: %s", loginResp.Code, loginResp.Message)
	}
	fmt.Println("loginResp:", loginResp) //FIXME: remove this line and provide a better logging solution
	return loginResp.Data.Token, nil
}

// LoginWithHash 进行哈希密码验证,并返回 JWT TOKEN
func LoginWithHash(c *client.Client, username, hasdedPassword, optcode string) (string, error) {
	//设置请求endpoint
	endpoint := "/api/auth/login/hash"

	//构建请求body
	body := loginReqest{
		Username: username,
		Password: hasdedPassword,
		Optcode:  optcode,
	}

	//发送请求
	resp, respBody, err := c.SendRequest("POST", endpoint, body)
	if err != nil {
		return "", fmt.Errorf("failed to send request in login with hash func: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to login with hash, status: %d, response: %s", resp.StatusCode, string(respBody))
	}

	//处理response body
	var loginResp LoginResponse
	err = json.Unmarshal(respBody, &loginResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body in login func: %w", err)
	}

	if loginResp.Code != 200 {
		return "", fmt.Errorf("login failed, code: %d, message: %s", loginResp.Code, loginResp.Message)
	}

	fmt.Println("loginResp:", loginResp) //FIXME: remove this line and provide a better logging solution
	return loginResp.Data.Token, nil
}

// Generate2FA 生成2FA密钥
func Generate2FA(c *client.Client, token string) (*TwoFAResponse, error) {
	//设置请求endpoint
	endpoint := "/api/auth/2fa/generate"

	//发送请求
	resp, respBody, err := c.SendAuthRequest("POST", endpoint, token, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request in generate 2fa func: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to generate 2fa, status: %d, response: %s", resp.StatusCode, string(respBody))
	}

	//处理response body
	var twoFAResp TwoFAResponse
	err = json.Unmarshal(respBody, &twoFAResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body in generate 2fa func: %w", err)
	}

	if twoFAResp.Code != 200 {
		return nil, fmt.Errorf("generate 2fa failed, code: %d, message: %s", twoFAResp.Code, twoFAResp.Message)
	}

	fmt.Println("twoFAResp:", twoFAResp) //FIXME: remove this line and provide a better logging solution
	return &twoFAResp, nil
}

// VerifyTwoFA 验证2FA码
func VerifyTwoFA(c *client.Client, token, code, secret string) error {
	//设置请求endpoint
	endpoint := "/api/auth/2fa/verify"

	//构建请求body
	body := VerifyTwoFAReqest{
		Code:   code,
		Sercet: secret,
	}

	//发送请求
	resp, respBody, err := c.SendAuthRequest("POST", endpoint, token, body)
	if err != nil {
		return fmt.Errorf("failed to send request in verify 2fa func: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to verify 2fa, status: %d, response: %s", resp.StatusCode, string(respBody))
	}

	//处理response body
	var verifyResp GenericResponse
	err = json.Unmarshal(respBody, &verifyResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body in verify 2fa func: %w", err)
	}

	if verifyResp.Code != 200 {
		return fmt.Errorf("verify 2fa failed, code: %d, message: %s", verifyResp.Code, verifyResp.Message)
	}

	fmt.Println("genericResp:", verifyResp) //FIXME: remove this line and provide a better logging solution
	return nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *client.Client, token string) (*UserInfo, error) {
	//设置请求endpoint
	endpoint := "/api/me"

	//发送请求
	resp, respBody, err := c.SendAuthRequest("GET", endpoint, token, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request in get user info func: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get user info, status: %d, response: %s", resp.StatusCode, string(respBody))
	}

	//处理response body
	var userResp UserResponse
	err = json.Unmarshal(respBody, &userResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body in get user info func: %w", err)
	}

	if userResp.Code != 200 {
		return nil, fmt.Errorf("failed to fetch user info: %s", userResp.Message)
	}

	fmt.Println("userInfo:", &userResp.Data) //FIXME: remove this line and provide a better logging solution
	return &userResp.Data, nil
}

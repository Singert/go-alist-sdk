package auth

import (
	"github.com/Singert/go-alist-sdk/client"
	"testing"
)

func TestLogin(t *testing.T) {
	c := client.NewClient("https://api.alist.com", "")
	token, err := Login(c, "testuser", "testpassword", "")
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	if token == "" {
		t.Fatalf("Expected token, got empty string")
	}
}

func TestLoginWithHash(t *testing.T) {
	c := client.NewClient("https://api.alist.com", "")
	token, err := LoginWithHash(c, "testuser", "hashedpassword", "")
	if err != nil {
		t.Fatalf("Login with hash failed: %v", err)
	}
	if token == "" {
		t.Fatalf("Expected token, got empty string")
	}
}

func TestGenerate2FA(t *testing.T) {
	c := client.NewClient("https://api.alist.com", "valid_token")
	twoFA, err := Generate2FA(c, "valid_token")
	if err != nil {
		t.Fatalf("2FA generation failed: %v", err)
	}
	if twoFA.Data.Sercet == "" {
		t.Fatalf("Expected 2FA secret, got empty string")
	}
}

func TestGetUserInfo(t *testing.T) {
	c := client.NewClient("https://api.alist.com", "valid_token")
	user, err := GetUserInfo(c, "valid_token")
	if err != nil {
		t.Fatalf("Get user info failed: %v", err)
	}
	if user.Username == "" {
		t.Fatalf("Expected username, got empty string")
	}
}

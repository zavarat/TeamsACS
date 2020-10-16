package mfa

import (
	"fmt"
	"testing"
)


func initAuth(user string) (secret, code string) {
	ng := NewGoogleAuth()
    secret = ng.GetSecret()
    fmt.Println("Secret:", secret)
    // Dynamic code (a 6-digit number is dynamically generated every 30s)
    code, err := ng.GetCode(secret)
    fmt.Println("Code:", code, err)
    // Username
    qrCode := ng.GetQrcode(user, code, "TeamsAcsDemo")
    fmt.Println("Qrcode", qrCode)
    return
}

func TestGoogleAuth_VerifyCode(t *testing.T) {

    // fmt.Println("-----------------开启二次认证----------------------")
    user := "testxxx@google.com"
    secret, code := initAuth(user)
    t.Log(secret, code)
    t.Log("Information Validation")
    // Authentication, dynamic code (from Google Authenticator or freeotp)
    bool, err := NewGoogleAuth().VerifyCode(secret, code)
    if bool {
        t.Log("√")
    } else {
        t.Fatal("X", err)
    }
}


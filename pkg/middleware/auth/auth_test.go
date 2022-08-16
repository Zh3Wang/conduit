package auth

import "testing"

func TestGenerateJwtToken(t *testing.T) {
	token, err := GenerateJwtToken("secret", "wangzhe", "email")
	if err != nil {
		t.Errorf("jwt err: %s", err.Error())
		return
	}
	t.Logf("token: %s", token)
}

func TestParseJwtToken(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Indhbmd6aGUiLCJlbWFpbCI6ImVtYWlsIiwiZXhwIjoxNjYwNTQ5ODgwLCJpc3MiOiJ0ZXN0In0.iVu9V8kRPMZkCwgkRHzhQHVyB7yA0_uVmlvDrkDnE0M`
	claims, err := ParseJwtToken("secret", token)
	if err != nil {
		t.Errorf("jwt err: %s", err.Error())
		return
	}
	t.Log(claims.StandardClaims.ExpiresAt, claims.UserName, claims.Email)
}

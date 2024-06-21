package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

type Payload struct {
	Aud        string `json:"aud,omitempty"`
	UserId     string `json:"userId,omitempty"`
	DepList    string `json:"depList,omitempty"` //多个逗号分隔
	UserName   string `json:"userName,omitempty"`
	TenantId   string `json:"tenantId,omitempty"`
	TenantName string `json:"tenantName,omitempty"`
	Domain     string `json:"domain,omitempty"`
	jwt.RegisteredClaims
}

func (r *Payload) String() string {
	return fmt.Sprintf("clientId=%s userId=%s tuserName=%s depList=%s tenantId=%s tenantName=%s domain=%s exp=%s", r.Aud,
		r.UserId, r.UserName, r.DepList, r.TenantId, r.TenantName, r.Domain, r.RegisteredClaims.ExpiresAt)
}

func ParseJWT(token string) (*Payload, error) {
	claims := &Payload{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("123456"), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func TestParseJWT(t *testing.T) {
	pay, err := ParseJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIyMWZjZWExOGIwMzMiLCJkZXBMaXN0IjoiIiwiZG9tYWluIjoiIiwiZXhwIjoxNzE4OTQyNjA2LCJzdWIiOiIiLCJ0ZW5hbnRDb2RlIjoiNmUyYWQ4ZGYiLCJ0ZW5hbnRJZCI6IjEwMDAwMDUiLCJ0ZW5hbnROYW1lIjoic2t5Z3VhcmQiLCJ1c2VySWQiOiIiLCJ1c2VyTmFtZSI6InpoeXRlc3QifQ.cMewgIHPd8x4EA0bF6cXHUQdzuI9UdAv5K_hSC3-JSs")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", pay)
}

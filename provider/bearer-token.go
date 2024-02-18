package provider

import (
	"context"
	"net/http"
)

type BearerToken struct {
	token string
}

func NewAuthProvider(token string) *BearerToken {
	return &BearerToken{token: token}
}

func (s *BearerToken) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("X-API-Key", s.token)
	return nil
}

package auth

import (
	"context"
	"errors"
	"testing"
	"time"
)

var testSecret = []byte("0123456789abcdef0123456789abcdef")

func TestJWTServiceIssuesAndVerifiesToken(t *testing.T) {
	now := time.Date(2026, 5, 28, 10, 0, 0, 0, time.UTC)
	service, err := NewJWTService(JWTConfig{
		Secret: testSecret,
		TTL:    time.Hour,
		Issuer: "go-scaffold-test",
		Clock:  func() time.Time { return now },
	})
	if err != nil {
		t.Fatalf("NewJWTService() error = %v", err)
	}

	token, err := service.Issue(context.Background(), Claims{Subject: "42", Username: "ada"})
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if token.Value == "" || token.Type != BearerTokenType || !token.ExpiresAt.Equal(now.Add(time.Hour)) {
		t.Fatalf("token = %#v, want bearer token with configured ttl", token)
	}

	claims, err := service.Verify(context.Background(), token.Value)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if claims.Subject != "42" || claims.Username != "ada" {
		t.Fatalf("claims = %#v, want subject 42 and username ada", claims)
	}
}

func TestJWTServiceRejectsTamperedToken(t *testing.T) {
	service, err := NewJWTService(JWTConfig{Secret: testSecret, TTL: time.Hour})
	if err != nil {
		t.Fatalf("NewJWTService() error = %v", err)
	}
	token, err := service.Issue(context.Background(), Claims{Subject: "1"})
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	tampered := token.Value[:len(token.Value)-1] + "x"

	if _, err := service.Verify(context.Background(), tampered); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("Verify(tampered) error = %v, want ErrInvalidToken", err)
	}
}

func TestJWTServiceRejectsExpiredToken(t *testing.T) {
	now := time.Date(2026, 5, 28, 10, 0, 0, 0, time.UTC)
	service, err := NewJWTService(JWTConfig{
		Secret: testSecret,
		TTL:    time.Minute,
		Clock:  func() time.Time { return now },
	})
	if err != nil {
		t.Fatalf("NewJWTService() error = %v", err)
	}
	token, err := service.Issue(context.Background(), Claims{Subject: "1"})
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	now = now.Add(2 * time.Minute)

	if _, err := service.Verify(context.Background(), token.Value); !errors.Is(err, ErrExpiredToken) {
		t.Fatalf("Verify(expired) error = %v, want ErrExpiredToken", err)
	}
}

func TestJWTServiceRequiresStrongSecret(t *testing.T) {
	if _, err := NewJWTService(JWTConfig{Secret: []byte("short")}); !errors.Is(err, ErrTokenSecretTooWeak) {
		t.Fatalf("NewJWTService(short secret) error = %v, want ErrTokenSecretTooWeak", err)
	}
}

func TestGenerateSecretUsesMinimumSize(t *testing.T) {
	secret, err := GenerateSecret(1)
	if err != nil {
		t.Fatalf("GenerateSecret() error = %v", err)
	}
	if len(secret) != MinTokenSecretBytes {
		t.Fatalf("secret len = %d, want %d", len(secret), MinTokenSecretBytes)
	}
}

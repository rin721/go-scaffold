package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// DefaultTokenTTL is used when JWTConfig.TTL is not positive.
	DefaultTokenTTL = 24 * time.Hour
	// MinTokenSecretBytes is the minimum HMAC secret length accepted by JWTService.
	MinTokenSecretBytes = 32
	// BearerTokenType is the token type returned by JWTService.
	BearerTokenType = "Bearer"
)

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("token expired")
	ErrInvalidTokenClaims = errors.New("invalid token claims")
	ErrTokenSecretTooWeak = errors.New("token secret must be at least 32 bytes")
)

// Token is an issued authentication token.
type Token struct {
	Value     string
	Type      string
	ExpiresAt time.Time
}

// Claims are the stable public claims understood by token services.
type Claims struct {
	Subject   string
	Username  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

// TokenService signs and verifies authentication tokens.
type TokenService interface {
	Issue(ctx context.Context, claims Claims) (Token, error)
	Verify(ctx context.Context, token string) (Claims, error)
}

// JWTConfig configures JWTService.
type JWTConfig struct {
	Secret []byte
	TTL    time.Duration
	Issuer string
	Clock  func() time.Time
}

// JWTService signs and verifies HMAC SHA-256 JWTs.
type JWTService struct {
	secret []byte
	ttl    time.Duration
	issuer string
	clock  func() time.Time
}

type jwtClaims struct {
	Username string `json:"username,omitempty"`
	jwt.RegisteredClaims
}

// NewJWTService creates a JWT token service.
func NewJWTService(cfg JWTConfig) (*JWTService, error) {
	if len(cfg.Secret) < MinTokenSecretBytes {
		return nil, ErrTokenSecretTooWeak
	}
	ttl := cfg.TTL
	if ttl <= 0 {
		ttl = DefaultTokenTTL
	}
	clock := cfg.Clock
	if clock == nil {
		clock = func() time.Time { return time.Now().UTC() }
	}
	return &JWTService{
		secret: append([]byte(nil), cfg.Secret...),
		ttl:    ttl,
		issuer: strings.TrimSpace(cfg.Issuer),
		clock:  clock,
	}, nil
}

// GenerateSecret returns cryptographically random bytes suitable for JWTConfig.Secret.
func GenerateSecret(size int) ([]byte, error) {
	if size < MinTokenSecretBytes {
		size = MinTokenSecretBytes
	}
	secret := make([]byte, size)
	if _, err := rand.Read(secret); err != nil {
		return nil, fmt.Errorf("generate auth token secret: %w", err)
	}
	return secret, nil
}

// Issue signs claims into a bearer JWT.
func (s *JWTService) Issue(ctx context.Context, claims Claims) (Token, error) {
	if err := ctx.Err(); err != nil {
		return Token{}, err
	}
	subject := strings.TrimSpace(claims.Subject)
	if subject == "" {
		return Token{}, ErrInvalidTokenClaims
	}
	now := s.now()
	issuedAt := claims.IssuedAt.UTC()
	if issuedAt.IsZero() {
		issuedAt = now
	}
	expiresAt := claims.ExpiresAt.UTC()
	if expiresAt.IsZero() {
		expiresAt = issuedAt.Add(s.ttl)
	}
	if !expiresAt.After(issuedAt) {
		return Token{}, ErrInvalidTokenClaims
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		Username: strings.TrimSpace(claims.Username),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
	signed, err := jwtToken.SignedString(s.secret)
	if err != nil {
		return Token{}, fmt.Errorf("sign jwt token: %w", err)
	}
	return Token{Value: signed, Type: BearerTokenType, ExpiresAt: expiresAt}, nil
}

// Verify validates a bearer JWT and returns its public claims.
func (s *JWTService) Verify(ctx context.Context, tokenValue string) (Claims, error) {
	if err := ctx.Err(); err != nil {
		return Claims{}, err
	}
	tokenValue = strings.TrimSpace(tokenValue)
	if tokenValue == "" {
		return Claims{}, ErrInvalidToken
	}
	claims := &jwtClaims{}
	options := []jwt.ParserOption{
		jwt.WithExpirationRequired(),
		jwt.WithTimeFunc(s.now),
	}
	if s.issuer != "" {
		options = append(options, jwt.WithIssuer(s.issuer))
	}
	parsed, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, ErrInvalidToken
		}
		return s.secret, nil
	}, options...)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return Claims{}, ErrExpiredToken
		}
		return Claims{}, ErrInvalidToken
	}
	if parsed == nil || !parsed.Valid {
		return Claims{}, ErrInvalidToken
	}
	subject := strings.TrimSpace(claims.Subject)
	if subject == "" || claims.ExpiresAt == nil {
		return Claims{}, ErrInvalidToken
	}
	out := Claims{
		Subject:   subject,
		Username:  strings.TrimSpace(claims.Username),
		ExpiresAt: claims.ExpiresAt.Time.UTC(),
	}
	if claims.IssuedAt != nil {
		out.IssuedAt = claims.IssuedAt.Time.UTC()
	}
	return out, nil
}

func (s *JWTService) now() time.Time {
	if s == nil || s.clock == nil {
		return time.Now().UTC()
	}
	return s.clock().UTC()
}

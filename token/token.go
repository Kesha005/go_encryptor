// pkg/jwt/jwt.go
package jwt

import (
	"errors"
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidTokenType = errors.New("invalid token type")
)

type JWTConfig struct {
	Secret               string `mapstructure:"secret"`
	AccessTokenDuration  int    `mapstructure:"access_token_duration"`
	RefreshTokenDuration int    `mapstructure:"refresh_token_duration"`
	Issuer               string `mapstructure:"issuer"`
	Audience             string `mapstructure:"audience"`
}


// Token types
const (
	TypeAccess  = "access"
	TypeRefresh = "refresh"
)

type TokenMaker struct {
	secretKey string
	config    JWTConfig
}

type Payload struct {
	UserID    uint      `json:"user_id"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone"`
	TokenType string    `json:"token_type"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// New creates a new token maker
func New(config JWTConfig) *TokenMaker {
	return &TokenMaker{
		secretKey: config.Secret,
		config:    config,
	}
}

// CreateAccessToken creates a new access token
func (m *TokenMaker) CreateAccessToken(userID uint, phone, role string) (string, error) {
	now := time.Now()

	payload := &Payload{
		UserID:    userID,
		Phone:     phone,
		TokenType: TypeAccess,
		Role:      role,
		IssuedAt:  now,
		ExpiresAt: now.Add(time.Duration(m.config.AccessTokenDuration) * time.Second),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    payload.UserID,
		"phone":      payload.Phone,
		"role":       payload.Role,
		"token_type": payload.TokenType,
		"issued_at":  payload.IssuedAt.Unix(),
		"expires_at": payload.ExpiresAt.Unix(),
	})

	return token.SignedString([]byte(m.secretKey))
}

// CreateRefreshToken creates a new refresh token
func (m *TokenMaker) CreateRefreshToken(userID uint, phone, role string) (string, error) {
	now := time.Now()
	payload := &Payload{
		UserID:    userID,
		Phone:     phone,
		Role:      role,
		TokenType: TypeRefresh,
		IssuedAt:  now,
		ExpiresAt: now.Add(time.Duration(m.config.RefreshTokenDuration) * time.Second),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    payload.UserID,
		"phone":      payload.Phone,
		"role":       payload.Role,
		"token_type": payload.TokenType,
		"issued_at":  payload.IssuedAt.Unix(),
		"expires_at": payload.ExpiresAt.Unix(),
	})

	return token.SignedString([]byte(m.secretKey))
}

// Verify verifies and returns token payload
func (m *TokenMaker) Verify(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}

	tokenType, ok := claims["token_type"].(string)
	if !ok {
		tokenType = TypeAccess // Default to access for backward compatibility
	}

	phone, _ := claims["phone"].(string) // Optional for refresh tokens

	issuedAt, ok := claims["issued_at"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}

	expiresAt, ok := claims["expires_at"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}
	expTime := time.Unix(int64(expiresAt), 0)
	if expTime.Before(time.Now()) {
		return nil, ErrExpiredToken
	}

	return &Payload{
		UserID:    uint(userID),
		Phone:     phone,
		Role:      role,
		TokenType: tokenType,
		IssuedAt:  time.Unix(int64(issuedAt), 0),
		ExpiresAt: time.Unix(int64(expiresAt), 0),
	}, nil
}

// VerifyAccessToken verifies and ensures it's an access token
func (m *TokenMaker) VerifyAccessToken(tokenString string) (*Payload, error) {
	payload, err := m.Verify(tokenString)
	if err != nil {
		return nil, err
	}

	if payload.TokenType != TypeAccess {
		return nil, ErrInvalidTokenType
	}

	return payload, nil
}

// VerifyRefreshToken verifies and ensures it's a refresh token
func (m *TokenMaker) VerifyRefreshToken(tokenString string) (*Payload, error) {
	payload, err := m.Verify(tokenString)
	if err != nil {
		return nil, err
	}

	if payload.TokenType != TypeRefresh {
		return nil, ErrInvalidTokenType
	}

	return payload, nil
}

// IsAccessToken checks if token is an access token
func (m *TokenMaker) IsAccessToken(tokenString string) (bool, error) {
	payload, err := m.Verify(tokenString)
	if err != nil {
		return false, err
	}
	return payload.TokenType == TypeAccess, nil
}

// IsRefreshToken checks if token is a refresh token
func (m *TokenMaker) IsRefreshToken(tokenString string) (bool, error) {
	payload, err := m.Verify(tokenString)
	if err != nil {
		return false, err
	}
	return payload.TokenType == TypeRefresh, nil
}

func (m *TokenMaker) GetAccessTokenDuration() int {
	return m.config.AccessTokenDuration
}

func (m *TokenMaker) GetRefreshTokenDuration() int {
	return m.config.RefreshTokenDuration
}

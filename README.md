# Go JWT

Lightweight JWT access & refresh token package for Go, built on `golang-jwt/jwt/v5`.

# Installing
go get github.com/Kesha005/go_jwt

*(adjust the module path to match your actual repo)*

# Features
- Access & refresh token generation
- Token type validation (prevents using a refresh token as an access token and vice versa)
- Expiration checking
- Config-driven (secret, durations, issuer, audience)

# Config

`TokenMaker` is built from a `JWTConfig` struct:

```go
type JWTConfig struct {
    Secret               string
    AccessTokenDuration  int // in seconds
    RefreshTokenDuration int // in seconds
    Issuer               string
    Audience             string
}
```

No `.env` loading is done inside the package — just build the config yourself (from env vars, a config file, viper, whatever) and pass it in.

# Usage

## Init

```go
import "github.com/Kesha005/go_jwt"

cfg := jwt.JWTConfig{
    Secret:               "your-secret-key",
    AccessTokenDuration:  900,    // 15 min
    RefreshTokenDuration: 604800, // 7 days
}

maker := jwt.New(cfg)
```

## Creating tokens

```go
accessToken, err := maker.CreateAccessToken(userID, phone, role)
if err != nil {
    // handle error
}

refreshToken, err := maker.CreateRefreshToken(userID, phone, role)
if err != nil {
    // handle error
}
```

## Verifying tokens

```go
// Generic verify — accepts either token type
payload, err := maker.Verify(accessToken)

// Strict verify — errors out if the wrong token type is passed
payload, err := maker.VerifyAccessToken(accessToken)
payload, err := maker.VerifyRefreshToken(refreshToken)
```

## Checking token type

```go
isAccess, err := maker.IsAccessToken(token)
isRefresh, err := maker.IsRefreshToken(token)
```

## Payload

```go
type Payload struct {
    UserID    uint
    Role      string
    Phone     string
    TokenType string
    IssuedAt  time.Time
    ExpiresAt time.Time
}
```

## Errors

```go
jwt.ErrInvalidToken     // malformed/invalid signature/invalid claims
jwt.ErrExpiredToken     // token expired
jwt.ErrInvalidTokenType // wrong token type passed to a strict verify method
```

# Notes to self
- `Verify` defaults `token_type` to `"access"` if it's missing, for backward compatibility with old tokens — keep that in mind if debugging weird type mismatches.
- Only HMAC signing methods are accepted; anything else is rejected as `ErrInvalidToken`.
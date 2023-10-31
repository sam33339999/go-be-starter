# JWT

## Base on golang-jwt @v5.0.0

- [registered_claims](https://github.com/golang-jwt/jwt/blob/v5/registered_claims.go)

```go
type RegisteredClaims struct {
	// the `iss` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1
	Issuer string `json:"iss,omitempty"`

	// the `sub` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2
	Subject string `json:"sub,omitempty"`

	// the `aud` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3
	Audience ClaimStrings `json:"aud,omitempty"`

	// the `exp` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
	ExpiresAt *NumericDate `json:"exp,omitempty"`

	// the `nbf` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5
	NotBefore *NumericDate `json:"nbf,omitempty"`

	// the `iat` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6
	IssuedAt *NumericDate `json:"iat,omitempty"`

	// the `jti` (JWT ID) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7
	ID string `json:"jti,omitempty"`
}
```

- iss: (Issuer) JWT 發行人
- sub: (Subject) JWT 所面向的用戶
- aud: (Audience) 接收 JWT 的一方
- exp: (Expires At) JWT 過期時間，過期時間必須大於發行時間 -> exp > iat
- nbf: (Not Before) JWT 可使用時間，在nbf 之前，該 JWT 都是不可用
- iat: (Issued At) JWT 發行時間
- jti: (JWT ID) JWT 的唯一身份標籤，主要用於一次性 token，從而迴避重放攻擊

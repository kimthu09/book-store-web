package tokenprovider

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	GenerateTokenForPayLoadEmail(data TokenPayloadEmail, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
	ValidateTokenForPayLoadEmail(token string) (*TokenPayloadEmail, error)
}

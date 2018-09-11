package main

// Authable is an interface for token Service
type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

// TokenService encodes and decodes tokens
type TokenService struct {
	repo Repository
}

// Decode decodes tokens
func (srv *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

// Encode encodes tokens
func (srv *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}

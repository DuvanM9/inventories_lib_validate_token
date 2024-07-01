package validate

import "errors"

var (
	ErrorInvalitToken  = errors.New("invalit token")
	ErrorInvalitClaims = errors.New("invalit claims")
)

package env

import "errors"

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

var ErrInvalidEnv = errors.New("invalid env")

func New(env string) (Env, error) {
	if env != string(Development) && env != string(Production) {
		return "", ErrInvalidEnv
	}

	return Env(env), nil
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsDev() bool {
	return e == Development
}

func (e Env) IsProd() bool {
	return e == Production
}

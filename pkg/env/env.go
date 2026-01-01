package env

import "errors"

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

var ErrInvalidEnv = errors.New("invalid env")

// New parses the provided env string and returns the corresponding Env.
// The input must match the string form of either Development or Production.
// On success it returns the Env and a nil error; otherwise it returns ErrInvalidEnv.
func New(env string) (Env, error) {
	if env != string(Development) && env != string(Production) {
		return "", ErrInvalidEnv
	}

	return Env(env), nil
}

// String returns the Env value as a string.
// It implements the fmt.Stringer interface.
func (e Env) String() string {
	return string(e)
}

// IsDev returns true if the Env is Development; false otherwise.
func (e Env) IsDev() bool {
	return e == Development
}

// IsProd returns true if the Env is Production; false otherwise.
func (e Env) IsProd() bool {
	return e == Production
}

package utils

type StringStringer string

func (s StringStringer) String() string {
	return string(s)
}

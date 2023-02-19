package ast

type Env map[string]interface{}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) Set(key string, value interface{}) {
	(*e)[key] = value
}

package monitors

type Monitor interface {
	Type() string
	ToMap() map[string]interface{}
	Start() error
}

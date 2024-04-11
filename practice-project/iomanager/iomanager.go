package iomanager

type IOmanager interface {
	GetInputFromFile() ([]string, []string, error)
	WriteJson(data interface{}) error
}

package domain

type Quiz struct {
	ID         string
	Code       string
	Name       string
	Categories []string
	Type       string
	Level      int
	Questions  interface{}
}

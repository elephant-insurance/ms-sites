package models

type Project struct {
	Name  string
	Files map[string][]byte
}

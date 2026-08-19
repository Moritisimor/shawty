package models

type URLAlias struct {
	ID       uint
	Alias    string
	URL      string
	DeleteAt int64
}

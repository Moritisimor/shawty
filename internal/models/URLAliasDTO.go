package models

type URLAliasDTO struct {
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

func (u URLAliasDTO) IsValid() bool {
	return u.Alias != "" && u.URL != ""
}

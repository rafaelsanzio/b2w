package models

//ResponseSwapi model
type ResponseSwapi struct {
	Result []Result `json:"results"`
}

//Result model in response swapi
type Result struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

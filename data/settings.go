package data

type Settings struct {
	PageRoot        string                  `json:"pageRoot"`
	Languages       map[string]LanguageInfo `json:"languages"`
	DefaultLanguage string                  `json:"defaultLanguage"`
}

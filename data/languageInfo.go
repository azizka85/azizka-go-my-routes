package data

import i18n "github.com/azizka85/azizka-go-i18n"

type LanguageInfo struct {
	Image      string           `json:"image"`
	Label      string           `json:"label"`
	Translator *i18n.Translator `json:"-"`
}

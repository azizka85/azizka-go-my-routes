package global

import (
	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/azizka85/azizka-go-my-routes/translations"
)

var Settings = data.Settings{
	PageRoot: "/",
	Languages: map[string]data.LanguageInfo{
		"kz": {
			Image:      "images/flags/kz.svg",
			Label:      "Қазақша",
			Translator: i18n.CreateTranslator(&translations.KZ),
		},
		"ru": {
			Image:      "images/flags/ru.svg",
			Label:      "Русский",
			Translator: i18n.CreateTranslator(&translations.RU),
		},
		"en": {
			Image:      "images/flags/en.svg",
			Label:      "English",
			Translator: i18n.CreateTranslator(&translations.EN),
		},
	},
	DefaultLanguage: "kz",
}

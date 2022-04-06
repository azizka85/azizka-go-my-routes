package settings

import "github.com/azizka85/azizka-go-my-routes/data"

var GlobalSettings = data.Settings{
	PageRoot: "/",
	Languages: map[string]data.LanguageInfo{
		"kz": {
			Image: "/images/flags/kz.svg",
			Label: "Қазақша",
		},
		"ru": {
			Image: "/images/flags/ru.svg",
			Label: "Русский",
		},
		"en": {
			Image: "/images/flags/en.svg",
			Label: "English",
		},
	},
	DefaultLanguage: "kz",
}

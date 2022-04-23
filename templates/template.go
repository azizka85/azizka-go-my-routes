package templates

import (
	"net/url"

	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/azizka85/azizka-go-my-routes/templates/layouts"
)

func RenderPage(
	ajax bool,
	lang string,
	content string,
	layoutNames []string,
	query url.Values,
	settings *data.Settings,
	translator *i18n.Translator,
) string {
	for _, layoutName := range layoutNames {
		switch layoutName {
		case "main-layout":
			content = layouts.Main(
				query,
				query.Has("main-layout-navigation"),
				settings.PageRoot,
				lang,
				content,
				translator,
			)
		}
	}

	if !ajax {
		content = layouts.Default(
			lang,
			settings.PageRoot,
			content,
			translator,
		)
	}

	return content
}

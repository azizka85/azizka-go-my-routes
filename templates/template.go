package templates

import (
	"net/http"
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
	request *http.Request,
	user *data.User,
	settings *data.Settings,
	translator *i18n.Translator,
) string {
	for _, layoutName := range layoutNames {
		switch layoutName {
		case "main-layout":
			content = layouts.Main(
				query.Has("main-layout-navigation"),
				settings.PageRoot,
				lang,
				content,
				query,
				request,
				user,
				settings,
				translator,
			)
		}
	}

	if !ajax {
		content = layouts.Default(
			lang,
			settings.PageRoot,
			content,
			settings,
			translator,
		)
	}

	return content
}

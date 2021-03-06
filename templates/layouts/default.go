// This file is generated by gorazor 1.2.2
// DON'T modified manually
// Should edit source file and re-generate: templates/layouts/default.gohtml

package layouts

import (
	"encoding/json"
	"io"
	"strings"

	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/sipin/gorazor/gorazor"
)

// Default generates templates/layouts/default.gohtml
func Default(lang string, pageRoot string, content string, settings *data.Settings, translator *i18n.Translator) string {
	var _b strings.Builder
	RenderDefault(&_b, lang, pageRoot, content, settings, translator)
	return _b.String()
}

// RenderDefault render templates/layouts/default.gohtml
func RenderDefault(_buffer io.StringWriter, lang string, pageRoot string, content string, settings *data.Settings, translator *i18n.Translator) {
	// Line: 15
	_buffer.WriteString("\n<!DOCTYPE html>\n<html lang=\"")
	// Line: 17
	_buffer.WriteString(gorazor.HTMLEscStr(lang))
	// Line: 17
	_buffer.WriteString("\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <title>\n    ")
	// Line: 23
	_buffer.WriteString(gorazor.HTMLEscStr(translator.Translate("My Routes")))
	// Line: 23
	_buffer.WriteString("\n  </title>\n  <link rel=\"icon\" type=\"image/png\" href='")
	// Line: 25
	_buffer.WriteString(gorazor.HTMLEscStr((pageRoot)))
	// Line: 25
	_buffer.WriteString("favicon.png'>\n  <link rel=\"stylesheet\" href=\"")
	// Line: 26
	_buffer.WriteString(gorazor.HTMLEscStr((pageRoot)))
	// Line: 26
	_buffer.WriteString("dist/css/main.css\">\n</head>\n<body>\n  <div class=\"splash\">\n    <div class=\"loader\" data-page=\"loader-page\">\n      <div class=\"loader-container\">\n        <img src=\"")
	// Line: 32
	_buffer.WriteString(gorazor.HTMLEscStr((pageRoot)))
	// Line: 32
	_buffer.WriteString("favicon.png\">\n        <div class=\"loader-container-progress\"></div>\n      </div>\n    </div>\n  </div>\n  <script>\n    ")

	settingsTxt, err := json.Marshal(settings)

	if err != nil {
		settingsTxt = []byte("{}")
	}

	// Line: 44
	_buffer.WriteString("\n    window.settings = ")
	// Line: 45
	_buffer.WriteString((string(settingsTxt)))
	// Line: 45
	_buffer.WriteString(";\n\n    const splashElem = document.querySelector('.splash');\n\n    splashElem?.classList.add('splash-open');\n  </script>\n  ")
	// Line: 51
	_buffer.WriteString((content))
	// Line: 51
	_buffer.WriteString("\n  <script src=\"")
	// Line: 52
	_buffer.WriteString(gorazor.HTMLEscStr((pageRoot)))
	// Line: 52
	_buffer.WriteString("dist/js/main.js\" type=\"module\"></script>\n</body>\n</html>")

}

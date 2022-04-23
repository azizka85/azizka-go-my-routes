package layouts

import (
	"bytes"
	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/sipin/gorazor/gorazor"
)

func Default(lang string, pageRoot string, content string, translator *i18n.Translator) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<!DOCTYPE html>\n<html lang=\"")
	_buffer.WriteString(gorazor.HTMLEscape(lang))
	_buffer.WriteString("\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <title>\n    ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("My Routes", nil, nil, nil)))
	_buffer.WriteString("\n  </title>\n  <link rel=\"icon\" type=\"image/png\" href='")
	_buffer.WriteString(gorazor.HTMLEscape((pageRoot)))
	_buffer.WriteString("favicon.png'>\n  <link rel=\"stylesheet\" href=\"")
	_buffer.WriteString(gorazor.HTMLEscape((pageRoot)))
	_buffer.WriteString("dist/css/main.css\">\n</head>\n<body>\n  <div class=\"splash\">\n    <div class=\"loader\" data-page=\"loader-page\">\n      <div class=\"loader-container\">\n        <img src=\"")
	_buffer.WriteString(gorazor.HTMLEscape((pageRoot)))
	_buffer.WriteString("favicon.png\">\n        <div class=\"loader-container-progress\"></div>\n      </div>\n    </div>\n  </div>\n  <script>\n    const splashElem = document.querySelector('.splash');\n\n    splashElem?.classList.add('splash-open');\n  </script>\n  ")
	_buffer.WriteString((content))
	_buffer.WriteString("\n  <script src=\"")
	_buffer.WriteString(gorazor.HTMLEscape((pageRoot)))
	_buffer.WriteString("dist/js/main.js\" type=\"module\"></script>\n</body>\n</html>")

	return _buffer.String()
}

package pages

import (
	"bytes"

	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/sipin/gorazor/gorazor"
)

func SignUp(pageRoot string, lang string, authServiceComponent string, translator *i18n.Translator) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<div data-page=\"signup-page\">\n  <div class=\"main-card\">\n    <div class=\"card main-card-body\">\n      <div class=\"card-body\">\n        <h2 \n          data-title=\"main\"\n          style=\"text-transform: uppercase; font-weight: lighter;\"\n        >\n          ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Sign Up")))
	_buffer.WriteString("\n        </h2>\n        <form method=\"post\" class=\"mb-1\">\n          <div class=\"form-item mb-1\">\n            <label \n              for=\"full-name\" \n              class=\"form-label\"            \n            >\n              <input \n                type=\"text\" \n                name=\"fullName\" \n                id=\"full-name\" \n                class=\"form-control\" \n                placeholder='")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Name")))
	_buffer.WriteString("*'\n                required\n              >\n              <span \n                id=\"full-name-label\"\n              >\n                ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Name")))
	_buffer.WriteString("*\n              </span>            \n            </label>          \n          </div>\n          <div class=\"form-item mb-1\">\n            <label \n              for=\"email\" \n              class=\"form-label\"\n            >\n              <input \n                type=\"email\" \n                name=\"email\" \n                id=\"email\" \n                class=\"form-control\" \n                placeholder=\"Email*\"\n                required\n              >          \n              <span>\n                Email*\n              </span>\n            </label>          \n          </div>\n          <div class=\"form-item mb-1\">\n            <label \n              for=\"password\" \n              class=\"form-label\"            \n            >\n              <input \n                type=\"password\" \n                name=\"password\" \n                id=\"password\" \n                class=\"form-control\" \n                placeholder='")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Password")))
	_buffer.WriteString("*'\n                required\n              >\n              <span \n                id=\"password-label\"\n              >\n                ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Password")))
	_buffer.WriteString("* \n              </span>\n            </label>          \n          </div>\n          <div class=\"form-item mb-1\">\n            <label \n              for=\"photo\" \n              class=\"form-label\"            \n            >\n              <input \n                type=\"text\" \n                name=\"photo\" \n                id=\"photo\" \n                class=\"form-control\" \n                placeholder=\"http://\"\n              >\n              <span \n                id=\"photo-label\"\n              >\n                ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Photo")))
	_buffer.WriteString("\n              </span>\n            </label>          \n          </div>\n          <div style=\"text-align: right;\" class=\"mb-1\">\n            <a \n              class=\"btn btn-light\" \n              href=\"")
	_buffer.WriteString(gorazor.HTMLEscape((pageRoot)))
	_buffer.WriteString(gorazor.HTMLEscape((lang)))
	_buffer.WriteString("/sign-in\"\n              data-button=\"sign-in\"\n            >\n              ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Sign In")))
	_buffer.WriteString("\n            </a>\n          </div>\n          <div style=\"text-align: right;\">\n            <button \n              type=\"submit\" \n              class=\"btn btn-success\"\n              data-button=\"sign-up\"\n            >\n              ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Sign Up")))
	_buffer.WriteString("\n            </button>\n            <a \n              class=\"btn btn-danger\" \n              href=\"")
	_buffer.WriteString(gorazor.HTMLEscape((pageRoot)))
	_buffer.WriteString(gorazor.HTMLEscape((lang)))
	_buffer.WriteString("\"\n              data-button=\"cancel\"\n            >\n              ")
	_buffer.WriteString(gorazor.HTMLEscape(translator.Translate("Cancel")))
	_buffer.WriteString("\n            </a>\n          </div>\n        </form>\n        ")
	_buffer.WriteString((authServiceComponent))
	_buffer.WriteString(" \n      </div>\n    </div>\n  </div>  \n</div>")

	return _buffer.String()
}

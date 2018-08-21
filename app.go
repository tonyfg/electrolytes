// +build !dev

package main

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/gobuffalo/packr"
	"github.com/zserge/webview"
)

func main() {
	box := packr.NewBox("web/public")
	html := box.String("index.html")
	js := box.String("app.js")
	css := box.String("app.css")

	// Open the packed assets inside webview
	window := webview.New(webview.Settings{
		URL:       "data:text/html," + url.PathEscape(html),
		Width:     800,
		Height:    600,
		Resizable: true,
	})

	window.Dispatch(func() {
		// Inject CSS
		window.Eval(fmt.Sprintf(`(function(css) {
			var style = document.createElement('style');
			var head = document.head || document.getElementsByTagName('head')[0];
			style.setAttribute('type', 'text/css');
			if (style.styleSheet) {
				style.styleSheet.cssText = css;
			} else {
				style.appendChild(document.createTextNode(css));
			}
			head.appendChild(style);
		})("%s")`, template.JSEscapeString(css)))

		// Inject JS
		window.Eval(js)
		window.Eval("require('index')")
	})

	window.Run()
}

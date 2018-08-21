# Electrolytes - it's what plants crave!

## What?

Electrolytes is a lightweight HTML/JS/CSS environment for developing desktop apps. By the power of the great go-webview library we can have something electron-like, but without the heavy resource usage.

## How (can I use it)?

You need to install the following dependencies:
- Node.js
- Yarn
- Golang 1.X
- Mage (optional)

Extra dependencies needed for Linux:
- GTK3
- GtkWebkit2


After installing dependencies, just clone this repo and run
`mage`
If you haven't installed mage, you can run it like this:
`go run mage.go`

This will start a development version of the app where any HTML/JS/CSS code is automatically reloaded when you save it.

Try `mage -l` for a list of build targets.


Submitting code to Electrolytes:

This repo uses a directory/file structure according to: https://github.com/golang-standards/project-layout

TODO:
- Cleanup web app to remove cruft that's unrelated to the skeleton
- Get a production build with all JS/CSS/HTML/Etc assets bundled inside the go binary

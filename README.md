Creating your app:

1. git clone this repo
2. go get -u -d github.com/magefile/mage
X2. go get github.com/mattn/goreman
X3. go get -u github.com/gobuffalo/packr/...

To launch the development version of the app:
- goreman start

To create a production build:
- ...

Developing:

This repo uses a directory/file structure according to: https://github.com/golang-standards/project-layout

TODO:
- Cleanup web app to remove cruft that's unrelated to the skeleton
- Get a production build with all JS/CSS/HTML/Etc assets bundled inside the go binary

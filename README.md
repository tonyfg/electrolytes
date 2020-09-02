你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
- Cross compilation support
- Neatly package production binaries for each platform
- Have an app generator like create-react-app instead of requiring people to clone the repo

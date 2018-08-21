// +build mage

package main

import (
	// "errors"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

////////////////////
// Public targets //
////////////////////

var Default = Dev

// Runs the webapp in development mode (changes to webapp files are auto-reloaded)
func Dev() error {
	cmd := exec.Command("yarn")
	cmd.Dir = "web"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	retval := cmd.Run()
	if retval != nil {
		return retval
	}

	cmd = exec.Command("go", "run", "app_dev.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go cmd.Run()

	cmd = exec.Command("yarn", "start")
	cmd.Dir = "web"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Builds the production app
func Build() error {
	mg.Deps(Clean, yarnBuild)
	os.MkdirAll("bin", 0755)

	cmd := exec.Command("go", "build", "-o", "bin/app", "app.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Installs the app in $GOPATH/bin
func Install() error {
	mg.Deps(Clean, yarnBuild)

	cmd := exec.Command("go", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Removes temporary and build files
func Clean() {
	os.RemoveAll("bin")
	os.RemoveAll("web/public")
}

/////////////////////
// Private targets //
/////////////////////

// Builds the HTML/JS/CSS bundles
func yarnBuild() error {
	cmd := exec.Command("yarn", "build")
	cmd.Dir = "web"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

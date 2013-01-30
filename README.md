blink1-go
=========
this project demonstrates interfacing to a usb device in go, specifically the [ThingM blink(1)](http://thingm.com/products/blink-1.html).
the original interfacing examples are written in c, so this project uses [CGO](http://golang.org/cmd/cgo/) to integrate c code into go.

## the setup / install ##

when writing this I was running Mac OSX 10.8.2

to run, you'll first need to install *libusb*.

1. install [Xcode](https://developer.apple.com/xcode/) and the [Command Line Tools for Xcode](https://developer.apple.com/downloads/index.action)
1. install [Homebrew](http://mxcl.github.com/homebrew/)
3. now use Homebrew to install *libusb*:

	$ brew install libusb-compat
	
4. make sure your $GOPATH is set to writable location by your user (i use the following location $HOME/go)
5. get and install the remote package

	$ go get github.com/heatxsink/blink1-go/blink1

4. *cd* to the directory with this project's files, then to build and run

	$ go run example.go

## derivatives ##
the code that interfaces *libusb* which is in C to Go is manually forked from GaryBoone's [GoBlink](https://github.com/GaryBoone/GoBlink) project.

## the license ##
this code is completely free under the mit license: [http://mit-license.org/](http://mit-license.org/).

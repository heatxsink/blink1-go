blink1-go
=========
This project demonstrates interfacing to a USB device in Go, specifically the [ThingM blink(1)](http://thingm.com/products/blink-1.html).
The original interfacing examples are written in C, so this project uses [CGO](http://golang.org/cmd/cgo/) to integrate C code into Go.

## Setup ##

When writing this I was running Mac OSX 10.8.2

To run, you'll first need to install *libusb*.

1. Install [Xcode](https://developer.apple.com/xcode/) and the [Command Line Tools for Xcode](https://developer.apple.com/downloads/index.action)
1. Install [Homebrew](http://mxcl.github.com/homebrew/)
3. Now use Homebrew to install *libusb*:

	$ brew install libusb-compat
	
4. *cd* to the directory with this project's files, then to build and run

	$ go run blink1.go

## Sources ##
The code that interfaces *libusb* which is in C to Go is manually forked from GaryBoone's [GoBlink](https://github.com/GaryBoone/GoBlink) project.

## License ##
This code is completely free under the MIT License: [http://mit-license.org/](http://mit-license.org/).
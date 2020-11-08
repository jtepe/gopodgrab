.PHONY: all install

VERSION = $(shell git log -n1 --no-merges --pretty="%h %cd")

all: install

install:
	go install -ldflags="-X 'github.com/jtepe/gopodgrab/cmd.Version=$(VERSION)'"

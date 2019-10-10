SHELL := /bin/bash
export GO111MODULE ?= on
export VERSION ?= 1.0.0-beta7
export GOBIN = $(shell pwd)/bin
BINARY := cloud-sdk-go

include scripts/Makefile.help
.DEFAULT_GOAL := help

include build/Makefile.swagger
include build/Makefile.deps
include build/Makefile.dev

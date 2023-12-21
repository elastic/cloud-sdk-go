SHELL := /bin/bash
export GO111MODULE ?= on
export VERSION ?= v1.16.1
export ECE_VERSION ?= $(shell cat ECE_VERSION)
export ECE_BRANCH ?= ms-101
ECE_DEF_FILE ?= api/version/$(ECE_VERSION).md
export GOBIN = $(shell pwd)/bin
BINARY := cloud-sdk-go

include scripts/Makefile.help
.DEFAULT_GOAL := help

include build/Makefile.swagger
include build/Makefile.deps
include build/Makefile.dev
include build/Makefile.build
include build/Makefile.apivalidation
include build/Makefile.version

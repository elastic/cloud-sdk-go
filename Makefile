SHELL := /bin/bash
export GO111MODULE ?= on
export VERSION ?= 1.0.0-bc10
export ECE_VERSION ?= 2.4.3
ECE_DEF_FILE ?= api/version/$(ECE_VERSION).md
export GOBIN = $(shell pwd)/bin
BINARY := cloud-sdk-go

include scripts/Makefile.help
.DEFAULT_GOAL := help

include build/Makefile.swagger
include build/Makefile.deps
include build/Makefile.dev

#!/usr/bin/env bash


go get
cd bin/
go clean
go build ../commands/*

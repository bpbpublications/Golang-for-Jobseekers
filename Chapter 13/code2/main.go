package main

import "embed"

// content holds our static web server content.
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS

func main() {}

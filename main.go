package main

import (
	"github/ppaulogustavo/aws-st-post/internal"
	"github/ppaulogustavo/aws-st-post/internal/infra"
)

func main() {
	internal.StartApp(infra.ServerStart)
}

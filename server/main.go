package main

import (
	_ "think-sso/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"think-sso/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}

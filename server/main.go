package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "think-sso/internal/logic"
	_ "think-sso/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"think-sso/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}

package main

import (
	_ "gooss/boot"
	_ "gooss/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}

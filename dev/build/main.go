package main

import (
	g "github.com/ysmood/gokit"
)

func main() {
	g.Mkdir("dist", nil)
	g.Remove("dist/**")

	g.E(g.Exec([]string{"go", "build", "-v", "../cmd/..."}, &g.ExecOptions{
		Dir: "dist",
	}))
}

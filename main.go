package main

import (
	"fmt"
	"github.com/khoazero123/gotunnelme/src/gotunnelme"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	file := filepath.Base(os.Args[0])

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, file, "<local port>")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	t := gotunnelme.NewTunnel()
	url, err := t.GetUrl("")
	if err != nil {
		panic(err)
	}
	println(url)
	err = t.CreateTunnel(i)
	if err != nil {
		panic(err)
	}
	t.StopTunnel()
}

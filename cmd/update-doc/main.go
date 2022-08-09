package main

import (
	"fmt"
	"github.com/spf13/cobra/doc"
	"github.com/wanewang/bbgo/pkg/cmd"
	"log"
	"path"
	"runtime"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	root := path.Join(path.Dir(path.Dir(path.Dir(b))), "doc", "commands")
	fmt.Println(root)
	if err := doc.GenMarkdownTree(cmd.RootCmd, root); err != nil {
		log.Fatal(err)
	}
}

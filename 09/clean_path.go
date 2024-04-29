package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// ~ も環境変数も展開したうえでパスをクリーン化
func Clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	// パスをそのままクリーンにする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go

	// パスを絶対パスに整形
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)
	// /usr/local/go/src/path/filepath/path_unix.go

	// パスを相対パスに整形
	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)
	// path/filepath/path.go
}

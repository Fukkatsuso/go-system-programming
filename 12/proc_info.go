package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// 実行ファイル名
	path, _ := os.Executable()
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("実行ファイルパス: %s\n", path)

	// プロセスID
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	// プロセスグループ、セッショングループ
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", syscall.Getpgrp(), sid)

	// ユーザーID、グループID
	fmt.Printf("ユーザーID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)

	// 実効ユーザーID、実効グループID
	fmt.Printf("実効ユーザーID: %d\n", os.Geteuid())
	fmt.Printf("実効グループID: %d\n", os.Getegid())

	// 作業フォルダ
	wd, _ := os.Getwd()
	fmt.Println(wd)
}

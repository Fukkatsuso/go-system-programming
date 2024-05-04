package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

type lockType int16

const (
	readLock  lockType = syscall.LOCK_SH // 共有ロック
	writeLock lockType = syscall.LOCK_EX // 排他ロック
)

func lock(f os.File, lt lockType) (err error) {
	for {
		err = syscall.Flock(int(f.Fd()), int(lt))
		if err != syscall.EINTR {
			break
		}
	}
	return err
}

func unlock(f os.File) error {
	return lock(f, syscall.LOCK_UN) // ロック解除
}

func main() {
	l, err := os.Open("flock.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("try locking...")
	lock(*l, syscall.LOCK_EX)
	fmt.Println("locked!")
	time.Sleep(10 * time.Second)
	unlock(*l)
	fmt.Println("unlock")
}

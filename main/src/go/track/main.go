package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	terminal "golang.org/x/term"
)

func main() {
}

// 输入一个字符后立刻执行
func inputAndExec() {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("设置终端模式失败：", err)
		return
	}
	defer terminal.Restore(int(os.Stdin.Fd()), oldState)

	var char byte
	fmt.Print("请输入一个字母：")
	ch := []byte{char}
	_, err = os.Stdin.Read(ch)
	if err != nil {
		fmt.Println("读取终端输入失败：", err)
		return
	}
	fmt.Printf("\r你输入的字母是：%c %s", ch, string(ch))
}

func progress() {
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/49*100, h)
	}
}

func print() {
	for {
		fmt.Println("\rtest" + strings.Repeat(" ", 50))
		time.Sleep(time.Second)
	}
}

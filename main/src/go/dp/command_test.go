package main

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Command(t *testing.T) {
	Convey("command and exec by sequence", t, func() {
		// 用于测试, 模拟来自客户端的事件
		eventChan := make(chan string)
		events := []string{"start", "archive", "archive", "start"}
		go func([]string) {
			for _, e := range events {
				eventChan <- e
			}
		}(events)
		defer close(eventChan)

		// 使用命令队列缓存命令; 注意这里的命令是一个对象
		commands := make(chan ICommand, 1000)
		defer close(commands)

		go func() {
			for {
				// 从请求或者其他地方获取相关事件参数
				event, ok := <-eventChan
				if !ok {
					return
				}
				var command ICommand
				switch event {
				case "start":
					command = NewStartCommand()
				case "archive":
					command = NewArchiveCommand()
				}

				// 将命令入队
				commands <- command
			}
		}()

		var rs []string
		for {
			select {
			case c := <-commands:
				rs = append(rs, c.Execute())
			case <-time.After(1 * time.Second):
				fmt.Println("timeout 1s")
				So(rs, ShouldResemble, events)
				return
			}
		}
	})
}

func Test_Command_Func(t *testing.T) {
	Convey("command2 and exec by sequence", t, func() {
		// 用于测试, 模拟来自客户端的事件
		eventChan := make(chan string)
		events := []string{"start", "archive", "archive", "start"}
		go func([]string) {
			for _, e := range events {
				eventChan <- e
			}
		}(events)
		defer close(eventChan)

		// 使用命令队列缓存命令; 注意这里的命令是一个对象
		commands := make(chan Command, 1000)
		defer close(commands)

		go func() {
			for {
				// 从请求或者其他地方获取相关事件参数
				event, ok := <-eventChan
				if !ok {
					return
				}

				var command Command
				switch event {
				case "start":
					command = StartCommandFunc()
				case "archive":
					command = ArchiveCommandFunc()
				}

				// 将命令入队
				commands <- command
			}
		}()

		var rs []string
		for {
			select {
			case c := <-commands:
				rs = append(rs, c())

			case <-time.After(1 * time.Second):
				fmt.Println("timeout 1s")
				So(rs, ShouldResemble, events)
				return
			}
		}
	})
}

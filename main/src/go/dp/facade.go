package main

import "fmt"

// 原始接口
type Istarter interface {
	start() error
}

// 门面模式接口
type Ifacade interface {
	on() *PC
}

type CPU struct {
}

func (CPU) start() error {
	fmt.Println("CPU start")
	return nil
}

type Memory struct {
}

func (Memory) start() error {
	fmt.Println("memory start")
	return nil
}

type Disk struct {
}

func (Disk) start() error {
	fmt.Println("disk start")
	return nil
}

// PC 门面模式 综合调用其他类的方法
type PC struct {
}

// 开机
func (pc *PC) on() (error, *PC) {
	cpu := &CPU{}
	cpu.start()
	memory := &Memory{}
	memory.start()
	disk := &Disk{}
	disk.start()
	return nil, pc
}

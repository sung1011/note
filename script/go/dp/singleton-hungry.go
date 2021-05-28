package main

type Instance struct{}

var sltHungry *Instance

func init() {
	sltHungry = &Instance{}
}

func GetInstanceHungry() *Instance {
	return sltHungry
}

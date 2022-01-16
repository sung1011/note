package main

import (
	"encoding/json"
	"time"
)

type Word struct {
	s     string
	visit int
	UpdAt *time.Time
}

func (w *Word) Clone() *Word {
	var nw Word
	b, _ := json.Marshal(w)
	// fmt.Println(w, string(b))
	json.Unmarshal(b, &nw)
	return &nw
}

type WordMap map[string]*Word

func (wm WordMap) Clone(ws []*Word) WordMap {
	NewWM := WordMap{}

	for k, v := range wm {
		NewWM[k] = v // 浅拷贝, 直接拷贝了地址
	}

	for _, w := range ws {
		// fmt.Println("init", NewWM[w.s])
		NewWM[w.s] = w
		// NewWM[w.s] = w.Clone() // 深拷贝, 替换掉需要更新的字段
	}

	return NewWM
}

package main

import (
	"fmt"
)

// 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

// 策略的Map
var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

// NewStorageStrategy
func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}

	return s, nil
}

// 保存文件
type fileStorage struct{}

// 保存到文件
func (s *fileStorage) Save(name string, data []byte) error {
	// return ioutil.WriteFile(name, data, os.ModeAppend)
	return nil
}

// 加密保存到文件
type encryptFileStorage struct{}

// 加密保存
func (s *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return err
	}

	// return ioutil.WriteFile(name, data, os.ModeAppend)
	return nil
}

func encrypt(data []byte) ([]byte, error) {
	return data, nil
}

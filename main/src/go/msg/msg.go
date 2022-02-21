package main

import (
	"encoding/binary"
	"errors"
	"io"
)

const msg_header = "12345678"

func encode(w io.Writer, s string) error {
	// header 8 + len 4 + s len
	if err := binary.Write(w, binary.BigEndian, []byte(msg_header)); err != nil {
		return err
	}

	len := int32(len([]byte(s)))
	if err := binary.Write(w, binary.BigEndian, len); err != nil {
		return err
	}
	// 打印不出来, 只能用debug工具

	if err := binary.Write(w, binary.BigEndian, []byte(s)); err != nil {
		return err
	}

	return nil
}

func decode(r io.Reader) ([]byte, error) {
	header := make([]byte, len(msg_header))
	// binary.Read(r, binary.BigEndian, header)
	io.ReadFull(r, header)
	if string(header) != msg_header {
		return nil, errors.New("invalid msg_header")
	}
	len := make([]byte, 4)
	if _, err := io.ReadFull(r, len); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(len)
	body := make([]byte, length)
	if _, err := io.ReadFull(r, body); err != nil {
		return nil, err
	}

	return body, nil
}

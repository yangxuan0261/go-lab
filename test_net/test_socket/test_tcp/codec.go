package test_tcp

import (
	"encoding/binary"
	"errors"
	"net"
)

// protobuf 编码解码器

var ErrOverSize = errors.New("conn read buff size over limit")

func ReadBuff(c net.Conn) ([]byte, error) {
	lenBuff := make([]byte, 2)
	_, err := c.Read(lenBuff)
	if err != nil {
		return nil, err
	}

	size := binary.BigEndian.Uint16(lenBuff)
	if size > 51200 { // 丢弃大于 50K 的数据包,同时关闭连接
		return nil, ErrOverSize
	}

	pbBuff := make([]byte, size)
	_, err = c.Read(pbBuff)
	if err != nil {
		return nil, err
	}

	return pbBuff, nil
}

func WriteBuff(c net.Conn, pbBuff []byte) error {
	pbLen := len(pbBuff)
	buff := make([]byte, pbLen+2)
	binary.BigEndian.PutUint16(buff[:2], uint16(pbLen))
	copy(buff[2:], pbBuff)
	_, err := c.Write(buff)
	if err != nil {
		return err
	}
	return nil
}

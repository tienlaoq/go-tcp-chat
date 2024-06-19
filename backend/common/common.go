package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func FromBytes(b []byte) (int32, error) {
	buf := bytes.NewReader(b)
	var result int32
	err := binary.Read(buf, binary.BigEndian, &result)
	return result, err
}

func ToBytes(i int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes(), err
}

func WriteMsg(conn net.Conn, msg string) error {
	//send the size of the message to be sent
	bytes, err := ToBytes(int32(len(msg)))
	if err != nil {
		return err
	}
	_, err = conn.Write(bytes)
	if err != nil {
		return err
	}
	//send the message
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func ReadMsg(conn net.Conn) (string, error) {
	// make a buffer to hold length of data
	lenBuf := make([]byte, 4)
	_, err := conn.Read(lenBuf)
	if err != nil {
		return "", err
	}
	lenData, err := FromBytes(lenBuf)
	if err != nil {
		return "", err
	}

	buf := make([]byte, lenData)
	reqLen := 0

	for reqLen < int(lenData) {
		tempreqLen, err := conn.Read(buf[reqLen:])
		reqLen += tempreqLen
		if err == io.EOF {
			return "", fmt.Errorf("Recieved EOF before recieving all promised data.")
		}
		if err != nil {
			return "", fmt.Errorf("Error reading: %s", err.Error())
		}
	}
	return string(buf[:reqLen]), nil

}

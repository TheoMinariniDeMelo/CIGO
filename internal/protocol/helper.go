package protocol

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func WriteMessage(address string, message Message) error{
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Printf("Error while attempting to connect to the ip: %s", address)
		return fmt.Errorf("Wasn't possible to establish connection")
	}

	defer conn.Close()

	buf, err := json.Marshal(message)

	conn.Write(buf)

	return nil
}

func ReadMessage(conn net.Conn) (Message, error) {
	defer conn.Close()
	
	var message Message
	
	payload, err := readFrame(conn)

	if err != nil {
		return message, fmt.Errorf("error to extract payload")
	}

	json.Unmarshal(payload, message)

	return message, nil
}

func readFrame(r io.Reader) ([]byte, error) {
    var size uint32

    err := binary.Read(r, binary.BigEndian, &size)
    if err != nil {
        return nil, err
    }

    const maxSize = 1024 * 1024

    if size > maxSize {
        return nil, fmt.Errorf("frame muito grande")
    }

    payload := make([]byte, size)

    if _, err := io.ReadFull(r, payload); err != nil {
        return nil, err
    }

    return payload, nil
}


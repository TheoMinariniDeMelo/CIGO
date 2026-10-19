package protocol

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func WriteMessage(address string, message Message) error {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return fmt.Errorf("connect to %s: %w", address, err)
    }
    defer conn.Close()

    payload, err := json.Marshal(message)
    if err != nil {
        return fmt.Errorf("marshal message: %w", err)
    }

    if len(payload) > 1024*1024 {
        return fmt.Errorf("message too large")
    }

    frame := make([]byte, 4+len(payload))

    binary.BigEndian.PutUint32(
        frame[:4],
        uint32(len(payload)),
    )

    copy(frame[4:], payload)

    for len(frame) > 0 {
        n, err := conn.Write(frame)
        if err != nil {
            return fmt.Errorf("write message: %w", err)
        }

        if n == 0 {
            return io.ErrShortWrite
        }

        frame = frame[n:]
    }

    return nil
}

func ReadMessage(conn net.Conn) (Message, error) {
	defer conn.Close()

	var message Message

	payload, err := readFrame(conn)

	if err != nil {
		return message, err
	}

	json.Unmarshal(payload, &message)

	return message, nil
}

func SendPing(address string) error {
	message := Message{
		Type: Ping.String(),
	}

	if err := WriteMessage(address, message); err != nil {
		return err
	}
	
	return nil
}

func SendPong(conn net.Conn){
	message := Message{
		Type: Ping.String(),
	}

	
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


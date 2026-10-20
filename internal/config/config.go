package config

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type CIGOConfig struct {
	BranchList     []string
	BranchPatterns []string
	//Seconds until next commit verification
	WaitingTime    uint32
}

func GetCIGOConfig(fileName string) (CIGOConfig, error) {
	strConfig := getConfigFile(fileName)
	
	var config CIGOConfig
	
	err := json.Unmarshal([]byte(strConfig), &config)
	if err != nil {
		return config, err
	}
	
	return config, nil
}

func getConfigFile(fileName string) string {
    fi, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    r := bufio.NewReader(fi)

	var builder strings.Builder

    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        // write a chunk
		builder.WriteString(string(buf[:]))
	}

	return builder.String()
}

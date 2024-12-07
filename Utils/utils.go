package utils

import (
	"fmt"
	"os"
)

func ThrowErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	ThrowErr(err)
	return data
}

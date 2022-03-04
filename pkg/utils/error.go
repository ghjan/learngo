package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Checkerr(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		if msg != "" {
			fmt.Println(msg)
		}
		if strings.Contains(err.Error(), "duplicate key value") ||
			strings.Contains(err.Error(), "a failed transaction") {
			return
		}
		os.Exit(-1)
	}
}

func Checkerr2(err error, msg string) (shouldExit bool) {
	shouldExit = true
	if err != nil {
		if msg != "" {
			fmt.Println(msg)
		}
		fmt.Println(err)
		if strings.Index(err.Error(), "duplicate key value") >= 0 ||
			strings.Index(err.Error(), "a failed transaction") >= 0 {
			shouldExit = false
			return
		}
		//os.Exit(-1)
	}
	return
}

func ShowError(err error) bool {
	if err != nil {
		log.Println(err)
	}
	return err != nil
}

func NoneError(err error) {
	if ShowError(err) {
		panic(err)
	}
}

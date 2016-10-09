package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dst, src string) error {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("%s dstfile srcfile\n", os.Args[0])
		os.Exit(1)
	}
	err := CopyFile(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println("copy error:", err.Error())
		os.Exit(1)
	}
}

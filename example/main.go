package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "file", "", "file to process")
}

func main() {
	flag.Parse()

	fmt.Println("file: ", file)
	file, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for {
		str := "[INFO] " + time.Now().Format(time.RFC3339Nano) + " Golang flag 获取多个值! \n"
		log.Print(str)
		file.WriteString(str)
		file.Sync()
		time.Sleep(1 * time.Second)
	}
}

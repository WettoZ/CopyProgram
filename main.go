package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var from string
var to string
var offset int
var limit int

func init() {
	flag.StringVar(&from, "from", "", "")
	flag.StringVar(&to, "to", "", "")
	flag.IntVar(&offset, "offset", 0, "")
	flag.IntVar(&limit, "limit", 0, "")
}

func Copy(from string, to string, limit int, offset int) error {
	to += "newfile.txt"
	file, err := os.Open(from)
	if err != nil {
		fmt.Println("[ERROR][Open]", err)
	}
	defer file.Close()

	newfile, err := os.Create(to)
	if err != nil {
		fmt.Println("[ERROR][Create]", err)
	}
	defer newfile.Close()

	n, err := file.Seek(int64(offset), 0)
	if err != nil {
		fmt.Println("[ERROR][Seek]", err)
	}

	_ = n

	f, err := io.CopyN(newfile, file, int64(limit))
	if err != nil {
		if f <= 0 {
			fmt.Println("[ERROR] Выход за io.EOF")
		}

	}
	fmt.Println("Data copy:", f)
	return nil
}

func main() {
	flag.Parse()
	Copy(from, to, limit, offset)
	fmt.Println(from, to, offset, limit)

}

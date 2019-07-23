package main

import (
	"log"
	"fmt"
	"io"
	"os"
	"time"
	flag "github.com/spf13/pflag"
)

var src, dst string
var offset, limit int64
var stopCh = make(chan struct{})



func init() {
	flag.StringVarP(&src, "src", "s", "", "path to source file")
	flag.StringVarP(&dst, "dst", "d", "", "path to destination file")
	flag.Int64Var(&offset, "offset", 0, "offset from the beginning  of the source file")
	flag.Int64Var(&limit, "limit", 0, "sets how much copy data in bytes from the offset")
  }
  
func Godd(src, dst string, offset, limit int64 ) int64  {
	source, err := os.Open(src)
	if  err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	
	fi, err := source.Stat()
	if err != nil {
		log.Fatal(err)
	}

	newPosition, err := source.Seek(offset, 0)
	if err != nil {
		log.Fatal(err)
	}

	if limit == 0 {
		limit = fi.Size() - newPosition
	}

	if offset > fi.Size() {
		log.Fatal("offset can't be greater then size of file")
		return 1
	}

	if offset + limit > fi.Size() {
		log.Fatal("offset + limit can't be greater then size of file")
		return 1
	}
	

	destination, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()

	
	nBytes, err:= io.CopyN(destination, source, limit)
	if err != nil {
		log.Fatal(err)
	}
	close(stopCh)
	return nBytes
	
}
 
func main() {
	flag.Parse()
	go Godd(src, dst, offset, limit)

	LOOP:
	for {
		select {
		case <-stopCh:
			fmt.Println("[100%]")
			break LOOP
		default:
			time.Sleep(1000 * time.Millisecond)
			fmt.Print("#")

		}
	}
}



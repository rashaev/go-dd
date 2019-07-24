package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"os"
)

var src, dst string
var offset, limit int64

func init() {
	flag.StringVarP(&src, "src", "s", "", "path to source file")
	flag.StringVarP(&dst, "dst", "d", "", "path to destination file")
	flag.Int64Var(&offset, "offset", 0, "offset from the beginning  of the source file")
	flag.Int64Var(&limit, "limit", 0, "sets how much copy data in bytes from the offset")
}

func Godd(src, dst string, offset, limit int64) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	fi, err := source.Stat()
	if err != nil {
		return 0, err
	}

	newPosition, err := source.Seek(offset, 0)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		limit = fi.Size() - newPosition
	}

	if offset > fi.Size() {
		err := errors.New("offset can't be greater then size of file")
		return 0, err
	}

	if offset+limit > fi.Size() {
		err := errors.New("offset + limit can't be greater then size of file")
		return 0, err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	bar := pb.Full.Start64(limit)
	srcProxy := bar.NewProxyReader(source)

	nBytes, err := io.CopyN(destination, srcProxy, limit)
	bar.Finish()

	return nBytes, err

}

func main() {
	flag.Parse()
	_, err := Godd(src, dst, offset, limit)
	if err != nil {
		log.Fatal(err)
	}
}

package reader

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"io"
	"path"
	"strings"
)

type Reader struct {
	Path string
}

func NewReader() *Reader {
	s := &Reader{}
	return s
}

func (r *Reader) Read(p string) (e error) {

  filepath := path.Base(p)
	filepath = strings.Replace(filepath,".vm",".asm",-1)
	f, err := os.Create(filepath)
	r.Path = p
	fd, err := os.Open(p)
	if err != nil {
		log.Println(err)
		return
	}
	reader := bufio.NewReader(fd)

	for {
	line, _, err := reader.ReadLine()
	if err == io.EOF{
		f.Sync()
		return nil
	}
	if err != nil {
		log.Println(err)
		return
	}

	p := NewParser()


		l, err := p.Read(string(line))
		if err != nil {
			log.Println(err)
			return
		}
		if l != "delete"{
		f.WriteString(l)
		fmt.Println(l)
		}
	}
	return err

}

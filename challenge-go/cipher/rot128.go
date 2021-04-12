package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Rot128Reader implements io.Reader that transforms
type Rot128Reader struct{ reader io.Reader }

func NewRot128Reader(r io.Reader) (*Rot128Reader, error) {
	return &Rot128Reader{reader: r}, nil
}

func (r *Rot128Reader) Read(p []byte) (int, error) {
	if n, err := r.reader.Read(p); err != nil {
		return n, err
	} else {
		rot128(p[:n])
		return n, nil
	}
}

type Rot128Writer struct {
	writer io.Writer
	buffer []byte // not thread-safe
}

func NewRot128Writer(w io.Writer) (*Rot128Writer, error) {
	return &Rot128Writer{
		writer: w,
		buffer: make([]byte, 4096, 4096),
	}, nil
}

func (w *Rot128Writer) Write(p []byte) (int, error) {
	n := copy(w.buffer, p)
	rot128(w.buffer[:n])
	return w.writer.Write(w.buffer[:n])
}

func rot128(buf []byte) {
	for idx := range buf {
		buf[idx] += 128
	}
}

func DeRot128(buf []byte) {
	for idx := range buf {
		buf[idx] -= 128
	}
}

func main() {
	a := []byte("test")
	fmt.Println(a)


	csvfile, err := os.Open("../data/fng.1000.csv.rot128")
	if err != nil{
		log.Fatalln("could not open")
	}

	r := bufio.NewReader(csvfile)
	rot, err := NewRot128Reader(r)
	if err != nil{
		log.Fatalln("could not open")
	}
	out, err := rot.reader.Read([]byte(r))
	if err != nil{
		log.Fatalln("could not open")
	}

	fmt.Println(out)
}

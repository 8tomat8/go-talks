package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Error int

func (e Error) Error() string {
	return strconv.Itoa(int(e))
}

const (
	ErrStopIt Error = iota + 100
)

func main() {
	pReader, pWriter := io.Pipe()

	stop := make(chan struct{})
	go ping(stop, pReader)

	for i := 0; i < 100; i++ {
		pWriter.Write([]byte("Write: " + strconv.Itoa(i)))
	}
	pWriter.CloseWithError(ErrStopIt)
	<-stop
}

func ping(stop chan struct{}, r io.Reader) {
	data := make([]byte, 2<<20)

	for {
		n, err := r.Read(data)
		if err != nil {
			if err == ErrStopIt {
				fmt.Println("WAT!?")
				break
			}
		}
		fmt.Println(string(data[:n]))
	}
	stop <- struct{}{}
}

// ReadFromFile example reading from file
func ReadFromFile() {
	file, err := ioutil.TempFile(os.TempDir(), "lesson")
	if err != nil {
		log.Fatal(err)
	}

	// Write smthg to file here!!!

	var resultData []byte
	data := make([]byte, 0, 1<<30)
	for {
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				resultData = append(resultData, data[:n]...)
				break
			}
			log.Fatal(err)
		}

		resultData = append(resultData, data[:n]...)
	}
	fmt.Println(string(resultData))

	resultData, err = ioutil.ReadAll(io.LimitReader(file, 1<<31))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resultData))
}

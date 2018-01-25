package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.WithField("func", "main")

	d, _ := NewDrammer(&Storage{make(map[string]int)})

	d.doTheDramma(logger)
}

type Storage struct {
	data map[string]int
}

func NewDrammer(s *Storage) (drammer, error) {
	if s == nil {
		return drammer{}, errors.New("AAAA")
	}
	return drammer{s}, nil
}

type drammer struct {
	*Storage
}

func (d drammer) doTheDramma(logger *logrus.Entry) {
	logger.WithField("func", "doTheDramma")
	val, ok := d.data["key"]
	if !ok {
		return
	}
	fmt.Println(val)
}

func doMoreDramma(s *Storage, l *logrus.Entry, data interface{}) {
	func(data interface{}) {
		l.WithField("func", "doMoreDramma")
		val, ok := s.data["key"]
		if !ok {
			return
		}
		fmt.Println(val, data)
	}(data)
}

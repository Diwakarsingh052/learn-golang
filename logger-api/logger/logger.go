package logger

import (
	"fmt"
	"io"
)

type Logger struct {
	ch chan string
}

func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	go func() {

		for v := range l.ch {
			fmt.Fprintln(w, v)
		}

	}()
	return &l
}

func (l *Logger) Println(v string) {
	select {
	case l.ch <- v:
	default:
		fmt.Println("Drop")

	}
}

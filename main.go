package main

import (
	"flag"
	"io"
	"os"
	"time"
)

// Configuration for slomo
type config struct {
	rate int
	buf  int
}

func main() {
	c := parseFlags()
	run(c, os.Stdin, os.Stdout)
}

// Parse the flags and set them in the global vars
func parseFlags() *config {
	c := &config{}
	flag.IntVar(&c.rate, "rate", 20, "rate in ms")
	flag.IntVar(&c.buf, "buf", 2, "size of buffer in bytes")
	flag.Parse()
	return c
}

// Reads from r and writes to w every few micros
func run(c *config, r io.Reader, w io.Writer) {
	ch := make(chan []byte, c.buf)
	go read(ch, c, r)
	write(ch, c, w)
}

// Reads from reader and sends it to the channel
func read(ch chan []byte, c *config, r io.Reader) {
	for {
		b := make([]byte, c.buf)
		_, err := r.Read(b)
		if err == io.EOF {
			close(ch)
			return
		}

		ch <- b
	}
}

// Writes output from the channel every few ms
func write(ch chan []byte, c *config, w io.Writer) {
	for {
		b, more := <-ch
		w.Write(b)

		if !more {
			return
		}

		time.Sleep(time.Millisecond * time.Duration(c.rate))
	}
}

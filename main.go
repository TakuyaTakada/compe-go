package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	io := NewIO()
	defer io.Output()

	// do something
}

type IO struct {
	sc  *bufio.Scanner
	out *bufio.Writer
}

func NewIO() *IO {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	bufSize := math.MaxInt64
	sc.Buffer([]byte{}, bufSize)

	out := bufio.NewWriter(os.Stdout)

	return &IO{
		sc:  sc,
		out: out,
	}
}

func (io *IO) Read() string {
	ok := io.sc.Scan()
	if !ok {
		panic(io.sc.Err())
	}
	return io.sc.Text()
}

func (io *IO) Write(val ...interface{}) {
	for _, v := range val {
		if _, err := fmt.Fprintf(io.out, "%v", v); err != nil {
			panic(err)
		}
	}
}

func (io *IO) Output() {
	err := io.out.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *IO) ReadInt() int {
	i, err := strconv.Atoi(io.Read())
	if err != nil {
		panic(err)
	}
	return i
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1  S1
	buf *bytes.Buffer
}

func (s *S1) Read(p []byte) (n int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me your name: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = input[:len(input)-1]
	s.F2 = input
	copy(p, []byte(input))
	return len(input), nil
}

func (s *S1) Write(p []byte) (n int, err error) {
	if s.F1 < 0 {
		return 0, fmt.Errorf("invalid repetition count")
	}
	for i := 0; i < s.F1; i++ {
		fmt.Print(string(p), "")
	}
	fmt.Println()
	return len(p), nil
}

func (s *S2) eof() bool {
	return s.buf.Len() == 0
}

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		return 0, io.EOF
	}
	return s.buf.Read(p)
}

func main() {
	s1var := S1{F1: 4, F2: "Hello"}
	fmt.Println("Initial S1:", s1var)

	buf := make([]byte, 20)
	_, err := s1var.Read(buf)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("Read:", s1var.F2)
	_, _ = s1var.Write([]byte("Hello There!"))

	s2var := S2{F1: s1var, buf: bytes.NewBufferString("Hello world!!")}
	r := bufio.NewReader(&s2var)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}

}

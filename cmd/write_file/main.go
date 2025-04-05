package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const data = "Data to write\n"
	const filePath = "/tmp/f4.txt"

	f1, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Cannot create file:", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, data)
	// f2: WriteString
	f2, err := os.Create("/tmp/f2.txt")
	if err != nil {
		fmt.Println("Cannot create file:", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(data)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
	// f3: bufio.Writer
	f3, err := os.Create("/tmp/f3.txt")
	if err != nil {
		fmt.Println("Cannot create file:", err)
		return
	}
	defer f3.Close()
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(data)
	if err != nil {
		fmt.Println("Buffered write error:", err)
		return
	}
	if err := w.Flush(); err != nil {
		fmt.Println("Flush error:", err)
		return
	}
	// f4: Write multiple lines using io.WriteString
	f4, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Cannot create file:", err)
		return
	}
	for i := 0; i < 5; i++ {
		n, err = io.WriteString(f4, data)
		if err != nil {
			fmt.Println("Write error:", err)
			f4.Close()
			return
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	f4.Close() // Закрываем до повторного открытия

	// f4 (append): Open with O_APPEND
	f4Append, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("OpenFile error:", err)
		return
	}
	defer f4Append.Close()

	n, err = f4Append.Write([]byte("Put some more data at the end.\n"))
	if err != nil {
		fmt.Println("Append write error:", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}

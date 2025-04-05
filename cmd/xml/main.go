package main

import (
	"encoding/xml"
	"fmt"
)

type Employee struct {
	XMLName   xml.Name `xml:"employee"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
	Address   Address  `xml:"address"`
	Comment   string   `xml:",comment"`
}

type Address struct {
	City    string `xml:"city"`
	Country string `xml:"country"`
}

func main() {
	// Создание объекта Employee
	r := Employee{
		ID:        7,
		FirstName: "Vasilii",
		LastName:  "Zazdravnykh",
		Comment:   "Technical Writer + DevOps",
		Address:   Address{City: "SomeWhere 12", Country: "12312, Greece"},
	}

	// Сериализация в XML с отступами
	output, err := xml.MarshalIndent(&r, "  ", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Добавляем стандартный header XML
	output = append([]byte(xml.Header), output...)

	// Выводим результат
	fmt.Printf("%s\n", output)
}

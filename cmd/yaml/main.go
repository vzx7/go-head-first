package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Person struct {
	Name    string `yaml:"name"`
	Age     int    `yaml:"age"`
	Country string `yaml:"country"`
}

func main() {
	// Создаем пример данных
	p := Person{Name: "John", Age: 30, Country: "USA"}

	// Сериализация в YAML
	data, err := yaml.Marshal(&p)
	if err != nil {
		fmt.Println("Error marshaling to YAML:", err)
		return
	}

	// Записываем в файл
	err = os.WriteFile("person.yaml", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Чтение из YAML-файла
	fileData, err := os.ReadFile("person.yaml")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Десериализация YAML в структуру
	var p2 Person
	err = yaml.Unmarshal(fileData, &p2)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return
	}

	// Выводим результат
	fmt.Println("Decoded from YAML:", p2)
}

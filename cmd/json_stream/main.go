package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

// random генерирует случайное число в диапазоне от min до max.
func random(min, max int) int {
	if max <= min {
		return min
	}
	return rand.Intn(max-min) + min
}

// getString генерирует случайную строку заданной длины.
func getString(l int) string {
	startChar := "A"
	var temp string
	for i := 0; i < l; i++ {
		myRand := random(0, 26)
		newChar := string(startChar[0] + byte(myRand))
		temp += newChar
	}
	return temp
}

// Serialize сериализует данные в JSON.
func Serialize(enc *json.Encoder, data any) error {
	return enc.Encode(data)
}

// DeSerialize десериализует данные из JSON.
func DeSerialize(e *json.Decoder, slice any) error {
	return e.Decode(slice)
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем тестовые данные
	for i := 0; i < 2; i++ {
		DataRecords = append(DataRecords, Data{
			Key: getString(5),
			Val: random(1, 100),
		})
	}

	// Буфер для хранения сериализованных данных
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)

	// Сериализация
	if err := Serialize(encoder, DataRecords); err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("После сериализации:", buf)

	// Десериализация
	decoder := json.NewDecoder(buf)
	var temp []Data
	if err := DeSerialize(decoder, &temp); err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	// Выводим результат десериализации
	fmt.Println("После десериализации:")
	for index, value := range temp {
		fmt.Printf("%d: %+v\n", index, value)
	}
}

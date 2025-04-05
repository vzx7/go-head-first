package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string  `json:"username"`
	Surname *string `json:"surname,omitempty"`
}

// Кастомная реализация UnmarshalJSON
func (u *User) UnmarshalJSON(data []byte) error {
	fmt.Println("📥 Raw JSON input:", string(data))

	// Шаг 1: временная структура (чтобы не вызывать рекурсию)
	type Alias User
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	// Шаг 2: обычный unmarshal во временную структуру
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Шаг 3: проверка, есть ли ключ "surname" в JSON
	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	if _, ok := rawMap["surname"]; ok {
		fmt.Println("✅ Поле 'surname' ПРИСУТСТВУЕТ в JSON")
	} else {
		fmt.Println("❌ Поля 'surname' НЕТ в JSON")
	}

	return nil
}

func main() {
	// Вариант 1: с фамилией
	jsonWithSurname := `{
		"username": "Max",
		"surname": "Payne"
	}`

	var u1 User
	json.Unmarshal([]byte(jsonWithSurname), &u1)
	fmt.Printf("Parsed user: %+v\n\n", u1)

	// Вариант 2: без фамилии
	jsonWithoutSurname := `{
		"username": "Lara"
	}`

	var u2 User
	json.Unmarshal([]byte(jsonWithoutSurname), &u2)
	fmt.Printf("Parsed user: %+v\n", u2)
}

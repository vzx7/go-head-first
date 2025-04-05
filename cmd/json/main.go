package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Age  int    `json:"age"`
	City string `json:"city"`
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Name    string  `json:"username"`
	Surname string  `json:"surname"`
	Profile Profile `json:"profile"`
	Posts   []Post  `json:"posts"`
}

func main() {
	// Создаём пользователя с вложенными структурами
	user := User{
		Name:    "Anna",
		Surname: "Petrova",
		Profile: Profile{
			Age:  29,
			City: "Moscow",
		},
		Posts: []Post{
			{Title: "Go tips", Content: "Use pointers wisely"},
			{Title: "JSON in Go", Content: "encoding/json is powerful"},
		},
	}

	// Кодируем в JSON с отступами (читаемый формат)
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("=== Encoded JSON ===")
	fmt.Println(string(jsonData))

	// Пример входящего JSON (может быть получен от API)
	input := `
{
  "username": "Ivan",
  "surname": "Ivanov",
  "profile": {
    "age": 35,
    "city": "Kazan"
  },
  "posts": [
    {
      "title": "Welcome",
      "content": "Hello from Go"
    },
    {
      "title": "Thoughts",
      "content": "JSON decoding is fun"
    }
  ]
}`

	var decoded User
	err = json.Unmarshal([]byte(input), &decoded)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("=== Decoded Struct ===")
	fmt.Printf("%+v\n", decoded)
}

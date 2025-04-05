package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	// Нормализуем имена флагов
	switch n {
	case "pass", "ps":
		n = "password"
	}
	return pflag.NormalizedName(n)
}

func main() {
	// Определение флагов
	pflag.StringP("name", "n", "Mike", "Name parameter")
	pflag.StringP("password", "p", "hardToGuess", "Password")

	// Настройка функции нормализации флагов
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)

	// Парсинг флагов из командной строки
	pflag.Parse()

	// Привязка флагов к viper
	viper.BindPFlags(pflag.CommandLine)

	// Получаем значения из viper
	name := viper.GetString("name")
	password := viper.GetString("password")

	fmt.Println("Name:", name)
	fmt.Println("Password:", password)

	// Чтение переменной окружения
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS:", val)
	}

	// Устанавливаем переменную окружения
	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS set to:", val)
}

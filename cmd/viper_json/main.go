package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigStructure struct {
	MacPass     string `mapstructure:"macos"`
	LinuxPass   string `mapstructure:"linux"`
	WindowsPass string `mapstructure:"windows"`
	PostHost    string `mapstructure:"postgres"`
	MySQLHost   string `mapstructure:"mysql"`
	MongoHost   string `mapstructure:"mongodb"`
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

const DefaultConfig = ".config.json"

func main() {
	// Получаем текущий рабочий каталог
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	// Используем файл конфигурации из аргументов или используем дефолтный
	configFile := workingDir + "/cmd/viper_json/" + DefaultConfig
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(configFile)
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	// Чтение конфигурации
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	// Проверка наличия параметров и вывод их значений
	if viper.IsSet("macos") {
		fmt.Println("macos:", viper.GetString("macos"))
	} else {
		fmt.Println("macos not set!")
	}

	if viper.IsSet("active") {
		value := viper.GetBool("active")
		if value {
			postgres := viper.GetString("postgres")
			mysql := viper.GetString("mysql")
			mongo := viper.GetString("mongodb")
			fmt.Println("P:", postgres, "My:", mysql, "Mo:", mongo)
		}
	} else {
		fmt.Println("active is not set!")
	}

	if !viper.IsSet("DoesNotExist") {
		fmt.Println("DoesNotExist is not set!")
	}

	// Десериализация конфигурации в структуру
	var config ConfigStructure
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshaling config:", err)
		return
	}

	// Печать данных конфигурации
	PrettyPrint(config)
}

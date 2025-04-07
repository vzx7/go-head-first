package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var visited = map[string]int{}

func walkFunction(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("error accessing %s: %w", path, err)
	}

	// Если файл - это директория
	if info.IsDir() {
		abs, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("error getting absolute path of %s: %w", path, err)
		}

		if _, ok := visited[abs]; ok {
			fmt.Println("Found cycle:", abs)
			return nil
		}
		visited[abs]++
		return nil
	}

	// Если файл - это символическая ссылка
	if info.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(path)
		if err != nil {
			return fmt.Errorf("error reading symlink %s: %w", path, err)
		}

		// Разрешение символической ссылки
		newPath, err := filepath.EvalSymlinks(target)
		if err != nil {
			return fmt.Errorf("error resolving symlink %s: %w", target, err)
		}

		// Получаем информацию о файле по новому пути
		linkInfo, err := os.Stat(newPath)
		if err != nil {
			return fmt.Errorf("error accessing %s: %w", newPath, err)
		}

		// Если это директория, следуем за ссылкой
		if linkInfo.IsDir() {
			fmt.Println("Following...", path, "-->", newPath)
			abs, err := filepath.Abs(newPath)
			if err != nil {
				return fmt.Errorf("error getting absolute path of %s: %w", newPath, err)
			}

			// Проверка на цикл
			if _, ok := visited[abs]; ok {
				fmt.Println("Found cycle!", abs)
				return nil
			}

			visited[abs]++
			return filepath.Walk(newPath, walkFunction)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments!")
		return
	}

	path := os.Args[1]
	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println("Error walking the file tree:", err)
		return
	}

	for k, v := range visited {
		if v > 1 {
			fmt.Printf("Cycle detected at %s, visited %d times\n", k, v)
		}
	}
}

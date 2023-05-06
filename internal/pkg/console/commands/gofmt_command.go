package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const GofmtCommandSignature = "gofmt"

// gofmtCommand - сильно упрощенный gofmt.
// На вход принимает *.txt файл, на выходе перед каждым абзацем вставляет таб и ставит точку в конце предложений.
// В текущей реализации конец предложения тогда, когда следующее после него слово начинается с заглавной буквы.
// Абзацем является \r\n, а также каждая новая строка
type gofmtCommand struct {
	*baseCommand
}

func NewGofmtCommand(baseCommand *baseCommand) *gofmtCommand {
	return &gofmtCommand{baseCommand: baseCommand}
}

func (c *gofmtCommand) Execute(args ...string) error {
	if len(args) == 0 {
		return errors.New("no file path provided")
	}

	path := args[0]
	if len(path) < 4 || path[len(path)-4:] != ".txt" {
		return errors.New("gofmt handles only .txt files")
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = myGofmt(file); err != nil {
		return err
	}

	return nil
}

func myGofmt(file io.Reader) error {
	scanner := bufio.NewScanner(file)

	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString("\n\t")
		line := scanner.Text()

		// Заменяем все \r\n на табы
		line = strings.Replace(line, "\r\n", "\t", -1)

		// Разбиваем строку на слова
		words := strings.Fields(line)

		for i, word := range words {
			// Добавлыем точку в конец слова, если следующее за ним слово начинается с заглавной буквы
			if i < len(words)-1 && unicode.IsUpper([]rune(words[i+1])[0]) {
				word += "."
			}

			sb.WriteString(" " + word)
		}

		// Нужно чтобы избеэать постановки точки в пустой строке
		if len(words) != 0 {
			sb.WriteRune('.')
		}
	}

	fmt.Println(sb.String())
	return nil
}

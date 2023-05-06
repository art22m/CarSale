package commands

import (
	"errors"
	"fmt"
	"strings"
)

// Spell

const SpellCommandSignature = "spell"

type spellCommand struct {
	*baseCommand
}

func NewSpellCommand(baseCommand *baseCommand) *spellCommand {
	return &spellCommand{baseCommand: baseCommand}
}

func (c *spellCommand) Execute(args ...string) error {
	if len(args) == 0 {
		return errors.New("no arguments")
	}

	res := spell(args[0])
	fmt.Println(res)

	return nil
}

func spell(word string) string {
	if word == "" {
		return word
	}

	var sb strings.Builder
	for _, c := range []byte(word) {
		sb.WriteByte(c)
		sb.WriteByte(' ')
	}

	spelledWord := sb.String()
	return spelledWord[:len(spelledWord)-1]
}

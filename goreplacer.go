package goreplace

import (
	"errors"
	"log"
	"os"
	"strings"
)

var (
	debugEnabled = true
)

func debug(format string, a ...interface{}) {
	if debugEnabled {
		log.Printf(format+"\n", a...)
	}
}

func ReplaceFunc(text, startText, endText string, f func(string) string) (string, error) {
	// check empty values
	if len(text) == 0 || len(startText) == 0 || len(endText) == 0 {
		return text, errors.New("Text, start text and end text cannot be empty")
	}

	// get start position
	startTextPosition := strings.Index(text, startText)

	if startTextPosition == -1 {
		return text, errors.New("Start text was not found on text")
	}

	// get end position
	endTextPosition := strings.Index(text, endText)

	if endTextPosition == -1 {
		return text, errors.New("End text was not found on text")
	}

	// check inverted position
	if startTextPosition >= endTextPosition {
		return text, errors.New("Start text position is bigger than end text position")
	}

	debug("Text start: %v, Text end: %v", startTextPosition, endTextPosition)

	// key start/endtext position
	keyStartTextPosition := startTextPosition + len(startText)
	keyEndTextPosition := endTextPosition

	// key text
	debug("Key start: %v, Key end: %v", keyStartTextPosition, keyEndTextPosition)
	debug("Text to be used: %v, Diff: %v", text, (keyEndTextPosition - keyStartTextPosition))
	keyText := text[keyStartTextPosition:keyEndTextPosition]
	debug("Key text: %v", keyText)

	if len(keyText) > 0 {
		textToReplace := f(keyText)
		text = text[0:startTextPosition] + textToReplace + text[endTextPosition+len(endText):len(text)]
	}

	debug("Final text: %v", text)

	// continue searching again
	startTextPosition = strings.Index(text, startText)

	if startTextPosition > -1 {
		return ReplaceFunc(text, startText, endText, f)
	}

	return text, nil
}

func ReplaceByMapOfString(text, startText, endText string, data map[string]string) string {
	result, _ := ReplaceFunc(text, startText, endText, func(keyText string) string {
		if len(keyText) > 0 {
			mapValue, found := data[keyText]

			if !found {
				mapValue = ""
			}

			return mapValue
		}

		return ""
	})

	return result
}

func ReplaceByEnvVars(text, startText, endText string) string {
	result, _ := ReplaceFunc(text, startText, endText, os.Getenv)
	return result
}

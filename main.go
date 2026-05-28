package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func applyUp(text string) string {
	re := regexp.MustCompile(`(\w+)\s\(up\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		return strings.ToUpper(word)
	})
}

func applyLow(text string) string {
	re := regexp.MustCompile(`(\w+)\s\(low\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		return strings.ToLower(word)
	})
}

func applyCap(text string) string {
	re := regexp.MustCompile(`(.+)\s\(cap\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		return strings.Title(strings.ToLower(word))
	})
}

func applyHex(text string) string {
	re := regexp.MustCompile(`(\w+)\s\(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		num, _ := strconv.ParseInt(word, 16, 64)
		return fmt.Sprintf("%d", num)
	})
}

func applyBin(text string) string {
	re := regexp.MustCompile(`(\w+)\s\(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		word := re.FindStringSubmatch(match)[1]
		num, _ := strconv.ParseInt(word, 2, 64)
		return fmt.Sprintf("%d", num)
	})
}

func applyUpN(text string) string {
	re := regexp.MustCompile(`((?:\w+[ \t]+){1,100})\(up,\s*(\d+)\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		sub := re.FindStringSubmatch(match)
		n, _ := strconv.Atoi(sub[2])
		words := strings.Fields(sub[1])
		if n > len(words) {
			n = len(words)
		}
		start := len(words) - n
		for i := start; i < len(words); i++ {
			words[i] = strings.ToUpper(words[i])
		}
		return strings.Join(words, " ")
	})
}

func applyLowN(text string) string {
	re := regexp.MustCompile(`((?:\w+[ \t]+){1,100})\(low,\s*(\d+)\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		sub := re.FindStringSubmatch(match)
		n, _ := strconv.Atoi(sub[2])
		words := strings.Fields(sub[1])
		if n > len(words) {
			n = len(words)
		}
		start := len(words) - n
		for i := start; i < len(words); i++ {
			words[i] = strings.ToLower(words[i])
		}
		return strings.Join(words, " ")
	})
}

func applyCapN(text string) string {
	re := regexp.MustCompile(`((?:\w+[ \t]+){1,100})\(cap,\s*(\d+)\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		sub := re.FindStringSubmatch(match)
		n, _ := strconv.Atoi(sub[2])
		words := strings.Fields(sub[1])
		if n > len(words) {
			n = len(words)
		}
		start := len(words) - n
		for i := start; i < len(words); i++ {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
		return strings.Join(words, " ")
	})
}

func applyPunctuation(text string) string {
	re := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = re.ReplaceAllString(text, "$1")
	re2 := regexp.MustCompile(`([.,!?:;]+)(\w)`)
	text = re2.ReplaceAllString(text, "$1 $2")
	return text
}

func applyQuotes(text string) string {
	re := regexp.MustCompile(`'\s+(.+?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}

func applyAN(text string) string {
	re := regexp.MustCompile(`\b[aA]\s+([hHaAeEiIoOuU]\w*)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		if match[0] == 'A' {
			return "An " + re.FindStringSubmatch(match)[1]
		}
		return "an " + re.FindStringSubmatch(match)[1]
	})
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur lecture:", err)
		os.Exit(1)
	}

	text := string(content)
	text = applyHex(text)
	text = applyBin(text)
	text = applyUpN(text)
	text = applyLowN(text)
	text = applyCapN(text)
	text = applyUp(text)
	text = applyLow(text)
	text = applyCap(text)
	text = applyPunctuation(text)
	text = applyQuotes(text)
	text = applyAN(text)

	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println("Erreur écriture:", err)
		os.Exit(1)
	}
}

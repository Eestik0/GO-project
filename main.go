package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func shiftEN(ch rune, shift int) rune {
	if ch >= 'a' && ch <= 'z' {
		return (ch-'a'+rune(shift)+26)%26 + 'a'
	} else if ch >= 'A' && ch <= 'Z' {
		return (ch-'A'+rune(shift)+26)%26 + 'A'
	}
	return ch
}

func encryptEN(text string, shift int) string {
	var result strings.Builder
	for _, ch := range text {
		result.WriteRune(shiftEN(ch, shift))
	}
	return result.String()
}

func bruteForceEN(text string) {
	for shift := 1; shift < 26; shift++ {
		fmt.Printf("Shift %2d: %s\n", shift, encryptEN(text, -shift))
	}
}

var uaSmall = []rune("абвгґдеєжзиіїйклмнопрстуфхцчшщьюя")
var uaBig = []rune("АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ")

const uaLen = 33

func shiftUA(ch rune, shift int) rune {
	for i, r := range uaSmall {
		if ch == r {
			return uaSmall[(i+shift+uaLen)%uaLen]
		}
	}
	for i, r := range uaBig {
		if ch == r {
			return uaBig[(i+shift+uaLen)%uaLen]
		}
	}
	return ch
}

func encryptUA(text string, shift int) string {
	var result strings.Builder
	for _, ch := range text {
		result.WriteRune(shiftUA(ch, shift))
	}
	return result.String()
}

func bruteForceUA(text string) {
	for shift := 1; shift < uaLen; shift++ {
		fmt.Printf("Shift %2d: %s\n", shift, encryptUA(text, -shift))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" <== Шифр Цезаря ==>")
	fmt.Println("1. Зашифрувати англ")
	fmt.Println("2. Розшифрувати англ (brute-force)")
	fmt.Println("3. Зашифрувати укр")
	fmt.Println("4. Розшифрувати укр   (brute-force)")

	fmt.Print("Виберіть опцію: ")
	var choice int
	fmt.Scanln(&choice)

	fmt.Print("Введіть текст: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	switch choice {
	case 1:
		fmt.Print("Зсув (0-25): ")
		var shift int
		fmt.Scanln(&shift)
		fmt.Println("Зашифровано:", encryptEN(text, shift))
	case 2:
		fmt.Println("Brute-force розшифрування (англ):")
		bruteForceEN(text)
	case 3:
		fmt.Print("Зсув (0-32): ")
		var shift int
		fmt.Scanln(&shift)
		fmt.Println("Зашифровано:", encryptUA(text, shift))
	case 4:
		fmt.Println("Brute-force розшифрування (укр):")
		bruteForceUA(text)
	default:
		fmt.Println(" Невірна опція")
	}
}

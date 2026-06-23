package chat

import (
	"fmt"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorDim    = "\033[2m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
)

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func PrintHeader(roomCode, myName, myIP string, userCount int) {
	content := fmt.Sprintf("  chat  •  room: %s  •  %s  •  %s  •  %d online  ",
		roomCode, myName, myIP, userCount)
	border := strings.Repeat("═", len(content))

	fmt.Println(colorCyan + "╔" + border + "╗")
	fmt.Println("║" + content + "║")
	fmt.Println("╚" + border + "╝" + colorReset)
	fmt.Println()
}

func PrintMessage(m Message) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("%s[%s]%s  %-12s: %s\n",
		colorGreen, timestamp, colorReset,
		m.From,
		m.Body,
	)
	PrintPrompt()
}

func PrintHistory(m Message) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("%s↑ [%s]  %-12s: %s%s\n",
		colorDim, timestamp,
		m.From,
		m.Body,
		colorReset,
	)
}

func PrintSystem(text string) {
	fmt.Printf("%s  ── %s ──%s\n", colorDim, text, colorReset)
	PrintPrompt()
}

func PrintUserList(users []string) {
	fmt.Println(colorYellow + "\n  Online users:" + colorReset)
	for _, u := range users {
		fmt.Println("   •", u)
	}
	fmt.Println()
	PrintPrompt()
}

func PrintPrompt() {
	fmt.Print("> ")
}

func PrintError(msg string) {
	fmt.Printf("\033[31mError: %s\033[0m\n", msg)
}
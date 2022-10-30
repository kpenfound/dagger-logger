package logger

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	defaultLight = "147"
	defaultDark  = "17"
)

type Logger struct {
	Light string
	Dark  string
}

func (l Logger) Write(p []byte) (int, error) {
	line := string(p)
	// Check for newlines
	nl := strings.Contains(line, "\n")
	line = strings.ReplaceAll(line, "\n", "")

	// Default colors
	light := l.Light
	if l.Light == "" {
		light = defaultLight
	}
	dark := l.Dark
	if l.Dark == "" {
		dark = defaultDark
	}

	// Get the lipgloss style
	style := getStyle(line, light, dark)

	// Render with or without newline
	if nl {
		fmt.Println(style.Render(line))
	} else {
		fmt.Print(style.Render(line))
	}

	return len(p), nil
}

func getStyle(line, light, dark string) lipgloss.Style {
	bg := light
	fg := dark

	step := strings.Split(line, " ")[0]
	step = strings.ReplaceAll(step, "#", "")
	istep, err := strconv.Atoi(step)
	if err != nil {
		// Unknown line, return blank style
		return lipgloss.NewStyle().
			UnsetForeground().
			UnsetBackground()
	}

	// Alternating colors
	if istep%2 == 0 {
		bg = dark
		fg = light
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(fg)).
		Background(lipgloss.Color(bg))

	return style
}

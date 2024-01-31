package ui

import "github.com/charmbracelet/lipgloss"

func FontColor(text, color string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(text)
}

// GenMask generate a mask string of the specified length
func GenMask(l int) string {
	return GenStr(l, "*")
}

// GenStr generate a string of the specified length, the string is composed of the given characters
func GenStr(l int, s string) string {
	var ss string
	for i := 0; i < l; i++ {
		ss += s
	}
	return ss
}

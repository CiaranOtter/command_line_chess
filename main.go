package main

import (
	"chess_board/board"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(board.NewBoard(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

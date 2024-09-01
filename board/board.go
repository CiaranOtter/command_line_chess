package board

import (
	"chess_board/pieces"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Board struct {
	Board     [8][8]pieces.PieceInterface
	cursor_x  int
	cursor_y  int
	move_list []pieces.Move
}

type (
	MoveMsg struct{}
)

func NewBoard() Board {
	b := Board{
		cursor_x: 0,
		cursor_y: 0,
	}

	for i := 0; i < 8; i++ {

		for j := 0; j < 8; j++ {
			if i == 1 || i == 6 {
				p := pieces.Pawn{}
				p.SetPos(pieces.NewMove(i, j))
				b.Board[i][j] = p

			} else if i == 0 || i == 7 {
				if j == 0 || j == 7 {
					p := pieces.Rook{}
					p.SetPos(pieces.NewMove(i, j))
					b.Board[i][j] = p
				} else if j == 1 || j == 6 {
					p := pieces.Knight{}
					p.SetPos(pieces.NewMove(i, j))
					b.Board[i][j] = p
				} else if j == 2 || j == 5 {
					p := pieces.Bishop{}
					p.SetPos(pieces.NewMove(i, j))
					b.Board[i][j] = p
				} else if j == 3 {
					p := pieces.Queen{}
					p.SetPos(pieces.NewMove(i, j))
					b.Board[i][j] = p
				} else if j == 4 {
					p := pieces.King{}
					p.SetPos(pieces.NewMove(i, j))
					b.Board[i][j] = p
				} else {
					b.Board[i][j] = pieces.EmptySpace{}
				}
			} else {
				b.Board[i][j] = pieces.EmptySpace{}
			}
		}
	}

	return b
}

func (b Board) Init() tea.Cmd {
	return nil
}

func (b Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmd []tea.Cmd = []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return b, tea.Quit
		case tea.KeyUp:
			b.cursor_y--
			if b.cursor_y < 0 {
				b.cursor_y = 7
			}
		case tea.KeyDown:
			b.cursor_y++
			if b.cursor_y > 7 {
				b.cursor_y = 0
			}
		case tea.KeyLeft:
			b.cursor_x--
			if b.cursor_x < 0 {
				b.cursor_x = 7
			}
		case tea.KeyRight:
			b.cursor_x++
			if b.cursor_x > 7 {
				b.cursor_x = 0
			}

			cmd = append(cmd, tea.Tick(time.Nanosecond, func(t time.Time) tea.Msg {
				return MoveMsg{}
			}))
		}
	case MoveMsg:
		b.move_list = b.Board[b.cursor_x][b.cursor_y].GetMoves()
	}

	return b, tea.Batch(cmd...)
}

func Elem(f func(any, any) bool, check any, cont ...any) bool {
	for _, i := range cont {
		if f(check, i) {
			return true
		}
	}

	return false
}

func (b Board) View() string {
	board_string := ""
	board_string = fmt.Sprintf("%s\t+-+-+-+-+-+-+-+-+\n", board_string)
	for i := 0; i < 8; i++ {
		board_string = fmt.Sprintf("%s%d\t|", board_string, 8-i)
		for j := 0; j < 8; j++ {
			s := lipgloss.NewStyle()
			if i == b.cursor_y && j == b.cursor_x {
				s = s.Background(lipgloss.Color("#b5e48c"))
			} else if Elem(pieces.Compare, b.Board[b.cursor_x][b.cursor_y].GetPos(), b.move_list) {
				s = s.Background(lipgloss.Color("#ffd166"))
			} else if (i+j)%2 != 0 {
				s = s.Background(lipgloss.Color("#000000")).Foreground(lipgloss.Color("ffffff"))
			} else {
				s = s.Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("000000"))
			}

			board_string = fmt.Sprintf("%s%s|", board_string, s.Render(b.Board[i][j].Print()))

		}

		board_string = fmt.Sprintf("%s\n\t+-+-+-+-+-+-+-+-+\n", board_string)
	}

	board_string = fmt.Sprintf("%s\n\n\t A B C D E F G H\n", board_string)
	return board_string
}

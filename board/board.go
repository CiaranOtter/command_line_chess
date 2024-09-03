package board

import (
	"chess_board/pieces"
	"chess_board/utils"
	"fmt"
	"reflect"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Board struct {
	utils.BoardInterface
	Board          [8][8]utils.PieceInterface
	cursor_x       int
	cursor_y       int
	move_list      []utils.Move
	move_pos       int
	selected       bool
	selected_piece utils.PieceInterface
}

func (b *Board) HasPiece(x, y int) bool {
	switch b.Board[y][x].(type) {
	case *pieces.EmptySpace:
		return false
	default:
		return true
	}
}

func (b *Board) IsOpponent(x, y, col int) bool {
	return b.HasPiece(x, y) && b.Board[y][x].GetColor() != col
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

			var color = 0
			if i > 4 {
				color = 1
			}
			if i == 1 || i == 6 {
				p := &pieces.Pawn{}
				p.SetPos(utils.NewMove(i, j))
				b.Board[i][j] = p
				fmt.Printf("Adding a pawn\n")
			} else if i == 0 || i == 7 {
				if j == 0 || j == 7 {
					p := &pieces.Rook{}
					p.SetPos(utils.NewMove(i, j))
					b.Board[i][j] = p

				} else if j == 1 || j == 6 {
					p := &pieces.Knight{}
					p.SetPos(utils.NewMove(i, j))
					b.Board[i][j] = p

				} else if j == 2 || j == 5 {
					p := &pieces.Bishop{}
					p.SetPos(utils.NewMove(i, j))
					b.Board[i][j] = p

				} else if j == 3 {
					p := &pieces.Queen{}
					p.SetPos(utils.NewMove(i, j))
					b.Board[i][j] = p

				} else if j == 4 {
					p := &pieces.King{}
					p.SetPos(utils.NewMove(i, j))
					b.Board[i][j] = p

				} else {
					p := &pieces.EmptySpace{}
					p.SetPos(utils.NewMove(i, j))
					b.Board[i][j] = p

				}
			} else {
				p := &pieces.EmptySpace{}
				p.SetPos(utils.NewMove(i, j))
				b.Board[i][j] = p
			}

			b.Board[i][j].SetColor(color)
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
		case tea.KeyEnter:
			if len(b.move_list) > 0 && !b.selected {
				b.selected = true
				b.selected_piece = b.Board[b.cursor_y][b.cursor_x]
			} else {
				old_x, old_y := b.selected_piece.GetPos().GetMove()
				b.selected_piece.SetPos(b.Board[b.cursor_y][b.cursor_x].GetPos())
				b.Board[b.cursor_y][b.cursor_x] = b.selected_piece

				b.Board[old_y][old_x] = &pieces.EmptySpace{}
				b.Board[old_y][old_x].SetPos(utils.NewMove(old_y, old_x))
				cmd = append(cmd, tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
					return tea.KeyMsg{Type: tea.KeyEscape}
				}))
			}
		case tea.KeyEsc:
			b.move_pos = 0
			b.selected = false
			b.selected_piece = nil
		case tea.KeyCtrlC:
			return b, tea.Quit
		case tea.KeyUp:
			if b.selected {
				b.move_pos = b.move_pos + 1
				if b.move_pos >= len(b.move_list) {
					b.move_pos = 0
				}
				b.cursor_x, b.cursor_y = b.move_list[b.move_pos].GetMove()
			} else {
				b.cursor_y--
				if b.cursor_y < 0 {
					b.cursor_y = 7
				}
			}
		case tea.KeyDown:
			if b.selected {
				b.move_pos = b.move_pos + 1
				if b.move_pos >= len(b.move_list) {
					b.move_pos = 0
				}
				b.cursor_x, b.cursor_y = b.move_list[b.move_pos].GetMove()
			} else {
				b.cursor_y++
				if b.cursor_y > 7 {
					b.cursor_y = 0
				}
			}
		case tea.KeyLeft:
			if b.selected {
				b.move_pos = b.move_pos + 1
				if b.move_pos >= len(b.move_list) {
					b.move_pos = 0
				}
				b.cursor_x, b.cursor_y = b.move_list[b.move_pos].GetMove()
			} else {
				b.cursor_x--
				if b.cursor_x < 0 {
					b.cursor_x = 7
				}
			}
		case tea.KeyRight:
			if b.selected {
				b.move_pos = b.move_pos + 1
				if b.move_pos >= len(b.move_list) {
					b.move_pos = 0
				}
				b.cursor_x, b.cursor_y = b.move_list[b.move_pos].GetMove()
			} else {
				b.cursor_x++
				if b.cursor_x > 7 {
					b.cursor_x = 0
				}
			}
		}

		cmd = append(cmd, tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
			return MoveMsg{}
		}))
	case MoveMsg:
		if !b.selected {
			b.move_list = b.Board[b.cursor_y][b.cursor_x].GetMoves(&b)
		}
	}

	return b, tea.Batch(cmd...)
}

func Elem(f func(any, any) bool, check any, cont []utils.Move) bool {
	for _, i := range cont {
		// fmt.Print(i)
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
			} else if Elem(utils.Compare, b.Board[i][j].GetPos(), b.move_list) {
				// fmt.Printf("There is a piece that is an element")
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

	t := reflect.TypeOf(b.Board[b.cursor_y][b.cursor_x])

	// If it's a pointer, get the underlying type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	board_string = fmt.Sprintf("%sCurrent piece: %s\n", board_string, t.Name())
	board_string = fmt.Sprintf("%sThis pices has %d moves available.\n", board_string, len(b.move_list))
	board_string = fmt.Sprintf("%sThe possible moves are:", board_string)
	for _, m := range b.move_list {
		x, y := m.GetMove()
		board_string = fmt.Sprintf("%s (%d, %d)", board_string, x, y)
	}
	board_string = fmt.Sprintf("%s\nSelecting a move: ", board_string)
	if b.selected {
		board_string = fmt.Sprintf("%sTrue\n", board_string)
	} else {
		board_string = fmt.Sprintf("%sFalse\n", board_string)
	}
	return board_string
}

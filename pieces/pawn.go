package pieces

import "chess_board/utils"

type Pawn struct {
	EmptySpace
}

func (p Pawn) Print() string {
	return "p"
}

func (p Pawn) GetMoves(b utils.BoardInterface) []utils.Move {
	var moves []utils.Move = []utils.Move{}

	x, y := p.Pos.GetMove()

	forward := func(pos int, dist int) int {
		if p.color == 0 {
			return pos + dist
		} else {
			return pos - dist
		}
	}
	if y == 1 || y == 6 {
		moves = append(moves, utils.NewMove(forward(y, 2), x))
	}

	moves = append(moves, utils.NewMove(forward(y, 1), x))

	check_pos := moves[len(moves)-1]

	c_x, c_y := check_pos.GetMove()

	if c_x < 7 && forward(c_y, 1) <= 7 && forward(c_y, 1) >= 0 && b.HasPiece(c_x+1, forward(c_y, 1)) {
		moves = append(moves, utils.NewMove(forward(c_y, 1), c_x+1))
	}

	if c_x > 0 && forward(c_y, 1) <= 7 && forward(c_y, 1) >= 0 && b.HasPiece(c_x-1, forward(c_y, 1)) {
		moves = append(moves, utils.NewMove(forward(c_y, 1), c_x-1))
	}

	return moves
}

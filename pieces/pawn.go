package pieces

import "chess_board/utils"

type Pawn struct {
	EmptySpace
}

func (p Pawn) Print() string {
	return "p"
}

func (p Pawn) GetMoves(b utils.BoardInterface) []utils.Move {
	var temp_moves []utils.Move = []utils.Move{}
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
		temp_moves = append(temp_moves, utils.NewMove(forward(y, 2), x))
	}

	temp_moves = append(temp_moves, utils.NewMove(forward(y, 1), x))

	c_x, c_y := p.GetPos().GetMove()

	for _, i := range temp_moves {
		t_x, t_y := i.GetMove()

		if t_x >= 0 && t_x <= 7 && t_y >= 0 && t_y <= 7 && !b.HasPiece(t_x, t_y) {
			moves = append(moves, i)
		}
	}

	if c_x < 7 && forward(c_y, 1) <= 7 && forward(c_y, 1) >= 0 && b.HasPiece(c_x+1, forward(c_y, 1)) {
		moves = append(moves, utils.NewMove(forward(c_y, 1), c_x+1))
	}

	if c_x > 0 && forward(c_y, 1) <= 7 && forward(c_y, 1) >= 0 && b.HasPiece(c_x-1, forward(c_y, 1)) {
		moves = append(moves, utils.NewMove(forward(c_y, 1), c_x-1))
	}

	return moves
}

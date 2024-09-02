package pieces

import "chess_board/utils"

type King struct {
	EmptySpace
}

func (k King) Print() string {
	return "K"
}

func (r King) GetMoves(b utils.BoardInterface) []utils.Move {
	moves := []utils.Move{}

	var dir_funcs []FuncType
	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x + 1, y
	})
	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x - 1, y
	})

	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x, y + 1
	})
	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x, y - 1
	})
	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x + 1, y + 1
	})
	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x - 1, y + 1
	})

	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x + 1, y - 1
	})
	dir_funcs = append(dir_funcs, func(x, y int) (int, int) {
		return x - 1, y - 1
	})

	for _, f := range dir_funcs {
		var x, y int
		x, y = f(r.GetPos().GetMove())
		if x >= 0 && y >= 0 && x <= 7 && y <= 7 && !b.HasPiece(x, y) {
			moves = append(moves, utils.NewMove(y, x))
		}
	}

	return moves
}

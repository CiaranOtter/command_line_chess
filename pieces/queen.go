package pieces

import "chess_board/utils"

type Queen struct {
	EmptySpace
}

func (q Queen) Print() string {
	return "q"
}

func (r Queen) GetMoves(b utils.BoardInterface) []utils.Move {
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
		for x, y = r.GetPos().GetMove(); x >= 0 && y >= 0 && x <= 7 && y <= 7 && !b.HasPiece(x, y); x, y = f(x, y) {
			moves = append(moves, utils.NewMove(y, x))
		}
	}

	return moves
}

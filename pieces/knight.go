package pieces

import "chess_board/utils"

type Knight struct {
	EmptySpace
}

func (k Knight) Print() string {
	return "k"
}

func (k Knight) GetMoves(b utils.BoardInterface) []utils.Move {

	x, y := k.GetPos().GetMove()
	var temp_moves = [8]utils.Move{
		utils.NewMove(y+2, x+1),
		utils.NewMove(y-2, x+1),
		utils.NewMove(y+2, x-1),
		utils.NewMove(y-2, x-1),
		utils.NewMove(y+1, x+2),
		utils.NewMove(y-1, x+2),
		utils.NewMove(y+1, x-2),
		utils.NewMove(y-1, x-2),
	}

	moves := []utils.Move{}

	for _, i := range temp_moves {
		t_x, t_y := i.GetMove()

		if t_x >= 0 && t_x <= 7 && t_y >= 0 && t_y <= 7 && !b.HasPiece(t_x, t_y) {
			moves = append(moves, i)
		}
	}

	return moves
}

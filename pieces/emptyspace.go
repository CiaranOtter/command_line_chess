package pieces

import (
	"chess_board/utils"
)

type EmptySpace struct {
	utils.PieceInterface
	Pos utils.Move
	// 0 for white, 1 for black
	color int
}

func (e *EmptySpace) SetPos(m utils.Move) {
	e.Pos = m
}

func (e *EmptySpace) SetColor(i int) {
	e.color = i
}

func (e EmptySpace) GetColor() int {
	return e.color
}

func (e EmptySpace) GetPos() utils.Move {
	return e.Pos
}

func (e EmptySpace) GetMoves(utils.BoardInterface) []utils.Move {
	return []utils.Move{}
}

func (e EmptySpace) Print() string {
	return " "
}

func (e EmptySpace) PrintName() string {
	return "Empty space"
}

func (e EmptySpace) GetMove(b utils.BoardInterface) utils.Move {
	return e.Pos
}

func NewSpace(x, y int) EmptySpace {
	return EmptySpace{
		Pos: utils.NewMove(x, y),
	}
}

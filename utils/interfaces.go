package utils

import (
	"log"
)

type PieceInterface interface {
	GetMoves(BoardInterface) []Move
	SetPos(m Move)
	SetColor(int)
	GetColor() int
	Print() string
	GetPos() Move
	PrintName() string
}

type BoardInterface interface {
	HasPiece(int, int) bool
}

type Move struct {
	x int
	y int
}

func Compare(a, b any) bool {

	a_move, A_ok := a.(Move)
	b_move, B_ok := b.(Move)

	if A_ok && B_ok {
		return (a_move.x == b_move.x) && (a_move.y == b_move.y)
	}

	log.Fatal("This is not a move")
	return false
}

func NewMove(y, x int) Move {
	return Move{
		x: x,
		y: y,
	}
}

func (m Move) GetMove() (int, int) {
	return m.x, m.y
}

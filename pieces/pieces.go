package pieces

type Move struct {
	x int
	y int
}

func Compare(a, b Move) bool {
	return (a.x == b.x) && (a.y == b.y)
}

type PieceInterface interface {
	GetMoves() []Move
	SetPos(m Move)
	Print() string
	GetPos() Move
}

type (
	EmptySpace struct {
		PieceInterface
		Pos Move
		// 0 for white, 1 for black
		color int
	}
	Pawn struct {
		EmptySpace
	}
	Rook struct {
		EmptySpace
	}
	Knight struct {
		EmptySpace
	}
	Bishop struct {
		EmptySpace
	}
	King struct {
		EmptySpace
	}
	Queen struct {
		EmptySpace
	}
)

func NewMove(x, y int) Move {
	return Move{
		x: x,
		y: y,
	}
}

func (m Move) GetMove() (int, int) {
	return m.x, m.y
}

func (e EmptySpace) SetPos(m Move) {
	e.Pos = m
}

func (e EmptySpace) GetMoves() []Move {
	return []Move{}
}

func (e EmptySpace) Print() string {
	return " "
}

func (e EmptySpace) GetMove() Move {
	return e.Pos
}

func NewSpace(x, y int) EmptySpace {
	return EmptySpace{
		Pos: Move{
			x: x,
			y: y,
		},
	}
}

func (p Pawn) Print() string {
	return "p"
}

func (p Pawn) GetMoves() []Move {
	var moves []Move = []Move{}

	forward := func(dist int) int {
		if p.color == 0 {
			return p.Pos.y + dist
		} else {
			return p.Pos.y - dist
		}
	}
	if p.Pos.y == 1 || p.Pos.y == 6 {
		moves = append(moves, Move{
			x: p.Pos.x,
			y: forward(2),
		})
	}

	moves = append(moves, Move{
		x: p.Pos.x,
		y: forward(1),
	})

	return moves
}

func (r Rook) Print() string {
	return "r"
}

func (b Bishop) Print() string {
	return "b"
}

func (k King) Print() string {
	return "K"
}

func (q Queen) Print() string {
	return "q"
}

func (k Knight) Print() string {
	return "k"
}

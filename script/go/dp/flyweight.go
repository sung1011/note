package main

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// ... 其他棋子
}

// 棋子静态只读; 享元
type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

// 工厂
func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

// 棋子静态+动态
type ChessPiece struct {
	Unit *ChessPieceUnit // 静态只读对象
	X    int             // 动态
	Y    int             // 动态
}

// 棋局
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// 初始化棋盘
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id), // 创建和复用只读对象
			X:    0,
			Y:    0,
		}
	}
	return board
}

// 移动棋子
func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}

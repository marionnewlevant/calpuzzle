package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type OrientedPiece [4][4]rune
type Piece []OrientedPiece
type Board [14][14]rune

func turn(piece OrientedPiece) OrientedPiece {
	var turned OrientedPiece
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			turned[i][j] = piece[3-j][i]
		}
	}
	return turned
}

func flip(piece OrientedPiece) OrientedPiece {
	var flipped OrientedPiece
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			flipped[i][3-j] = piece[i][j]
		}
	}
	return flipped
}

func shift(piece OrientedPiece) OrientedPiece {
	var shifted OrientedPiece
	// copy into shifted
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			shifted[i][j] = piece[i][j]
		}
	}
	// shift up
	for c := 0; c < 4; c++ {
		if shifted[0][0] == ' ' && shifted[0][1] == ' ' && shifted[0][2] == ' ' && shifted[0][3] == ' ' {
			for i := 1; i < 4; i++ {
				for j := 0; j < 4; j++ {
					shifted[i-1][j] = shifted[i][j]
				}
			}
			for j := 0; j < 4; j++ {
				shifted[3][j] = ' '
			}
		} else {
			break
		}
	}
	// shift left
	for c := 0; c < 4; c++ {
		if shifted[0][0] == ' ' && shifted[1][0] == ' ' && shifted[2][0] == ' ' && shifted[3][0] == ' ' {
			for i := 0; i < 4; i++ {
				for j := 1; j < 4; j++ {
					shifted[i][j-1] = shifted[i][j]
				}
			}
			for i := 0; i < 4; i++ {
				shifted[i][3] = ' '
			}
		} else {
			break
		}
	}
	return shifted
}

func orientations(piece OrientedPiece) Piece {
	o := Piece{
		shift(piece), shift(turn(piece)), shift(turn(turn(piece))), shift(turn(turn(turn(piece)))),
		shift(flip(piece)), shift(turn(flip(piece))), shift(turn(turn(flip(piece)))), shift(turn(turn(turn(flip(piece))))),
	}
	oUniq := Piece{}
	for _, v := range o {
		found := false
		for _, u := range oUniq {
			if u == v {
				found = true
				break
			}
		}
		if !found {
			oUniq = append(oUniq, v)
		}
	}
	return oUniq
}

func (b *Board) place(piece OrientedPiece, mark rune, ipos int, jpos int) bool {
	canPlace := true
	for i := 0; i < 4 && canPlace; i++ {
		for j := 0; j < 4 && canPlace; j++ {
			if piece[i][j] != ' ' && (*b)[i+ipos][j+jpos] != ' ' {
				canPlace = false
			}
		}
	}
	if canPlace {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if piece[i][j] != ' ' {
					(*b)[i+ipos][j+jpos] = mark
				}
			}
		}
	}
	return canPlace
}

func (b *Board) remove(mark rune) {
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			if (*b)[i][j] == mark {
				(*b)[i][j] = ' '
			}
		}
	}
}

func (b *Board) month(m string, mark rune) error {
	m = strings.ToLower(m)
	switch {
	case m == "jan":
		(*b)[3][3] = mark
		return nil
	case m == "feb":
		(*b)[3][4] = mark
		return nil
	case m == "mar":
		(*b)[3][5] = mark
		return nil
	case m == "apr":
		(*b)[3][6] = mark
		return nil
	case m == "may":
		(*b)[3][7] = mark
		return nil
	case m == "jun":
		(*b)[3][8] = mark
		return nil
	case m == "jul":
		(*b)[4][3] = mark
		return nil
	case m == "aug":
		(*b)[4][4] = mark
		return nil
	case m == "sep":
		(*b)[4][5] = mark
		return nil
	case m == "oct":
		(*b)[4][6] = mark
		return nil
	case m == "nov":
		(*b)[4][7] = mark
		return nil
	case m == "dec":
		(*b)[4][8] = mark
		return nil
	}
	return fmt.Errorf("unknown month %q", m)
}

func (b *Board) date(d int, mark rune) error {
	switch {
	case d == 1:
		(*b)[5][3] = mark
		return nil
	case d == 2:
		(*b)[5][4] = mark
		return nil
	case d == 3:
		(*b)[5][5] = mark
		return nil
	case d == 4:
		(*b)[5][6] = mark
		return nil
	case d == 5:
		(*b)[5][7] = mark
		return nil
	case d == 6:
		(*b)[5][8] = mark
		return nil
	case d == 7:
		(*b)[5][9] = mark
		return nil
	case d == 8:
		(*b)[6][3] = mark
		return nil
	case d == 9:
		(*b)[6][4] = mark
		return nil
	case d == 10:
		(*b)[6][5] = mark
		return nil
	case d == 11:
		(*b)[6][6] = mark
		return nil
	case d == 12:
		(*b)[6][7] = mark
		return nil
	case d == 13:
		(*b)[6][8] = mark
		return nil
	case d == 14:
		(*b)[6][9] = mark
		return nil
	case d == 15:
		(*b)[7][3] = mark
		return nil
	case d == 16:
		(*b)[7][4] = mark
		return nil
	case d == 17:
		(*b)[7][5] = mark
		return nil
	case d == 18:
		(*b)[7][6] = mark
		return nil
	case d == 19:
		(*b)[7][7] = mark
		return nil
	case d == 20:
		(*b)[7][8] = mark
		return nil
	case d == 21:
		(*b)[7][9] = mark
		return nil
	case d == 22:
		(*b)[8][3] = mark
		return nil
	case d == 23:
		(*b)[8][4] = mark
		return nil
	case d == 24:
		(*b)[8][5] = mark
		return nil
	case d == 25:
		(*b)[8][6] = mark
		return nil
	case d == 26:
		(*b)[8][7] = mark
		return nil
	case d == 27:
		(*b)[8][8] = mark
		return nil
	case d == 28:
		(*b)[8][9] = mark
		return nil
	case d == 29:
		(*b)[9][3] = mark
		return nil
	case d == 30:
		(*b)[9][4] = mark
		return nil
	case d == 31:
		(*b)[9][5] = mark
		return nil
	}
	return fmt.Errorf("unknown date %d", d)
}

func (b *Board) weekDay(w string, mark rune) error {
	w = strings.ToLower(w)
	switch {
	case w == "sun":
		(*b)[9][6] = mark
		return nil
	case w == "mon":
		(*b)[9][7] = mark
		return nil
	case w == "tue":
		(*b)[9][8] = mark
		return nil
	case w == "wed":
		(*b)[9][9] = mark
		return nil
	case w == "thu":
		(*b)[10][7] = mark
		return nil
	case w == "fri":
		(*b)[10][8] = mark
		return nil
	case w == "sat":
		(*b)[10][9] = mark
		return nil
	}
	return fmt.Errorf("unknown weekDay %q", w)
}

func emptyBoard() Board {
	b := Board{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#'},
		{'#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', ' ', ' ', ' ', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	return b
}

func piece(n int) (OrientedPiece, error) {
	switch {
	case n == 0:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
		}, nil
	case n == 1:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', '*', ' ', ' '},
		}, nil
	case n == 2:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', '*', ' ', ' '},
			{' ', '*', ' ', ' '},
		}, nil
	case n == 3:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', '*', '*', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	case n == 4:
		return OrientedPiece{
			{'*', '*', ' ', ' '},
			{'*', '*', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	case n == 5:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', '*', '*', ' '},
			{'*', ' ', ' ', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	case n == 6:
		return OrientedPiece{
			{'*', '*', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', '*', ' ', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	case n == 7:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', '*', ' ', ' '},
			{' ', '*', ' ', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	case n == 8:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', '*', '*', ' '},
			{' ', ' ', '*', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	case n == 9:
		return OrientedPiece{
			{'*', ' ', ' ', ' '},
			{'*', ' ', ' ', ' '},
			{'*', '*', ' ', ' '},
			{' ', ' ', ' ', ' '},
		}, nil
	default:
		return OrientedPiece{}, fmt.Errorf("piece index %d out of bounds", n)
	}
}

func pieces(seed int64) []Piece {
	rand.Seed(seed)
	ret := []Piece{}
	for i := 0; i < 10; i++ {
		op, err := piece(i)
		if err != nil {
			panic(err)
		}
		p := orientations(op)
		rand.Shuffle(len(p), func(i, j int) {
			p[i], p[j] = p[j], p[i]
		})
		ret = append(ret, p)
	}
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

func (p *OrientedPiece) output(title string) {
	fmt.Println(title)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%c ", (*p)[i][j])
		}
		fmt.Println()
	}
}
func (o *Piece) output(title string) {
	fmt.Println(title)
	for i := 0; i < 4; i++ {
		for _, p := range *o {
			for j := 0; j < 4; j++ {
				fmt.Printf("%c ", p[i][j])
			}
			fmt.Print("    ")
		}
		fmt.Println()
	}

}

func (b *Board) output(title string) {
	fmt.Println(title)
	for i := 2; i < 12; i++ {
		for j := 2; j < 11; j++ {
			fmt.Printf("%c ", (*b)[i][j])
		}
		fmt.Println()
	}
}

func (b *Board) solve(remainingPieces []Piece) bool {
	if len(remainingPieces) == 0 {
		return true
	}
	mark := rune('0' + len(remainingPieces) - 1)
	orientations := remainingPieces[0]
	for _, o := range orientations {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if b.place(o, mark, i, j) {
					if b.solve(remainingPieces[1:]) {
						return true
					}
					b.remove(mark)
				}
			}
		}
	}
	return false
}

func run(args []string) error {
	if len(args) != 4 {
		return fmt.Errorf("usage: %s weekDay month date\nexample: %s mon apr 1", args[0], args[0])
	}
	b := emptyBoard()
	if len(args[1]) < 3 {
		return fmt.Errorf("invalid weekDay %s", args[1])
	}
	if err := b.weekDay(args[1][:3], '-'); err != nil {
		return err
	}
	if len(args[2]) < 3 {
		return fmt.Errorf("invalid month %s", args[2])
	}
	if err := b.month(args[2][:3], '-'); err != nil {
		return err
	}
	d, err := strconv.Atoi(args[3])
	if err != nil {
		return fmt.Errorf("usage: %q is not an int", args[3])
	}
	if err := b.date(d, '-'); err != nil {
		return err
	}

	if b.solve(pieces(time.Now().UnixNano())) {
		b.output("solution:")
	}
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

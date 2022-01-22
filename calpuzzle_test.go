package main

import (
	"fmt"
	"testing"
)

func TestTurn(t *testing.T) {
	p := OrientedPiece{
		{' ', ' ', ' ', '*'},
		{' ', ' ', ' ', '*'},
		{' ', ' ', ' ', '*'},
		{' ', ' ', ' ', '*'},
	}
	pt := OrientedPiece{
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{'*', '*', '*', '*'},
	}
	ptt := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
	}
	if pt != turn(p) {
		t.Errorf("turn(%q) == %q, want %q", p, turn(p), pt)
	}
	if ptt != turn(turn(p)) {
		t.Errorf("turn(turn(%q)) == %q, want %q", p, turn(turn(p)), ptt)
	}
	if ptt != turn(pt) {
		t.Errorf("turn(%q)) == %q, want %q", pt, turn(pt), ptt)
	}
}

func TestFlip(t *testing.T) {
	p := OrientedPiece{
		{' ', ' ', ' ', '*'},
		{' ', '*', '*', '*'},
		{' ', ' ', ' ', '*'},
		{' ', ' ', ' ', '*'},
	}
	pf := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', '*', '*', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
	}
	if pf != flip(p) {
		t.Errorf("flip(%q) == %q, want %q", p, flip(p), pf)
	}
}

func TestShift(t *testing.T) {
	p := OrientedPiece{
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', '*'},
	}
	ps := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
	}
	if ps != shift(p) {
		t.Errorf("shift(%q) == %q, want %q", p, shift(p), ps)
	}
}

func TestOrientations(t *testing.T) {
	p := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
	}
	o := orientations(p)
	if len(o) != 2 {
		t.Errorf("expect 2 orientations for %q, found %d", p, len(o))
		o.output("")
	}
}
func TestOrientations2(t *testing.T) {
	p := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', '*', ' ', ' '},
	}
	o := orientations(p)
	if len(o) != 8 {
		t.Errorf("expect 8 orientations for %q, found %d", p, len(o))
		o.output("")
	}
	// oAll := Piece{
	// 	shift(p), shift(turn(p)), shift(turn(turn(p))), shift(turn(turn(turn(p)))),
	// 	shift(flip(p)), shift(turn(flip(p))), shift(turn(turn(flip(p)))), shift(turn(turn(turn(flip(p))))),
	// }
	// oAll.output("debug")
}

func TestPlaceFails(t *testing.T) {
	p := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
	}
	b := emptyBoard()
	placed := b.place(p, '?', 0, 0)
	if placed {
		t.Errorf("placing %q on %q at 0, 0 should fail", p, emptyBoard())
	}
}

func TestPlaceSucceeds(t *testing.T) {
	p := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
		{'*', ' ', ' ', ' '},
	}
	b := emptyBoard()
	placed := b.place(p, '?', 3, 3)
	if !placed {
		t.Errorf("placing %q on %q at 3, 3 should succeed", p, emptyBoard())
	}
	if b[3][3] != '?' {
		t.Errorf("placed %q on %q, no ? at 3, 3", p, b)
	}
}

func TestPlace(t *testing.T) {
	p := OrientedPiece{
		{'*', ' ', ' ', ' '},
		{'*', '*', '*', ' '},
		{' ', ' ', '*', ' '},
		{' ', ' ', ' ', ' '},
	}
	b := Board{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', ' ', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', ' ', ' ', ' ', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', ' ', '#', ' ', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	placed := b.place(p, '.', 5, 4)
	if !placed {
		t.Errorf("failed to place %q on board", p)
		b.output("board")
	}

}

func TestPieces(t *testing.T) {
	p := pieces(0)
	if len(p) != 10 {
		t.Errorf("expected 10 pieces, got %q, pieces: %q", len(p), p)
	}
	for i, o := range p {
		o.output(fmt.Sprintf("orientation %d", i))
	}
}

package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

type level struct {
	num       int
	tiles     [][]tile
	landmarks map[position]tile // TODO Rename
}

// Tiles
type tile = byte

const (
	EMPTY        = ' '
	RUNNER       = '&'
	GUARD        = '0'
	BRICK        = '#'
	SOLIDBRICK   = '@'
	FALSEBRICK   = 'X'
	LADDER       = 'H'
	ESCAPELADDER = 'S'
	ROPE         = '-'
	GOLD         = '$'
)

// Position
type position struct{ x, y int }

func manhattanDist(a, b position) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}

func newLevel(num int) (*level, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("levels/%03d.lvl", num))
	if err != nil {
		return nil, err
	}

	lvl := &level{
		num, bytes.Split(content, []byte("\n")), make(map[position]tile),
	}
	for i, tiles := range lvl.tiles {
		for j, tile := range tiles {
			if tile == RUNNER || tile == GUARD || tile == GOLD {
				lvl.landmarks[position{j, i}] = tile
				lvl.tiles[i][j] = EMPTY
			}
		}
	}

	return lvl, nil
}

func (l *level) String() string {
	return string(bytes.Join(l.getTiles(), []byte("\n")))
}

func (l *level) emptyBelow(pos position) bool {
	return l.getTiles()[pos.y+1][pos.x] == EMPTY
}

// TODO Rewrite + Rename
func (l *level) getTiles() [][]tile {
	tiles := make([][]tile, len(l.tiles))
	for i := range tiles {
		tiles[i] = make([]tile, len(l.tiles[i]))
		copy(tiles[i], l.tiles[i])
	}

	// TODO Comment
	for pos, tile := range l.landmarks {
		tiles[pos.y][pos.x] = tile
	}

	return tiles
}

func (l *level) validMove(orig, dest position, dir direction) bool {
	if dest.x < 0 || dest.x >= 28 || /*dest.y < 0 ||*/ dest.y >= 16 {
		return false
	}

	origTile := l.tiles[orig.y][orig.x]
	if dir == DOWN && origTile == ROPE {
		return false
	}
	if dest.y < 0 && origTile == ESCAPELADDER {
		return true
	}

	switch destTile := l.tiles[dest.y][dest.x]; destTile {
	case EMPTY, ROPE:
		if dir == UP {
			return origTile == LADDER || origTile == ESCAPELADDER
		}
		return true
	/*
		case BRICK:
			return dir != UP && false // TODO && bricksbrokenat(dest)
	*/
	case BRICK, SOLIDBRICK:
		return false
	case LADDER:
		return dir != DOWN
	case ESCAPELADDER:
		if l.getTiles()[orig.y][orig.x] == RUNNER {
			return dir != DOWN
		}
		return dir != UP
	}

	return false
}

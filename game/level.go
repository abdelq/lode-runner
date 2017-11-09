package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

type tile = byte
type position struct{ x, y int }
type level struct {
	num       int
	tiles     [][]tile
	landmarks map[position]tile // TODO Rename
	game      *Game             // TODO Temporary
}

// Tiles
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

func newLevel(num int) (*level, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("levels/%03d.lvl", num))
	if err != nil {
		return nil, err
	}

	// TODO Rewrite this section
	lvl := &level{
		num:       num,
		tiles:     bytes.Split(content, []byte("\n")),
		landmarks: make(map[position]tile),
	}
	for i, tiles := range lvl.tiles {
		for j, tile := range tiles {
			/*if tile == GOLD {
				lvl.landmarks[position{j, i}] = tile
			} else */if tile == GUARD || tile == RUNNER {
				lvl.landmarks[position{j, i}] = tile
				lvl.tiles[i][j] = EMPTY // TODO
			}
		}
	}

	return lvl, nil
}

// TODO Rename function
// TODO Try to get rid of game usage
func (l *level) getTiles() [][]tile {
	// TODO Rewrite
	tiles := make([][]tile, len(l.tiles))
	for i := range tiles {
		tiles[i] = make([]tile, len(l.tiles[i]))
		copy(tiles[i], l.tiles[i])
	}

	// Runner
	tiles[l.game.Runner.pos.y][l.game.Runner.pos.x] = RUNNER

	// Guards
	for guard := range l.game.Guards {
		tiles[guard.pos.y][guard.pos.x] = GUARD
	}

	return tiles
}

func (l *level) emptyBelow(pos position) bool {
	return l.getTiles()[pos.y+1][pos.x] == EMPTY
}

func manhattanDist(a, b position) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}

// TODO Rename + Interface
/*func (l *level) toString() string {
	return bytes.Join(lvl.getTiles(), []byte("\n"))
}*/

/*func (l *level) print() {
// TODO One-liner using join ?
	for _, row := range l.grid {
		fmt.Println(string(row))
	}
}*/

/*func (l *level) replaceTiles(positions []position, tile tile) {
	for _, pos := range positions {
		l.tiles[pos.y][pos.x] = tile
	}
}*/

func (l *level) validMove(orig position, dest position, dir direction) bool {
	if dest.x < 0 || dest.x >= 28 || /*dest.y < 0 ||*/ dest.y >= 16 {
		return false
	}

	// TODO valid_decor_move
	origTile := l.tiles[orig.y][orig.x]
	destTile := l.tiles[dest.y][dest.x]
	if dir == DOWN && origTile == ROPE {
		return false
	}
	if dest.y < 0 && origTile == ESCAPELADDER {
		return true
	}

	switch destTile {
	case EMPTY, ROPE:
		if dir == UP {
			return origTile == LADDER || origTile == ESCAPELADDER
		} else {
			return true
		}
	case BRICK:
		return dir != UP && false // TODO && bricksbrokenat(dest)
	case SOLIDBRICK:
		return false
	case LADDER:
		return dir != DOWN
	case ESCAPELADDER:
		if l.getTiles()[orig.y][orig.x] == RUNNER {
			return dir != DOWN
		} else {
			return dir != UP
		}
	}

	return false
}

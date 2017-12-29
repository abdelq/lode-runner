package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

type tile = byte
type level struct {
	num     int
	tiles   [][]tile
	players map[position]tile
	gold    []position
	escape  []position
	holes   map[position]uint8
}

// Tiles
const (
	EMPTY   = ' '
	RUNNER  = '&'
	GUARD   = '0'
	BRICK   = '#'
	BLOCK   = '@'
	TRAP    = 'X'
	LADDER  = 'H'
	HLADDER = 'S'
	ROPE    = '-'
	GOLD    = '$'
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

	// TODO Specify right len/cap
	lvl := &level{
		num, bytes.Split(content, []byte("\n")),
		make(map[position]tile), make([]position, 0),
		make([]position, 0), make(map[position]uint8),
	}

	// Collect data for players/gold
	for i, tiles := range lvl.tiles {
		for j, tile := range tiles {
			switch tile {
			case RUNNER, GUARD:
				lvl.players[position{j, i}] = tile
				lvl.tiles[i][j] = EMPTY
			case GOLD:
				lvl.gold = append(lvl.gold, position{j, i})
				lvl.tiles[i][j] = EMPTY
			case HLADDER:
				lvl.escape = append(lvl.escape, position{j, i})
				lvl.tiles[i][j] = EMPTY
			}
		}
	}

	return lvl, nil
}

func (l *level) String() string {
	return string(bytes.Join(l.getTiles(), []byte("\n")))
}

func (l *level) stringTiles() []string {
	getTiles := l.getTiles()
	tiles := make([]string, len(l.tiles))
	for i := range tiles {
		tiles[i] = string(getTiles[i])
	}
	return tiles
}

func (l *level) width() int {
	return len(l.tiles[0])
}

func (l *level) height() int {
	return len(l.tiles)
}

func (l *level) emptyBelow(pos position) bool {
	if pos.y+1 >= l.height() /*-1*/ {
		return false
	}

	// XXX
	return l.getTiles()[pos.y+1][pos.x] == EMPTY ||
		l.tiles[pos.y+1][pos.x] == ROPE ||
		l.getTiles()[pos.y+1][pos.x] == GOLD
}

func (l *level) goldCollected() bool {
	return len(l.gold) == 0
}

// TODO Rewrite + Rename
func (l *level) getTiles() [][]tile {
	tiles := make([][]tile, len(l.tiles))
	for i := range tiles {
		tiles[i] = make([]tile, len(l.tiles[i]))
		copy(tiles[i], l.tiles[i])
	}

	// Gold
	for _, pos := range l.gold {
		tiles[pos.y][pos.x] = GOLD
	}

	// Escape ladders
	if l.goldCollected() {
		for _, pos := range l.escape {
			tiles[pos.y][pos.x] = HLADDER
		}
	}

	// Players
	for pos, tile := range l.players {
		tiles[pos.y][pos.x] = tile
	}

	return tiles
}

func (l *level) validMove(orig, dest position, dir uint8) bool {
	if dest.x < 0 || dest.x >= l.width() /*|| dest.y < 0*/ || dest.y >= l.height() /*-1*/ {
		return false
	}

	origTile, destTile := l.tiles[orig.y][orig.x], l.tiles[dest.y][dest.x]
	switch destTile {
	case EMPTY, ROPE:
		if dir == UP {
			return origTile == LADDER || origTile == HLADDER
		}
		return true
	case BRICK, BLOCK:
		return false
	case LADDER, HLADDER:
		return true
	}

	return false
}

func (l *level) validDig(pos position) bool {
	if pos.x < 0 || pos.x >= l.width() || pos.y >= l.height() {
		return false
	}

	return l.tiles[pos.y][pos.x] == BRICK
}

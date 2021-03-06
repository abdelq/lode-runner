package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sync"
)

type tile = byte
type level struct {
	num     int
	tiles   [][]tile
	players sync.Map
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

	// Randomly flip the level horizontally
	lines := bytes.Split(content, []byte("\n"))
	if rand.Intn(2) == 0 {
		for i, line := range lines {
			for j := 0; j < len(line)/2; j++ {
				lines[i][j], lines[i][len(line)-1-j] = lines[i][len(line)-1-j], lines[i][j]
			}
		}
	}

	// TODO Specify right len/cap
	lvl := &level{
		num: num, tiles: lines, gold: make([]position, 0),
		escape: make([]position, 0), holes: make(map[position]uint8),
	}

	// Collect data for players/gold
	for i, tiles := range lvl.tiles {
		for j, tile := range tiles {
			switch tile {
			case RUNNER:
				lvl.players.Store(position{j, i}, tile)
				lvl.tiles[i][j] = EMPTY
			case GUARD:
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

	// Keep only 3 random pieces of gold for the first levels
	if num < 9 {
		for i := range lvl.gold {
			j := rand.Intn(i + 1)
			lvl.gold[i], lvl.gold[j] = lvl.gold[j], lvl.gold[i]
		}
		lvl.gold = lvl.gold[:3]
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
	if pos.y+1 >= l.height() {
		return false
	}

	// XXX
	return l.getTiles()[pos.y+1][pos.x] == EMPTY ||
		l.tiles[pos.y+1][pos.x] == ROPE ||
		l.getTiles()[pos.y+1][pos.x] == GOLD ||
		l.getTiles()[pos.y+1][pos.x] == TRAP
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
	// if l.goldCollected() {
	for _, pos := range l.escape {
		tiles[pos.y][pos.x] = HLADDER
	}
	// }

	// Players
	l.players.Range(func(pos, tile interface{}) bool {
		tiles[pos.(position).y][pos.(position).x] = tile.(uint8)
		return true
	})

	return tiles
}

func (l *level) validMove(orig, dest position, dir uint8) bool {
	if dest.x < 0 || dest.x >= l.width() ||
		dest.y < 0 || dest.y >= l.height() {
		return false
	}

	origTile := l.tiles[orig.y][orig.x]
	switch l.tiles[dest.y][dest.x] {
	case EMPTY, ROPE, TRAP:
		if dir == UP {
			return origTile == LADDER || origTile == HLADDER
		}
		return true
	case BRICK, BLOCK:
		return false
	case LADDER:
		return true
	}

	return false
}

func (l *level) validDig(pos position) bool {
	if pos.x < 0 || pos.x >= l.width() ||
		pos.y < 0 || pos.y >= l.height() {
		return false
	}

	// FIXME If not falling

	return l.tiles[pos.y][pos.x] == BRICK
}

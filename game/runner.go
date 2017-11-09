package game

type Runner struct {
	Name  string
	pos   position
	state state
}

// TODO Maybe just use landmarks directly
func (r *Runner) init(game *Game) {
	for pos, tile := range game.Level.landmarks {
		if tile == RUNNER {
			r.pos = pos
			delete(game.Level.landmarks, pos) // TODO Test
			return
		}
	}
}

func (r *Runner) Move(lvl *level, dir direction) {
	if r.state == DIGGING {
		return
	}

	if r.state == FALLING && dir != DOWN {
		dir = DOWN
	}

	// TODO Improve
	var newPos position
	switch dir {
	case NONE:
		newPos = position{r.pos.x, r.pos.y}
	case UP:
		newPos = position{r.pos.x, r.pos.y - 1}
	case LEFT:
		newPos = position{r.pos.x - 1, r.pos.y}
	case DOWN:
		newPos = position{r.pos.x, r.pos.y + 1}
	case RIGHT:
		newPos = position{r.pos.x + 1, r.pos.y}
	}

	if !lvl.validMove(r.pos, newPos, dir) {
		if r.state == FALLING {
			r.state = ALIVE
		}
		return
	}

	r.pos.x, r.pos.y = newPos.x, newPos.y // TODO Improve

	if dir == DOWN || lvl.emptyBelow(r.pos) {
		r.state = FALLING
	}

	// TODO
	//gfx_move_sprite(HERO, orig, hero.pos)
	//game.check_collisions()
}

func (r *Runner) Dig(lvl *level, dir direction) {} // TODO TODO

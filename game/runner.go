package game

type Runner struct {
	Name  string
	pos   *position
	state state
}

// TODO Replace in original lvl by empty
func (r *Runner) init(game *Game) {
	for i := len(game.Lvl.grid) - 1; i >= 0; i-- {
		for j := len(game.Lvl.grid[i]) - 1; j >= 0; j-- {
			if game.Lvl.grid[i][j] == RUNNER {
				game.Lvl.grid[i][j] = EMPTY // TODO Try using cell = ?
				r.pos.x, r.pos.y = j, i
				return
			}
		}
	}
}

// TODO Timeout + Direction
func (r *Runner) Move(direction uint8, game *Game) {
	if r.state == DIGGING {
		return
	}

	if r.state == FALLING && direction != DOWN {
		direction = DOWN
	}

	// TODO Check if position changes when getting into function (passage valeur)
	if !game.Lvl.validMove(*(r.pos), direction) {
		if r.state == FALLING {
			// TODO r.state &= ~STATE_FALLING
		}
		return
	} else {
		r.pos.set(direction)
	}

	if direction == DOWN || game.Lvl.emptyBelow(*(r.pos)) {
		//r.state = FALLING
		// TODO r.state |= FALLING
	}

	// TODO
	//gfx_move_sprite(HERO, orig, hero.pos)
	//game.check_collisions()
}

// TODO
func (r *Runner) Dig(direction uint8, game *Game) {}

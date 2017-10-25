package main

type runner struct {
	name  string
	pos   position
	state state
}

func (r *runner) init(grid [][]byte) {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == RUNNER {
				r.pos.x, r.pos.y = j, i
				return
			}
		}
	}
}

// TODO
func (r *runner) move(direction string, game *game) {
	// TODO Timeout ?
	if r.state == DIGGING {
		return
	}

	if r.state == FALLING && direction != "down" {
		direction = "down"
	}

	/*if !game.lvl.valid_move(r.pos, direction) {
		if r.state == FALLING {
			hero.state &= ~STATE_FALLING
		}
		return
	}*/

	switch direction {
	case "up":
		r.pos.y--
	case "left":
		r.pos.x--
	case "down":
		r.pos.y++
	case "right":
		r.pos.x++
	}

	if direction == "down" || game.lvl.emptyBelow(r.pos) {
		r.state = FALLING
	}

	//gfx_move_sprite(HERO, orig, hero.pos)
	//game.check_collisions()
}

// TODO
func (r *runner) dig(direction string, game *game) {
	// TODO Timeout ?
	/*var digPos *position
	switch direction {
	case "left":
		digPos = &position{r.pos.x - 1, r.pos.y + 1}
	case "right":
		digPos = &position{r.pos.x + 1, r.pos.y + 1}
	default:
		// TODO Error
	}*/

	/*if game.lvl.validDig(digPos) {
		r.state = DIGGING
		// bricks_break(&digPos)
	}*/
}

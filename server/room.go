package main

type Room struct {
	players map[*Player]bool
	join    chan *Player
	leave   chan *Player
}

func newRoom() *Room {
	return &Room{
		join:    make(chan *Player),
		leave:   make(chan *Player),
		players: make(map[*Player]bool),
	}
}

func (room *Room) run() {
	for {
		select {
		case player := <-room.join:
			player.room = room
			room.players[player] = true
		case player := <-room.leave:
			if _, ok := room.players[player]; ok {
				player.room = nil
				delete(room.players, player)
			}
		}
	}
}

package main

var rooms = map[string]*Room{}

type Room struct {
	join  chan *Player
	leave chan *Player
	//broadcast chan []byte
	players map[*Player]bool
}

func newRoom() *Room {
	return &Room{
		join:  make(chan *Player),
		leave: make(chan *Player),
		//broadcast: make(chan []byte),
		players: make(map[*Player]bool),
	}
}

func (room *Room) run() {
	for {
		select {
		case player := <-room.join:
			room.players[player] = true
			player.room = room
		case player := <-room.leave:
			if _, ok := room.players[player]; ok {
				delete(room.players, player)
				player.room = nil
				//close(player.send)
			}
			/*case message := <-room.broadcast:
			for player := range room.players {
				select {
				case player.send <- message:
				default:
					close(player.send)
					delete(room.players, player)
				}
			}*/
		}
	}
}

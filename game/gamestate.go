package game

// GameState is a struct containing global information about the game
type GameState struct {
	roomPlayers map[*Room][]*Player // map of all rooms to their players
	rooms       map[string]*Room    // all the rooms in the game
	players     map[string]*Player  //all the players in the game
}

// AddRoom adds the given room to the game.
func (s *GameState) AddRoom(r *Room) {
	s.rooms[r.Name] = r
}

// AddPlayer adds the given player to the given room.
func (s *GameState) AddPlayer(player *Player) {
	s.players[player.Name] = player
}

// SetPlayer places player in the given room.
func (s *GameState) SetPlayer(player string, room string) {
	p := s.players[player]
	r := s.rooms[room]
	s.roomPlayers[r] = append(s.roomPlayers[r], p)
	p.Room = r
}

// GameState is the global game state.
var State GameState

// Init initializes the game state.
func Init() {
	State.roomPlayers = make(map[*Room][]*Player)
	State.rooms = make(map[string]*Room)
	State.players = make(map[string]*Player)
}

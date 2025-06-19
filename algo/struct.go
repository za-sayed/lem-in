package algo

type Graph struct {
	NumberOfAnts int
	StartRoom    *Room
	EndRoom      *Room
	Rooms        map[string]*Room
	Links        []string
}

type Room struct {
	Name           string
	X, Y           int
	IsStart        bool
	IsEnd          bool
	ConnectedRooms []*Room
}

type Path struct {
	Rooms []*Room
}

type Solution struct {
	Paths []*Path
}

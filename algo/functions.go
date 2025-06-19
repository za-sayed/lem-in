package algo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to parse and create a room
func parseRoom(line string, rooms map[string]*Room, isStart bool, isEnd bool) string {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		fmt.Println("ERROR: Invalid room format", line)
		os.Exit(1)
	}

	name, xStr, yStr := fields[0], fields[1], fields[2]
	if strings.HasPrefix(name, "#") || strings.HasPrefix(name, "L") {
		fmt.Println("ERROR: Invalid room name", name)
		os.Exit(1)
	}

	x, err1 := strconv.Atoi(xStr)
	y, err2 := strconv.Atoi(yStr)
	if err1 != nil || err2 != nil {
		fmt.Println("ERROR: Invalid room coordinates", line)
		os.Exit(1)
	}

	// Check for duplicate coordinates
	for _, room := range rooms {
		if room.X == x && room.Y == y {
			fmt.Println("ERROR: Duplicate room coordinates", x, y)
			os.Exit(1)
		}
	}

	if _, exists := rooms[name]; exists {
		fmt.Println("ERROR: Duplicate room name", name)
		os.Exit(1)
	}

	// Create and initialize the room
	room := &Room{Name: name, X: x, Y: y, ConnectedRooms: []*Room{}}
	room.IsStart = isStart
	room.IsEnd = isEnd

	rooms[name] = room
	return name
}

// Function to parse and create a tunnel between rooms
func parseTunnel(line string, rooms map[string]*Room, tunnels *[]string) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 || parts[0] == parts[1] {
		fmt.Println("ERROR: Invalid tunnel definition", line)
		os.Exit(1)
	}
	room1, ok1 := rooms[parts[0]]
	room2, ok2 := rooms[parts[1]]
	if !ok1 || !ok2 {
		fmt.Println("ERROR: Tunnel references unknown room(s)", line)
		os.Exit(1)
	}

	// Check for duplicate tunnels
	for _, tunnel := range *tunnels {
		if tunnel == line || tunnel == (parts[1]+"-"+parts[0]) {
			fmt.Println("ERROR: Duplicate tunnel", line)
			os.Exit(1)
		}
	}

	*tunnels = append(*tunnels, line)

	// Create bidirectional connections
	room1.ConnectedRooms = append(room1.ConnectedRooms, room2)
	room2.ConnectedRooms = append(room2.ConnectedRooms, room1)
}

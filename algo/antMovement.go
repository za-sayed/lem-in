package algo

import (
	"fmt"
	"strings"
)

func PrintAntMovements(solution *Solution, graph *Graph) {
	DirPathIndex, hasDirPath := FindDirectPath(solution, graph.StartRoom, graph.EndRoom)
	antPaths, antPositions := make([]int, graph.NumberOfAnts), make([]int, graph.NumberOfAnts)
	roomAvailable := make(map[*Room]bool)
	
	// Assign ants to paths
	for i := 0; i < graph.NumberOfAnts; i++ {
		antPaths[i] = i % len(solution.Paths)
		if i == graph.NumberOfAnts-1 && hasDirPath {
			antPaths[i] = DirPathIndex
		}
	}

	// Initialize room availability
	for _, path := range solution.Paths {
		for _, room := range path.Rooms {
			roomAvailable[room] = true
		}
	}

	for {
		var moves []string
		antMoved := false
		dirPathUsed := false

		for antNo := 0; antNo < graph.NumberOfAnts; antNo++ {
			path, pos := antPaths[antNo], antPositions[antNo]
			if pos >= len(solution.Paths[path].Rooms)-1 {
				continue
			}

			if len(solution.Paths[path].Rooms) == 2 {
				if dirPathUsed {
					continue
				}
				dirPathUsed = true
			}

			currentRoom, nextRoom := solution.Paths[path].Rooms[pos], solution.Paths[path].Rooms[pos+1]
			if roomAvailable[nextRoom] {
				if currentRoom != graph.StartRoom {
					roomAvailable[currentRoom] = true
				}
				if nextRoom != graph.EndRoom {
					roomAvailable[nextRoom] = false
				}

				antPositions[antNo]++
				moves = append(moves, fmt.Sprintf("L%d-%s", antNo+1, nextRoom.Name))
				antMoved = true
			}
		}

		if !antMoved {
			break
		}
		fmt.Println(strings.Join(moves, " "))
	}

}

func FindDirectPath(solution *Solution, startRoom, endRoom *Room) (int, bool) {
	for i, path := range solution.Paths {
		if len(path.Rooms) == 2 && path.Rooms[0] == startRoom && path.Rooms[1] == endRoom {
			return i, true
		}
	}
	return -1, false
}

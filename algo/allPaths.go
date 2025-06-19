package algo

import "fmt"

func FindAllPaths(startRoom, endRoom string, rooms map[string]*Room) ([]Path, error) {
	type Node struct {
		room    *Room
		path    []*Room
		visited map[string]bool
	}

	// If start or end room does not exist
	if _, ok := rooms[startRoom]; !ok {
		return nil, fmt.Errorf("ERROR: Missing start room")
	}
	if _, ok := rooms[endRoom]; !ok {
		return nil, fmt.Errorf("ERROR: Missing end room")
	}

	// queue initialization
	queue := []Node{{
		room:    rooms[startRoom],
		path:    []*Room{rooms[startRoom]},
		visited: map[string]bool{startRoom: true},
	}}

	var allPaths []Path

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If we reach the end room, store the path
		if curr.room.Name == endRoom {
			allPaths = append(allPaths, Path{Rooms: curr.path})
			continue
		}

		// Explore all connected rooms
		for _, neighbor := range curr.room.ConnectedRooms {
			if !curr.visited[neighbor.Name] {
				// Create a new visited map
				newVisited := make(map[string]bool, len(curr.visited)+1)
				for k := range curr.visited {
					newVisited[k] = true
				}
				newVisited[neighbor.Name] = true

				// Create a new path slice
				newPath := append(curr.path[0:len(curr.path):len(curr.path)], neighbor)

				// Add to queue
				queue = append(queue, Node{
					room:    neighbor,
					path:    newPath,
					visited: newVisited,
				})
			}
		}
	}

	if len(allPaths) == 0 {
		return nil, fmt.Errorf("ERROR: No path found from start room to end room")
	}

	return allPaths, nil
}

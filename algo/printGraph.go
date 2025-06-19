package algo

import "fmt"

func PrintGraph(graph *Graph) {
	if graph.StartRoom == nil {
		fmt.Println("ERROR: Missing start room")
		return
	}

	if graph.EndRoom == nil {
		fmt.Println("ERROR: Missing end room")
		return
	}

	if len(graph.Rooms) == 0 {
		fmt.Println("ERROR: No rooms found")
		return
	}

	if len(graph.Links) == 0 {
		fmt.Println("ERROR: No links between rooms")
		return
	}

	// Print number of ants
	fmt.Printf("%d\n", graph.NumberOfAnts)

	// Print start room
	fmt.Printf("##start\n%s %d %d\n", graph.StartRoom.Name, graph.StartRoom.X, graph.StartRoom.Y)

	// Print rooms (excluding start and end)
	for name, room := range graph.Rooms {
		if !room.IsStart && !room.IsEnd {
			fmt.Printf("%s %d %d\n", name, room.X, room.Y)
		}
	}

	// Print end room
	fmt.Printf("##end\n%s %d %d\n", graph.EndRoom.Name, graph.EndRoom.X, graph.EndRoom.Y)

	// Print links
	for _, link := range graph.Links {
		fmt.Println(link)
	}
}

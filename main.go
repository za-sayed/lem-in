package main

import (
	"fmt"
	"lem-in/algo"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: No input file provided")
		return
	}

	graph, err := algo.ParseInputFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	algo.PrintGraph(graph)

	paths, err := algo.FindAllPaths(graph.StartRoom.Name, graph.EndRoom.Name, graph.Rooms)
	if err != nil {
		fmt.Println(err)
		return
	}

	solution := algo.UniquePaths(graph, paths)

	if len(solution.Paths) == 0 {
		fmt.Println("ERROR: No unique path found")
		return
	}

	//algo.PrintPaths(*solution)

	fmt.Println()

	algo.PrintAntMovements(solution, graph)

}

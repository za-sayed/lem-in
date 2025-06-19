package algo

import "fmt"

func PrintPaths(result Solution) {
	fmt.Println()
	for _, path := range result.Paths {
		for i, room := range path.Rooms {
			if i > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(room.Name)
		}
		fmt.Println()
	}
}

func PrintPaths2(paths []Path) {
	fmt.Println()
	for _, path := range paths {
		for i, room := range path.Rooms {
			if i > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(room.Name)
		}
		fmt.Println()
	}
}

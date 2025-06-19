package algo

func UniquePaths(graph *Graph, allPaths []Path) *Solution {
	if len(allPaths) == 0 {
		return &Solution{Paths: []*Path{}}
	}

	var bestPathsSet []Path
	maxUniqueCount := 0

	// Store sets of rooms for each path
	pathSets := make([]map[string]bool, len(allPaths))
	for i, path := range allPaths {
		pathSets[i] = getPathSet(graph, &path)
	}

	// Iterate over paths to find the best set of unique paths
	for i, pathSet := range pathSets {
		currentSet := []Path{allPaths[i]}
		currentSetRoomMap := make(map[string]bool, len(pathSet))

		// Copy path rooms into the map
		for room := range pathSet {
			currentSetRoomMap[room] = true
		}

		count := 1
		for j := 0; j < len(allPaths); j++ {
			if i == j {
				continue
			}

			otherSet := pathSets[j]
			if areSetsUnique(currentSetRoomMap, otherSet) {
				currentSet = append(currentSet, allPaths[j])
				count++

				// Add new path rooms to the current set
				for room := range otherSet {
					currentSetRoomMap[room] = true
				}
			}
		}

		if count > maxUniqueCount {
			maxUniqueCount = count
			bestPathsSet = currentSet
		}

		// Early exit if we find the maximum possible unique count
		if maxUniqueCount == len(allPaths) {
			break
		}
	}

	// If no unique paths were found, return at least one path
	if maxUniqueCount == 0 {
		return &Solution{Paths: []*Path{&allPaths[0]}}
	}

	// Convert bestPathsSet to []*Path for the Solution struct
	selectedPaths := make([]*Path, len(bestPathsSet))
	for i := range bestPathsSet {
		selectedPaths[i] = &bestPathsSet[i]
	}

	return &Solution{Paths: selectedPaths}
}

func getPathSet(graph *Graph, path *Path) map[string]bool {
	nodeSet := make(map[string]bool, len(path.Rooms))
	for _, node := range path.Rooms {
		if node != graph.StartRoom && node != graph.EndRoom {
			nodeSet[node.Name] = true
		}
	}
	return nodeSet
}

func areSetsUnique(existingSet, newSet map[string]bool) bool {
	for room := range newSet {
		if existingSet[room] {
			return false
		}
	}
	return true
}

package algo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInputFile(filename string) (*Graph, error) {
	if !strings.HasSuffix(filename, ".txt") {
		return nil, fmt.Errorf("ERROR: Invalid file type, only .txt files are accepted")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numAnts int = -1
	rooms := make(map[string]*Room)
	tunnels := []string{}
	var startRoom, endRoom string
	endDefined := false
	roomPattern := regexp.MustCompile(`^\S+ -*\d+ -*\d+$`)
	tunnelPattern := regexp.MustCompile(`^\S+-\S+$`)

	// Read number of ants
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		numAnts, err = strconv.Atoi(line)
		if err == nil && numAnts > 0 {
			break
		}
	}

	if numAnts <= 0 || err != nil || numAnts > 10000 {
		return nil, fmt.Errorf("ERROR: Couldn't find a valid number of ants")
	}

	// Read and parse the rest of the file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "##") {
			if strings.ToLower(line) == "##start" {
				if startRoom != "" {
					return nil, fmt.Errorf("ERROR: Multiple start rooms defined")
				}
				if endDefined {
					return nil, fmt.Errorf("ERROR: Start must be defined before End")
				}
				if !scanner.Scan() || strings.TrimSpace(scanner.Text()) == "" {
					return nil, fmt.Errorf("ERROR: Missing room after ##start")
				}
				startRoom = parseRoom(strings.TrimSpace(scanner.Text()), rooms, true, false)
			} else if strings.ToLower(line) == "##end" {
				if endRoom != "" {
					return nil, fmt.Errorf("ERROR: Multiple end rooms defined")
				}
				endDefined = true
				if !scanner.Scan() {
					return nil, fmt.Errorf("ERROR: Missing room after ##end")
				}
				endRoom = parseRoom(strings.TrimSpace(scanner.Text()), rooms, false, true)
			} 
			continue
		}
		if tunnelPattern.MatchString(line) {
			parseTunnel(line, rooms, &tunnels)
		} else if roomPattern.MatchString(line) {
			parseRoom(line, rooms, false, false)
		} else if strings.HasPrefix(line, "#") {
			continue
		} else {
			return nil, fmt.Errorf("ERROR: Invalid data Format: %s", line)
		}
	}
	
	if startRoom == "" || endRoom == "" {
		return nil, fmt.Errorf("ERROR: Missing start or end room")
	}
	if startRoom == endRoom {
		return nil, fmt.Errorf("ERROR: Start room and end room cannot be the same")
	}
	if len(rooms) == 0 {
		return nil, fmt.Errorf("ERROR: No rooms found")
	}
	if len(tunnels) == 0 {
		return nil, fmt.Errorf("ERROR: No links between rooms")
	}

	// Construct and return the graph
	graph := &Graph{
		NumberOfAnts: numAnts,
		StartRoom:    rooms[startRoom],
		EndRoom:      rooms[endRoom],
		Rooms:        rooms,
		Links:        tunnels,
	}

	return graph, nil
}

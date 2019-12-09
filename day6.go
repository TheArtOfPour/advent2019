package main

import (
	"strings"
)

type planet struct {
	name     string
	orbitals []*planet
}

func countOrbits(planets map[string][]string, current string, depth int) int {
	orbits := depth
	// for d := 0; d < depth; d++ {
	// 	print(" ")
	// }
	//print(current)
	// leaf node
	if len(planets[current]) == 0 {
		return orbits
	}
	// branch
	for i := 0; i < len(planets[current]); i++ {
		// print("->")
		// print(planets[current][i])
		// println()
		orbits += countOrbits(planets, planets[current][i], depth+1)
	}
	return orbits
}

func planetaryTree(planets map[string][]string, current string, parent *planet) *planet {
	thisPlanet := planet{name: current}
	thisPlanet.orbitals = append(thisPlanet.orbitals, parent)
	// leaf
	if len(planets[current]) == 0 {
		return &thisPlanet
	}
	// branch
	for i := 0; i < len(planets[current]); i++ {
		thisPlanet.orbitals = append(thisPlanet.orbitals, planetaryTree(planets, planets[current][i], &thisPlanet))
	}
	return &thisPlanet
}

func advent6A(test string) (int, error) {
	planets := make(map[string][]string)
	for _, s := range strings.Split(test, "\n") {
		entry := strings.Split(s, ")")
		name := strings.TrimSpace(entry[0])
		childName := strings.TrimSpace(entry[1])
		_, ok := planets[name]
		if !ok {
			planets[name] = []string{childName}
		} else {
			planets[name] = append(planets[name], childName)
		}
		//println(fmt.Sprintf("%v", planets))
	}
	return countOrbits(planets, "COM", 0), nil
}

func advent6B(test string) (int, error) {
	planets := make(map[string][]string)
	for _, s := range strings.Split(test, "\n") {
		entry := strings.Split(s, ")")
		name := strings.TrimSpace(entry[0])
		childName := strings.TrimSpace(entry[1])
		_, ok := planets[name]
		if !ok {
			planets[name] = []string{childName}
		} else {
			planets[name] = append(planets[name], childName)
		}
		//println(fmt.Sprintf("%v", planets))
	}
	tree := planetaryTree(planets, "COM", &planet{})
	you, distance := search(tree.orbitals, []string{}, "YOU", 0)
	println(you.name)
	println(distance)
	santa, distance := search(you.orbitals, []string{}, "SAN", 0)
	println(santa.name)
	println(distance)
	return distance - 1, nil
}

func search(p []*planet, searched []string, target string, d int) (*planet, int) {
	for i := 0; i < len(p); i++ {
		found := false
		for j := 0; j < len(searched); j++ {
			if searched[j] == p[i].name {
				found = true
				break
			}
		}
		if found {
			continue
		}
		if p[i].name == target {
			return p[i], d
		}
		searched = append(searched, p[i].name)
		check, dist := search(p[i].orbitals, searched, target, d+1)
		if check.name == target {
			return check, dist
		}
	}
	return &planet{}, 0
}

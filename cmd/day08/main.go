package main

import (
	"container/heap"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/cory-miller/advent-of-code-2025/internal/minheap"
)

type junction struct {
	x int
	y int
	z int
}

type pair struct {
	a junction
	b junction
}

type circuit struct {
	j []junction
}

func parseInput(content string) []junction {
	junctions := make([]junction, 0)
	for _, line := range strings.Split(content, "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		j := junction{x, y, z}
		junctions = append(junctions, j)
	}
	return junctions
}

func distance(a, b junction) float64 {
	x := math.Pow(float64(a.x-b.x), 2)
	y := math.Pow(float64(a.y-b.y), 2)
	z := math.Pow(float64(a.z-b.z), 2)

	return math.Sqrt(x + y + z)
}

func partOne(junctions []junction, connections int) int {
	// In part one we will make a min heap based on the distances between two junctions.
	// Then we will make the specified number of connections between the shortest distances.
	// Then we will inspect all of the circuits created and return the product of the 3 largest.

	h := minheap.New(func(a, b pair) bool {
		return distance(a.a, a.b) < distance(b.a, b.b)
	})

	for i := 0; i < len(junctions)-1; i++ {
		for j := i + 1; j < len(junctions); j++ {
			heap.Push(h, pair{junctions[i], junctions[j]})
		}
	}

	circuits := make(map[junction]*circuit)

	for i := 0; i < connections; i++ {
		if h.Len() == 0 {
			break
		}

		p := heap.Pop(h).(pair)

		ja, hasA := circuits[p.a]

		jb, hasB := circuits[p.b]

		// If these are the same circuit we can continue along. In this puzzle,
		// a reverse connection still counts.
		if hasA && hasB && ja == jb {
			continue
		}

		// If we've seen both of the junctions but they aren't the same circuit, we
		// need to merge the two circuits.
		if hasA && hasB && ja != jb {
			for _, j := range jb.j {
				circuits[j] = ja
			}
			ja.j = append(ja.j, jb.j...)
			continue
		}

		// Otherwise the remaining cases are adding one or both nodes to a circuit.
		// The circuit already exists in the second and third case below.
		if !hasA && !hasB {
			l := &circuit{[]junction{p.a, p.b}}
			circuits[p.a] = l
			circuits[p.b] = l
		} else if !hasA {
			jb.j = append(jb.j, p.a)
			circuits[p.a] = jb
			circuits[p.b] = jb
		} else {
			ja.j = append(ja.j, p.b)
			circuits[p.a] = ja
			circuits[p.b] = ja
		}
	}

	// We will track the three largest circuit sizes, plus the circuits we've seen.
	f := -1
	s := -1
	t := -1
	x := make(map[*circuit]bool)

	for _, circuit := range circuits {
		if _, has := x[circuit]; has {
			continue
		}

		if len(circuit.j) > f {
			t = s
			s = f
			f = len(circuit.j)
		} else if len(circuit.j) > s {
			t = s
			s = len(circuit.j)
		} else if len(circuit.j) > t {
			t = len(circuit.j)
		}

		x[circuit] = true
	}

	return f * s * t
}

func partTwo(junctions []junction) int {
	// In part two, we are tasked with creating connections until there is a single
	// completed circuit. To do this, we can check the size of the circuit after making two
	// connections. If the size of the circuit matches the number of junctions, we
	// can return the product of their X coordinates.

	// Because we are asked to do this with the last pair needed to complete the circuit
	// based on shortest distance, we will re-use the min heap.

	h := minheap.New(func(a, b pair) bool {
		return distance(a.a, a.b) < distance(b.a, b.b)
	})

	for i := 0; i < len(junctions)-1; i++ {
		for j := i + 1; j < len(junctions); j++ {
			heap.Push(h, pair{junctions[i], junctions[j]})
		}
	}

	circuits := make(map[junction]*circuit)

	ans := 0

	for h.Len() > 0 {
		p := heap.Pop(h).(pair)

		ja, hasA := circuits[p.a]

		jb, hasB := circuits[p.b]

		if hasA && hasB && ja == jb {
			continue
		}

		if hasA && hasB && ja != jb {
			for _, j := range jb.j {
				circuits[j] = ja
			}
			ja.j = append(ja.j, jb.j...)
			continue
		}

		if !hasA && !hasB {
			l := &circuit{[]junction{p.a, p.b}}
			circuits[p.a] = l
			circuits[p.b] = l
		} else if !hasA {
			jb.j = append(jb.j, p.a)
			circuits[p.a] = jb
			circuits[p.b] = jb
		} else {
			ja.j = append(ja.j, p.b)
			circuits[p.a] = ja
			circuits[p.b] = ja
		}

		// The circuit has been completed. Return the product.
		if len(circuits[p.a].j) == len(junctions) {
			return p.a.x * p.b.x
		}
	}

	return ans
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/08.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(b)
	input := parseInput(contents)

	fmt.Printf("Part 1: %d\n", partOne(input, 1000))
	fmt.Printf("Part 2: %v\n", partTwo(input))
}

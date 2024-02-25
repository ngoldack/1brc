package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}

	ms := process(os.Args[1])
	finish(ms)

}

type Measurement struct {
	Count int
	Sum   float64
	Min   float64
	Max   float64
}

func finish(ms map[string]*Measurement) {
	cs := make([]string, 0, len(ms))
	for c := range ms {
		cs = append(cs, c)
	}
	slices.Sort(cs)

	ratio := math.Pow(10, float64(1))
	for _, c := range cs {
		fmt.Printf("%s=%v/%v/%v\n", c, ms[c].Min, math.Round(ms[c].Sum/float64(ms[c].Count)*ratio)/ratio, ms[c].Max)
	}
}

func process(filename string) map[string]*Measurement {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ms := make(map[string]*Measurement)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		c, rv, ok := strings.Cut(t, ";")
		if !ok {
			log.Fatalf("Could not split: %s", t)
		}

		v, err := strconv.ParseFloat(rv, 64)
		if err != nil {
			log.Fatalf("could not parse rv=%s; t=%s; %v", rv, t, err)
		}

		if m, ok := ms[c]; ok {
			// City already exists
			m.Count++
			m.Sum += v

			if v < m.Min {
				m.Min = v
			}

			if v > m.Max {
				m.Max = v
			}

			continue
		}

		ms[c] = &Measurement{
			Count: 1,
			Sum:   v,
			Min:   v,
			Max:   v,
		}
	}

	return ms
}

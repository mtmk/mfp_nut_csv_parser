package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type nut struct {
	Carbs    float64
	Protein  float64
	Fat      float64
	Calories float64
}

func main() {
	csvFile, err := os.Open("mfp.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//var nuts []nut
	m := make(map[string]nut)
	head := true
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if head {
			head = false
			continue
		}
		//fmt.Printf("%s -- %s -- %s\n", line[0], line[1], line[2])
		calories, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		carbs, err := strconv.ParseFloat(line[12], 64)
		if err != nil {
			log.Fatal(err)
		}
		protein, err := strconv.ParseFloat(line[15], 64)
		if err != nil {
			log.Fatal(err)
		}
		fat, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			log.Fatal(err)
		}
		day := line[0]

		n, ok := m[day]
		if !ok {
			n = nut{}
			m[day] = n
		}

		// fmt.Println(calories)

		n.Calories += calories
		n.Calories = math.Round(n.Calories)
		n.Protein += protein
		n.Protein = math.Round(n.Protein)
		n.Fat += fat
		n.Fat = math.Round(n.Fat)
		n.Carbs += carbs
		n.Carbs = math.Round(n.Carbs)

		m[day] = n

		// nuts = append(nuts, nut{
		// 	Day:      line[0],
		// 	Calories: calories,
		// })
	}

	// jsonStr, err := json.Marshal(m)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("JSON:\n%s", jsonStr)

	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("day,calories,protein,carbs,fat\n")
	for _, k := range keys {
		n := m[k]
		fmt.Printf("%s,%.0f,%.0f,%.0f,%.0f\n", k, n.Calories, n.Protein, n.Carbs, n.Fat)
	}
}

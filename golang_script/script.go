package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
)

var (
	mapp = make(map[int]string)
	wg   = sync.WaitGroup{}
)

type Body struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

func main() {

	min_id := 0
	step := 100
	max_id := step

	for {
		for i := min_id; i < max_id; i++ {
			wg.Add(1)
			go request(i)
		}

		wg.Wait()

		min_id += step
		max_id += step
	}
}

func request(id int) {
	defer wg.Done()
	url := fmt.Sprintf("http://localhost:8080/fragment?id=%d", id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	var body Body
	json.NewDecoder(resp.Body).Decode(&body)

	saveText(body)
}

func saveText(body Body) {
	mapp[body.Index] = body.Text
	validate()
}

func validate() {
	if len(mapp) >= 28 {
		showResult()
		os.Exit(0)
	}
}

func showResult() {
	var puzzle []string
	for _, value := range mapp {
		puzzle = append(puzzle, value)
	}
	// Sort the puzzle based on keys
	keys := make([]int, 0, len(mapp))
	for k := range mapp {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	sortedPuzzle := make([]string, len(puzzle))
	for i, k := range keys {
		sortedPuzzle[i] = mapp[k]
	}

	fmt.Println(strings.Join(sortedPuzzle, " "))
}

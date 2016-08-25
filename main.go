package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	input := flag.String("input", "", "Text file with review summaries")
	flag.Parse()

	if *input == "" {
		log.Fatal("Please specify the input file")
		return
	}

	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rankers := setupRankers()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			r, err := NewReviewFromString(line)
			if err != nil {
				log.Fatal(err)
			}
			if r.Source != SourceMonkey {
				for i := range rankers {
					rankers[i].Rank(r)
				}
			}
			fmt.Println(r)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func setupRankers() []Ranker {
	var rankers = []Ranker{}
	rankers = append(rankers, &LotsToSayRanker{})
	rankers = append(rankers, &BurstRanker{})
	rankers = append(rankers, &SameDeviceRanker{})
	rankers = append(rankers, &AllStarRanker{})
	rankers = append(rankers, &SolicitedRanker{})
	return rankers
}

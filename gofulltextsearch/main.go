package main

import (
	"fmt"

	"github.com/kotaroooo0/stalefish"
)

func main() {
	// documents := []stalefish.Document{
	// 	{
	// 		ID: 0,
	// 		Fields: map[string]string{
	// 			"name":    "Ichiro Suzuki",
	// 			"content": "I have a lot of TASKs. I am very sad :(",
	// 		},
	// 	},
	// 	{
	// 		ID: 1,
	// 		Fields: map[string]string{
	// 			"name":    "Taro Tanaka",
	// 			"content": "I have no tasks. I am HAPPY :)",
	// 		},
	// 	},
	// }
	analyzer := stalefish.Analyzer{
		[]stalefish.CharFilter{stalefish.MappingCharFilter{map[string]string{":)": "_happy_", ":(": "_sad_"}}},
		stalefish.StandardTokenizer{},
		[]stalefish.TokenFilter{stalefish.LowercaseFilter{}, stalefish.StopWordFilter{}, stalefish.StemmerFilter{}},
	}

	documents := []stalefish.Document{
		{
			ID: 0,
			Fields: map[string]string{
				"content": "fine fax fox fun :)",
			},
		},
		{
			ID: 1,
			Fields: map[string]string{
				"content": "fine fax fox hot :)",
			},
		},
		{
			ID: 2,
			Fields: map[string]string{
				"content": "fine fax hot pie :)",
			},
		},
	}

	// Indexing
	index := stalefish.Index{}
	index.Indexing(documents, analyzer)

	// Search
	searcher := stalefish.Searcher{index, analyzer}
	ids := searcher.Search(stalefish.Query{
		"fine FaX foxes happy",
		[]string{"content"},
	})
	fmt.Println(ids)
}

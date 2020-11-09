package main

import (
	"github.com/k0kubun/pp"
	"github.com/kotaroooo0/stalefish"
)

func main() {
	documents := []stalefish.Document{
		{
			ID: 0,
			Fields: map[string]string{
				"first_name": "tanaka",
				"last_name":  "taro",
				"content":    "fine fax fox fun",
			},
		},
		{
			ID: 1,
			Fields: map[string]string{
				"first_name": "tanaka",
				"last_name":  "ziro",
				"content":    "fine fax fox hot",
			},
		},
		{
			ID: 2,
			Fields: map[string]string{
				"first_name": "sato",
				"last_name":  "taro",
				"content":    "fine fax hot pie",
			},
		},

		{
			ID: 3,
			Fields: map[string]string{
				"first_name": "sato",
				"last_name":  "ichiro",
				"content":    "fine best snow",
			},
		},
	}
	analyzer := stalefish.Analyzer{
		stalefish.StandardTokenizer{},
		[]stalefish.CharFilter{},
		[]stalefish.Filter{stalefish.LowercaseFilter{}},
	}

	index := stalefish.Index{}
	index.Indexing(documents, analyzer)
	pp.Println(index)

	searcher := stalefish.Searcher{index, analyzer}
	ids := searcher.Search(stalefish.Query{
		"fine fax fox",
		[]string{"content"},
	})
	pp.Println(ids)

}

package main

import (
	"fmt"

	"github.com/kotaroooo0/stalefish"
)

func main() {
	// create index
	config := stalefish.NewDBConfig("root", "password", "127.0.0.1", "3306", "stalefish")
	db, _ := stalefish.NewDBClient(config) // omit error handling
	storage := stalefish.NewStorageRdbImpl(db)
	analyzer := stalefish.NewAnalyzer([]stalefish.CharFilter{}, stalefish.StandardTokenizer{}, []stalefish.TokenFilter{stalefish.StemmerFilter{}, stalefish.LowercaseFilter{}, stalefish.StopWordFilter{}})
	indexer := stalefish.NewIndexer(storage, analyzer, make(stalefish.InvertedIndexMap))

	indexer.AddDocument(stalefish.NewDocument("You can watch lots of interesting dramas on Amazon Prime."))
	indexer.AddDocument(stalefish.NewDocument("Forest phenomena in the Amazon are a prime concern."))
	indexer.AddDocument(stalefish.NewDocument("I watched amazon prime until late at night yesterday."))
	indexer.AddDocument(stalefish.NewDocument("Breaking Bad is a very jarring drama."))

	// search documents
	q := stalefish.NewPhraseQuery("amAzon PRime", analyzer) // Uppercase and lowercase notation fluctuation
	seacher := q.Searcher(storage)
	result, _ := seacher.Search() // omit error handling
	fmt.Println(result)
	// result: [{1 You can watch lots of interesting dramas on Amazon Prime.} {3 I watched amazon prime until late at night yesterday.}]

	q = stalefish.NewPhraseQuery("drama", analyzer) // Singular and plural notation fluctuations
	seacher = q.Searcher(storage)
	result, _ = seacher.Search() // omit error handling
	fmt.Println(result)
	// result: [{1 You can watch lots of interesting dramas on Amazon Prime.} {4 Breaking Bad is a very jarring drama.}]
}

package main

import (
	"fmt"

	"github.com/kotaroooo0/stalefish"
)

func main() {
	db, _ := stalefish.NewDBClient(stalefish.NewDBConfig("root", "password", "127.0.0.1", "3306", "stalefish"))
	storage := stalefish.NewStorageRdbImpl(db)
	analyzer := stalefish.NewAnalyzer([]stalefish.CharFilter{}, stalefish.NewStandardTokenizer(), []stalefish.TokenFilter{stalefish.NewLowercaseFilter()})

	indexer := stalefish.NewIndexer(storage, analyzer, 1)
	for _, body := range []string{"Ruby PHP JS", "Go Ruby", "Ruby Go PHP", "Go PHP"} {
		indexer.AddDocument(stalefish.NewDocument(body))
	}

	sorter := stalefish.NewTfIdfSorter(storage)
	mq := stalefish.NewMatchQuery("GO Ruby", stalefish.OR, analyzer, sorter)
	mseacher := mq.Searcher(storage)
	result, _ := mseacher.Search()
	fmt.Println(result) // [{2 Go Ruby 2} {3 Ruby Go PHP 3} {4 Go PHP 2} {1 Ruby PHP JS 3}]

	pq := stalefish.NewPhraseQuery("go RUBY", analyzer, nil)
	pseacher := pq.Searcher(storage)
	result, _ = pseacher.Search()
	fmt.Println(result) // [{2 Go Ruby 2}
}

// func main() {
// 	analyzer := stalefish.NewAnalyzer(
// 		[]stalefish.CharFilter{stalefish.NewMappingCharFilter(map[string]string{":(": "sad"})},
// 		stalefish.NewStandardTokenizer(),
// 		[]stalefish.TokenFilter{stalefish.NewLowercaseFilter(), stalefish.NewStemmerFilter(), stalefish.NewStopWordFilter([]string{"i", "my", "me", "the", "a", "for"})},
// 	)
// 	fmt.Println(analyzer.Analyze("I feel TIRED :(")) // {[{0 feel } {0 tire } {0 sad }]}
// }

// func main() {
// 	kagome, _ := morphology.NewKagome()
// 	analyzer := stalefish.NewAnalyzer(
// 		[]stalefish.CharFilter{},
// 		stalefish.NewMorphologicalTokenizer(kagome),
// 		[]stalefish.TokenFilter{stalefish.NewReadingformFilter(stalefish.Romaji)},
// 	)
// 	fmt.Println(analyzer.Analyze("東京と京都"))
// }

// func main() {
// 	db, _ := stalefish.NewDBClient(stalefish.NewDBConfig("root", "password", "127.0.0.1", "3306", "stalefish"))
// 	storage := stalefish.NewStorageRdbImpl(db)
// 	analyzer := stalefish.NewAnalyzer(
// 		[]stalefish.CharFilter{},
// 		stalefish.NewStandardTokenizer(),
// 		[]stalefish.TokenFilter{},
// 	)

// 	indexer := stalefish.NewIndexer(storage, analyzer, make(stalefish.InvertedIndex))
// 	indexer.AddDocument(stalefish.NewDocument("go ruby javascript"))
// 	indexer.AddDocument(stalefish.NewDocument("go python typescript"))
// 	indexer.AddDocument(stalefish.NewDocument("go perl flutter"))
// }

// func main() {
// mq := stalefish.NewMatchQuery("interesting night", stalefish.AND, analyzer)
// msearcher := mq.Searcher(storage)
// result, _ = msearcher.Search()
// fmt.Println(result)

// kagome, err := morphology.NewKagome()
// if err != nil {
// 	panic(err)
// }
// jpanalyzer := stalefish.NewAnalyzer([]stalefish.CharFilter{}, stalefish.NewMorphologicalTokenizer(kagome), []stalefish.TokenFilter{stalefish.NewReadingformFilter(stalefish.Romaji)})
// jpindexer := stalefish.NewIndexer(storage, jpanalyzer, make(stalefish.InvertedIndex))

// err = jpindexer.AddDocument(stalefish.NewDocument("俺は東京生まれ東京育ち"))
// if err != nil {
// 	panic(err)
// }
// err = jpindexer.AddDocument(stalefish.NewDocument("京都は鴨川がとても良い"))
// if err != nil {
// 	panic(err)
// }
// err = jpindexer.AddDocument(stalefish.NewDocument("東京と京都、どっちに住みたい？"))
// if err != nil {
// 	panic(err)
// }
// err = jpindexer.AddDocument(stalefish.NewDocument("北海道が一番良い"))
// if err != nil {
// 	panic(err)
// }

// jmq := stalefish.NewMatchQuery("kyoto", stalefish.OR, jpanalyzer)
// jmsearcher := jmq.Searcher(storage)
// jresult, _ := jmsearcher.Search()
// fmt.Println(jresult)
// }

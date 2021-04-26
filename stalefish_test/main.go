package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/kotaroooo0/gojaconv/jaconv"
	"github.com/kotaroooo0/stalefish"
)

// func main() {
// 	db, _ := stalefish.NewDBClient(stalefish.NewDBConfig("root", "password", "127.0.0.1", "3306", "stalefish"))
// 	storage := stalefish.NewStorageRdbImpl(db)
// 	analyzer := stalefish.NewAnalyzer(
// 		[]stalefish.CharFilter{}, stalefish.NewStandardTokenizer(), []stalefish.TokenFilter{stalefish.NewStemmerFilter(), stalefish.NewLowercaseFilter(), stalefish.NewStopWordFilter()},
// 	)

// 	indexer := stalefish.NewIndexer(storage, analyzer)
// 	indexer.AddDocument(stalefish.NewDocument("You can watch lots of interesting dramas on Amazon Prime."))
// 	indexer.AddDocument(stalefish.NewDocument("Forest phenomena in the Amazon are a prime concern."))

// 	pq := stalefish.NewPhraseQuery("amAzon PRime", *analyzer)
// 	pseacher := pq.Searcher(storage)
// 	result, _ := pseacher.Search()
// 	fmt.Println(result)
// 	// result: [{1 You can watch lots of interesting dramas on Amazon Prime.}]

// 	mq := stalefish.NewMatchQuery("amazon concerns", stalefish.AND, *analyzer)
// 	mseacher := mq.Searcher(storage)
// 	result, _ = mseacher.Search()
// 	fmt.Println(result)
// 	// result: [{2 Forest phenomena in the Amazon are a prime concern.}]
// }

func main() {

	hebon := jaconv.ToHebon("おはよう")
	fmt.Println(hebon) // ohayo

	db, _ := stalefish.NewDBClient(stalefish.NewDBConfig("root", "password", "127.0.0.1", "3306", "stalefish"))
	storage := stalefish.NewStorageRdbImpl(db)
	analyzer := stalefish.NewAnalyzer([]stalefish.CharFilter{}, stalefish.NewStandardTokenizer(), []stalefish.TokenFilter{stalefish.NewLowercaseFilter()})

	indexer := stalefish.NewIndexer(storage, analyzer)
	indexer.AddDocument(stalefish.NewDocument("Go Ruby PHP"))
	indexer.AddDocument(stalefish.NewDocument("Go PHP Python"))
	indexer.AddDocument(stalefish.NewDocument("Go Python Ruby"))

	pp.Println(storage.GetInvertedIndexByTokenIDs([]stalefish.TokenID{stalefish.TokenID(0), stalefish.TokenID(1)}))

	mq := stalefish.NewMatchQuery("GO PHP", stalefish.AND, analyzer)
	mseacher := mq.Searcher(storage)
	result, _ := mseacher.Search()
	fmt.Println(result) // result: [{1 Go Ruby PHP} {2 Go PHP Python}]

	pq := stalefish.NewPhraseQuery("GO PHP", analyzer)
	pseacher := pq.Searcher(storage)
	result, _ = pseacher.Search()
	fmt.Println(result) // result: [{2 Go PHP Python}]

	// stalefish.InvertedIndex{
	// 	0x0000000000000001: stalefish.PostingList{
	// 	  Postings: &stalefish.Postings{
	// 		DocumentID: 0x0000000000000001,
	// 		Positions:  []uint64{
	// 		  0x0000000000000000,
	// 		},
	// 		Next: &stalefish.Postings{
	// 		  DocumentID: 0x0000000000000002,
	// 		  Positions:  []uint64{
	// 			0x0000000000000000,
	// 		  },
	// 		  Next: &stalefish.Postings{
	// 			DocumentID: 0x0000000000000003,
	// 			Positions:  []uint64{
	// 			  0x0000000000000000,
	// 			},
	// 			Next: (*stalefish.Postings)(nil),
	// 		  },
	// 		},
	// 	  },
	// 	},
	//   }
}

// func main() {
// 	analyzer := stalefish.NewAnalyzer(
// 		[]stalefish.CharFilter{stalefish.NewMappingCharFilter(map[string]string{":(": "sad"})},
// 		stalefish.NewStandardTokenizer(),
// 		[]stalefish.TokenFilter{stalefish.NewLowercaseFilter(), stalefish.NewStemmerFilter(), stalefish.NewStopWordFilter()},
// 	)
// 	fmt.Println(analyzer.Analyze("I feel TIRED :(")) // output: &{[{0 feel } {0 tire } {0 sad }]}
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

package main

import (
	"log"

	"github.com/olivere/elastic"
)

func main() {
	// fmt.Println("---------- elastic.NewClient() ----------")
	_, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Println(err)
	}

	// fmt.Println("---------- elastic.NewClient(elastic.SetSniff(false)) ----------")
	// _, err := elastic.NewClient(elastic.SetSniff(false))
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(`---------- http.Get("http://127.0.0.1:9200") ----------`)

	// res, err := http.Get("http://127.0.0.1:9200/?pretty=true")
	// if err != nil {
	// 	log.Println(err)
	// }
	// dr, _ := httputil.DumpResponse(res, true)
	// fmt.Println(string(dr))
}

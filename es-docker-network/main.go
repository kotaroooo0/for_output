package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/olivere/elastic"
)

func main() {
	// fmt.Println("---------- elastic.NewClient() ----------")
	// _, err := elastic.NewClient()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("---------- elastic.NewClient(elastic.SetSniff(false)) ----------")
	_, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(`---------- http.Get("http://127.0.0.1:9200") ----------`)

	res, err := http.Get("http://127.0.0.1:9200/_nodes/http?pretty=true")
	if err != nil {
		fmt.Println(err)
	}
	dr, _ := httputil.DumpResponse(res, true)
	fmt.Println(string(dr))
}

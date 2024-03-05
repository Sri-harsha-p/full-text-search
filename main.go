package main

import (
	"flag"
	utils "github.com/Sri-harsha-p/full-text-search.git/utils"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()
	log.Println("full text seach is in progress")
	start := time.Now()
	docs, err := utils.LoadDocument(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded %d document in %v ", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("index %d document in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("search found %d document in %v ", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}

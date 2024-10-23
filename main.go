package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/jaywantadh/text-SE/utils" //replace all local path with your own mod file.

)

func main(){
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "---your search database----", "---dump name----")
	flag.StringVar(&query, "q", "----query search-----", "search query")
	flag.Parse()
	
	log.Println("Full Text Search is in progress...")
	
	start := time.Now()
	
	docs, err := utils.Loaddocuments(dumpPath)
	if err != nil{
		log.Fatal(err)
	} 

	log.Printf("loaded %d documens in %v", len(docs), time.Since(start))
	
	start = time.Now() 
	idx	:= make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs{

		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)	
	}

}

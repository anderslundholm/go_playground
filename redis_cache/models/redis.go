package models

import (
	"log"
	"strings"
	"time"

	"github.com/mediocregopher/radix.v2/pool"
)

var db *pool.Pool

func init() {
	var err error
	db, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Panic(err)
	}
}

func GetAutocomplete(qtype string, query string, maxResults int) ([]string, error) {
	start := time.Now()

	elapsed := time.Since(start)
	log.Printf("Autocomplete took %s", elapsed)

	query = strings.ToLower(query)
	zset := "ix:" + qtype
	var step int = 50
	var autoCompletedWords []string
	var rangeChunk []string
	queryCheck := query
	start, err := db.Cmd("zrank", zset, query).Int()
	if err != nil {
		return nil, err
	}
	for queryCheck[0:len(query)] == query && len(autoCompletedWords) < maxResults {
		rangeChunk, err = db.Cmd("zrange", zset, start, start+step-1).List()
		start += step
		for _, word := range rangeChunk {
			queryCheck = word
			if word[len(word)-1:] == "*" {
				autoCompletedWords = append(autoCompletedWords, word[:len(word)-1])
			}
		}
	}
	return autoCompletedWords, nil
}

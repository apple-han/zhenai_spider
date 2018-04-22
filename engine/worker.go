package engine

import "learn/crawler/fetch"

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetch.Fetch(r.Url)
	if err != nil {
		//log.Print("Fetcher: error "+
		//	"fetching url %s: %v",
		//	r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}


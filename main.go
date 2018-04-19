package main

import (
	"learn/crawler/engine"
	"learn/crawler/zhenai/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})
}

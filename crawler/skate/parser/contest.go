package parser

import (
	"file/learngo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
const contestListRe  = `open\([0-9&#;]*[a-zA-Z0-9\.?]*\=([0-9]*)[0-9&#;]*\)`

func ParseContestList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(contestListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		println(string(m[1]))
		//result.Requests = append(result.Requests, engine.Request{
		//	Url:    string(m[1]),
		//	Parser: nil, //engine.NewFuncParser(ParseCity, "ParseCity"),
		//})
	}
	return result
}

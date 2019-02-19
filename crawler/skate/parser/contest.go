package parser

import (
	"file/learngo/crawler/engine"
	"file/learngo/crawler/model"
	"regexp"
	"strconv"
)

const contestListRe = `open\([0-9&#;]*[a-zA-Z0-9\.?]*\=([0-9]*)[0-9&#;]*\)`
const contestNameRe = `>(.*)</span>.*value\=\"`
const contestTypeRe = `.*>(.*赛)</td>`
const contestStationRe = `.*>(.*站)</td>`
const contestContestTimeRe = `.*>([\d]*\/[\d]*\/[\d]*—[\d]*\/[\d]*\/[\d]*)</td>`

func ParseContestList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(contestListRe)
	matches := re.FindAllSubmatch(contents, -1)
	var Ids []string
	for _, m := range matches {
		flag := true
		for _, value := range Ids {
			if value == string(m[1]) {
				flag = false
				break
			}
		}
		if flag {
			Ids = append(Ids, string(m[1]))
		}

	}

	result := engine.ParseResult{}
	for _, m := range Ids {
		result.Requests = append(result.Requests, engine.Request{
			Url:    `http://www.chinashorttrack.com/jscs/bsrc1.aspx?id=` + m,
			Parser: engine.NewFuncParser(ParseMatchList, "ParseMatchList"),
			//Parser: engine.NilParser{},
		})
		id, _ := strconv.Atoi(m)
		item := model.Contest{
			Id:          id,
			Name:        extractString(contents, regexp.MustCompile(contestNameRe+m)),
			Type:        extractString(contents, regexp.MustCompile(m+contestTypeRe)),
			Station:     extractString(contents, regexp.MustCompile(m+contestStationRe)),
			ContestTime: extractString(contents, regexp.MustCompile(m+contestContestTimeRe)),
		}
		result.Items = append(result.Items, engine.Item{
			Url:     `http://www.chinashorttrack.com/jscs/bsrc1.aspx?id=` + m,
			Type:    `contest`,
			Id:      m,
			Payload: item,
		})

	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

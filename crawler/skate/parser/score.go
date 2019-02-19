package parser

import (
	"file/learngo/crawler/engine"
	"file/learngo/crawler/model"
	"regexp"
)

const matchIdRe = `rcbh=([\d]*)`
const groupScoreRe = `<p[\s\S]*?</table>`
const groupRe = `<span class="tag bg-dot">(.*)</span>`
const scoreRe = `<td align="center" style="width:10%;">(.*)</td><td align="center" style="width:10%;">(.*)</td><td align="center" style="width:10%;">(.*)</td><td align="center" style="width:10%;">(.*)</td><td align="center" style="width:30%;">(.*)</td><td align="center" style="width:10%;">(.*)</td><td align="center" style="width:10%;">(.*)</td><td class="hidden-l hidden-m hidden-s hidden-b" style="border-color:#507CD1;">(.*)</td>`

func ParseScore(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(groupScoreRe)
	matches := re.FindAllSubmatch(contents, -1)
	matchId := extractString(contents, regexp.MustCompile(matchIdRe))
	result := engine.ParseResult{}
	for _, v := range matches {
		groupNum := extractString(v[0], regexp.MustCompile(groupRe))
		re2 := regexp.MustCompile(scoreRe)
		scores := re2.FindAllSubmatch(v[0], -1)
		for _, vv := range scores {
			score := model.Score{
				Group:     string(groupNum),
				No:        string(vv[1]),
				RowNum:    string(vv[2]),
				HeadNum:   string(vv[3]),
				Name:      string(vv[4]),
				Organize:  string(vv[5]),
				TimeScore: string(vv[6]),
				Remark:    string(vv[7]),
				MatchId:   string(matchId),
			}
			result.Items = append(result.Items, engine.Item{
				Url:     ``,
				Type:    `score`,
				Id:      string(matchId),
				Payload: score,
			})
		}
	}
	return result
}

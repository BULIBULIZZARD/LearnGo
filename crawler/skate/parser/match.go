package parser

import (
	"file/learngo/crawler/engine"
	"file/learngo/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

const MatchDecUrlRe = `ot=(.*)&amp;md=(.*)&amp;rcbh=(.*)&amp;id=(.*)&#39;,`
const MatchDecRe = `&amp;id=[\d]*&#39;,&#39;menu&#39;\)">.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdsj_[\d]*" width="10%"> ` +
	`<span id="ContentPlaceHolder1_Repeater1_Repeater2_\d_Label1_[\d]*">(.*)</span></td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdxm_[\d]*" width="15%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdzbbh_[\d]*" width="10%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdxb_[\d]*" width="10%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdsb_[\d]*" width="10%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdrs_[\d]*" width="10%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdzs_[\d]*" width="10%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdlqgz_[\d]*" width="10%">(.*)</td>.*\n` +
	`	<td id="ContentPlaceHolder1_Repeater1_Repeater2_\d_tdbz_[\d]*" width="10%">(.*)</td>`
const ContestIdRE = `./bsrc1.aspx\?id=([\d]*)`

func ParseMatchList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(MatchDecUrlRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	contestId := extractString(contents, regexp.MustCompile(ContestIdRE))
	cid, _ := strconv.Atoi(contestId)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    `http://www.chinashorttrack.com/jscs/jsfz.aspx?ot=` + string(m[1]) + `&md=` + string(m[2]) + `&rcbh=` + string(m[3]) + `&id=` + string(m[4]),
			Parser: engine.NewFuncParser(ParseScore, "ParseScore"),
		})
		data := GetMatchData(contents, string(m[3])+MatchDecRe)
		id, _ := strconv.Atoi(string(m[3]))
		matchData := model.Match{
			Id:        id,
			Time:      data[2],
			MatchName: data[3],
			Gender:    data[4],
			MatchType: data[5],
			PlayerNum: data[6],
			GroupNum:  data[7],
			Enter:     data[8],
			Remark:    data[9],
			ContestId: cid,
		}
		result.Items = append(result.Items, engine.Item{
			Url:     `http://www.chinashorttrack.com/jscs/jsfz.aspx?ot=` + string(m[1]) + `&md=` + string(m[2]) + `&rcbh=` + string(m[3]) + `&id=` + string(m[4]),
			Type:    `match`,
			Id:      string(m[4]),
			Payload: matchData,
		})
	}
	return result
}
func GetMatchData(contents []byte, rule string) []string {
	re := regexp.MustCompile(rule)
	data := re.FindSubmatch(contents)
	result := []string{""}
	for _, v := range data {
		result = append(result, strings.Replace(string(v), " ", "", -1))
	}
	return result
}

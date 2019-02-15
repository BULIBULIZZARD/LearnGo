package controller

import (
	"file/learngo/crawler/frontend/view"
	"github.com/olivere/elastic"
	"net/http"
	"strings"
	"strconv"
	"file/learngo/crawler/frontend/model"
	"context"
	"reflect"
	"file/learngo/crawler/engine"
	"regexp"
)

type SearchResultHandel struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandel{
	client,err:=elastic.NewClient(elastic.SetURL("http://192.168.99.100:9200/"),elastic.SetSniff(false))
	if err!= nil{
		panic(err)
	}
	return SearchResultHandel{
		view:view.CreateSearchResultView(template),
		client:client,
	}
}

func (h SearchResultHandel) ServeHTTP(w http.ResponseWriter,req *http.Request){
	q:=strings.TrimSpace(req.FormValue("q"))
	from,err:=strconv.Atoi(
		req.FormValue("from"))
	if err != nil{
		from = 0
	}
	var page model.SearchResult
	page ,err = h.getSearchResult(q,from)
	if err != nil{
		http.Error(w,string(err.Error()),http.StatusBadRequest)
	}
	err = h.view.Render(w,page)
	if err != nil{
		http.Error(w,string(err.Error()),http.StatusBadRequest)
	}
}
func (h SearchResultHandel) getSearchResult(q string,from int) (model.SearchResult,error){
	var result model.SearchResult
	result.Query = q
	resp,err :=h.client.
		Search("dating_show").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil{
		return result,err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result,err
}
func rewriteQueryString(q string) string{
	re :=regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q,"Payload.$1:")
}
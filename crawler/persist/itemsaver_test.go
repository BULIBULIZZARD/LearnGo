package persist

import (
	"testing"
	"file/learngo/crawler/model"
	"github.com/olivere/elastic"
	"context"
	"file/learngo/crawler/engine"
	"encoding/json"
)

func TestSave(t *testing.T) {

	profile := model.Profile{
		Age:34,
		Height:162,
		Weight:57,
		Income:"3001-5000元",
		Gender:"女",
		Name:"安静的雪",
		Xinzuo:"哈哈座",
		Occupation:"人事/行政",
		Marriage:"离异",
		House:"已购房",
		Hokou:"山东菏泽",
		Education:"大学本科",
		Car:"未购车",
	}
	expected := engine.Item{
		Url:"http://album.zhenai.com/u/1680817021",
		Type:"zhenai",
		Id:"1680817021",
		Payload:profile,
	}
	err :=save(expected)
	if err != nil{
		panic(err)
	}
	client,err:=elastic.NewClient(elastic.SetURL("http://192.168.99.100:9200/"),elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	resp,err :=client.Get().Index("dating_profile").
		Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil{
		panic(err)
	}
	t.Logf("%s" ,*resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source,&actual)
	if err != nil{
		panic(err)
	}
	if actual!=expected{
		t.Errorf("get %v; expected %v",actual,expected)
	}

}
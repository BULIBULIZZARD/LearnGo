package persist

import (
	"file/learngo/crawler/engine"
	"github.com/olivere/elastic"
	"file/learngo/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item,result *string) error{
	err :=persist.Save(s.Client,s.Index,item)
	log.Printf("Item %v saved",item)
	if err == nil{
		*result = "ok"
	}else{
		log.Printf("Error saving item %v: %v",item,err)
	}
	return err
}
package persist

import (
	"context"
	"database/sql"
	"file/learngo/crawler/engine"
	"file/learngo/crawler/model"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
)

import _ "github.com/go-sql-driver/mysql"

const DB_Driver = "root:fushihao@tcp(127.0.0.1:3306)/skate?charset=utf8"

func ItemSaver(index string) (chan engine.Item, error) {
	//client, err := elastic.NewClient(elastic.SetURL("http://192.168.99.100:9200/"), elastic.SetSniff(false))
	//if err != nil {
	//	return nil, err
	//}
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		panic(err)
	}
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			log.Printf("Item Saver:got item :%v", item)
			err := MysqlSaver(db, item)
			if err != nil {
				log.Print(err.Error())
			}
		}
	}()
	return out, nil
}
func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type ")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())

	return err
}

func MysqlSaver(db *sql.DB, item engine.Item) error {
	err := errors.New(``)
	switch data := item.Payload.(type) {
	case model.Contest:
		err = saveContest(db, data)
		break
	case model.Match:
		err = saveMatch(db, data)
		break
	case model.Score:
		err = saveScore(db, data)
		break
	default:
		break
	}
	return err
}
func saveContest(db *sql.DB, contest model.Contest) error {
	stmt, err := db.Prepare(`insert contest set id=?,name=?,contest_time=?,station=?,type=?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contest.Id, contest.Name, contest.ContestTime, contest.Station, contest.Type)
	return err
}
func saveMatch(db *sql.DB, match model.Match) error {
	stmt, err := db.Prepare(`insert s_match set id=?,time=?,match_name=?,gender=?,match_type=?,player_num=?,group_num=?,enter=?,remark=?,contest_id=?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(match.Id, match.Time, match.MatchName, match.Gender, match.MatchType, match.PlayerNum, match.GroupNum, match.Enter, match.Remark, match.ContestId)
	return err
}
func saveScore(db *sql.DB, score model.Score) error {
	stmt, err := db.Prepare(`insert score set s_group=?,no=?,row_num=?,head_num=?,name=?,organize=?,time_score=?,remark=?,match_id=?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(score.Group, score.No, score.RowNum, score.HeadNum, score.Name, score.Organize, score.TimeScore, score.Remark, score.MatchId)
	return err
}

package services

import (
	"cron/gocron"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	DATE_FORMAT = "2006-01-02 15:04"
)

//精度支持到分钟
type Task struct {
	//2012-06-12 12:22
	Time string `json:"time"`
	//2012-06-12 12:22
	EndTime string `json:"endTime"`

	Every uint64 `json:"every"`
	//Minute, Hour, Day, Week, Month, Year
	Unit string `json:"unit"`

	LastRun time.Time `json:"lastRun"`
}

type DbTask struct {
	Key   string
	Value Task
}

var db *leveldb.DB

func init() {
	_db, err := leveldb.OpenFile("data/db", nil)

	if err != nil {
		panic(err)
	}
	db = _db

	s := gocron.GetScheduler()
	s.Start()
	initJob()
}

func initJob() {
	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		key := iter.Key()
		task := TaskFromJson(iter.Value())
		task.Run(string(key[:]))
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}

}

func Tasks() []*DbTask {
	iter := db.NewIterator(nil, nil)
	list := []*DbTask{}
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		dbTask := new(DbTask)
		dbTask.Key = string(key[:])
		dbTask.Value = TaskFromJson(value)
		list = append(list, dbTask)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func DeleteScheduler(key string) {
	db.Delete([]byte(key), nil)
	s := gocron.GetScheduler()
	s.Remove(key)
}

func (sch *Task) Save(md5 string) {
	db.Put([]byte(md5), []byte(sch.toJson()), nil)
}

func (sch *Task) Run(md5 string) error {
	fmt.Println("加入任务-->" + md5)
	t, _ := time.ParseInLocation(DATE_FORMAT, sch.Time, time.Local)
	j := gocron.NewJob(md5, sch.Every, sch.Unit, t)
	j.Do(task, md5)

	s := gocron.GetScheduler()
	s.Add(j)

	return nil
}

func (sch *Task) toJson() string {
	js, err := json.Marshal(sch)
	if err != nil {
		log.Fatal(err)
	}

	return string(js)
}

// func (sch *Scheduler) getTime() (time.Time, error) {
// 	return time.Parse(time.RFC3339, sch.AtDate+"T"+sch.AtTime+"+00:00")
// }

func task(j *gocron.Job, md5 string) {
	fmt.Println("Run task : " + md5)
	value, _ := db.Get([]byte(md5), nil)
	var task Task
	_ = json.Unmarshal(value, &task)
	task.Time = j.NextRun().Format(DATE_FORMAT)
	task.LastRun = j.LastRun()
	task.Save(md5)

	fmt.Println("Value:" + task.toJson())
}

func TaskFromJson(value []byte) (sch Task) {
	json.Unmarshal(value, &sch)
	return
}

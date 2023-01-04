package process

import (
	"database/sql"
	"errors"
	"example/bookAPI/internal/models/job"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

type Queue struct {
	DB   *gorm.DB
	Lock *sync.Mutex
	Queuer
}

type callback func(q *Queue)

type Queuer interface {
	Spawn(group *sync.WaitGroup, callback callback)
}

func GetDB() (*gorm.DB, error) {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: db,
		}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return gormDB, err
}

func Process(process callback) error {
	gormDB, err := GetDB()
	if err != nil {
		return errors.New("failed to establish db connection")
	} else {
		fmt.Println("connected to database")
		var wg = &sync.WaitGroup{}
		fmt.Printf("starting queue process \n")
		q := new(Queue)
		q.DB = gormDB
		var m sync.Mutex
		q.Lock = &m
		wg.Add(1)
		go q.Spawn(wg, process)
		time.Sleep(2 * time.Second)
		q.DB.Create(&job.Job{Name: "Test Job", Command: "pkg/queue/jobs/test/test.go", Args: `{"payload": "something"}`})
		time.Sleep(8 * time.Second)
		q.DB.Create(&job.Job{Name: "Test Job 2", Command: "pkg/queue/jobs/test/test.go", Args: `{"payload": "something 2"}`})
		defer wg.Wait()
	}
	return nil
}

func (q *Queue) Spawn(wg *sync.WaitGroup, callback callback) {
	defer wg.Done()
	for true {
		time.Sleep(100 * time.Millisecond)
		go func() {
			q.Lock.Lock()
			callback(q)
			q.Lock.Unlock()
		}()
	}
}

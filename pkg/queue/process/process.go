package main

import (
	"database/sql"
	"example/bookAPI/internal/models/job"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type Queue struct {
	db *gorm.DB
}

func main() {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: db,
		}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected to database")
		var wg = &sync.WaitGroup{}
		fmt.Printf("starting queue process \n")
		q := new(Queue)
		q.db = gormDB
		wg.Add(1)
		go q.spawn(wg)
		time.Sleep(2 * time.Second)
		q.db.Create(&job.Job{Name: "New job"})
		time.Sleep(8 * time.Second)
		q.db.Create(&job.Job{Name: "New NEW job"})
		defer wg.Wait()
	}
}

func (q *Queue) spawn(wg *sync.WaitGroup) {
	defer wg.Done()
	loopFunc := func() string {
		return "hello"
	}
	fmt.Println(loopFunc)
	//for true {
	//	time.Sleep(100 * time.Millisecond)
	//	var newJob job.Job
	//	query := q.db.First(&newJob)
	//	if query.Error == nil {
	//		if query.RowsAffected > 0 {
	//			fmt.Println(newJob.Name)
	//			q.db.Delete(&newJob)
	//		}
	//	}
	//}
}

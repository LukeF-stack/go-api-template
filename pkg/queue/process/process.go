package main

import (
	"database/sql"
	"example/bookAPI/internal/models/job"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"os/exec"
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
		}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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
		q.db.Create(&job.Job{Name: "Test Job", Command: "pkg/queue/jobs/test/test.go", Args: `{"payload": "something"}`})
		time.Sleep(8 * time.Second)
		q.db.Create(&job.Job{Name: "Test Job 2", Command: "pkg/queue/jobs/test/test.go", Args: `{"payload": "something 2"}`})
		defer wg.Wait()
	}
}

func (q *Queue) spawn(wg *sync.WaitGroup) {
	defer wg.Done()
	for true {
		time.Sleep(100 * time.Millisecond)
		var newJob job.Job
		query := q.db.First(&newJob)
		if query.Error == nil {
			if query.RowsAffected > 0 {
				fmt.Println(newJob.Name)
				cmd := exec.Command("go", "run", newJob.Command, newJob.Args)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()

				if err != nil {
					fmt.Println(err)
				} else {
					q.db.Delete(&newJob)
				}
			}
		}
	}
}

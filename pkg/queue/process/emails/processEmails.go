package main

import (
	"example/bookAPI/internal/models/job"
	process2 "example/bookAPI/pkg/queue/process"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	process2.Process(process)
}

func process(q *process2.Queue) {
	var newJob job.Job
	query := q.DB.First(&newJob)
	if query.Error == nil {
		if query.RowsAffected > 0 {
			fmt.Println("processing job: " + newJob.Name)
			if fileExists(newJob.Command) {
				fmt.Println("executing command: " + newJob.Command)
				cmd := exec.Command("go", "run", newJob.Command, newJob.Args)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("job successfully completed")
					q.DB.Delete(&newJob)
				}
			} else {
				q.DB.Delete(&newJob)
			}
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

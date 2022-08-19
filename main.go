package main

import (
	"example/bookAPI/internal/database"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Run main!")
	awaitConn := make(chan bool)
	connection := new(database.Connection)
	go connection.Init(awaitConn)
	fmt.Println("Main: Waiting for db connection to finish")
	<-awaitConn
	fmt.Println("Main: Completed")
	cmd := exec.Command("go", "run", "pkg/migration.go")
	cmd.Stdout = os.Stdout
	out := cmd.Run()
	if out != nil {
		fmt.Println(out)
	}
}

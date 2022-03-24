package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/robsantossilva/gophercises-cli_task_manager/cmd"
	"github.com/robsantossilva/gophercises-cli_task_manager/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	//db.CreateTask("Lavar a louça")

	tasks, _ := db.AllTasks()

	fmt.Println(tasks[0].Value)

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

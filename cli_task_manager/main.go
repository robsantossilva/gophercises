package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robsantossilva/gophercises-cli_task_manager/cmd"
	"github.com/robsantossilva/gophercises-cli_task_manager/db"
)

func main() {
	pwd, err := os.Getwd()
	must(err)

	// home, _ := homedir.Dir()
	// fmt.Println(home)

	dbPath := filepath.Join(pwd, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

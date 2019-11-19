package main

import (
	"errors"
	"fmt"
	"github.com/robfig/cron"
	"github.com/zhshch2002/gitup/config"
	"os"
	"os/exec"
)

func UpdateRepo(r config.Repo) {
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = r.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	cmd = exec.Command("git", "reset", "--hard", r.Branch)
	cmd.Dir = r.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	c := cron.New()
	for _, R := range config.Repos {
		r := R
		if r.Mode == "ontime" {
			err := c.AddFunc(r.Time, func() {
				UpdateRepo(r)
			})
			if err != nil {
				panic(errors.New("Add OnTime Job " + fmt.Sprint(r) + " " + err.Error()))
			}
		}
	}
	c.Run()
	select {}
}

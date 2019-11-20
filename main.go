package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/zhshch2002/gitup/config"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func FetchRepo(r config.Repo) {
	log.Println("Fetch repo", r.Dir, r.Branch)

	log.Println(r.Dir, "-", "git", "fetch", "--all")
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = r.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	log.Println(r.Dir, "-", "git", "reset", "--hard", r.Branch)
	cmd = exec.Command("git", "reset", "--hard", r.Branch)
	cmd.Dir = r.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
func md5str(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func main() {
	c := cron.New()

	WebHooks := map[string]config.Repo{}

	for _, R := range config.Repos {
		r := R
		if r.Mode == "ontime" {
			log.Println("Set OnTime job", r.Dir)
			err := c.AddFunc(r.Time, func() {
				FetchRepo(r)
			})
			if err != nil {
				panic(errors.New("Add OnTime Job " + fmt.Sprint(r) + " " + err.Error()))
			}
		} else if r.Mode == "webhook" {
			k := md5str(r.Dir)
			WebHooks[k] = r
			log.Println(r.Dir + " - /a/" + k)
		}
	}
	go c.Run()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Any("/", func(ctx *gin.Context) {
		for k, r := range WebHooks {
			ctx.String(http.StatusOK, r.Dir+" - /a/"+k)
		}
	})
	router.Any("/a/:id", func(ctx *gin.Context) {
		r, ok := WebHooks[ctx.Param("id")]
		if !ok {
			return
		}
		log.Println("on webhook", r.Dir)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
				}
			}()
			FetchRepo(r)
		}()
	})
	go func() {
		log.Println("Listen on", config.ListenAddr)
		err := router.Run(config.ListenAddr)
		if err != nil {
			panic(err)
		}
	}()
	select {}
}

package main

import (
	"fmt"
	"testing"
	"time"

	cron "github.com/robfig/cron/v3"
)

func TestCron1(t *testing.T) {
	c := cron.New()

	c.AddFunc("0/1 * * * *", func() {
		fmt.Println("tick every 1 minute")
	})

	c.Start()
	time.Sleep(time.Hour * 5)
}

func TestCron2(t *testing.T) {
	// 支持秒级的表达式
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0/5 * * * * *", func() {
		fmt.Println("tick every 5 seconde")
	})

	c.Start()
	time.Sleep(time.Hour * 5)
}

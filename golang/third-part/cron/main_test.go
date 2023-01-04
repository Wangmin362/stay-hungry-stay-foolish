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

// 时区设置
func TestCron3(t *testing.T) {
	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.AddFunc("0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at New York")
	})

	c.AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at Tokyo")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type configBean struct {
	Test string
}

type config struct {
	LastModify time.Time
	Data       configBean
}

var Config *config

func loadConfig(path string) error {
	var locker = new(sync.RWMutex)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if Config != nil && fileInfo.ModTime().Before(Config.LastModify) {
		return errors.New("no need update")
	}

	var configBean configBean
	err = json.Unmarshal(data, &configBean)
	if err != nil {
		return err
	}

	config := &config{
		LastModify: fileInfo.ModTime(),
		Data:       configBean,
	}

	locker.Lock()
	Config = config
	locker.Unlock()

	return nil

}

func main() {
	fmt.Println("start main process")

	configPath := "/tmp/env.json"
	done := make(chan bool, 1)

	_ = loadConfig(configPath)
	fmt.Printf("current config value is: %s \n", Config.Data.Test)

	// 定义信号通道
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGUSR1)

	go func(path string) {
		for {
			select {
			case <-sig:
				// 收到信号, 加载配置文件
				err := loadConfig(path)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("")
				fmt.Println("received signal!")
				fmt.Printf("current config value is: %s \n", Config.Data.Test)
			}
		}
	}(configPath)

	// 挂起进程，直到获取到一个信号
	<-done
}

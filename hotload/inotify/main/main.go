/* ----------------------------------------------------------------------------------------
 *
 * Desc:	inotify 文件监控
 * Author:  liangqi000@gmail.com
 * Date:    2019-06-07
 *
 * ----------------------------------------------------------------------------------------
 */
package main

import (
	"fmt"
	"github.com/apptut/go-labs/hotload/inotify/watcher"
	"log"
	"syscall"
)


func main() {

	path := "/tmp/env.json"
	notify, err := watcher.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = notify.AddWatcher(path, syscall.IN_CLOSE_WRITE)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case event := <-notify.Events:
				if event & syscall.IN_CLOSE_WRITE == syscall.IN_CLOSE_WRITE {
					fmt.Printf(" file changed \n")

					// 加载配置文件函数
					// loadConfig(path)
				}
			}
		}
	}()

	<- done

}


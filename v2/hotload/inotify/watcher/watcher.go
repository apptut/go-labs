/* ----------------------------------------------------------------------------------------
 *
 * Desc:    文件监听封装
 * Author:  liangqi000@gmail.com
 * Date:    2019-06-08
 *
 * ----------------------------------------------------------------------------------------
 */

package watcher

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

type Watcher struct {
	Events chan uint32
	fd     int
}

func NewWatcher() (*Watcher, error) {
	fd, err := unix.InotifyInit()
	if err != nil {
		return nil, err
	}

	watcher := &Watcher{
		Events: make(chan uint32),
		fd: fd,
	}
	watcher.getEvents()

	return watcher, nil
}

func (w *Watcher) AddWatcher(file string, mask uint32) (error) {
	_, err := unix.InotifyAddWatch(w.fd, file, mask)
	if err != nil {
		return err
	}

	return nil
}

func (w *Watcher) getEvents() {
	go func() {
		var buf [unix.SizeofInotifyEvent * 4096]byte

		for {
			n, err := unix.Read(w.fd, buf[:])
			if err != nil {
				n = 0
				continue
			}

			var offset uint32
			for offset <= uint32(n-unix.SizeofInotifyEvent) {
				raw := (*unix.InotifyEvent)(unsafe.Pointer(&buf[offset]))

				mask := uint32(raw.Mask)
				nameLen := uint32(raw.Len)

				// 塞到事件队列
				w.Events <- mask
				offset += unix.SizeofInotifyEvent + nameLen
			}
		}
	}()
}

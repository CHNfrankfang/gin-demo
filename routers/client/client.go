package client

import (
	"gin/logs"
	"os"
	"time"
)

type Status uint8

const (
	Wait     Status = iota // TIME—WAIT状态
	Initing                // 初始化中
	Inited                 // 初始化成功
	Running                // 运行中
	Stopping               // 已发出停止命令，但是还没有停止
	Stoped                 // 停止
	Failed
)

type ScanClient struct {
	Uptime int64
	Name   string
	Status Status
	Path   string
	Files  []string
	done   chan struct{}
}

func (c *ScanClient) Start(path string) bool {
	logs.SL.Debug("start")
	if c.Status == Running {
		return true
	}
	_, err := os.ReadDir(path)
	if err != nil {
		c.Status = Failed
		return false
	}
	c.Path = path
	c.Status = Running
	c.Name = "scan"
	c.done = make(chan struct{})
	c.Files = make([]string, 0)
	c.Uptime = time.Now().Unix()

	t := time.NewTicker(3 * time.Second)
	go func() {
		for {
			select {
			case <-c.done:
				c.Status = Stoped
				c.Files = nil
				return
			case <-t.C:
				files, err := os.ReadDir(c.Path)
				if err != nil {
					c.Status = Failed
					c.Files = nil
					c.done <- struct{}{}
				}
				if len(c.Files) > 10 {
					for i := 0; i < len(c.Files); i++ {
						c.Files[i] = files[i%len(files)].Name()
					}
				} else {
					// 如果 c.Files 不足 10 个，直接追加文件名
					for _, v := range files {
						c.Files = append(c.Files, v.Name())
					}
				}

			}
		}
	}()
	return true
}

func (c *ScanClient) Stop() bool {
	c.done <- struct{}{}
	return true
}

func (c *ScanClient) Info() ScanClient {
	return *c
}

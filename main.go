package main

import (
	"sync"
	"time"
)

func main() {
	smsSender := NewSMSSender("0000000000")
	emailSender := NewEmailSender("cjconnorashton@gmail.com")

	wpoller := NewWPoller(smsSender, emailSender)
	interval := time.Second * 5

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		wpoller.start(interval)
	}()

	time.Sleep(time.Second * 10)
	wpoller.close()

	wg.Wait()
}

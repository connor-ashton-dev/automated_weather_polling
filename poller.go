package main

import (
	"fmt"
	"log"
	"time"
)

type WPoller struct {
	closeCh chan struct{}
	senders []Sender
}

func NewWPoller(senders ...Sender) *WPoller {
	return &WPoller{
		closeCh: make(chan struct{}),
		senders: senders,
	}
}

func (wp *WPoller) close() {
	close(wp.closeCh)
}

func (wp *WPoller) start(interval time.Duration) {
	fmt.Println("starting the wpoller")
	ticker := time.NewTicker(interval)

poller:
	for {
		select {
		case <-ticker.C:
			data, err := getWeatherResults(52.52, 13.41)
			if err != nil {
				log.Fatal(err)
			}
			if err := wp.handleData(data); err != nil {
				log.Fatal(err)
			}
		case <-wp.closeCh:
			// handle graceful shutdown
			break poller
		}
	}
	fmt.Println("wpoller stopped gracefully")
}

func (wp *WPoller) handleData(data *WeatherData) error {
	// handle the data
	for _, s := range wp.senders {
		go func(s Sender, d *WeatherData) {
			if err := s.Send(d); err != nil {
				fmt.Println(err)
			}
		}(s, data) // Pass data as an argument here
	}
	return nil
}

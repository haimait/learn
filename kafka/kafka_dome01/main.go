package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	addr  = "182.92.234.23:9092"
	topic = "web_log"
)

func main() {
	go Receive()
	go Publish()
	time.Sleep(time.Second * 100000)
}

func Receive() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	fmt.Println("receive connecting ...")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{addr},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		//MaxWait:   time.Hour,
	})
	//r.SetOffset(42)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	for {
		fmt.Println("receive start ...")
		m, err := r.ReadMessage(ctx)
		fmt.Println("receive one ...")
		if err != nil {
			fmt.Println("readmessave error", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
	fmt.Println("receive close")
	r.Close()
}

func Publish() {
	fmt.Println("publish connecting ...")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{addr},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	fmt.Println("publish start ...")
	for {
		message := "msg" + time.Now().String()
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte(message),
			},
		)

		if err != nil {
			print("publish err:", err)
			return
		} else {
			fmt.Println("publish sucess")
		}

		time.Sleep(time.Second)
	}
	fmt.Println("public close")
	w.Close()
}
package stream

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	// to produce messages
	topic := "aaaa"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "47.243.62.135:9092", string(topic), partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one3!")},
		kafka.Message{Value: []byte("two3!")},
		kafka.Message{Value: []byte("three3!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func TestRcv(t *testing.T) {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"47.243.62.135:9092"},
		GroupID:  "consumer-group-id",
		Topic:    string("aaaa"),
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
func TestRead(t *testing.T) {
	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	batchSize := int(10e6) // 10MB

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     string("aaa"),
		Partition: 0,
		MinBytes:  batchSize,
		MaxBytes:  batchSize,
	})

	r.SetOffsetAt(context.Background(), startTime)

	for {
		m, err := r.FetchMessage(context.Background())

		if err != nil {
			break
		}
		if m.Time.After(endTime) {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\\n", m.Offset, string(m.Key), string(m.Value))
		fmt.Println("------------")
		err = r.CommitMessages(context.Background(), m)
		if err != nil {
			println(err)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

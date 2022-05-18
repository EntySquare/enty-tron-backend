package stream

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

import "entysquare/enty-tron-backend/pkg/stream/message"

var leader = "47.243.62.135:9092"

// kafka client
func GenClient() (c *Client, err error) {
	return &Client{
		RWMutex: &sync.RWMutex{},
		readers: make(map[Topic]*Processor),
		writers: make(map[Topic]*kafka.Conn),
	}, nil
}

type Client struct {
	*sync.RWMutex
	readers map[Topic]*Processor
	writers map[Topic]*kafka.Conn
}

type Processor struct {
	reader *kafka.Reader
	handle func(m Message) error
}

type MessageInterface interface {
	Marshal() (dAtA []byte, err error)
}

type Message struct {
	Topic Topic
	// Payload is inferred into correspond message struct in protobuf
	Payload MessageInterface
}

func (c *Client) Send(streamMsg Message, retry ...int64) error {
	// retry 3 times by default
	var loop int64 = 3
	if retry != nil && retry[0] >= 0 {
		loop = retry[0]
	}
	topic := streamMsg.Topic
	conn, err := c.getWriter(topic)
	if err != nil {
		return err
	}
	bytes, err := streamMsg.Payload.Marshal()
	if err != nil {
		return err
	}

	_, err = conn.WriteMessages(kafka.Message{Value: bytes})
	if err != nil {
		if judgePipeBroken(err) {
			// reset client for this topic
			_, err = c.getWriter(topic, true)
		}
		if loop == 0 {
			return err
		}
		loop--
		return c.Send(streamMsg, loop)
	}
	fmt.Println("entypay messaging send: ", streamMsg)
	return nil
}

func judgePipeBroken(err error) bool {
	if strings.Contains(err.Error(), "broken pipe") {
		return true
	}
	return false
}

func (c *Client) Register(topic Topic, handle func(m Message) error) error {
	reader, err := c.getReader(topic)
	if err != nil {
		return err
	}
	c.Lock()
	c.readers[topic] = &Processor{
		reader: reader,
		handle: handle,
	}
	c.Unlock()
	return nil
}

// Process execute process functions in topic map according to subscribed messages
func (c *Client) Process() error {

	ctx, cancel := context.WithCancel(context.TODO())

	for topic, processor := range c.readers {
		go func(shut context.CancelFunc, topicIn Topic, processorIn *Processor) {
			err := processorIn.processTopic(ctx, topicIn)
			if err != nil {
				shut()
			}
		}(cancel, topic, processor)
	}

	<-ctx.Done()
	// process went stopped status for accident
	log.Error("oms shut done for some reason")
	return nil
}

func BuildMessageFromSpec(topic Topic, rawMsg MessageInterface) Message {
	m := Message{
		Topic: topic,
	}
	switch topic {
	case ORDER_RESP:
		this := (rawMsg).(*message.OrderResponse)
		m.Payload = this
	case DEPOSIT_ORDER:
		this := (rawMsg).(*message.DepositOrder)
		m.Payload = this
	case TRANSFER_ORDER:
		this := (rawMsg).(*message.TransferOrder)
		m.Payload = this
	case MAKE_ORDER:
		this := (rawMsg).(*message.GeneratePayOrder)
		m.Payload = this
	case PAY_ORDER:
		this := (rawMsg).(*message.PayOrder)
		m.Payload = this
	case WITHDRAW_ORDER:
		this := (rawMsg).(*message.WithdrawOrder)
		m.Payload = this
	case TRANSACTION:
		this := (rawMsg).(*message.Transaction)
		m.Payload = this
	case TRANSACTION_RESP:
		this := (rawMsg).(*message.TransactionResponse)
		m.Payload = this
	case TRANSACTION_UPDATE:
		this := (rawMsg).(*message.TransactionUpdate)
		m.Payload = this
	case WITHDRAW_COLLECT:
		this := (rawMsg).(*message.WithdrawCollectRequest)
		m.Payload = this
	case RISK_REQUEST:
		this := (rawMsg).(*message.RiskRequest)
		m.Payload = this
	case RISK_RESPONSE:
		this := (rawMsg).(*message.RiskResponse)
		m.Payload = this
	case COLLECT_CHECK:
		this := (rawMsg).(*message.CollectCheck)
		m.Payload = this
	case CHAIN_PANIC:
		this := (rawMsg).(*message.ChainHandlerPanicRequest)
		m.Payload = this
	default:
		panic("reflection went wrong")
	}
	return m
}

func BuildMessageFromRaw(topic Topic, rawMsg *kafka.Message) Message {
	var err error
	m := Message{
		Topic: topic,
	}
	switch topic {
	case ORDER_RESP:
		this := &message.OrderResponse{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case DEPOSIT_ORDER:
		this := &message.DepositOrder{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case TRANSFER_ORDER:
		this := &message.TransferOrder{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case MAKE_ORDER:
		this := &message.GeneratePayOrder{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case PAY_ORDER:
		this := &message.PayOrder{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case WITHDRAW_ORDER:
		this := &message.WithdrawOrder{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case TRANSACTION:
		this := &message.Transaction{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case TRANSACTION_RESP:
		this := &message.TransactionResponse{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case TRANSACTION_UPDATE:
		this := &message.TransactionUpdate{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case WITHDRAW_COLLECT:
		this := &message.WithdrawCollectRequest{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case RISK_REQUEST:
		this := &message.RiskRequest{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case RISK_RESPONSE:
		this := &message.RiskResponse{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case COLLECT_CHECK:
		this := &message.CollectCheck{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	case CHAIN_PANIC:
		this := &message.ChainHandlerPanicRequest{}
		err = proto.Unmarshal(rawMsg.Value, this)
		m.Payload = this
	default:
		panic("reflection went wrong")
	}
	if err != nil {
		panic("reflection went wrong")
	}
	return m
}

func (p *Processor) processTopic(ctx context.Context, topic Topic) error {
	for {
		m, err := p.reader.FetchMessage(ctx)
		if err != nil {
			if judgePipeBroken(err) {
				r := kafka.NewReader(kafka.ReaderConfig{
					Brokers:  []string{leader},
					GroupID:  "oms",
					Topic:    string(topic),
					MinBytes: 1,
					MaxBytes: 10e6,
				})
				p.reader = r
				continue
			}
			log.Error(err)
			return err
		}
		fmt.Println("entypay streaming received message: ", topic)
		msg := BuildMessageFromRaw(topic, &m)

		err = p.handle(msg)
		if err != nil {
			fmt.Println("panic error occurred :" + err.Error())
			panic("handle topic error :" + topic)
		}
		err = p.reader.CommitMessages(ctx, m)
		if err != nil {
			panic("commit error")
		}

	}
}

func (c *Client) getReader(topic Topic) (r *kafka.Reader, err error) {
	c.Lock()
	processor, ok := c.readers[topic]
	c.Unlock()
	if !ok {
		r = kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{leader},
			GroupID:  "oms",
			Topic:    string(topic),
			MinBytes: 1,
			MaxBytes: 10e6,
		})
		return r, nil
	}
	return processor.reader, nil
}

// pass reset flag in and reset correspond topic conn
func (c *Client) getWriter(topic Topic, reset ...bool) (conn *kafka.Conn, err error) {
	// reset writer
	if reset != nil {
		c.Lock()
		conn, err = kafka.DialLeader(context.Background(), "tcp", leader, string(topic), 0)
		if err != nil {
			return nil, err
		}
		c.writers[topic] = conn
		c.Unlock()
	}
	// if there is a conn return it, it may cause broken pipe error but covered at send layer
	c.RLock()
	conn, ok := c.writers[topic]
	c.RUnlock()
	if !ok {
		conn, err = kafka.DialLeader(context.Background(), "tcp", leader, string(topic), 0)
		if err != nil {
			return nil, err
		}
		c.Lock()
		c.writers[topic] = conn
		c.Unlock()
		return conn, nil
	}
	return conn, nil
}

func (c *Client) RetryLoop(retryCount int, streamMsg Message) error {
	writeMessageOk := false
	//retry when write messages has failed
writeDepositOrderMessageRetryLoop:
	for i := 0; i < retryCount; i++ {
		err := c.Send(streamMsg)
		if err != nil {
			print("err:", err)
			return err
		}
		if err == nil {
			writeMessageOk = true
			break writeDepositOrderMessageRetryLoop
		}
	}
	if !writeMessageOk {
		return errors.New("retry failed")
	}
	return nil
}

package RMQWrapper

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func CreateProducerInstance() (rocketmq.Producer, error) {

	p, err := rocketmq.NewProducer(
		producer.WithGroupName("new2TestGroup"),
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)

	return p, err

}

//SendMsg to send a message to a queue
func SendMsg(Pmsg string, p rocketmq.Producer) {

	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	topic := "SelfTest3P"
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(Pmsg),
	}
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	time.Sleep(time.Second * 30)

}

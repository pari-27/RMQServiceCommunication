package RMQWrapper

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func CreateConsumerInstance() (rocketmq.PushConsumer, error) {

	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("new2TestGroup"),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.BroadCasting),
		consumer.WithInstance("firstCreatedInstance"),
	)

	return c, err

}

//ReceiveMsg to get messages from the queue
func ReceiveMsg(rMsg chan string, c rocketmq.PushConsumer) {
	var message []*primitive.MessageExt

	err := c.Subscribe("SelfTest3P", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		message = msgs

		for m := range message {
			fmt.Printf("received msg: %v \n", message[m])

			msg := fmt.Sprintf("Topic: %s|||| Body:%s |||| MsgID:%s |||| QueueId:%s", message[m].Topic, string(message[m].Body), message[m].MsgId, strconv.Itoa(message[m].Queue.QueueId))

			rMsg <- msg
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println("helllloooooo" + err.Error())
	}

	// fmt.Println("subscription done")
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	time.Sleep(time.Hour)

	err = c.Shutdown()
	if err != nil {
		fmt.Printf("Shutdown Consumer error: %s", err.Error())
	}

	close(rMsg)

}

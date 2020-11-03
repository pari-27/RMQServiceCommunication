package service

import (
	rmq "RMQServiceCommunication/RMQWrapper"
	"fmt"
	"time"

	logger "github.com/sirupsen/logrus"
)

// @Title computeListOfData
// @Description process the list through message queue
// @Accept  String
func ComputeListOfData(deps Dependencies) {
	c, _ := rmq.CreateConsumerInstance()
	var msg = make(chan string, 100)
	go rmq.ReceiveMsg(msg, c)
	for {
		v, ok := <-msg
		if ok == false {
			time.Sleep(time.Minute)
			continue
		}
		fmt.Println("Received in compute service", v, ok)
		err := deps.Store.StoreComputedUsers(v)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in compute loop users")
			return
		}

	}
}

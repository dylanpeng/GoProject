package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	topics = "testgo"
	wg sync.WaitGroup
)

func main(){
	go oneConsumer()
	go anotherConsumer()
	exit := make(chan os.Signal)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	fmt.Printf("stop by exit signal '%s'", sign)

}

func oneConsumer(){
	groupID := "group-1"
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	//创建消费者
	c, err := cluster.NewConsumer(strings.Split("localhost:9092", ","), groupID, strings.Split(topics, ","), config)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	defer c.Close()

	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				fmt.Println(err)
			case <-noti:
			}
		}
	}(c)

	for msg := range c.Messages(){
		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

func anotherConsumer(){
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions(topics)

	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topics, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}
}

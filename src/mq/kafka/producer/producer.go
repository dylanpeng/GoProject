package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"strings"
)

func main(){
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil{
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:"testgo",
		Partition:int32(-1),
		Key:sarama.StringEncoder("key"),
	}

	var value string
	for {
		inputReader := bufio.NewReader(os.Stdin)
		value, err = inputReader.ReadString('\n')

		if err != nil{
			panic(err)
		}

		value = strings.Replace(value, "\n", "", -1)
		msg.Value = sarama.ByteEncoder(value)
		partition, offset, err := producer.SendMessage(msg)

		if err != nil{
			fmt.Println("Send Message Fail!")
		}

		fmt.Printf("Partion = %d, offset = %d\n", partition, offset)

	}

}

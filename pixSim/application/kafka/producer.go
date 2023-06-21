package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":"kafka:9092",
	}
	
	producer, err := ckafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}
	return producer
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic, 
			Partition: ckafka.PartitionAny,},
		Value: []byte(msg),
	}
	
	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch event := e.(type) {
		case *ckafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery Failed! ", event.TopicPartition)
			} else {
				fmt.Println("Delivery Succesfull! ", event.TopicPartition)
			}
		}
	}
}
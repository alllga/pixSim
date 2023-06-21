package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
	Database *gorm.DB
	Producer *ckafka.Producer
	DeliveryChan chan ckafka.Event

}


func NewKafkaProcessor (database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database: database,
		Producer: producer,
		DeliveryChan: deliveryChan,
	}
}

func (kProc *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":"kafka:9092",
		"group.id":"consumergroup",
		"auto.offset.reset":"earliest",
	}

	consumer, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")
	
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}

}
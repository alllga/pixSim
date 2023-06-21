package kafka

import (
	"fmt"

	"github.com/alllga/pixSim/application/factory"
	appmodel "github.com/alllga/pixSim/application/model"
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
			kProc.processMessage(msg)
		}
	}

}

func (kProc *KafkaProcessor) processMessage(msg *ckafka.Message) {
	transactionTopic := "transations"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *&msg.TopicPartition.Topic; topic {
	case transactionTopic:
		kProc.processTransaction(msg)
	case transactionConfirmationTopic:
	default:
		fmt.Println("not a valid topic!", string(msg.Value))
	}
}

func (kProc *KafkaProcessor) processTransaction(msg *ckafka.Message) error {
	transation := appmodel.NewTransaction()
	err := transation.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	transationUseCase := factory.TransactionUseCaseFactory(kProc.Database)

	createdTransaction, err := transationUseCase.Register(
		transation.AccountID,
		transation.Amount,
		transation.PixKeyTo,
		transation.PixKeyKindTo,
		transation.Description,
	)
	if err != nil {
		fmt.Println("error registering transaction")
		return err
	}

	topic := "bank"+createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID

}
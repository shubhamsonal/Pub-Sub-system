package factory

import (
	"strconv"

	"github.com/jaswanth05rongali/pub-sub/logger"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//GetMessage  returns the message format according to publishing type
func GetMessage(kafkaPubMessageType string, key string, pubPartition string, kafkaTopic string, value string) kafka.Message {
	factoryLogger := logger.Getlogger()
	var message kafka.Message
	switch kafkaPubMessageType {
	case "1":
		key := key
		message = kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(value),
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		}

	default:
		var part int32
		if kafkaPubMessageType == "0" {
			part = kafka.PartitionAny
		} else {
			par, er := strconv.Atoi(pubPartition)
			part = int32(par)
			if er != nil {
				factoryLogger.Errorf("error while converting partitionToPublish to int, exiting...")
			}
		}
		message = kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: part},
			Value:          []byte(value),
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		}
	}
	return message
}

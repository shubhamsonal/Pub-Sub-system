package client

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jaswanth05rongali/pub-sub/logger"
	"github.com/spf13/viper"
)

//Interface defines an interface for client Object
type Interface interface {
	Init()
	SendMessage(string) bool
	RetrySendingMessage(string) bool
	SaveToFile(string) error
}

//Object creates and holds the entire client package
type Object struct {
}

var (
	serverStatus   bool
	prevServerTime int64
	serverRunTime  int64
	serverDownTime int64
	waitTime       int64
)

var clientLogger logger.Logger

//Init will initialize the client variables
func (c Object) Init() {
	clientLogger = logger.Getlogger()
	serverStatus = true
	prevServerTime = time.Now().Unix()
	serverRunTime = viper.GetInt64("serverRunTime")
	serverDownTime = viper.GetInt64("serverDownTime")
	waitTime = viper.GetInt64("waitTime")
}

func (c Object) getServerStatus() bool {
	currentTime := time.Now().Unix()
	if serverStatus {
		if currentTime > (prevServerTime + serverRunTime) {
			serverStatus = !serverStatus
			prevServerTime = currentTime
		}
	} else {
		if currentTime > (prevServerTime + serverDownTime) {
			serverStatus = !serverStatus
			prevServerTime = currentTime
		}
	}

	return serverStatus
}

func (c Object) getMessageDetails(value string) (string, string, string, string, string) {
	dataStrings := strings.Split(strings.Split(strings.Split(value, "{")[1], "}")[0], ",")
	requestString := strings.Split(dataStrings[0], ":")[1]
	requestID := requestString[1 : len(requestString)-1]
	topicString := strings.Split(dataStrings[1], ":")[1]
	topicName := topicString[1 : len(topicString)-1]
	messageString := strings.Split(dataStrings[2], ":")[1]
	messageBody := messageString[1 : len(messageString)-1]
	emailString := strings.Split(dataStrings[4], ":")[1]
	emailID := emailString[1 : len(emailString)-1]
	phoneString := strings.Split(dataStrings[5], ":")[1]
	phoneNumber := phoneString[1 : len(phoneString)-1]

	return requestID, topicName, messageBody, emailID, phoneNumber
}

//SendMessage will check for the client. If it is up and running then the func sends the message to the client and returns a true value. If client is down, it returns a false.
func (c Object) SendMessage(message string) bool {
	requestID, topicName, messageBody, emailID, phoneNumber := c.getMessageDetails(message)

	var customerDetail string
	if topicName == "Phone" {
		customerDetail = phoneNumber
	} else {
		customerDetail = emailID
	}

	if c.getServerStatus() {
		fmt.Printf("Message: '%v' sent successfully to %v. Request ID: %v\n", messageBody, customerDetail, requestID)
		clientLogger.Infof("Message: %v sent successfully to %v. Request ID: %v", messageBody, customerDetail, requestID)
		return true
	}

	fmt.Printf("Message: '%v' delivery to %v failed. Server Down!! Request ID: %v\n", messageBody, customerDetail, requestID)
	clientLogger.Infof("Message: %v delivery to %v failed. Server Down!! Request ID: %v", messageBody, customerDetail, requestID)
	return false
}

//RetrySendingMessage will try resending the messages
func (c Object) RetrySendingMessage(message string) bool {
	numberOfRetries := viper.GetInt("numberOfRetries")
	for i := 0; i < numberOfRetries; i++ {
		time.Sleep(time.Duration(waitTime) * time.Second)
		fmt.Printf("Retry - %v:\n", i+1)
		clientLogger.Infof("Retry - %v:", i+1)
		sent := c.SendMessage(message)
		if sent {
			return true
		}
	}

	return false
}

//SaveToFile will save a discarded message to a file
func (c Object) SaveToFile(message string) error {
	requestID, _, _, _, _ := c.getMessageDetails(message)

	f, err := os.OpenFile("failed.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if _, err := f.Write([]byte(message)); err != nil {
		log.Fatal(err)
		return err
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("The Delivery of message with Request ID:%v, to the client has been failed. Storing it in a file - failed.log\n", requestID)
	clientLogger.Infof("The Delivery of message with Request ID:%v, to the client has been failed. Storing it in a file - failed.log", requestID)
	return nil
}

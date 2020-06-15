package client

import (
	"log"
	"testing"
	"time"

	"github.com/jaswanth05rongali/pub-sub/logger"
	"github.com/spf13/viper"
)

var cli Object

func TestSendMessage(t *testing.T) {
	FileLocation := "./testlogs/client_test.log"
	err := logger.NewLogger(FileLocation)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	cli.Init()
	if cli.getServerStatus() == cli.SendMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}") {
		t.Logf("Send Message working fine.\n")
	} else {
		t.Errorf("Expected true from SendMessage but not successful.\n")
	}

	time.Sleep(time.Duration(viper.GetInt("serverRunTime")+5) * time.Second)

	if cli.getServerStatus() == cli.SendMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}") {
		t.Logf("Send Message working fine.\n")
	} else {
		t.Errorf("Expected false from SendMessage but not successful.\n")
	}
}

func TestRetrySendingMessage(t *testing.T) {
	viper.Set("numberOfRetries", 1)
	cli.Init()
	for i := 0; i < 2; i++ {
		if cli.getServerStatus() {
			if cli.RetrySendingMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}") {
				t.Logf("Retry Sending Message working fine.\n")
			} else {

				t.Errorf("Expected success from RetrySendingMessage but failed.\n")
			}
		} else {
			if cli.RetrySendingMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}") {
				t.Errorf("Expected failure from RetrySendingMessage but succeded.\n")
			} else {
				t.Logf("Retry Sending Message working fine.\n")
			}
		}
		time.Sleep(time.Duration(viper.GetInt("serverRunTime")+5) * time.Second)
	}
}

func TestSaveToFile(t *testing.T) {
	err := cli.SaveToFile("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}")
	if err != nil {
		t.Errorf("Expected Success but got error:%v.\n", err)
	}
}

package worker

import (
	"fmt"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswanth05rongali/pub-sub/logger"
	"github.com/jaswanth05rongali/pub-sub/mocks"
)

func TestInit(t *testing.T) {
	FileLocation := "./testlogs/worker_test.log"
	err := logger.NewLogger(FileLocation)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	var c *ConsumerObject
	broker := "localhost:19092"
	group := "testGroup"
	c.Init(broker, group)
	switch c.GetConsumer().String() {
	case "rdkafka#consumer-1":
		if broker == "localhost:19092" {
			t.Logf("Init Working Fine")
		} else {
			t.Errorf("Expected Init to return a nil value but got a consumer instance")
		}
	default:
		if broker == "localhost:19092" {
			t.Errorf("Init should create a legitimate Consumer. But not happened...")
		} else {
			t.Logf("Init Working Fine")
		}
	}
}

func TestConsume(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockInterface := mocks.NewMockInterface(mockCtrl)
	testConsumer := &ConsumerObject{ClientInterface: mockInterface}

	// testConsumer.Init("localhost:19092", "testGroup")

	mockInterface.EXPECT().Init().Times(1)
	// mockInterface.EXPECT().SendMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}").Return(true).AnyTimes()
	mockInterface.EXPECT().SendMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}").Return(false).AnyTimes()
	mockInterface.EXPECT().RetrySendingMessage("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}").Return(false).AnyTimes()
	mockInterface.EXPECT().SaveToFile("{\"request_id\":\"1\",\"topic_name\":\"foo\",\"message_body\":\"Transaction Successful\",\"transaction_id\":\"987456321\",\"email\":\"kafka@gopostman.com\",\"phone\":\"9876543210\",\"customer_id\":\"1\",\"key\":\"1254\",\"pubMessageType\":\"0\",\"pubPartition\":\"3\"}").Return(nil).AnyTimes()

	fmt.Println(testConsumer.Consume(true))
}

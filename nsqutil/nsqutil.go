package nsqutil

import (
	"encoding/json"

	"github.com/bitly/go-nsq"
	"github.com/kobeld/goutils"
)

var (
	nsqAddr    string
	lookupAddr string
)

func SetupNsqConfig(nsqAddr, lookupAddr string) {
	nsqAddr = nsqAddr
	lookupAddr = lookupAddr
	return
}

type Consumer interface {
	TopicAndChannel() (string, string)
	HandleMessage(*nsq.Message) error
}

func RegisterNsqConsumers(consumers []Consumer) (err error) {

	config := nsq.NewConfig()
	for _, consumer := range consumers {
		topic, channel := consumer.TopicAndChannel()

		nsqConsumer, err := nsq.NewConsumer(topic, channel, config)
		if err != nil {
			panic(err)
		}

		nsqConsumer.AddHandler(consumer)

		err = nsqConsumer.ConnectToNSQLookupd(lookupAddr)
		if err != nil {
			goutils.PrintStackAndError(err)
			return err
		}
	}

	return
}

// The default NSQ Producer
var defaultNsqProducer *nsq.Producer

// Get the the singleton NSQ Producer
func getNsqProducer() (*nsq.Producer, error) {
	var err error
	if defaultNsqProducer == nil {
		config := nsq.NewConfig()
		defaultNsqProducer, err = nsq.NewProducer(nsqAddr, config)
		if err != nil {
			goutils.PrintStackAndError(err)
			return nil, err
		}
	}
	return defaultNsqProducer, nil
}

// The general publish method that put data into NSQ as a producer
func PuslishNsqTopicWithData(topic string, topicData interface{}) (err error) {
	producer, err := getNsqProducer()
	if err != nil {
		goutils.PrintStackAndError(err)
		return
	}

	data, err := json.Marshal(topicData)
	if err != nil {
		goutils.PrintStackAndError(err)
		return
	}

	err = producer.Publish(topic, data)
	if err != nil {
		goutils.PrintStackAndError(err)
		return
	}

	return
}

package kafkaproducer

import (
	"github.com/Shopify/sarama"
	"time"
	"log"
	"github.com/spf13/viper"
)
var(
	//config file reader
	viperInstance *viper.Viper
)

func init(){
	viperInstance = viper.New()
	//add search path for the config file name
	viperInstance.AddConfigPath("/Users/tylerlisowski/Documents/gopath/src/github.com/relyt0925/company_news_reader/kafkaproducer/")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("json")
	viperInstance.ReadInConfig()
}

//NewProducer creates a new Async Kafka Producer based on the given configuration with
//connections to the ips/domains and port combinations specified in brokerList. Entry
//in broker list is <domain OR ip>:<port>
func NewProducer(brokerList []string, config *sarama.Config ) sarama.AsyncProducer {

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
}


func getDefaultConfig() *sarama.Config{
	// For the access log, we are looking for AP semantics, with high throughput.
	// By creating batches of compressed messages, we reduce network I/O at a cost of more latency.
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	return config
}

//NewDefaultProducer creates a Async Kafka Producer with default parameters
func NewDefaultProducer() sarama.AsyncProducer{
	//read from environment variable to get broker list
	brokerList := viperInstance.GetStringSlice("broker_list")
	config := getDefaultConfig()
	return NewProducer(brokerList,config)
}




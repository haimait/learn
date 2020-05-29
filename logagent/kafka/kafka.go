package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"os"
)

// send log msg to kafka

var (
	client  sarama.SyncProducer // 全局的kafka producer对象
	msgChan chan *Message
	log     *logrus.Logger
)

// Message 发送到kafka的message
type Message struct {
	Data  string
	Topic string
}

func init() {
	log = logrus.New()
	// 设置日志输出为os.Stdout
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	log.Info("kafka:init log success")
}

// Init 是初始化kafka的函数
func Init(addrs []string, chanSize int) (err error) {
	//配置kafka
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		log.Errorf("producer closed, err:", err)
		return
	}
	//初使化消息要发送的chan消息结构体(有缓冲区的)
	msgChan = make(chan *Message, chanSize)

	//死循环把chan消息发送到kafka的消息队列里
	go sendKafka()
	return
}

// SendLog 往msgChan发送消息的函数
func SendLog(msg *Message) (err error) {
	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("msgChan id full")
	}
	return
}

func sendKafka() {
	for msg := range msgChan {

		// 构造一个消息
		kafkaMsg := &sarama.ProducerMessage{}
		kafkaMsg.Topic = msg.Topic
		kafkaMsg.Value = sarama.StringEncoder(msg.Data)

		// 发送消息
		pid, offset, err := client.SendMessage(kafkaMsg)
		if err != nil {
			log.Warnf("send msg failed, err:%v\n", err)
			continue
		}
		log.Infof("send msg success, pid:%v offset:%v\n", pid, offset)
	}
}

package msq

import (
	"context"
	"encoding/json"
	"net"

	"github.com/mrehanabbasi/company-inc/config"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type MsqConn struct {
	Conn *kafka.Conn
}

func InitKafka() *MsqConn {
	broker := net.JoinHostPort(viper.GetString(config.KafkaHost), viper.GetString(config.KafkaPort))
	log.Info("Broker: ", broker)
	topicName := viper.GetString(config.KafkaTopic)

	conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topicName, 0)
	if err != nil {
		log.Panic(err)
		panic(err.Error())
	}

	return &MsqConn{Conn: conn}
}

func (w *MsqConn) ProduceCompanyEvent(company *models.Company, httpMethod string) error {
	// Serialize the Company struct to JSON
	eventData, err := json.Marshal(company)
	if err != nil {
		return err
	}

	// Produce the event to Kafka
	msg := kafka.Message{
		Key:   []byte(company.ID + "-" + httpMethod),
		Value: eventData,
	}
	_, err = w.Conn.WriteMessages(msg)
	return err
}

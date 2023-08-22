// nome  do pacote
package main

// dependências
import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// função de entrypoint do go run
func main() {
	// configurações do consumer
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp-consumer-" + time.Now().UTC().String(),
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}
	// criando o consumer
	c, err := kafka.NewConsumer(configMap)
	// se existirem  erros, retorna-os
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	// lista de tópicos a serem lidos
	topics := []string{"teste"}
	// inscrevendo o consumer nos tópicos
	c.SubscribeTopics(topics, nil)
	// laço infinito para leitura de mensagens
	for {
		// lendo as mensagens
		msg, err := c.ReadMessage(-1)
		// caso não existam erros
		if err == nil {
			// imprime a mensagem
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}

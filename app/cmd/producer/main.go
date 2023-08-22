// nome  do pacote
package main

// dependências
import (
	"fmt"
	"time"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// função de entrypoint do go run
func main() {
	// criando o canal de publicação para confirmação de envio
	deliveryChan := make(chan kafka.Event)
	// criando o producer
	producer := NewKafkaProducer()
	// publicando a mensagem
	Publish(time.Now().UTC().String(), "teste", producer, nil, deliveryChan)
	fmt.Println("Mensagem enviada. Aguardando confirmação...")
	// // método assíncrono de confirmação da publicação
	// go DeliveryReport(deliveryChan)
	// método síncrono de confirmação da publicação
	// aguardando a confirmação de envio
	e := <-deliveryChan
	// obtendo a mensagem de confirmação
	msg := e.(*kafka.Message)
	// se existirem erros, retorna-os
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar")
	// senão, imprime os dados da mensagem
	} else {
		fmt.Println("Mensagem publicada: ", "Headers=", msg.Headers,
			" Key=", msg.Key,
			" Topic[Partition]@Offset=", msg.TopicPartition)
	}
	// esperando o retorno da publicação
	producer.Flush(10000)
}

// função de criação de um producer
func NewKafkaProducer() *kafka.Producer {
	// configurações do producer
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}
	// criando o producer
	p, err := kafka.NewProducer(configMap)
	// se existirem  erros, retorna-os
	if err != nil {
		log.Println(err.Error())
	}
	// retorna o producer
	return p
}

// função para publicação de uma mensagem em um tópico
func Publish(content string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	// preparando a mensagem a ser publicada
	message := &kafka.Message{
		// payload, deve ser um slice de bytes
		Value:          []byte(content),
		// partition, foi configurado como aleatória
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		// key
		Key:            key,
	}
	// publicando a mensagem
	err := producer.Produce(message, deliveryChan)
	// se existirem erros, retorna-os
	if err != nil {
		return err
	}
	// não retorna erro
	return nil
}

// função de recebimento das mensagems de confirmação
func DeliveryReport(deliveryChan chan kafka.Event) {
	// loop infinito no canal que receberá as mensagens
	for e := range deliveryChan {
		// itera sobre o tipo da mensagem
		switch ev := e.(type) {
		// caso seja mensagem do kafka
		case *kafka.Message:
			// se existirem erros, retorna-os
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			// senão, imprime os dados da mensagem
			} else {
				fmt.Println("Mensagem publicada: ", "Headers=", ev.Headers,
					" Key=", ev.Key,
					" Topic[Partition]@Offset=", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"

	"github.com/IBM/sarama"
)

func main() {
	// Initialize Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatal("Error initializing Kafka producer:", err)
	}
	defer producer.Close()

	fmt.Println("salam hasti")

	producer.SendMessage(&sarama.ProducerMessage{
		Topic: "your-topic",
		Value: sarama.StringEncoder(string("yo yo from the kafka")),
	})

	// Create a queue with a limited length
	// q := queue.New()

	// // Start a goroutine to enqueue random bytes
	// go func() {
	// 	for {
	// 		data := generateRandomData()
	// 		q.Add(data)

	// 		// Check if the queue length is a multiple of 1024
	// 		if q.Length()%1024 == 0 {
	// 			// Dequeue and produce data
	// 			msg := q.Get(1024)
	// 			for i := 0; i < 4; i++ {
	// 				// Assert the underlying data type (byte slice) and then index it
	// 				if data, ok := msg.([]byte); ok {
	// 					// Produce messages using SendMessage
	// 					producer.SendMessage(&sarama.ProducerMessage{
	// 						Topic: "your-topic",
	// 						Value: sarama.StringEncoder(string(data[i])),
	// 					})
	// 				} else {
	// 					log.Println("Error: Unable to assert data as []byte")
	// 				}
	// 			}
	// 		}
	// 	}
	// }()

}

func generateRandomData() []byte {
	sentences := []string{
		"This is the first random sentence.",
		"Here is the second random sentence.",
		"And this is the third random sentence.",
	}

	// Choose one of the sentences randomly
	chosenSentence := sentences[rand.Intn(len(sentences))]

	// Hash the chosen sentence with MD5
	hasher := md5.New()
	hasher.Write([]byte(chosenSentence))
	md5Hash := hasher.Sum(nil)

	// Convert the MD5 hash to a hexadecimal string
	md5String := hex.EncodeToString(md5Hash)

	// Return the MD5 hash as []byte
	return []byte(md5String)
}

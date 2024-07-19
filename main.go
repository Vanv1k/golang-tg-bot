package main

import (
	"flag"
	"log"

	tgClient "github.com/Vanv1k/golang-tg-bot/clients/telegram"
	event_consumer "github.com/Vanv1k/golang-tg-bot/consumer/event-consumer"
	"github.com/Vanv1k/golang-tg-bot/events/telegram"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()))

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

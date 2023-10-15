package main

import (
	tgClient "family-bot/clients/telegram"
	event_consumer "family-bot/consumer/event-consumer"
	"family-bot/events/telegram"
	"family-bot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-token",
		"",
		"token argument")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

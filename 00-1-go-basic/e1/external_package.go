package main

import "github.com/dooray-go/dooray"

func main() {
	dooray.PostWebhook(
		"https://hook.dooray.com/services/3036349505739914786/3770555218093552684/autJQopeRTiVWUNxrgfaFA",
		&dooray.WebhookMessage{
			BotName: "Manty",
			Text:    "Hello, World!"},
	)
}

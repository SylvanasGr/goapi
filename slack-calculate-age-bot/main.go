package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Printf("Command events: %v \n Event timestamp: %v \n %v Event parameteres %v \n Event %v \n", event.Timestamp, event.Command, event.Parameters, event.Event)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6380590892000-6380597518672-onglJbiHA4d90h9IijqCIWfo")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06AWD3FTLH-6370445890577-d49c04b2e78f1e8bb952edb2f7d27fce3583922ee982a41f3e539a5e9336d6e1")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	// yob stands for year of birth
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := time.Now().Year() - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
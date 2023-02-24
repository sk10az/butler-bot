package main

import (
	"context"
	"github.com/sk10az/butler-bot/config"
	"log"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
)

type QueueItem struct {
	StartTime   time.Time
	Entities    tg.Entities
	Updates     *tg.UpdateNewChannelMessage
	IsNSFW      bool
	IsLandscape bool
	Input       string
}

func main() {
	log.Println("Starting...")
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	//var queues chan *QueueItem
	//queues = make(chan *QueueItem, 100)

	dispatcher := tg.NewUpdateDispatcher()
	client := telegram.NewClient(config.ApiId, config.ApiHash, telegram.Options{
		// Logger:        logger,
		UpdateHandler: dispatcher,
	})
	if err := client.Run(context.Background(), func(ctx context.Context) error {
		api := client.API()

		auth, err := client.Auth().Bot(ctx, config.BotToken)
		if err != nil {
			return err
		}

		log.Println(auth.User.GetID())

		up := uploader.NewUploader(api)
		sender := message.NewSender(api).WithUploader(up)

		//go processQueue(config, up, sender, queues)

		dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, u *tg.UpdateNewMessage) error {
			m, ok := u.Message.(*tg.Message)
			if !ok || m.Out {
				return nil
			}

			_, err := sender.Reply(e, u).Text(context.Background(), m.Message)
			return err
		})
		return nil
	}); err != nil {
		panic(err)
	}
}

package butler_bot

import (
	"log"
)

type QueueItem struct {
}

func main() {
	log.Println("Starting...")
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

}

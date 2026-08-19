package helpers

import (
	"context"
	"log"
	"time"

	"github.com/Moritisimor/shawty/internal/repo"
)

func StartCleanerRoutine(r *repo.URLAliasRepo, sleepTimeMinutes int) {
	ctx := context.Background()

	go func() {
		for {
			log.Println("Cleaner Routine Running!")

			aliases, err := r.GetAllAliases(ctx)
			if err != nil {
				log.Printf("Error while getting aliases: %s\nThis may be worth investigating\n", err.Error())
			}

			rn := time.Now().Unix()
			for _, alias := range aliases {
				if alias.DeleteAt <= rn {
					if err := r.DeleteAliasWithID(alias.ID, ctx); err != nil {
						log.Printf("Error while deleting alias: %s\nThis may be worth investigating\n", err.Error())
					}
				}
			}

			log.Printf("Cleaning done! Took %d milliseconds.\n", time.Now().UnixMilli() - rn * 1000)
			time.Sleep(time.Minute * time.Duration(sleepTimeMinutes))
		}
	}()
}

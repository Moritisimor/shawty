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
			time.Sleep(time.Minute * time.Duration(sleepTimeMinutes))
			log.Println("Cleaner Routine Running!")

			aliases, err := r.GetAllExpiredAliasIDs(ctx)
			if err != nil {
				log.Printf("Error while getting aliases: %s\n", err.Error())
				continue
			}

			rn := time.Now().Unix()
			for _, id := range aliases {
				if err := r.DeleteAliasWithID(id, ctx); err != nil {
					log.Printf("Error while deleting alias: %s", err.Error())
					continue
				}
			}

			log.Printf("Cleaning done! Took %d milliseconds.\n", time.Now().UnixMilli() - rn * 1000)
		}
	}()
}

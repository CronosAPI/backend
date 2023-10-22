package database

import (
	"backend/internal/types"
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Network:  os.Getenv("NETWORK"),
		Addr:     os.Getenv("ADDRESS"),
		Password: os.Getenv("PASSWORD"),
		DB:       0,
	})
}

func CreateDatabase() {
	ctx := context.Background()

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func InsertIntoTable(tableName string, data []types.ROCKETS) error {
	ctx := context.Background()

	for _, v := range data {
		// fmt.Printf("%s: %s %s, %s, %s\n", v.ID, v.Country, v.LaunchDate, v.Mass, v.Launcher)
		fv := map[string]interface{}{
			"Payload_ID":   v.Payload_Id,
			"Manufacturer": v.Manufacturer,
			"Payload_Mass": v.Payload_Mass,
			"Orbit":        v.Orbit,
		}

		for key, each := range fv {
			Client.HSet(ctx, fmt.Sprintf("%s:%s", tableName, v.Payload_Id), key, each)
		}

	}

	// fmt.Printf("Data inserted into table '%s'\n", tableName)
	return nil
}

/* Query Tables */
func QueryTable(tableName string) error {
	ctx := context.Background()

	// data, err := Client.Get(ctx, tableName).Result()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Data in table '%d'\n", len(data))
	// for field, _ := range data {
	// fmt.Println(data)
	// }

	iter := Client.Scan(ctx, 0, "solarflare:*", 0).Iterator()

	for iter.Next(ctx) {
		fmt.Println(iter.Val())
	}

	return nil
}

package isro

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InsertData_ISRO_Spacecrafts(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var spacecrafts SpaceCrafts
	if err := json.Unmarshal(data, &spacecrafts); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, spacecraft := range spacecrafts.SpaceCraftArray {
		key := KEY_SPACECRAFT + fmt.Sprint(spacecraft.ID)

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "id", spacecraft.ID)
		pipe.HSet(ctx, key, "name", spacecraft.Name)
		_, err := pipe.Exec(ctx)

		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return err
		}

	}

	return nil
}

func GetAllValues_ISRO_Spacecrafts(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_SPACECRAFT+"*"))

	var spacecraft SpaceCraft

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {

		if err := Client.HGetAll(ctx, key).Scan(&spacecraft); err != nil {
			panic(err)
		}
		if print {
			fmt.Printf("%d : %s,\n", spacecraft.ID, spacecraft.Name)
		}
	}
	if print {
		fmt.Println("}")
	}

	return nil
}

func InsertData_ISRO_Launchers(Client *redis.Client, data []byte) error {

	ctx := context.Background()

	var launchers Launchers
	if err := json.Unmarshal(data, &launchers); err != nil {
		fmt.Println("Error unmarshalling json object", err)
		return err
	}

	for _, launcher := range launchers.LauncherArray {
		key := KEY_LAUNCHER + fmt.Sprint(launcher.ID)

		pipe := Client.TxPipeline()

		pipe.HSet(ctx, key, "id", launcher.ID)
		_, err := pipe.Exec(ctx)

		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return err
		}

	}

	return nil
}

func GetAllValues_ISRO_Launchers(Client *redis.Client, print bool) error {
	ctx := context.Background()

	keys := Client.Keys(ctx, fmt.Sprint(KEY_LAUNCHER+"*"))

	var launcher Launcher

	if print {
		fmt.Println("{ ")
	}
	for _, key := range keys.Val() {

		if err := Client.HGetAll(ctx, key).Scan(&launcher); err != nil {
			panic(err)
		}
		if print {
			fmt.Printf("%s\n", launcher.ID)
		}
	}
	if print {
		fmt.Println("}")
	}

	return nil
}

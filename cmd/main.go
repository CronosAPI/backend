package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "github.com/gin-gonic/gin"
)

type SpaceCraft struct {
	Name string `json:"name"`
	ID   int    `json:"ID"`
}

type parseJSON struct {
	SpaceCrafts []SpaceCraft
}

func main() {
	response, err := http.Get("https://isro.vercel.app/api/spacecrafts")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var resp parseJSON // json resp
	json.Unmarshal(responseData, &resp)

	if err != nil {
		log.Fatalln(err)
	}

	for _, each := range resp.SpaceCrafts {
		fmt.Printf("ID %d: Name %s: \n", each.ID, each.Name)
	}

	// fmt.Println(string(responseData))
}

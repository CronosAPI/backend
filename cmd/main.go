package main

import (
	"backend/internal/parse"
	"backend/internal/types"
	"fmt"
)

func main() {
	ret := parse.GrabAPI("https://isro.vercel.app/api/spacecrafts")
	ans := types.CreateSpaceCraft(ret)
	fmt.Println(ans)
}

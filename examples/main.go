package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/wrkit/xcoin_client"
)

// main is the entry point for the application
func main() {
	// Create a new client with configuration
	cfg := xcoin_client.NewConfig("http://127.0.0.1:5678")

	xcoinClient := xcoin_client.NewClient(cfg)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	fmt.Println("GetVersion  -------------------------")
	// Get version information (internally non-blocking)
	version, err := xcoinClient.GetVersion(ctx)
	if err != nil {
		log.Printf("Failed to get version: %v\n", err)
		return
	}
	fmt.Printf("Version: %++v\n", version)

	fmt.Println("GetPairList2  -------------------------")

	pairList, err := xcoinClient.GetPairList2(ctx)
	if err != nil {
		log.Printf("Failed to get pair list: %v\n", err)
		return
	}

	sPairList, _ := json.MarshalIndent(pairList, " ", "	")

	fmt.Println(string(sPairList))

}

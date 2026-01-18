package main

import (
	"fmt"
	"log"

	"github.com/gen2brain/malgo"
)

func main() {

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		log.Fatalf("Failed to initialize malgo context: %v", err)
	}

	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	// Fetch Playback Devices
	playback, err := ctx.Devices(malgo.Playback)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Playback Devices ---")
	for i, d := range playback {
		// Use d.Name() to get the string name
		fmt.Printf("%d: %s\n", i, d.Name())
	}

	// Fetch Capture Devices
	capture, err := ctx.Devices(malgo.Capture)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n--- Capture Devices ---\n")
	for i, d := range capture {
		fmt.Printf("%d: %s\n", i, d.Name())
	}
}
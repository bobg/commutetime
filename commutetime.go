package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"googlemaps.github.io/maps"
)

func main() {
	var (
		origin = flag.String("origin", "", "origin")
		dest   = flag.String("dest", "", "destination")
		apiKey = flag.String("apikey", "", "API key")
	)

	flag.Parse()

	client, err := maps.NewClient(maps.WithAPIKey(*apiKey))
	if err != nil {
		log.Fatalf("creating maps client: %s", err)
	}

	req := &maps.DirectionsRequest{
		Origin:        *origin,
		Destination:   *dest,
		Mode:          maps.TravelModeDriving,
		DepartureTime: "now",
		Units:         maps.UnitsImperial,
		TrafficModel:  maps.TrafficModelBestGuess,
	}

	ctx := context.Background()
	routes, _, err := client.Directions(ctx, req)
	if err != nil {
		log.Fatalf("getting directions: %s", err)
	}

	var dur time.Duration
	for _, leg := range routes[0].Legs {
		dur += leg.DurationInTraffic
	}

	now := time.Now()
	day := now.Weekday().String()
	day = day[:3]
	fmt.Printf("%s %d-%02d-%02d %d:%02d %s\n", day, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), dur)
}

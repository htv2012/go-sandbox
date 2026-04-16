package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type BerryFlavorMap struct {
	Potency int              `json:"potency"`
	Flavor  NamedAPIResource `json:"flavor"`
}

type Berry struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	GrowthTime       int              `json:"growth_time"`
	MaxHarvest       int              `json:"max_harvest"`
	NaturalGiftPower int              `json:"natural_gift_power"`
	Size             int              `json:"size"`
	Smoothness       int              `json:"smoothness"`
	SoilDryness      int              `json:"soil_dryness"`
	Firmness         NamedAPIResource `json:"firmness"`
	Flavors          []BerryFlavorMap `json:"flavors"`
	Item             NamedAPIResource `json:"item"`
	NaturalGiftType  NamedAPIResource `json:"natural_gift_type"`
}

func (b Berry) String() string {
	return fmt.Sprintf("[%d] %s", b.ID, b.Name)
}

func fetchBerry(id int, client *http.Client) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/berry/%d", id)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[%d] Connection Error: %v\n", id, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[%d] HTTP Error: %s\n", id, resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[%d] Parse Error: %v\n", id, err)
		return
	}

	var b Berry
	if err := json.Unmarshal(body, &b); err != nil {
		return
	}

	fmt.Println(b)
}

func main() {
	startTime := time.Now()

	client := &http.Client{
		Timeout: time.Second * 30, // Increased timeout for bulk fetching
	}

	fmt.Println("Starting batch fetch (1-64)...")
	fmt.Println("--------------------------------")

	for i := 1; i <= 64; i++ {
		fetchBerry(i, client)
	}

	elapsed := time.Since(startTime)

	fmt.Println("--------------------------------")
	fmt.Printf("Success! Fetched 64 berries.\n")
	fmt.Printf("Total execution time: %s\n", elapsed)

	avg := elapsed / 64
	fmt.Printf("Average time per request: %s\n", avg)
}

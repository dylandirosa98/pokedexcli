package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func explore(fig *config, arg ...string) error {
	fmt.Printf("Exploring %s...\n", arg[0])
	link := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", arg[0])
	val, ok := fig.cache.Get(link)
	if ok {
		print(val)
		return nil
	} else {
		fig.cache.Add(link, val)
	}
	res, err := http.Get(link)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	data := Exploration{}
	json.Unmarshal(body, &data)
	if len(data.PokemonEncounters) == 0 {
		return fmt.Errorf("no data")
	}
	fmt.Print("Found Pokemon:\n")
	for _, i := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", i.Pokemon.Name)
	}
	return nil
}

type Exploration struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

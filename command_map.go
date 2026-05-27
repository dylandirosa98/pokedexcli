package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func mapCommand(fig *config) error {
	link := "https://pokeapi.co/api/v2/location-area"
	if fig.nextURL != nil {
		link = *fig.nextURL
	}
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
	data := Data{}
	json.Unmarshal(body, &data)
	x, ok := data.Previous.(string)
	if ok {
		fig.previousURL = &x
	}
	fig.nextURL = &data.Next
	if len(data.Results) == 0 {
		return fmt.Errorf("no data")
	}
	for i := 0; i < 20; i++ {
		name := data.Results[i].Name
		if name == "" {
			return fmt.Errorf("not a string")
		}
		fmt.Printf("%s\n", name)
	}
	return nil
}

type Data struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapb(fig *config) error {
	link := "https://pokeapi.co/api/v2/location-area"
	if fig.previousURL != nil {
		link = *fig.previousURL
	}
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
	data := Data{}
	json.Unmarshal(body, &data)
	x, ok := data.Previous.(string)
	if ok {
		fig.previousURL = &x
	}
	fig.nextURL = &data.Next
	if len(data.Results) == 0 {
		return fmt.Errorf("no data")
	}
	for i := 0; i < 20; i++ {
		name := data.Results[i].Name
		if name == "" {
			return fmt.Errorf("not a string")
		}
		fmt.Printf("%s\n", name)
	}
	return nil
}

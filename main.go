package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mrbri/pokedexcli/catch"
	"github.com/mrbri/pokedexcli/explore"
	"github.com/mrbri/pokedexcli/inspect"
	"github.com/mrbri/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Areas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type stack []string

func (s *stack) push(u string) {
	*s = append(*s, u)
}

func (s *stack) pop() string {
	if len(*s) == 0 {
		return ""
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func (s *stack) peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len((*s))-1]
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		// callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		// callback:    commandExit,
	},
}

// Name: pidgey
// Height: 3
// Weight: 18
// Stats:
//   -hp: 40
//   -attack: 45
//   -defense: 40
//   -special-attack: 35
//   -special-defense: 35
//   -speed: 56
// Types:
//   - normal
//   - flying

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  struct {
		hp      int
		attack  int
		defense int
	}
}

var pokedex = map[string]Pokemon

func main() {
	cache := pokecache.NewCache()

	var history stack
	// history.push("https://pokeapi.co/api/v2/location-area/?limit=20")
	nextUrl := ""

	for {
		fmt.Print("pokedex>")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		line := scanner.Text()

		if line == "exit" {
			return
		} else if line == "help" {
			fmt.Println("\nWelcome to Pokedex!")
			fmt.Printf("Usage:\n\n")
			for k, v := range commands {
				fmt.Printf("%s: %s\n", k, v.description)
			}
			fmt.Println("")
		} else if line == "map" {

			currentUrl := ""
			if nextUrl != "" {
				currentUrl = nextUrl
			} else if history.peek() != "" {
				currentUrl = history.peek()
			} else {
				currentUrl = "https://pokeapi.co/api/v2/location-area/?limit=20"
			}

			fmt.Println(currentUrl)
			res, err := http.Get(currentUrl)
			if err != nil {
				log.Fatal(err)
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				log.Fatal(err)
			}

			cache.Add(currentUrl, body)
			areas := Areas{}
			err = json.Unmarshal(body, &areas)
			if err != nil {
				log.Fatal(err)
			}
			// conf.Previous = conf.Next
			// conf.Next = areas.Next
			// history.push(areas.Next)
			history.push(currentUrl)
			nextUrl = areas.Next

			for _, s := range areas.Results {
				fmt.Println(s.Name)
			}
		} else if line == "mapb" {

			nextUrl = history.pop()
			prevUrl := history.peek()
			fmt.Printf("%s\n%s\n", nextUrl, prevUrl)
			res, err := http.Get(prevUrl)
			if err != nil {
				log.Fatal(err)
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				log.Fatal(err)
			}

			areas := Areas{}
			err = json.Unmarshal(body, &areas)
			if err != nil {
				log.Fatal(err)
			}

			for _, s := range areas.Results {
				fmt.Println(s.Name)
			}
		} else if line == "history" {
			for _, h := range history {
				fmt.Printf("%v+\n", h)
			}
		} else if strings.HasPrefix(line, "explore") {
			explComm := strings.Split(line, " ")
			exploreUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", explComm[1])

			fmt.Println(exploreUrl)
			res, err := http.Get(exploreUrl)
			if err != nil {
				log.Fatal(err)
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				log.Fatal(err)
			}

			explArea := explore.ExploreArea{}
			err = json.Unmarshal(body, &explArea)
			if err != nil {
				log.Fatal(err)
			}

			for _, enc := range explArea.PokemonEncounters {
				fmt.Println(enc.Pokemon.Name)
			}

		} else if strings.HasPrefix(line, "catch") {
			catchName := strings.Split(line, " ")
			catchUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", catchName[1])

			fmt.Println(catchUrl)
			res, err := http.Get(catchUrl)
			if err != nil {
				log.Fatal(err)
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				log.Fatal(err)
			}

			poke := catch.Pokemon{}
			err = json.Unmarshal(body, &poke)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Base Experience: %d\n", poke.BaseExperience)
		} else if strings.HasPrefix(line, "inspect") {
		}
	}
}

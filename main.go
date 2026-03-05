package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	i "pokedexcli/internal"
	"pokedexcli/internal/pokeapi"
	"time"
)

var ErrExit = errors.New("Closing the Pokedex... Goodbye!")

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

type Config struct {
	Next          *string // URL for the next page of results
	Previous      *string // URL for the previous page of results
	pokeapiClient pokeapi.Client
	Pokedex       map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Get the pokemon at a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your pokedex",
			callback:    commandPokedex,
		},
	}
}

func main() {
	httpClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		Next:          nil,
		Previous:      nil,
		pokeapiClient: httpClient,
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		words := i.CleanInput(line)
		firstWord := words[0]
		args := words[1:]

		commands := getCommands()
		if command, ok := commands[firstWord]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				if errors.Is(err, ErrExit) {
					os.Exit(0)
				}
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func commandExit(cfg *Config, args ...string) error {
	return ErrExit
}

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *Config, args ...string) error {
	if cfg.Next == nil && cfg.Previous != nil {
		return fmt.Errorf("No next location to map")
	}
	locations, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Prev

	for _, name := range locations.Results {
		fmt.Println(name.Name)
	}
	return nil
}

func commandMapBack(cfg *Config, args ...string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("You are on the first page.")
	}
	locations, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Prev

	for _, name := range locations.Results {
		fmt.Println(name.Name)
	}
	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No area provided")
	}
	area := args[0]

	fmt.Println("Exploring %s...", area)
	fmt.Println("Found Pokemon:")
	pokemon := []string{}
	exploreArea, err := cfg.pokeapiClient.ExploreArea(area)
	if err != nil {
		return err
	}
	for _, encounter := range exploreArea.PokemonEncounters {
		pokemon = append(pokemon, encounter.Pokemon.Name)
	}
	for _, name := range pokemon {
		fmt.Println(name)
	}
	return nil
}

func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon to catch")
	}
	pokemon := args[0]

	// get information about the pokemon from the API
	pokemonInfo, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	// Get the experience required to catch the pokemon
	experienceRequired := pokemonInfo.BaseExperience

	// Logic to catch or not the pokemon higher experience should be harder to catch
	if 40 > rand.Intn(experienceRequired) {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.Pokedex[pokemon] = *pokemonInfo
		fmt.Println("You may now inspect it with the inspect command.")
	}
	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon to inspect")
	}
	pokemonName := args[0]

	// get in the pokedex if exist
	pokemonInfo, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("Pokemon not found in pokedex")
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("	-%s: %d \n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, typ := range pokemonInfo.Types {
		fmt.Printf("	-%s \n", typ.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your pokedex is empty.")
		return nil
	}

	fmt.Println("Your pokedex:")
	for pokemon := range cfg.Pokedex {
		fmt.Println("	- " + pokemon)
	}

	return nil
}

# PokedexCLI
A CLI for interacting with the Pokedex API.
This project is an exercise from [Boot.dev](https://boot.dev). 
It's actually very nice exercice and learning.

## Installation
You need to install Go before and check if it's installed correctly:
```
go version
```
Then, clone the repository and run the following command:
```
go build
```
Finally, run the executable (it will be named `pokedexcli`):
```
./pokedexcli
```

## Usage
Run the executable and follow the prompts to interact with the Pokedex API.
This is an interactive CLI, so type commands and press enter to execute them.
> The prompt starts with `Pokedex > `. It's from the cli.

Exemple of a command `help`
```
Pokedex > help
Welcome to the Pokedex!
Usage:

pokedex: Displays your pokedex
exit: Exit the Pokedex
help: Displays a help message
map: Displays the next 20 location areas
mapb: Displays the previous 20 location areas
explore <area_name>: Get the pokemon at a location
catch <pokemon_name>: Catch a pokemon
inspect <pokemon_name>: Inspect a pokemon
```

### Commands
- `pokedex`: Displays your pokedex
- `exit`: Exit the Pokedex
- `help`: Displays a help message
- `map`: Displays the next 20 location areas
- `mapb`: Displays the previous 20 location areas
- `explore <area_name>`: Get the pokemon at a location
- `catch <pokemon_name>`: Catch a pokemon
- `inspect <pokemon_name>`: Inspect a pokemon

Then press enter to execute the command.

# Future and Ideas for Extending the Project
Few ideas from Lane

- [] Update the CLI to support the "up" arrow to cycle through previous commands
- [] Simulate battles between pokemon
- [] Add more unit tests
- [] Refactor the code to organize it better and make it more testable
- [] Keep pokemon in a "party" and allow them to level up
- [] Allow for pokemon that are caught to evolve after a set amount of time
- [] Persist a user's Pokedex to disk so they can save progress between sessions
- [] Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- [] Random encounters with wild pokemon
- [] Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

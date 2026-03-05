package pokeapi

type APIPokeReponseLocation struct {
	Count   int        `json:"count"`
	Next    *string    `json:"next"`
	Prev    *string    `json:"previous"`
	Results []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type APIPokeResponseArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int64                 `json:"game_index"`
	ID                   int64                 `json:"id"`
	Location             Location              `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod Location                           `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64    `json:"rate"`
	Version Location `json:"version"`
}

type Name struct {
	Language Location `json:"language"`
	Name     string   `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        Location                        `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int64             `json:"max_chance"`
	Version          Location          `json:"version"`
}

type EncounterDetail struct {
	Chance          int64      `json:"chance"`
	ConditionValues []Location `json:"condition_values"`
	MaxLevel        int64      `json:"max_level"`
	Method          Location   `json:"method"`
	MinLevel        int64      `json:"min_level"`
}

type Pokemon struct {
	Abilities              []Ability     `json:"abilities"`
	BaseExperience         int           `json:"base_experience"`
	Cries                  Cries         `json:"cries"`
	Forms                  []Species     `json:"forms"`
	GameIndices            []GameIndex   `json:"game_indices"`
	Height                 int64         `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int64         `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Move        `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int64         `json:"order"`
	PastAbilities          []PastAbility `json:"past_abilities"`
	PastStats              []PastStat    `json:"past_stats"`
	PastTypes              []interface{} `json:"past_types"`
	Species                Species       `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Stats                  []Stat        `json:"stats"`
	Types                  []Type        `json:"types"`
	Weight                 int64         `json:"weight"`
}

type Ability struct {
	Ability  *Species `json:"ability"`
	IsHidden bool     `json:"is_hidden"`
	Slot     int64    `json:"slot"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type GameIndex struct {
	GameIndex int64   `json:"game_index"`
	Version   Species `json:"version"`
}

type Move struct {
	Move                Species              `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int64   `json:"level_learned_at"`
	MoveLearnMethod Species `json:"move_learn_method"`
	Order           *int64  `json:"order"`
	VersionGroup    Species `json:"version_group"`
}

type PastAbility struct {
	Abilities  []Ability `json:"abilities"`
	Generation Species   `json:"generation"`
}

type PastStat struct {
	Generation Species `json:"generation"`
	Stats      []Stat  `json:"stats"`
}

type Stat struct {
	BaseStat int64   `json:"base_stat"`
	Effort   int64   `json:"effort"`
	Stat     Species `json:"stat"`
}

type GenerationV struct {
	BlackWhite Sprites `json:"black-white"`
}

type GenerationIv struct {
	DiamondPearl        Sprites `json:"diamond-pearl"`
	HeartgoldSoulsilver Sprites `json:"heartgold-soulsilver"`
	Platinum            Sprites `json:"platinum"`
}

type Versions struct {
	GenerationI    GenerationI     `json:"generation-i"`
	GenerationIi   GenerationIi    `json:"generation-ii"`
	GenerationIii  GenerationIii   `json:"generation-iii"`
	GenerationIv   GenerationIv    `json:"generation-iv"`
	GenerationIx   GenerationIx    `json:"generation-ix"`
	GenerationV    GenerationV     `json:"generation-v"`
	GenerationVi   map[string]Home `json:"generation-vi"`
	GenerationVii  GenerationVii   `json:"generation-vii"`
	GenerationViii GenerationViii  `json:"generation-viii"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Sprites         `json:"showdown"`
}

type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
	Other            *Other      `json:"other,omitempty"`
	Versions         *Versions   `json:"versions,omitempty"`
	Animated         *Sprites    `json:"animated,omitempty"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  RedBlue `json:"yellow"`
}

type RedBlue struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}

type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Gold    `json:"silver"`
}

type Crystal struct {
	BackDefault           string `json:"back_default"`
	BackShiny             string `json:"back_shiny"`
	BackShinyTransparent  string `json:"back_shiny_transparent"`
	BackTransparent       string `json:"back_transparent"`
	FrontDefault          string `json:"front_default"`
	FrontShiny            string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent      string `json:"front_transparent"`
}

type Gold struct {
	BackDefault      string  `json:"back_default"`
	BackShiny        string  `json:"back_shiny"`
	FrontDefault     string  `json:"front_default"`
	FrontShiny       string  `json:"front_shiny"`
	FrontTransparent *string `json:"front_transparent,omitempty"`
}

type GenerationIii struct {
	Emerald          OfficialArtwork `json:"emerald"`
	FireredLeafgreen Gold            `json:"firered-leafgreen"`
	RubySapphire     Gold            `json:"ruby-sapphire"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIx struct {
	ScarletViolet DreamWorld `json:"scarlet-violet"`
}

type DreamWorld struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}

type Home struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type GenerationVii struct {
	Icons             DreamWorld `json:"icons"`
	UltraSunUltraMoon Home       `json:"ultra-sun-ultra-moon"`
}

type GenerationViii struct {
	BrilliantDiamondShiningPearl DreamWorld `json:"brilliant-diamond-shining-pearl"`
	Icons                        DreamWorld `json:"icons"`
}

type Type struct {
	Slot int64   `json:"slot"`
	Type Species `json:"type"`
}

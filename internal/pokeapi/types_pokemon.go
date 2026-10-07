package pokeapi

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability
}

type PokemonForm struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	BaseExperience int              `json:"base_experience"`
	Height         int              `json:"height"`
	IsDefault      bool             `json:"is_default"`
	Order          int              `json:"order"`
	Weight         int              `json:"weight"`
	Abilities      []PokemonAbility `json:"ability"`
	Forms          []PokemonForm    `json:"forms"`
}

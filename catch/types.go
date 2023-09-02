package catch

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

type Stats struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

// type Pokemon struct {
// 	ID             int    `json:"id"`
// 	Name           string `json:"name"`
// 	BaseExperience int    `json:"base_experience"`
// 	Height         int    `json:"height"`
// 	IsDefault      bool   `json:"is_default"`
// 	Order          int    `json:"order"`
// 	Weight         int    `json:"weight"`
// 	Abilities      []struct {
// 		IsHidden bool `json:"is_hidden"`
// 		Slot     int  `json:"slot"`
// 		Ability  struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"ability"`
// 	} `json:"abilities"`
// }

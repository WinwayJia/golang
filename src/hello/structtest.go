package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:",string"`         // The "string" option signals that a field is stored as JSON inside a JSON-encoded string
	Year   int      `json:"released"`        // Field appears in JSON as key "released".
	Color  bool     `json:"color,omitempty"` // Field appears in JSON as key "myName" and the field is omitted from the object if its value is empty, as defined above.
	Actors []string `json:"-"`               // Field is ignored by this package.
	Star   float32  `json:",omitempty"`      // Field appears in JSON as key "Field" (the default), but the field is skipped if empty. Note the leading comma.
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}

	bytes, err := json.MarshalIndent(movies, "", " ")
	if err == nil {
		fmt.Println(string(bytes))
	}
}

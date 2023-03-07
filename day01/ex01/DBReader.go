package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type DBReader interface {
	read() (Recipes, error)
	convert(cakes Recipes) string
}

type XMLname string
type JSONname string
type Recipes struct {
	Cakes []Cakes `json:"cake" xml:"cake"`
}
type Cakes struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}
type Ingredients struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

func (filename XMLname) read() (Recipes, error) {
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename, err)
		return Recipes{nil}, err
	}
	var cakes Recipes
	err = xml.Unmarshal(file, &cakes)
	if err != nil {
		fmt.Println("can't read parse file", err)
	}
	return cakes, err
}
func (filename XMLname) convert(cakes Recipes) string {
	b, err := json.MarshalIndent(cakes, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (filename JSONname) read() (Recipes, error) {
	file, err := os.ReadFile(string(filename))
	if err != nil {
		fmt.Println("cannot read input file:", filename)
		return Recipes{}, err
	}
	var cakes Recipes
	err = json.Unmarshal(file, &cakes)
	if err != nil {
		fmt.Println("can't read parse file", err)
	}
	return cakes, err
}
func (filename JSONname) convert(cakes Recipes) string {
	b, err := xml.MarshalIndent(cakes, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}


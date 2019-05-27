package entities

import "fmt"

type Result struct {
	Geometry  Geometry `json:"geometry"`
	Icon      string   `json:"icon"`
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Photos    []*Photo `json:"photos"`
	PlaceId   string   `json:"place_id"`
	Reference string   `json:"reference"`
	Scope     string   `json:"scope"`
	Types     []string `json:"types"`
	Vicinity  string   `json:"vicinity"`
}

func (e *Result) String() string {
	return fmt.Sprintf("%+v", *e)
}

type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

func (e *Geometry) String() string {
	return fmt.Sprintf("%+v", *e)
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (e *Location) String() string {
	return fmt.Sprintf("%+v", *e)
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

func (e *Viewport) String() string {
	return fmt.Sprintf("%+v", *e)
}

type Photo struct {
	Height           int      `json:"height"`
	HtmlAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

func (e *Photo) String() string {
	return fmt.Sprintf("%+v", *e)
}

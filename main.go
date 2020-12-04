package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/reujab/wallpaper"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Rendition struct {
	Uri   string `json:"uri"`
	Width int    `json:"width,string"`
}

type Image struct {
	Renditions []Rendition `json:"renditions"`
}

type Photo struct {
	Image Image `json:"image"`
}

type Gallery struct {
	GalleryTitle string  `json:"galleryTitle"`
	Items        []Photo `json:"items"`
}

func main() {
	gallery := new(Gallery)
	ymString := time.Now().Format("2006-01")
	url := fmt.Sprintf("https://www.nationalgeographic.com/content/photography/en_US/photo-of-the-day/_jcr_content/.gallery.%s.json", ymString)
	getJson(url, gallery)
	renditions := gallery.Items[0].Image.Renditions
	// Sort renditions to get the biggest one
	sort.Slice(renditions, func(i, j int) bool { return renditions[i].Width > renditions[j].Width })
	biggest := renditions[0]
	wallpaper.SetFromURL(biggest.Uri)
}

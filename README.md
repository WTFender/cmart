# cmart
API client for Cleveland Museum of Art written in Go.

### CLI
Usage
```
cmart.exe {artwork,creator,exhibit} {get,list}
```
Example
```jsonc
// cmart.exe artwork get -id 111811
{
  "Id": 111811,
  "Accession_number": "1930.331",
  "Collection": "Indian Art",
  "Creation_date": "1000s",
  "Creation_date_earliest": 1000,
  "Creation_date_latest": 1099,
  "Creditline": "Purchase from the J. H. Wade Fund",
  "Current_location": "244 Indian and Southeast Asian",
  "Fun_fact": "The trampled figure holds a serpent in his left hand and with his right points up to Shiva."
  // ...snip...
}
```

### Package
```golang
// Create API client
CMA := cmart.NewCmaApi()

// Get an Artwork, Creator, or Exhibit by ID
artwork := CMA.GetArtworkById(111811)
fmt.Println(artwork.Title)

// List Artworks, Creators, and Exhibits; optionally apply filters
artworks, meta := CMA.ListArtworks(cmart.ArtworksFilter{
    Created_after:            2019,
    African_american_artists: true,
})
if meta.Total > 0 {
    fmt.Println(artworks[0].Title)
}
```
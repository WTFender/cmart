# cmart
API client for Cleveland Museum of Art written in Go.

### CLI
Usage
```
cmart.exe {artwork,creator,exhibit} {get,list}
```
Example
```json
$ cmart.exe artwork get -id 111811
{
  "Id": 111811,
  "Accession_number": "1930.331",
  "Collection": "Indian Art",
  "Creation_date": "1000s",
  "Creation_date_earliest": 1000,
  "Creation_date_latest": 1099,
  "Creditline": "Purchase from the J. H. Wade Fund",
  "Current_location": "244 Indian and Southeast Asian",
  "Fun_fact": "The trampled figure holds a serpent in his left hand and with his right points up to Shiva.",
  ...snip...
}
```
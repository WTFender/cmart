package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wtfender/cmart"
)

func printUsage(bin string, cmd string, sub string, err string) {
	usage := fmt.Sprintf("Usage: %v %v %v\n", bin, cmd, sub)
	if err != "" {
		usage = err + "\n" + usage
	}
	fmt.Println(usage)
}

func main() {
	// Setup RestApi Handler
	CMA := cmart.NewRestApi()
	bin := os.Args[0]

	// CLI Arguments
	artworkGet := flag.NewFlagSet(bin+" artwork get", flag.ExitOnError)
	artworkGetId := artworkGet.Int("id", 0, "CMA artwork ID to retrieve; takes precedence over -accession")
	artworkGetAccNbr := artworkGet.Float64("accession", 0.0, "CMA artwork accession number to retrieve")

	artworkFilter := cmart.ArtworksFilter{}
	artworkList := flag.NewFlagSet(bin+" artwork list", flag.ExitOnError)
	artworkList.BoolVar(&artworkFilter.African_american_artists, "african-american-artists", false, "Filters by works created by African American artists.")
	artworkList.StringVar(&artworkFilter.Artists, "artists", "", "Filter by name of artist.")
	artworkList.StringVar(&artworkFilter.Catalogue_raisonne, "catalogue-raisonne", "", "Filter by catalogue raisonne.")
	artworkList.BoolVar(&artworkFilter.Cc0, "cc0", false, "Filters by works that have share license cc0.")
	artworkList.BoolVar(&artworkFilter.Cia_alumni_artists, "cia-alumni-artists", false, "Filters by works created by Cleveland Institute of Art alumni.")
	artworkList.StringVar(&artworkFilter.Citations, "citations", "", "Keyword search against the citations field.")
	artworkList.BoolVar(&artworkFilter.Copyrighted, "copyrighted", false, "Filters by works that have some sort of copyright.")
	artworkList.IntVar(&artworkFilter.Created_after, "created-after", 0, "Returns artworks created after the year specified. Negative years are BCE.")
	artworkList.IntVar(&artworkFilter.Created_after_age, "created-after-age", 0, "Filters by artworks that were created by artists older than the provided value in years at time of creation.")
	artworkList.IntVar(&artworkFilter.Created_before, "created-before", 0, "Returns artworks created before the year specified. Negative years are BCE.")
	artworkList.IntVar(&artworkFilter.Created_before_age, "created-before-age", 0, "Filters by artworks that were created by artists younger than the provided value in years at time of creation.")
	artworkList.StringVar(&artworkFilter.Credit, "credit", "", "Filter by credit line.")
	artworkList.BoolVar(&artworkFilter.Currently_on_loan, "currently-on-loan", false, "Filters by works that are currently on loan.")
	artworkList.BoolVar(&artworkFilter.Currently_on_view, "currently-on-view", false, "Filters by works that are currently on view at CMA.")
	artworkList.StringVar(&artworkFilter.Department, "department", "", "Filter by department. List of valid departments in Appendix B.")
	artworkList.StringVar(&artworkFilter.Exhibition_history, "exhibition-history", "", "Filter by exhibition history of artwork.")
	artworkList.BoolVar(&artworkFilter.Female_artists, "female-artists", false, "Filters by artworks created by female artists.")
	artworkList.IntVar(&artworkFilter.Has_image, "has-image", 0, "0 or 1. Filter to return only artworks that have a web image asset. (synonymous with the deprecated field web_image)")
	artworkList.IntVar(&artworkFilter.Limit, "limit", 10, "Limit the number of results")
	artworkList.BoolVar(&artworkFilter.May_show_artists, "may-show-artists", false, "Filters by works exhibited in Cleveland Museum of Art May Shows")
	artworkList.StringVar(&artworkFilter.Medium, "medium", "", "Filter by artwork medium.")
	artworkList.BoolVar(&artworkFilter.Nazi_era_provenance, "nazi-era-provenance", false, "Filters by nazi-era provenance.")
	artworkList.StringVar(&artworkFilter.Provenance, "provenance", "", "Filter by provenance of artwork")
	artworkList.StringVar(&artworkFilter.Q, "q", "", "Any keyword or phrase that searches against title, creator, artwork description, and several other meaningful fields related to the artwork.")
	artworkList.BoolVar(&artworkFilter.Recently_acquired, "recently-acquired", false, "Filters by artworks acquired by the museum in the last three years.")
	artworkList.IntVar(&artworkFilter.Skip, "skip", 0, "Offset index for results.")
	artworkList.StringVar(&artworkFilter.Title, "title", "", "Filter by title of artwork.")
	artworkList.StringVar(&artworkFilter.Type, "type", "", "Filter by artwork types. List of valid types in Appendix C.")

	creatorGet := flag.NewFlagSet(bin+" creator get", flag.ExitOnError)
	creatorGetId := creatorGet.Int("id", 0, "CMA creator ID to retrieve")

	creatorFilter := cmart.CreatorsFilter{}
	creatorList := flag.NewFlagSet(bin+" creator list", flag.ExitOnError)
	creatorList.IntVar(&creatorFilter.Limit, "limit", 10, "Limit the number of results")
	creatorList.StringVar(&creatorFilter.Biography, "bio", "", "Filter creators by biography text")
	creatorList.IntVar(&creatorFilter.Birth_year, "birth-year", 0, "Filter creators by year born")
	creatorList.IntVar(&creatorFilter.Birth_year_after, "birth-year-after", 0, "Filter creators by year born")
	creatorList.IntVar(&creatorFilter.Birth_year_before, "birth-year-before", 0, "Filter creators by year born")
	creatorList.IntVar(&creatorFilter.Death_year, "death-year", 0, "Filter creators by year deceased")
	creatorList.IntVar(&creatorFilter.Death_year_after, "death-year-after", 0, "Filter creators by year deceased")
	creatorList.IntVar(&creatorFilter.Death_year_before, "death-year-before", 0, "Filter creators by year deceased")
	creatorList.StringVar(&creatorFilter.Name, "name", "", "Filter creators by name")
	creatorList.StringVar(&creatorFilter.Nationality, "nationality", "", "Filter creators by nationality")
	creatorList.IntVar(&creatorFilter.Skip, "skip", 0, "Skip the first n results")

	exhibitGet := flag.NewFlagSet(bin+" exhibit get", flag.ExitOnError)
	exhibitGetId := exhibitGet.Int("id", 0, "CMA exhibition ID to retrieve")

	exhibitFilter := cmart.ExhibitionsFilter{}
	exhibitList := flag.NewFlagSet(bin+" artwork list", flag.ExitOnError)
	exhibitList.StringVar(&exhibitFilter.Closed_after, "closed-after", "", "Filter exhibits closed after yyyy-mm-dd")
	exhibitList.StringVar(&exhibitFilter.Closed_before, "closed-before", "", "Filter exhibits closed before yyyy-mm-dd")
	exhibitList.IntVar(&exhibitFilter.Limit, "limit", 10, "Limit the number of results")
	exhibitList.StringVar(&exhibitFilter.Opened_after, "opened-after", "", "Filter exhibits opened after yyyy-mm-dd")
	exhibitList.StringVar(&exhibitFilter.Opened_before, "opened-before", "", "Filter exhibits opened before yyyy-mm-dd")
	exhibitList.StringVar(&exhibitFilter.Organizer, "organizer", "", "Filter by exhibition organizer")
	exhibitList.IntVar(&exhibitFilter.Skip, "skip", 0, "Skip the first n results")
	exhibitList.StringVar(&exhibitFilter.Title, "title", "", "Filter by exhibit title")
	exhibitList.StringVar(&exhibitFilter.Venue, "venue", "", "Filter by venue")

	isCmd := func(cmd string, cmds []string) bool {
		for _, c := range cmds {
			if cmd == c {
				return true
			}
		}
		return false
	}

	if len(os.Args) < 2 || !isCmd(os.Args[1], []string{"artwork", "creator", "exhibit"}) {
		printUsage(os.Args[0], "{artwork,creator,exhibit}", "{get,list}", "")
		os.Exit(1)
	} else if len(os.Args) < 3 || !isCmd(os.Args[2], []string{"get", "list"}) {
		printUsage(os.Args[0], os.Args[1], "{get,list}", "")
		os.Exit(1)
	}

	cmd := os.Args[1]
	sub := os.Args[2]

	switch cmd {
	case "artwork":

		switch sub {
		case "get":
			artworkGet.Parse(os.Args[3:])
			var artwork cmart.Artwork
			if *artworkGetId > 0 {
				artwork = CMA.GetArtworkById(*artworkGetId)
			} else if *artworkGetAccNbr > 0 {
				artwork = CMA.GetArtworkByAccessionNbr(*artworkGetAccNbr)
			} else {
				printUsage(bin, cmd, sub, "Missing required flag: id or accession (must be a non-zero number)")
				artworkGet.PrintDefaults()
				os.Exit(1)
			}
			artworkJson, _ := json.MarshalIndent(artwork, "", "  ")
			fmt.Println(string(artworkJson))

		case "list":
			artworkList.Parse(os.Args[3:])
			artworks, _ := CMA.ListArtworks(artworkFilter)
			artworksJson, _ := json.MarshalIndent(artworks, "", "  ")
			fmt.Println(string(artworksJson))
		}

	case "creator":

		switch sub {
		case "get":
			creatorGet.Parse(os.Args[3:])
			if *creatorGetId == 0 {
				printUsage(bin, cmd, sub, "Missing required flag: id  (must be a non-zero number)")
				creatorGet.PrintDefaults()
				os.Exit(1)
			}
			creator := CMA.GetCreatorById(*creatorGetId)
			creatorJson, _ := json.MarshalIndent(creator, "", "  ")
			fmt.Println(string(creatorJson))

		case "list":
			creatorList.Parse(os.Args[3:])
			creators, _ := CMA.ListCreators(creatorFilter)
			creatorsJson, _ := json.MarshalIndent(creators, "", "  ")
			fmt.Println(string(creatorsJson))
		}

	case "exhibit":

		switch sub {
		case "get":
			exhibitGet.Parse(os.Args[3:])
			if *exhibitGetId == 0 {
				printUsage(bin, cmd, sub, "Missing required flag: id  (must be a non-zero number)")
				exhibitGet.PrintDefaults()
				os.Exit(1)
			}
			exhibition := CMA.GetExhibitionById(*exhibitGetId)
			exhibitionJson, _ := json.MarshalIndent(exhibition, "", "  ")
			fmt.Println(string(exhibitionJson))

		case "list":
			exhibitList.Parse(os.Args[3:])
			exhibits, _ := CMA.ListExhibitions(exhibitFilter)
			exhibitsJson, _ := json.MarshalIndent(exhibits, "", "  ")
			fmt.Println(string(exhibitsJson))

		}
	}
}

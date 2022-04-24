package cmart

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
)

type ResponseInfo struct {
	Total int
	// TODO Parameters []Something
}

type ArtworksResponse struct {
	Data []Artwork
	Info ResponseInfo
}

type ArtworkResponse struct {
	Data Artwork
}

type CreatorsResponse struct {
	Data []Creator
	Info ResponseInfo
}

type CreatorResponse struct {
	Data Creator
}

type ExhibitionsResponse struct {
	Data []Exhibition `json:",omitempty"`
	Info struct {
		Total int
		// TODO Parameters []Something
	}
}

type ExhibitionResponse struct {
	Data Exhibition   `json:",omitempty"`
	Info ResponseInfo `json:",omitempty"`
}

type Creator struct {
	Id          int       `json:",omitempty"`
	Artworks    []Artwork `json:",omitempty"`
	Biography   string    `json:",omitempty"`
	Birth_year  string    `json:",omitempty"`
	Created_at  string    `json:",omitempty"`
	Death_year  string    `json:",omitempty"`
	Description string    `json:",omitempty"`
	Name        string    `json:",omitempty"`
	Nationality string    `json:",omitempty"`
	Role        string    `json:",omitempty"`
	Updated_at  string    `json:",omitempty"`
}

type Artwork struct {
	Id                          int          `json:",omitempty"`
	Accession_number            string       `json:",omitempty"`
	Catalogue_raisonne          string       `json:",omitempty"`
	Citations                   []Citation   `json:",omitempty"`
	Collection                  string       `json:",omitempty"`
	Copyright                   string       `json:",omitempty"`
	Creation_date               string       `json:",omitempty"`
	Creation_date_earliest      int          `json:",omitempty"`
	Creation_date_latest        int          `json:",omitempty"`
	Creators                    []Creator    `json:",omitempty"`
	Creditline                  string       `json:",omitempty"`
	Culture                     []string     `json:",omitempty"`
	Current_location            string       `json:",omitempty"`
	Department                  string       `json:",omitempty"`
	Digital_description         string       `json:",omitempty"`
	Edition_of_the_work         string       `json:",omitempty"`
	Find_spot                   string       `json:",omitempty"`
	Former_accession_numbers    []string     `json:",omitempty"`
	Fun_fact                    string       `json:",omitempty"`
	Measurements                string       `json:",omitempty"`
	Provenance                  []Provenance `json:",omitempty"`
	Series                      string       `json:",omitempty"`
	Series_in_original_language string       `json:",omitempty"`
	Share_license_status        string       `json:",omitempty"`
	Sketchfab_id                string       `json:",omitempty"`
	Sketchfab_url               string       `json:",omitempty"`
	State_of_the_work           string       `json:",omitempty"`
	Technique                   string       `json:",omitempty"`
	Title                       string       `json:",omitempty"`
	Title_in_original_language  string       `json:",omitempty"`
	Tombstone                   string       `json:",omitempty"`
	Type                        string       `json:",omitempty"`
	Updated_at                  string       `json:",omitempty"`
	Url                         string       `json:",omitempty"`
	Wall_description            string       `json:",omitempty"`
	Dimensions                  struct {
		Base    Dimension `json:",omitempty"`
		Overall Dimension `json:",omitempty"`
	}
	Exhibitions struct {
		Current []Exhibition `json:",omitempty"`
		Legacy  []string     `json:",omitempty"`
	}
	External_resources struct {
		Internet_archive []string `json:",omitempty"`
		Wikidata         []string `json:",omitempty"`
	}
	Images struct {
		Full  Image `json:",omitempty"`
		Print Image `json:",omitempty"`
		Web   Image `json:",omitempty"`
	}
	Inscriptions []struct {
		Inscription             string `json:",omitempty"`
		Inscription_translation string `json:",omitempty"`
		Inscription_remark      string `json:",omitempty"`
	}
	Support_materials []struct {
		Description string   `json:",omitempty"`
		Watermarks  []string `json:",omitempty"`
	}
	// TODO
	// Related_works []SomeType {}
}

type Citation struct {
	Citation    string `json:",omitempty"`
	Page_number string `json:",omitempty"`
	Url         string `json:",omitempty"`
}

type Dimension struct {
	Depth  float64 `json:",omitempty"`
	Height float64 `json:",omitempty"`
	Width  float64 `json:",omitempty"`
}

type Exhibition struct {
	Id           int       `json:",omitempty"`
	Artworks     []Artwork `json:",omitempty"`
	Closing_date string    `json:",omitempty"`
	Created_at   string    `json:",omitempty"`
	Description  string    `json:",omitempty"`
	Opening_date string    `json:",omitempty"`
	Organizer    string    `json:",omitempty"`
	Title        string    `json:",omitempty"`
	Updated_at   string    `json:",omitempty"`
	Venues       []Venue   `json:",omitempty"`
}

type Image struct {
	Filename string `json:",omitempty"`
	Filesize string `json:",omitempty"`
	Height   string `json:",omitempty"`
	Url      string `json:",omitempty"`
	Width    string `json:",omitempty"`
}

type Provenance struct {
	Citations   []string `json:",omitempty"`
	Date        string   `json:",omitempty"`
	Description string   `json:",omitempty"`
	Footnotes   []string `json:",omitempty"`
}

type Venue struct {
	Venue_id   int    `json:",omitempty"`
	Name       string `json:",omitempty"`
	Start_date string `json:",omitempty"`
	End_date   string `json:",omitempty"`
}

type ArtworksFilter struct {
	African_american_artists bool   //Filters by works created by African American artists.
	Artists                  string //Filter by name of artist.
	Catalogue_raisonne       string //Filter by catalogue raisonne.
	Cc0                      bool   //Filters by works that have share license cc0.
	Cia_alumni_artists       bool   //Filters by works created by Cleveland Institute of Art alumni.
	Citations                string //Keyword search against the citations field.
	Copyrighted              bool   //Filters by works that have some sort of copyright.
	Created_after            int    //Returns artworks created after the year specified. Negative years are BCE.
	Created_after_age        int    //Filters by artworks that were created by artists older than the provided value in years at time of creation.
	Created_before           int    //Returns artworks created before the year specified. Negative years are BCE.
	Created_before_age       int    //Filters by artworks that were created by artists younger than the provided value in years at time of creation.
	Credit                   string //Filter by credit line.
	Currently_on_loan        bool   //Filters by works that are currently on loan.
	Currently_on_view        bool   //Filters by works that are currently on view at CMA.
	Department               string //Filter by department. List of valid departments in Appendix B.
	Exhibition_history       string //Filter by exhibition history of artwork.
	Female_artists           bool   //Filters by artworks created by female artists.
	Has_image                int    //0 or 1. Filter to return only artworks that have a web image asset. (synonymous with the deprecated field web_image)
	Limit                    int    //Limit for number of results. If no limit provided, API will return the maximum (1000) number of records.
	May_show_artists         bool   //Filters by works exhibited in Cleveland Museum of Art May Shows
	Medium                   string //Filter by artwork medium.
	Nazi_era_provenance      bool   //Filters by nazi-era provenance.
	Provenance               string //Filter by provenance of artwork
	Q                        string //Any keyword or phrase that searches against title, creator, artwork description, and several other meaningful fields related to the artwork.
	Recently_acquired        bool   //Filters by artworks acquired by the museum in the last three years.
	Skip                     int    //Offset index for results.
	Title                    string //Filter by title of artwork.
	Type                     string //Filter by artwork types. List of valid types in Appendix C.
	/* TODO
	dimensions	float64,float64,float64	Filter artworks by dimensions with the unit of measurement being meters. This filter is somewhat tricky, as the terminolgy for describing object dimensions varies from object to object (for example coins have diameters, swords have lengths, and necklaces have heights). An object's most descriptive dimension (whatever you think is the best way to describe it in meters) is generally put in the first part of the comma seperated list of dimensions. A default value of 20cm will be used if no value is provided for a dimension in the list. The second and third dimensions places are interchangable and describe a square that an object's remaining dimensions could fit inside. The dimensions filter returns objects with a fault tolerance of 20cm on all dimensions.
	dimensions_max	float64,float64,float64	Filter artworks to return all works that can fit inside a box defined by provided 3 values with the unit of measurement being meters. Place the most descriptive dimension in the first value, and any remaining dimensions in the second two values. If no value is provided for a dimension, a default value of 20cm is used. The dimensions_max filter has a fault tolerance of 0 on all dimensions, and will not return objects that cannot fit in the described box.
	dimensions_min	float64,float64,float64	Filter artworks to return all works that cannot fit inside a box defined by provided 3 values with the unit of measurement being meters. Place the most descriptive dimension in the first value, and any remaining dimensions in the second two values. If no value is provided for a dimension, a default value of 20cm is used. The dimensions_min filter has a fault tolerance of 0 on all dimensions, and will not return objects that can fit in the described box.
	*/
}

type CreatorsFilter struct {
	Biography         string //Filter by a keyword in creator biography.
	Birth_year        int    //Filter by exact match on creator's birth year.
	Birth_year_after  int    //Filter by creators born after a certain year.
	Birth_year_before int    //Filter by creators born before a certain year.
	Death_year        int    //Filter by exact match on creator's death year.
	Death_year_after  int    //Filter by creators who have died after a certain year.
	Death_year_before int    //Filter by creators who have died before a certain year.
	Limit             int    //Limit for number of results. If no limit provided, API will return the maximum (100) number of records.
	Name              string //Filter by matches or partial matches to the name of any creator.
	Nationality       string //Filter by a keyword in creator nationality, e.g. 'French'.
	Skip              int    //Offset index for results.
}

type ExhibitionsFilter struct {
	Closed_after  string //Filter exhibitions closed after a certain data. (date in YYYY-MM-DD format, e.g. 1974-01-01)
	Closed_before string //Filter exhibitions closed before a certain data. (date in YYYY-MM-DD format, e.g. 1974-01-01)
	Limit         int    //Limit for number of results. If no limit provided, API will return the maximum (100) number of records.
	Opened_after  string //Filter exhibitions opened after a certain data. (date in YYYY-MM-DD format, e.g. 1974-01-01)
	Opened_before string //Filter exhibitions opened before a certain data. (date in YYYY-MM-DD format, e.g. 1974-01-01)
	Organizer     string //Filter by exhibition organizer.
	Skip          int    //Offset index for results.
	Title         string //Filter by matches or partial matches to the title of an exhibition.
	Venue         string //Filter by exhibitioned opened in certain venues.
}

type RestApi struct {
	Baseurl string
	Get     func(string, interface{})
}

func NewRestApi() RestApi {
	var api RestApi
	api.Baseurl = "https://openaccess-api.clevelandart.org/api"
	api.Get = api.call
	return api
}

func (api RestApi) call(path string, response interface{}) {
	err := requests.
		URL(api.Baseurl + path).
		ToJSON(&response).
		CheckStatus(200).
		Fetch(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func (api RestApi) ListArtworks(filter ArtworksFilter) ([]Artwork, ResponseInfo) {
	params := buildParams(filter)
	var resp ArtworksResponse
	path := fmt.Sprintf("/artworks/?%v", params)
	api.Get(path, &resp)
	return resp.Data, resp.Info
}

func (api RestApi) GetArtworkById(artwork_id int) Artwork {
	var resp ArtworkResponse
	path := fmt.Sprintf("/artworks/%v", artwork_id)
	api.Get(path, &resp)
	return resp.Data
}

func (api RestApi) GetArtworkByAccessionNbr(accession_nbr float64) Artwork {
	var resp ArtworkResponse
	path := fmt.Sprintf("/artworks/%v", accession_nbr)
	api.Get(path, &resp)
	return resp.Data
}

func (api RestApi) ListCreators(filter CreatorsFilter) ([]Creator, ResponseInfo) {
	params := buildParams(filter)
	var resp CreatorsResponse
	path := fmt.Sprintf("/creators/?%v", params)
	api.Get(path, &resp)
	return resp.Data, resp.Info
}

func (api RestApi) GetCreatorById(creator_id int) Creator {
	var resp CreatorResponse
	path := fmt.Sprintf("/creators/%v", creator_id)
	api.Get(path, &resp)
	return resp.Data
}

func (api RestApi) ListExhibitions(filter ExhibitionsFilter) ([]Exhibition, ResponseInfo) {
	params := buildParams(filter)
	var resp ExhibitionsResponse
	path := fmt.Sprintf("/exhibitions/?%v", params)
	api.Get(path, &resp)
	return resp.Data, resp.Info
}

func (api RestApi) GetExhibitionById(exhibit_id int) Exhibition {
	var resp ExhibitionResponse
	path := fmt.Sprintf("/exhibitions/%v", exhibit_id)
	api.Get(path, &resp)
	return resp.Data
}

func isDateParam(paramKey string) bool {
	dateParams := []string{"opened_after", "opened_before", "closed_after", "closed_before"}
	for _, p := range dateParams {
		if p == paramKey {
			return true
		}
	}
	return false
}

func buildParams(filter interface{}) string {
	params := ""
	v := reflect.ValueOf(filter)
	prop := v.Type()
	for i := 0; i < v.NumField(); i++ {
		param := ""
		val := v.Field(i).Interface()
		switch val {
		case false: //skip
		case nil: //skip
		case "": //skip
		case 0: //skip
		case true:
			param = strings.ToLower(prop.Field(i).Name)
		default:
			paramVal := fmt.Sprintf("%v", val)
			paramKey := strings.ToLower(prop.Field(i).Name)
			if isDateParam(paramKey) {
				d, err := time.Parse("2006-01-02", fmt.Sprintf("%v", val))
				if err != nil {
					fmt.Printf("Bad %v, expected yyyy-mm-dd: '%v'\n", paramKey, val)
					os.Exit(1)
				}
				paramVal = d.Format("2006-01-02")
			}
			param = paramKey + "=" + url.QueryEscape(paramVal)
		}
		if param != "" && params == "" {
			params += param
		} else if param != "" && params != "" {
			params += "&" + param
		}
	}
	return params
}

func test(CMA RestApi) {
	// Get an Artwork, Creator, or Exhibit by ID
	artwork := CMA.GetArtworkById(111811)
	fmt.Println(artwork.Title)

	creator := CMA.GetCreatorById(1859)
	fmt.Println(creator.Name)

	exhibit := CMA.GetExhibitionById(312462)
	fmt.Println(exhibit.Title)

	// List Artworks, Creators, and Exhibits; optionally apply filters
	artworks, _ := CMA.ListArtworks(ArtworksFilter{
		Created_after:            2019,
		African_american_artists: true,
	})
	if len(artworks) > 0 {
		fmt.Println(artworks[0].Title)
	}

	creators, _ := CMA.ListCreators(CreatorsFilter{
		Birth_year_after: 1990,
	})
	if len(creators) > 0 {
		fmt.Println(creators[0].Description)
	}

	exhibits, _ := CMA.ListExhibitions(ExhibitionsFilter{
		Opened_after: "2017-01-01",
	})
	if len(exhibits) > 0 {
		fmt.Println(exhibits[0].Title)
	}
}

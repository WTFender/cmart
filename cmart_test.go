package cmart

import (
	"context"
	"fmt"
	"testing"

	"github.com/carlmjohnson/requests"
)

func NewMockRestApi() RestApi {
	var api RestApi
	api.Baseurl = "https://openaccess-api.clevelandart.org/api"
	api.Get = api.mockCall
	return api
}

func (api RestApi) mockCall(path string, response interface{}) {
	fmt.Println(path)
	err := requests.URL(api.Baseurl + path).
		Transport(requests.Replay("tests")).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func TestListArtworks(t *testing.T) {
	limit := 3
	mockCMA := NewMockRestApi()
	artworks, _ := mockCMA.ListArtworks(ArtworksFilter{
		Limit: limit,
	})
	if len(artworks) != 3 {
		t.Logf("Expected %v artworks, got %v", limit, len(artworks))
		t.Fail()
	}
}

func TestListCreators(t *testing.T) {
	limit := 3
	mockCMA := NewMockRestApi()
	creators, _ := mockCMA.ListCreators(CreatorsFilter{
		Limit: limit,
	})
	if len(creators) != 3 {
		t.Logf("Expected %v creators, got %v", limit, len(creators))
		t.Fail()
	}
}

func TestListExhibits(t *testing.T) {
	limit := 3
	mockCMA := NewMockRestApi()
	exhibits, _ := mockCMA.ListExhibitions(ExhibitionsFilter{
		Limit: limit,
	})
	if len(exhibits) != 3 {
		t.Logf("Expected %v exhibits, got %v", limit, len(exhibits))
		t.Fail()
	}
}

func TestGetArtwork(t *testing.T) {
	mockCMA := NewMockRestApi()
	artworkIds := []int{110180, 92937, 94979}
	for _, aid := range artworkIds {
		artwork := mockCMA.GetArtworkById(aid)
		if artwork.Id != aid {
			t.Logf("Expected artwork id %v, got %v", aid, artwork.Id)
			t.Fail()
		}
	}
}

func TestGetCreator(t *testing.T) {
	mockCMA := NewMockRestApi()
	creatorIds := []int{1859, 290228, 1833}
	for _, cid := range creatorIds {
		creator := mockCMA.GetCreatorById(cid)
		if creator.Id != cid {
			t.Logf("Expected creator id %v, got %v", cid, creator.Id)
			t.Fail()
		}
	}
}

func TestGetExhibit(t *testing.T) {
	mockCMA := NewMockRestApi()
	exhibitsIds := []int{290449, 375733, 312462}
	for _, cid := range exhibitsIds {
		exhibit := mockCMA.GetExhibitionById(cid)
		if exhibit.Id != cid {
			t.Logf("Expected exhibition id %v, got %v", cid, exhibit.Id)
			t.Fail()
		}
	}
}

package main

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wtfender/cmart"
)

func getRandomArtwork() (cmart.Artwork, error) {
	CMA := cmart.NewCmaApi()
	artworks, _ := CMA.ListArtworks(cmart.ArtworksFilter{
		Limit: 1,
		Skip:  rand.Intn(100),
	})
	if len(artworks) > 0 {
		return artworks[0], nil
	}
	return cmart.Artwork{}, errors.New("no artwork found")
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, code := "", 0
	artwork, err := getRandomArtwork()
	if err != nil {
		body = string(err.Error())
		code = 501
	} else {
		artworkJson, _ := json.Marshal(artwork)
		body = string(artworkJson)
		code = 200
	}
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: code,
	}, nil
}

func main() {
	lambda.Start(handler)
}

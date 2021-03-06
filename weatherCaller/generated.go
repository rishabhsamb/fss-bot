package weatherCaller

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __getCityInput is used internally by genqlient
type __getCityInput struct {
	City string `json:"city"`
}

// getCityGetCityByNameCity includes the requested fields of the GraphQL type City.
type getCityGetCityByNameCity struct {
	Weather getCityGetCityByNameCityWeather `json:"weather"`
}

// getCityGetCityByNameCityWeather includes the requested fields of the GraphQL type Weather.
type getCityGetCityByNameCityWeather struct {
	Summary     getCityGetCityByNameCityWeatherSummary     `json:"summary"`
	Temperature getCityGetCityByNameCityWeatherTemperature `json:"temperature"`
	Wind        getCityGetCityByNameCityWeatherWind        `json:"wind"`
}

// getCityGetCityByNameCityWeatherSummary includes the requested fields of the GraphQL type Summary.
type getCityGetCityByNameCityWeatherSummary struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// getCityGetCityByNameCityWeatherTemperature includes the requested fields of the GraphQL type Temperature.
type getCityGetCityByNameCityWeatherTemperature struct {
	Actual    float64 `json:"actual"`
	FeelsLike float64 `json:"feelsLike"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
}

// getCityGetCityByNameCityWeatherWind includes the requested fields of the GraphQL type Wind.
type getCityGetCityByNameCityWeatherWind struct {
	Speed float64 `json:"speed"`
}

// getCityResponse is returned by getCity on success.
type getCityResponse struct {
	GetCityByName getCityGetCityByNameCity `json:"getCityByName"`
}

func getCity(
	ctx context.Context,
	client graphql.Client,
	city string,
) (*getCityResponse, error) {
	__input := __getCityInput{
		City: city,
	}
	var err error

	var retval getCityResponse
	err = client.MakeRequest(
		ctx,
		"getCity",
		`
query getCity ($city: String!) {
	getCityByName(name: $city) {
		weather {
			summary {
				title
				description
			}
			temperature {
				actual
				feelsLike
				min
				max
			}
			wind {
				speed
			}
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

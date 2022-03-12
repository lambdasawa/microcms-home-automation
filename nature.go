package main

import (
	"context"
	"os"

	nature "github.com/lambdasawa/microcms-home-automation/openapi/nature"
)

type (
	DeviceDetail struct {
		Device     nature.Device      `json:"deice"`
		Appliances []nature.Appliance `json:"appliances"`
	}

	Nature struct {
		User          *nature.User   `json:"user"`
		DeviceDetails []DeviceDetail `json:"deviceDetails"`
	}
)

func newNature() (*nature.DefaultApiService, context.Context) {
	api := nature.NewAPIClient(nature.NewConfiguration()).DefaultApi

	ctx := context.WithValue(context.Background(), nature.ContextAccessToken, os.Getenv("NATURE_ACCESS_TOKEN"))

	return api, ctx
}

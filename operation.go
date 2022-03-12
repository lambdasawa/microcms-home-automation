package main

import (
	"context"
	"fmt"

	nature "github.com/lambdasawa/microcms-home-automation/openapi/nature"
)

func handleEvent(event *WebhookEvent) error {
	api, ctx := newNature()

	n, err := fetch(ctx, api)
	if err != nil {
		return err
	}

	switch event.API {
	case "office":
		if err := handleOfficeEvent(ctx, api, n, event); err != nil {
			return err
		}
	case "living-room":
		if err := handleLivingRoomEvent(ctx, api, n, event); err != nil {
			return err
		}
	case "bedroom":
		if err := handleBedRoomEvent(ctx, api, n, event); err != nil {
			return err
		}
	}

	return nil
}

func handleOfficeEvent(ctx context.Context, api *nature.DefaultApiService, n *Nature, event *WebhookEvent) error {
	var light, aircon string
	for _, d := range n.DeviceDetails {
		for _, a := range d.Appliances {
			if *a.Nickname == "オフィスの照明" {
				light = *a.Id
			}
			if *a.Nickname == "オフィスのエアコン" {
				aircon = *a.Id
			}
		}
	}

	value := event.Contents.New.PublishValue

	if value.LightOn == true {
		if _, err := api.Call1AppliancesApplianceLightPost(ctx, light).Button("on").Execute(); err != nil {
			return err
		}
	} else {
		if _, err := api.Call1AppliancesApplianceLightPost(ctx, light).Button("off").Execute(); err != nil {
			return err
		}
	}

	if value.AirconOn == true {
		if _, err := api.
			Call1AppliancesApplianceAirconSettingsPost(ctx, aircon).
			OperationMode(value.AirModeValue()).
			AirDirection(value.AirDirValue()).
			AirVolume(value.AirVolValue()).
			Temperature(fmt.Sprint(value.Temp)).
			Execute(); err != nil {
			return err
		}
	} else {
		if _, err := api.
			Call1AppliancesApplianceAirconSettingsPost(ctx, aircon).
			Button("power-off").
			Execute(); err != nil {
			return err
		}
	}

	return nil

}

func handleLivingRoomEvent(ctx context.Context, api *nature.DefaultApiService, n *Nature, event *WebhookEvent) error {
	var light, aircon string
	for _, d := range n.DeviceDetails {
		for _, a := range d.Appliances {
			if *a.Nickname == "リビングの照明" {
				light = *a.Id
			}
			if *a.Nickname == "リビングのエアコン" {
				aircon = *a.Id
			}
		}
	}

	value := event.Contents.New.PublishValue

	if value.LightOn == true {
		if _, err := api.Call1AppliancesApplianceLightPost(ctx, light).Button("on").Execute(); err != nil {
			return err
		}
	} else {
		if _, err := api.Call1AppliancesApplianceLightPost(ctx, light).Button("off").Execute(); err != nil {
			return err
		}
	}

	if value.AirconOn == true {
		if _, err := api.
			Call1AppliancesApplianceAirconSettingsPost(ctx, aircon).
			OperationMode(value.AirModeValue()).
			AirDirection(value.AirDirValue()).
			AirVolume(value.AirVolValue()).
			Temperature(fmt.Sprint(value.Temp)).
			Execute(); err != nil {
			return err
		}
	} else {
		if _, err := api.
			Call1AppliancesApplianceAirconSettingsPost(ctx, aircon).
			Button("power-off").
			Execute(); err != nil {
			return err
		}
	}

	return nil
}

func handleBedRoomEvent(ctx context.Context, api *nature.DefaultApiService, n *Nature, event *WebhookEvent) error {
	var light string
	for _, d := range n.DeviceDetails {
		for _, a := range d.Appliances {
			if *a.Nickname == "寝室の照明" {
				light = *a.Id
			}
		}
	}

	value := event.Contents.New.PublishValue

	if value.LightOn == true {
		if _, err := api.Call1AppliancesApplianceLightPost(ctx, light).Button("on").Execute(); err != nil {
			return err
		}
	} else {
		if _, err := api.Call1AppliancesApplianceLightPost(ctx, light).Button("off").Execute(); err != nil {
			return err
		}
	}

	return nil
}

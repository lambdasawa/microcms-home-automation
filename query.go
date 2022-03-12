package main

import (
	"context"
	"encoding/json"
	"fmt"

	nature "github.com/lambdasawa/microcms-home-automation/openapi/nature"
)

func fetch(ctx context.Context, api *nature.DefaultApiService) (*Nature, error) {
	me, _, err := api.Call1UsersMeGet(ctx).Execute()
	if err != nil {
		return nil, err
	}

	devices, _, err := api.Call1DevicesGet(ctx).Execute()
	if err != nil {
		return nil, err
	}

	appliances, _, err := api.Call1AppliancesGet(ctx).Execute()
	if err != nil {
		return nil, err
	}

	details := []DeviceDetail{}
	for _, d := range devices {
		as := []nature.Appliance{}

		for _, a := range appliances {
			if d.Id == a.Device.Id {
				as = append(as, a)
			}
		}

		details = append(details, DeviceDetail{
			Device:     d,
			Appliances: as,
		})
	}

	n := &Nature{
		User:          me,
		DeviceDetails: details,
	}

	return n, nil
}

func show(n *Nature) error {
	bytes, err := json.Marshal(n)
	if err != nil {
		return err
	}
	fmt.Print(string(bytes))

	return nil
}

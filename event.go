package main

type WebhookEvent struct {
	Service  string   `json:"service,omitempty"`
	API      string   `json:"api,omitempty"`
	ID       string   `json:"id,omitempty"`
	Type     string   `json:"type,omitempty"`
	Contents Contents `json:"contents,omitempty"`
}

type Contents struct {
	Old New `json:"old,omitempty"`
	New New `json:"new,omitempty"`
}

type New struct {
	ID           string       `json:"id,omitempty"`
	Status       []string     `json:"status,omitempty"`
	DraftKey     interface{}  `json:"draftKey"`
	PublishValue PublishValue `json:"publishValue,omitempty"`
	DraftValue   interface{}  `json:"draftValue"`
}

type PublishValue struct {
	CreatedAt   string   `json:"createdAt,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
	PublishedAt string   `json:"publishedAt,omitempty"`
	RevisedAt   string   `json:"revisedAt,omitempty"`
	LightOn     bool     `json:"lightOn,omitempty"`
	AirconOn    bool     `json:"airconOn,omitempty"`
	AirMode     []string `json:"airMode,omitempty"`
	AirDir      []string `json:"airDir,omitempty"`
	AirVol      []string `json:"airVol,omitempty"`
	Temp        float64  `json:"temp,omitempty"`
}

func (value PublishValue) AirModeValue() string {
	switch v := value.AirMode[0]; v {
	case "暖房":
		return "warm"
	case "冷房":
		return "cool"
	default:
		return v
	}
}

func (value PublishValue) AirDirValue() string {
	switch v := value.AirDir[0]; v {
	case "自動":
		return "auto"
	case "スイング":
		return "swing"
	default:
		return v
	}
}

func (value PublishValue) AirVolValue() string {
	switch v := value.AirVol[0]; v {
	case "自動":
		return "auto"
	default:
		return v
	}
}

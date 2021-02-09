package schemas

// SettingJSON struct
type SettingJSON struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Label   string `json:"label"`
	Default string `json:"default"`
	Info    string `json:"info"`
}

// BlockJSON struct
type BlockJSON struct {
	ID       string        `json:"id"`
	Type     string        `json:"type"`
	Label    string        `json:"label"`
	Settings []SettingJSON `json:"settings"`
}

// SchemaJSON struct
type SchemaJSON struct {
	Name       string        `json:"name"`
	Class      string        `json:"class"`
	Settings   []SettingJSON `json:"settings"`
	BlocksJSON []BlockJSON   `json:"blocks"`
}

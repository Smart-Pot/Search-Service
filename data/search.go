package data

type Post struct {
	ID      string   `json:"id"`
	UserID  string   `json:"userId"`
	Plant   string   `json:"plant"`
	Info    string   `json:"info"`
	EnvData EnvData  `json:"envData"`
	Images  []string `json:"images"`
	Like    []string `json:"like"`
	Date    string   `json:"date"`
}

type EnvData struct {
	Humidity    string `json:"humidity" validate:"required"`
	Temperature string `json:"temperature" validate:"required"`
	Light       string `json:"light" validate:"required"`
}

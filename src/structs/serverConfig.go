package structs

type ServerConfig struct {
	Name        string `json:"serverName"`
	websiteName string `json:"websiteName"`
	Port        uint   `json:"port"`
}

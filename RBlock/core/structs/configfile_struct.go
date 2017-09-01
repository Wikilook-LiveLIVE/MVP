package structs

type ConfigFile struct {
	HttpServerPort string `json:"HttpServerPort"`
	SiteLocation   string `json: "SiteLocation"`
	LoggerConsole  bool   `json:"LoggerConsole"`
	LogFileSize    uint64 `json:"LogFileSize"`
	LogThreshold   rune   `json:"LogThreshold"` //D:68, I:73, E:69

	EncGlobalPass string `json:"EncGlobalPass"`
	EncGlobalSalt string `json:"EncGlobalSalt"`

	EmailHost string `json:"EmailHost"`
	EmailPort string `json:"EmailPort'`
	EmailUser string `json:"EmailUser'`
	EmailPass string `'json:"EmailPass"`

	DBhost    string `json:"DBhost"`
	DBport    uint16 `json:"DBport"`
	DBuser    string `json:"DBuser"`
	DBpass    string `json:"DBpass"`
	DBschema  string `json:"DBschema"`
	DBmaxConn int    `json:"DBmaxConn"`
}
package config

type Config struct {
	Debug int64 `json:"Debug"`
	Count int64 `json:"count"`
	HTTP  struct {
		Get  []string `json:"get"`
		Post []string `json:"post"`
	} `json:"http"`
	HTTPS struct {
		Get  []string `json:"get"`
		Post []string `json:"post"`
	} `json:"https"`
	Method      string `json:"method"`
	Millisecond int64  `json:"millisecond"`
	TLS         bool   `json:"tls"`
}

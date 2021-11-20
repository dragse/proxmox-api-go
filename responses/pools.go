package responses

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment"`
}

type PoolDetail struct {
	Members []interface{} `json:"members"`
	Comment string
}

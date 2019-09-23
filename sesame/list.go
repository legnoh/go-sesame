package sesame

const (
	listPath = "/sesames"
)

type Sesame struct {
	DeviceId string `json:"device_id"`
	Serial   string
	Nickname string
}

type Sesames struct {
	Sesames []Sesame
}

func (c *Client) GetSesameList() (*Sesames, error) {
	resp, err := c.get(listPath)
	if err != nil {
		return nil, err
	}

	var result Sesames
	return &result, c.decodeJSON(resp, &result)
}

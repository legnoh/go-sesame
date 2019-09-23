package sesame

const (
	listPath = "/sesames"
)

type Sesames struct {
	Sesames []Sesame
}

type Sesame struct {
	DeviceId string
	Serial   string
	Nickname string
}

func (c *Client) GetSesameList() (*Sesames, error) {
	resp, err := c.get(listPath)
	if err != nil {
		return nil, err
	}

	var result Sesames
	return &result, c.decodeJSON(resp, &result)
}

package sesame

const (
	listPath = "/sesames"
)

type Sesame struct {
	DeviceId string `json:"device_id"`
	Serial   string
	Nickname string
}

func (c *Client) GetSesameList() ([]Sesame, error) {
	resp, err := c.get(listPath)
	if err != nil {
		return nil, err
	}

	result := make([]Sesame, 0)
	return result, c.decodeJSON(resp, &result)
}

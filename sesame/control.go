package sesame

import "path"

const (
	controlPath = "/sesame"
)

type SesameStatus struct {
	Locked     bool
	Battery    int
	Responsive bool
}

type Command struct {
	Command string
}

type TaskId struct {
	TaskId string
}

func (c *Client) GetSesameStatus(deviceId string) (*SesameStatus, error) {

	apiPath := path.Join(controlPath, deviceId)

	resp, err := c.get(apiPath)
	if err != nil {
		return nil, err
	}

	var result SesameStatus
	return &result, c.decodeJSON(resp, &result)
}

func (c *Client) ControlSesame(deviceId string, command string) (*TaskId, error) {

	apiPath := path.Join(controlPath, deviceId)
	var data map[string]string = map[string]string{"command": command}

	resp, err := c.post(apiPath, data)
	if err != nil {
		return nil, err
	}

	var result TaskId
	return &result, c.decodeJSON(resp, &result)
}

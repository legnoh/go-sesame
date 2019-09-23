package sesame

const (
	taskPath = "/action-result"
)

type QueryExecutionResult struct {
	Status     string
	Successful bool
	Error      string
}

func (c *Client) GetQueryExecutionResult(taskId string) (*QueryExecutionResult, error) {

	resp, err := c.get(taskPath + "?task_id=" + taskId)
	if err != nil {
		return nil, err
	}

	var result QueryExecutionResult
	return &result, c.decodeJSON(resp, &result)
}

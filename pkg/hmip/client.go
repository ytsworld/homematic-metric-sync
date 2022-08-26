package hmip

func CreateClient() (*HmIPClient, error) {

	state := HmIPCurrentStatus{}

	c := &HmIPClient{
		AuthToken:    "",
		AccessPoint:  "",
		UserAgent:    "",
		CurrentState: &state,
	}

	return c, nil
}

package hmip

func (c *HmIPClient) SearchRoomForDevice(deviceId string) string {
	for _, v := range c.CurrentState.Groups {
		if v.Type == "META" {
			for _, channel := range v.Channels {
				if channel.DeviceId == deviceId {
					return v.Label
				}
			}
		}
	}

	return ""
}

package synthetics

import (
	"encoding/base64"
	"fmt"
)

// GetMonitorScript is used to retrieve the script that belongs
// to a New Relic Synthetics scripted monitor.
func (s *Synthetics) GetMonitorScript(monitorID string) (*MonitorScript, error) {
	resp := MonitorScript{}
	url := fmt.Sprintf("/monitors/%s/script", monitorID)
	_, err := s.client.Get(url, nil, &resp)

	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(resp.Text)

	if err != nil {
		return nil, err
	}

	resp.Text = string(decoded)

	return &resp, nil
}

// UpdateMonitorScript is used to add a script to an existing New Relic Synthetics monitor_script.
func (s *Synthetics) UpdateMonitorScript(monitorID string, script MonitorScript) error {
	script.Text = base64.StdEncoding.EncodeToString([]byte(script.Text))

	url := fmt.Sprintf("/monitors/%s/script", monitorID)
	_, err := s.client.Put(url, nil, &script, nil)

	if err != nil {
		return err
	}

	return nil
}

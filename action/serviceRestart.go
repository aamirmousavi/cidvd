package action

import (
	"os/exec"
)

func ServiceRestart(serviceName string) ([]byte, error) {
	cmd := exec.Command("systemctl", "restart", serviceName)
	rp, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return rp, nil
}

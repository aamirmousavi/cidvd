package action

import (
	"os/exec"
)

func GitPull(address string) ([]byte, error) {
	cmd := exec.Command("git", "pull", "--all")
	cmd.Dir = address
	rp, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return rp, nil
}

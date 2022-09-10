package action

import "os/exec"

func GoBuild(address, target string) ([]byte, error) {
	cmd := exec.Command("go", "build", "-o", target)
	cmd.Dir = address
	rp, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return rp, nil
}

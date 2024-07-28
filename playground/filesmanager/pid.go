package filesmanager

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Assumes pgrep installed and bash spawing only one child
func GetChildPID(parentID int) (int, error) {
	psCmd := exec.Command("pgrep", "-P", strconv.Itoa(parentID)) //nolint:gosec
	output, err := psCmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error finding child process: %v", err)
	}
	childPids := strings.Fields(string(output))
	if len(childPids) == 0 {
		return 0, fmt.Errorf("no child process found for PID %d", parentID)
	}
	childPid, err := strconv.Atoi(childPids[0])
	if err != nil {
		return 0, fmt.Errorf("invalid child PID: %v", err)
	}
	return childPid, nil
}

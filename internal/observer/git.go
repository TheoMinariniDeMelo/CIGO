package observer

import (
	"fmt"
	"os/exec"
	"strings"
)

type BranchState struct {
	Name 	   string
	LastCommit string
}

func GetAllBranchs(dir string) ([]string, error) {
	cmd := exec.Command("../scripts/get_branchs.sh")
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("sh error: %s", output)
	}

	return strings.Split(string(output[:]), "\n"), nil
}

func UpdateRepo(dir string) error {
	cmd := exec.Command("../scripts/update_repo.sh")
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sh error: %s", output)
	}
	return nil
}

func UpdateBranch(dir string) error {

	cmd := exec.Command("../scripts/update_branch.sh")
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sh error: %s", output)
	}
	return nil
}

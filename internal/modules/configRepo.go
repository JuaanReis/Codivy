package modules

import (
	"os/exec"
)

func SetupRepo(repoURL, repoName, language string) error {
	clonePath := "tmp/" + repoName
	templatePath := "template/" + language

	exec.Command("git", "clone", repoURL, clonePath).Run()

	err := CopyDir(templatePath, clonePath)
	if err != nil {
		return err
	}

	exec.Command("git", "-C", clonePath, "add", ".").Run()
	exec.Command("git", "-C", clonePath, "commit", "-m", "initial commit").Run()
	exec.Command("git", "-C", clonePath, "push").Run()

	return nil
}

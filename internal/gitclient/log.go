package gitclient

import (
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func (gc *GitClient) GetCommitHistory() ([]*object.Commit, error) {
	cIter, err := gc.repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, err
	}

	var commits []*object.Commit
	cIter.ForEach(func(c *object.Commit) error {
		commits = append(commits, c)
		return nil
	})

	return commits, nil
}

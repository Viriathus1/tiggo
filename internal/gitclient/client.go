package gitclient

import "github.com/go-git/go-git/v6"

type GitClient struct {
	repo     *git.Repository
	worktree *git.Worktree
}

func NewGitClient() (*GitClient, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return &GitClient{
		repo:     repo,
		worktree: worktree,
	}, nil
}

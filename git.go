package main

import (
	git "github.com/go-git/go-git/v5"
	"go.opentelemetry.io/otel/attribute"
)

type GitScm struct {
	repositoryPath string
}

// contributeAttributes this method never fails, returning the current state of the contributed attributes
// at the moment of the failure
func (scm *GitScm) contributeAttributes() []attribute.KeyValue {
	repository, err := git.PlainOpen(scm.repositoryPath)
	if err != nil {
		return []attribute.KeyValue{}
	}

	// from now on, this is a Git repository
	gitAttributes := []attribute.KeyValue{
		attribute.Key(ScmType).String("git"),
	}

	origin, err := repository.Remote("origin")
	if err != nil {
		return gitAttributes
	}
	gitAttributes = append(gitAttributes, attribute.Key(ScmRepository).StringSlice(origin.Config().URLs))

	branch, err := repository.Head()
	if err != nil {
		return gitAttributes
	}
	gitAttributes = append(gitAttributes, attribute.Key(ScmBranch).String(branch.Name().String()))

	return gitAttributes
}

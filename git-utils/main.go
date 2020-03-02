package main

import (
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	repoPath := os.Args[1]

	Info("git clone some random repo")

	repo, err := git.PlainOpen(repoPath)
	CheckIfError(err)

	branch, tag, err := GetCurrentBranchAndTag(repo)
	CheckIfError(err)

	Info("%s - %s", branch, tag)

	// Print the branch names and tags in a table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Path", "Branch", "Tag"})
	t.AppendRows([]table.Row{
		{1, repoPath, branch, tag},
	})
	t.Render()
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func GetCurrentBranchAndTag(repository *git.Repository) (string, string, error) {
	branchRefs, err := repository.Branches()
	if err != nil {
		return "", "", err
	}

	headRef, err := repository.Head()
	if err != nil {
		return "", "", err
	}

	var currentBranchName string
	branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchName = branchRef.Name().Short()
			return nil
		}
		return nil
	})

	var currentTagName string
	tagRefs, err := repository.Tags()
	tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		if tagRef.Hash() == headRef.Hash() {
			currentTagName = tagRef.Name().Short()
			return nil
		}
		return nil
	})

	return currentBranchName, currentTagName, nil
}

func GetCurrentCommitFromRepository(repository *git.Repository) (string, error) {
	headRef, err := repository.Head()
	if err != nil {
		return "", err
	}
	headSha := headRef.Hash().String()

	return headSha, nil
}

func GetLatestTagFromRepository(repository *git.Repository) (string, error) {
	tagRefs, err := repository.Tags()
	if err != nil {
		return "", err
	}

	var latestTagCommit *object.Commit
	var latestTagName string
	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		revision := plumbing.Revision(tagRef.Name().String())
		tagCommitHash, err := repository.ResolveRevision(revision)
		if err != nil {
			return err
		}

		commit, err := repository.CommitObject(*tagCommitHash)
		if err != nil {
			return err
		}

		if latestTagCommit == nil {
			latestTagCommit = commit
			latestTagName = tagRef.Name().String()
		}

		if commit.Committer.When.After(latestTagCommit.Committer.When) {
			latestTagCommit = commit
			latestTagName = tagRef.Name().String()
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return latestTagName, nil
}

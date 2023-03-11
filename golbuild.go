package golbuild

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gol-gol/golbin"
	"github.com/gol-gol/golfiles"
)

type GolBuild struct {
	Name    string     `json:"name"`
	Version string     `json:"version"`
	Date    string     `json:"date"`
	Git     GitDetails `json:"git"`
}

type GitDetails struct {
	Commit string `json:"commit"`
	Branch string `json:"branch"`
	Date   string `json:"date"`
}

var (
	DateFormat = "2006-01-02 15:04:05"
)

/*
FetchDetails prepares GolBuild struct with details.
*/
func (build *GolBuild) FetchDetails() {
	gitHash, errCommit := golbin.Exec("git log -1 --pretty=format:'%H'")
	gitBranch, errBranch := golbin.Exec("git rev-parse --abbrev-ref HEAD")
	gitDate, errDate := golbin.Exec("git log -1 --pretty=format:'%cd'")
	if errCommit != nil {
		gitHash = "undef"
	}
	if errBranch != nil {
		gitBranch = "undef"
	}
	if errDate != nil {
		gitDate = "undef"
	}

	if build.Name == "" {
		build.Name = golfiles.CwdBasename()
	}
	if build.Version == "" || build.Version[0] != 'v' {
		build.Version = gitHash
	}
	build.Date = time.Now().Format(DateFormat)
	build.Git = GitDetails{
		Commit: gitHash,
		Branch: gitBranch,
		Date:   gitDate,
	}
}

/*
Print simply prints details from GolBuild in a pre-determined format.
*/
func (build *GolBuild) Print() {
	msg := `* %s
	* Version: %s
	* Date: %s
	* From Git Commit
		Hash: %s
		Branch: %s
		Date: %s
`
	fmt.Printf(
		msg,
		build.Name,
		build.Version,
		build.Date,
		build.Git.Commit,
		build.Git.Branch,
		build.Git.Date,
	)
}

/*
Unmarshal simply wraps around json.Unmarshal for GolBuild.
*/
func Unmarshal(buildData []byte, build *GolBuild) error {
	if err := json.Unmarshal(buildData, build); err != nil {
		return err
	}
	return nil
}

/*
Print simply prints details from GolBuild Bytes.
*/
func Print(buildBytes []byte) {
	var build GolBuild
	if err := Unmarshal(buildBytes, &build); err != nil {
		log.Println("[WARN] There have been issues to deserialize .golbuild.json\n", err.Error())
		return
	}

	build.Print()
}

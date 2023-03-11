package golbuild

import (
	"encoding/json"
	"testing"
	"time"
)

func TestGolBuildFetchDetails(t *testing.T) {
	build := GolBuild{}
	build.FetchDetails()

	if build.Name != "golbuild" {
		t.Error("FAILED for build name, or you have clone under different name.")
	}
	if build.Version == "" {
		t.Error("FAILED for build version.")
	}

	buildDate, errDate := time.Parse(DateFormat, build.Date)
	if errDate != nil || buildDate.Day() != time.Now().Day() {
		t.Error("FAILED for build date.")
	}

	if build.Git.Commit == "" || build.Git.Branch == "" || build.Git.Date == "" {
		t.Error("FAILED to try Git Details.")
	}
}

func TestUnmarshal(t *testing.T) {
	var buildX, buildY GolBuild
	buildX.FetchDetails()
	buildBytes, _ := json.Marshal(buildX)
	if err := Unmarshal(buildBytes, &buildY); err != nil {
		t.Error("FAILED to Unmarshal.")
	}

	if buildX.Name != buildY.Name {
		t.Error("FAILED to match data from Unmarshal.")
	}
}

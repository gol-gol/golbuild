package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gol-gol/golbuild"
	"github.com/gol-gol/golfiles"
)

var (
	//go:embed .golbuild.json
	BuildBytes []byte

	GoBuildJson = flag.String("f", ".golbuild.json", "build file that would be embedded")
)

func main() {
	flag.Parse()
	if len(BuildBytes) > 0 {
		fmt.Println("GolBuild Details:")
		golbuild.Print(BuildBytes)
	}

	if err := UpdateGobuildJson(*GoBuildJson); err != nil {
		log.Printf("Failed to persist build details. %s", err.Error())
	}

	fmt.Println("Written Details:")
	buildBytes, _ := ioutil.ReadFile(*GoBuildJson)
	golbuild.Print(buildBytes)
}

/*
UpdateGobuildJson prepare GoBuildJson file to be embedded..
	it picks Name as current dir if not available in file already,
	sets Version as current Git Hash if empty or first char is not 'v',
	the Date as current time,
	also sets Git Commit, Branch & Date details.
*/
func UpdateGobuildJson(goBuildJson string) error {
	build := golbuild.GolBuild{}
	if golfiles.PathExists(goBuildJson) {
		buildValue, errJson := ioutil.ReadFile(goBuildJson)
		if errJson == nil {
			json.Unmarshal(buildValue, &build)
		}
	}

	build.FetchDetails()

	buildBytes, errMarshal := json.Marshal(build)
	if errMarshal != nil {
		return errMarshal
	}
	return ioutil.WriteFile(goBuildJson, buildBytes, os.ModePerm)
}

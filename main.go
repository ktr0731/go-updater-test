package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	semver "github.com/ktr0731/go-semver"
	updater "github.com/ktr0731/go-updater"
	"github.com/ktr0731/go-updater/github"
)

var version = semver.MustParse("0.1.0")

var (
	v        = flag.Bool("v", false, "show version")
	doUpdate = flag.Bool("update", false, "do update")
)

func main() {
	flag.Parse()

	if *v {
		fmt.Println(version)
		return
	}

	means, err := updater.NewMeans(github.GitHubReleaseMeans("ktr0731", "go-updater-test", github.TarGZIPDecompresser))
	if err != nil {
		log.Fatal(err)
	}
	u := updater.New(version, means)
	u.UpdateIf = updater.FoundPatchUpdate

	fmt.Printf("current: %s\n", version)
	fmt.Println("checking updates...")
	updatable, latest, err := u.Updatable(context.Background())
	if updatable {
		fmt.Printf("updatable: %s to %s\n", version, latest)
	} else {
		fmt.Println("not found")
	}

	if *doUpdate {
		fmt.Println("updating...")
	}
}

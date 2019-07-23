package logging

import (
	"fmt"
	"runtime"
)

var (
	// from ldflags
	gitRevision string
	gitBranch   string
	buildNumber string
	commitID    string
	commitUnix  string
	buildUnix   string

	// calculated
	buildVersion   string
	runtimeVersion string
)

type Version struct {
	GitRevision    string `json:"gitRevision"`
	GitBranch      string `json:"gitBranch"`
	BuildNumber    string `json:"buildNumber"`
	CommitID       string `json:"commitId"`
	CommitUnix     string `json:"commitUnix"`
	BuildUnix      string `json:"buildUnix"`
	BuildVersion   string `json:"buildVersion"`
	RuntimeVersion string `json:"runtimeVersion"`
}

// Build version
//
// Example:
// 0.9.1-49-g8c4849d 1510923415-1511108364 665
//
// Description:
// 0.9.1-49-g8c4849d       latest git tag a commit hash and a dirtiness mark
// 1510923415-1511108364   build time and latest commit time in unix
// 665                     CI job build number
func getBuildVersion() string {
	version := fmt.Sprintf("%s %s-%s", commitID, buildUnix, commitUnix)
	if buildNumber != "" {
		version = fmt.Sprintf("%s %s", version, buildNumber)
	}
	return version
}

func getRuntimeVersion() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func GetVersion() Version {
	return Version{
		BuildNumber:    buildNumber,
		GitBranch:      gitBranch,
		GitRevision:    gitRevision,
		BuildUnix:      buildUnix,
		CommitID:       commitID,
		CommitUnix:     commitUnix,
		BuildVersion:   buildVersion,
		RuntimeVersion: runtimeVersion,
	}
}

package vcs

import (
	"runtime/debug"
)

// VCS related GO build constants
const (
	Vcs      = "vcs"
	Revision = "vcs.revision"
	Time     = "vcs.time"
	Modified = "vcs.modified"

	Default = "N/A"
)

type VCS struct {
	VCS      string
	Revision string
	Time     string
	Modified string
}

func NewVCS() VCS {

	vcs := VCS{}

	vcs.VCS = Default
	vcs.Revision = Default
	vcs.Time = Default
	vcs.Modified = Default

	return vcs

}

func GetVCS() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == Vcs {
				return setting.Value
			}
		}
	}
	return ""
}

func GetRevision() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == Revision {
				return setting.Value
			}
		}
	}
	return ""
}

func GetTime() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == Time {
				return setting.Value
			}
		}
	}
	return ""
}

func GetModified() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == Modified {
				return setting.Value
			}
		}
	}
	return ""
}

func Read(vcs *VCS) {

	if vcs == nil {
		return
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	for _, s := range info.Settings {
		switch s.Key {
		case Vcs:
			vcs.VCS = s.Value
		case Revision:
			vcs.Revision = s.Value
		case Time:
			vcs.Time = s.Value
		case Modified:
			vcs.Modified = s.Value
		}
	}

}

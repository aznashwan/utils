// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package configurer

const (
	// PackageManagerLoopFunction is a bash function that executes its arguments
	// in a loop with a delay until either the command either returns
	// with an exit code other than 100.
	PackageManagerLoopFunction = `
function package_manager_loop {
    local rc=
    while true; do
        if ($*); then
                return 0
        else
                rc=$?
        fi
        if [ $rc -eq 100 ]; then
		sleep 10s
                continue
        fi
        return $rc
	done
}`
)

// DefaultPackages is a list of the default packages Juju'd like to see
// installed on all it's machines.
var DefaultPackages = []string{
// TODO (everyone): populate this list with all required packages.
}

// SeriesRequiringCloudTools is a map of series which require the addition of a
// cloud archive cloud tools source.
var SeriesRequiringCloudTools = map[string]bool{
	// TODO (everyone): add any and all LTS's or other OS which require cloud
	// tools' series here.
	"precise": true,
}
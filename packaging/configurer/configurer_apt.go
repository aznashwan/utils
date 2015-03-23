// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package configurer

import (
	"github.com/juju/utils/packaging"
)

// aptConfigurer is the PackagingConfigurer implementation for apt-based systems.
type aptConfigurer struct {
	*baseConfigurer
}

// RenderSource implements PackagingConfigurer.
func (c *aptConfigurer) RenderSource(src packaging.PackageSource) string {
	return src.RenderSourceFile(AptSourceTemplate[1:])
}

// RenderPreferences implements PackagingConfigurer.
func (c *aptConfigurer) RenderPreferences(prefs packaging.PackagePreferences) string {
	return prefs.RenderPreferenceFile(AptPreferenceTemplate[1:])
}
/*
Copyright 2020 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package errors

import (
	"fmt"
	"regexp"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
	"github.com/GoogleContainerTools/skaffold/proto"
)

// re is a shortcut around regexp.MustCompile
func re(s string) *regexp.Regexp {
	return regexp.MustCompile(s)
}

type problem struct {
	regexp      *regexp.Regexp
	description string
	suggestion  func(opts config.SkaffoldOptions) string
}

// Build Problems are Errors in build phase
var knownBuildProblems = map[proto.StatusCode]problem{
	proto.StatusCode_BUILD_PUSH_ACCESS_DENIED: {
		regexp:      re(fmt.Sprintf(".* %s.* denied: .*", PushImageErrPrefix)),
		description: "Build Failed. No push access to specified image repository",
		suggestion:  suggestBuildPushAccessDeniedAction,
	},
	proto.StatusCode_BUILD_PROJECT_NOT_FOUND: {
		regexp:      re(fmt.Sprintf("build failed: %s.* unknown: Project", PushImageErrPrefix)),
		description: "Build Failed",
		suggestion:  func(config.SkaffoldOptions) string { return "Check your GCR project." },
	},
}

/*
Copyright 2018-2020, Arm Limited and affiliates.
Copyright 2018 The Kubernetes Authors.

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

package options

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"

	// libs that provide registration functions
	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/kubernetes/pkg/version/verflag"

	// ensure libs have a chance to globally register their flags
	_ "github.com/golang/glog"
)

// AddGlobalFlags explicitly registers flags that libraries (glog, verflag, etc.) register
// against the global flagsets from "flag" and "github.com/spf13/pflag".
// We do this in order to prevent unwanted flags from leaking into the Kubelet's flagset.
func AddGlobalFlags(fs *pflag.FlagSet) {
	addGlogFlags(fs)
	addCadvisorFlags(fs)
	verflag.AddFlags(fs)
	logs.AddFlags(fs)
}

// normalize replaces underscores with hyphens
// we should always use hyphens instead of underscores when registering kubelet flags
func normalize(s string) string {
	return strings.Replace(s, "_", "-", -1)
}

// register adds a flag to local that targets the Value associated with the Flag named globalName in global
func register(global *flag.FlagSet, local *pflag.FlagSet, globalName string) {
	if f := global.Lookup(globalName); f != nil {
		pflagFlag := pflag.PFlagFromGoFlag(f)
		pflagFlag.Name = normalize(pflagFlag.Name)
		local.AddFlag(pflagFlag)
	} else {
		panic(fmt.Sprintf("failed to find flag in global flagset (flag): %s", globalName))
	}
}

// registerDeprecated registers the flag with register, and then marks it deprecated
func registerDeprecated(global *flag.FlagSet, local *pflag.FlagSet, globalName, deprecated string) {
	register(global, local, globalName)
	local.Lookup(normalize(globalName)).Deprecated = deprecated
}

// addGlogFlags adds flags from github.com/golang/glog
func addGlogFlags(fs *pflag.FlagSet) {
	// lookup flags in global flag set and re-register the values with our flagset
	global := flag.CommandLine
	local := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	register(global, local, "logtostderr")
	register(global, local, "alsologtostderr")
	register(global, local, "v")
	register(global, local, "stderrthreshold")
	register(global, local, "vmodule")
	register(global, local, "log_backtrace_at")
	register(global, local, "log_dir")

	fs.AddFlagSet(local)
}

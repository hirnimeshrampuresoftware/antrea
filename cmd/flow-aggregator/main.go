// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main under directory cmd parses and validates user input,
// instantiates and initializes objects imported from pkg, and runs
// the process.
package main

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"antrea.io/antrea/pkg/log"
	"antrea.io/antrea/pkg/version"
)

func main() {
	command := newFlowAggregatorCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func newFlowAggregatorCommand() *cobra.Command {
	opts := newOptions()

	cmd := &cobra.Command{
		Use:  "flow-aggregator",
		Long: "The Flow Aggregator.",
		Run: func(cmd *cobra.Command, args []string) {
			log.InitLogs(cmd.Flags())
			defer log.FlushLogs()
			if err := opts.complete(args); err != nil {
				klog.Fatalf("Failed to complete args: %v", err)
			}
			if err := opts.validate(args); err != nil {
				klog.Fatalf("Failed to validate args: %v", err)
			}
			if err := run(opts); err != nil {
				klog.Fatalf("Error running flow aggregator: %v", err)
			}
		},
		Version: version.GetFullVersionWithRuntimeInfo(),
	}

	flags := cmd.Flags()
	opts.addFlags(flags)
	log.AddFlags(flags)
	return cmd
}

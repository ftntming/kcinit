// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"github.com/ftntming/kcinit/console"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output version of kcinit",
	Run: version,
}

func version(cmd *cobra.Command, args []string) {
	console.Writeln("kcinit version \"0.5\"")
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

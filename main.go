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

package main

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	bogglebotCmd().Execute()
}

// initialize the root command
func bogglebotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "boggle-bot",
		Short: "Get all possible words for a particular boggle board",
		Run:   run,
	}
	// Global flags
	pflags := cmd.PersistentFlags()
	pflags.StringP("gridSize", "S", "4", "number of columns in the grid")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	gridSize, err := strconv.Atoi(cmd.Flag("gridSize").Value.String())
	if err != nil {
		log.Fatalf("cannot convert flag 'gridSize' to int: %v", err)
	}
	expectedArgs := gridSize * gridSize
	if len(args) != expectedArgs {
		log.Fatalf("Not enough characters provided. Should be %v, got: %v", expectedArgs, len(args))
	}
	dict := LoadDict("./english-words/words.txt")
	defer dict.Close()
}

package core

import "github.com/spf13/cobra"

type CobraClosure = func(cmd *cobra.Command, args []string)
type StrFunc = func() string
type VariablesFunc = map[string]func() string

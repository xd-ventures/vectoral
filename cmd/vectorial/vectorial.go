package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xd-ventures/vectoral/internal/embeder"
	"github.com/xd-ventures/vectoral/internal/reader"
)

func main() {
	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable is not set")
		os.Exit(1)
	}

	var rootCmd = &cobra.Command{
		Use:   "vectoral [file]",
		Short: "Process file content or standard input",
		Args:  cobra.MaximumNArgs(1),
		Run:   run,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	var input []byte
	var err error

	if len(args) == 0 {
		input, err = parser.ReadStdin()
	} else {
		input, err = parser.ReadFile(args[0])
	}

	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	embedding, err := openai.CalculateOpenAIAdaEmbedding(string(input))
	if err != nil {
		fmt.Printf("Error calculating embedding: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(embedding)
}

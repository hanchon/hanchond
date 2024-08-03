package playground

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hanchon/hanchond/playground/database"
	"github.com/hanchon/hanchond/playground/sql"
	"github.com/spf13/cobra"
)

// changeVersionCmd represents the changeVersion command
var changeVersionCmd = &cobra.Command{
	Use:   "change-version [id] [version]",
	Args:  cobra.ExactArgs(2),
	Short: "Change the binary version of the given node",
	Long:  `It will update the database entry for the node, you need to manually stop and re-start the node for it to take effect on the running chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		queries := sql.InitDBFromCmd(cmd)
		id := args[0]
		idNumber, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("could not parse the ID:", err.Error())
			os.Exit(1)
		}

		_, err = queries.GetNode(context.Background(), idNumber)
		if err != nil {
			fmt.Println("could not get the node:", err.Error())
			os.Exit(1)
		}

		err = queries.SetBinaryVersion(context.Background(), database.SetBinaryVersionParams{
			BinaryVersion: strings.TrimSpace(args[1]),
			ID:            idNumber,
		})

		if err != nil {
			fmt.Println("could not update the binary version:", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Node %d update to version %s\n", idNumber, args[1])
	},
}

func init() {
	PlaygroundCmd.AddCommand(changeVersionCmd)
}

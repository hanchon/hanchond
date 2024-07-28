package converter

import (
	"fmt"
	"os"

	"github.com/hanchon/hanchond/lib/converter"
	"github.com/spf13/cobra"
)

// AddrCmd represents the addr command
var AddrCmd = &cobra.Command{
	Use:   "addr",
	Args:  cobra.ExactArgs(1),
	Short: "Convert between bech32 and hex addresses",
	Long:  `Convert between cosmos and ethereum encoded addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		if converter.Has0xPrefix(input) {
			prefix, err := cmd.Flags().GetString("prefix")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			value, err := converter.HexToBech32(input, prefix)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println(value)
		} else {
			addr, err := converter.Bech32ToHex(input)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(addr)
		}
	},
}

func init() {
	ConverterCmd.AddCommand(AddrCmd)
	AddrCmd.Flags().StringP("prefix", "p", "evmos", "bech32 prefix")
}

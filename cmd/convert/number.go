package convert

import (
	"fmt"
	"strings"

	converterutil "github.com/hanchon/hanchond/lib/converter"
	"github.com/spf13/cobra"
)

// numberCmd represents the number command
var numberCmd = &cobra.Command{
	Use:   "number",
	Args:  cobra.ExactArgs(1),
	Short: "Convert numbers between decimal and hex",
	Long:  `No flags required, but hex numbers MUST have the 0x prefix.`,
	Run: func(_ *cobra.Command, args []string) {
		number := args[0]
		input := strings.TrimSpace(number)
		if converterutil.Has0xPrefix(input) {
			fmt.Println(converterutil.HexStringToDecimal(input))
		} else {
			fmt.Println(converterutil.DecimalStringToHex(input))
		}
	},
}

func init() {
	ConvertCmd.AddCommand(numberCmd)
}

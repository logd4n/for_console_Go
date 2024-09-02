/*  # # #   # # #  # # #  # # # #       # # # #     #       #        # # #   # # # #
  #     # #     # #    # #             #    # #   #        #       # * * #  #
 #       #     # #    # # # # #       # # #   # #         #       #  *  #  #  # # #
#     # #     # #    # #             #    #   #          #       # * * #  #      #
# # #   # # #  # # #  # # # #       # # #    #          # # # #  # # #    # # # # */

package console

import (
	"bufio"
	"os"
)

func ConsolePauseWithoutMessage() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func ConsolePauseWithMessage(message string) {
	scanner := bufio.NewScanner(os.Stdin)
	print("\n", message)
	scanner.Scan()
}

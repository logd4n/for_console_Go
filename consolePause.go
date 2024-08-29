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

func ConsolePause() {
	scanner := bufio.NewScanner(os.Stdin)
	print("\nНажмите на любую клавишу...\n")
	scanner.Scan()
	//fmt.Scanf(" ")
}

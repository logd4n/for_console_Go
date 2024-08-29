/*  # # #   # # #  # # #  # # # #       # # # #     #       #        # # #   # # # #
  #     # #     # #    # #             #    # #   #        #       # * * #  #
 #       #     # #    # # # # #       # # #   # #         #       #  *  #  #  # # #
#     # #     # #    # #             #    #   #          #       # * * #  #      #
# # #   # # #  # # #  # # # #       # # #    #          # # # #  # # #    # # # # */

package console

import (
	"os"
	"os/exec"
	"runtime"
)

var systems map[string]func()

func init() {
	systems = make(map[string]func())

	systems["windows"] = func() {
		command := exec.Command("cmd", "/c", "cls")
		command.Stdout = os.Stdout
		command.Run()
	}

	systems["linux"] = func() {
		command := exec.Command("clear")
		command.Stdout = os.Stdout
		command.Run()
	}
}

func ConsoleClear() {
	value, ok := systems[runtime.GOOS]

	if ok {
		value()
	} else {
		panic("Your OS (" + runtime.GOOS + ") is unsupported!")
	}
}

// main.go

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const defaultWorkDurationInMins = 25
const defaultRestDurationInMins = 5

func main() {
	workDurationInMins := defaultWorkDurationInMins
	restDurationInMins := defaultRestDurationInMins
	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of arguments, using default values of 25minutes working and 5minutes resting")
	}
	if len(os.Args) == 3 {
		wdur, err1 := strconv.Atoi(os.Args[1])
		rdur, err2 := strconv.Atoi(os.Args[2])
		workDurationInMins = wdur
		restDurationInMins = rdur
		if err1 != nil || err2 != nil {
			fmt.Println("Problem setting interval values, using default values of 25minutes working and 5minutes resting")
		}
	}
	t := timer{
		start:        time.Now(),
		inWorkMode:   true,
		workDuration: workDurationInMins * 60,
		restDuration: restDurationInMins * 60,
	}
	m := model{
		timerMain:    t,
		workDuration: workDurationInMins * 60,
		restDuration: restDurationInMins * 60,
		progress:     progress.New(progress.WithDefaultGradient()),
	}
	fmt.Println(m)
	prevElapsed := 0
	for {
		elapsed := t.getElapsedTimeInSeconds()
		if elapsed != prevElapsed {
			m.timerMain.printTimeRemaining(elapsed)
			if _, err := tea.NewProgram(m).Run(); err != nil {
				fmt.Println("Oh no!", err)
				os.Exit(1)
			}
			prevElapsed = elapsed
			if t.shouldSwitchMode(elapsed) {
				t.alert()
				t.switchMode()
			}
		}
	}
}

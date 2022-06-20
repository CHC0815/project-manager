package main

import (
	"fmt"
	"os"

	"cophee.team/project-manager/config"
	"cophee.team/project-manager/project"
	"cophee.team/project-manager/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	project.Configuration = config.ReadConfig()
	if len(project.Configuration.ImportDirs) == 0 {
		fmt.Println("No import directories configured")
		os.Exit(1)
	}

	tui.Prog = tea.NewProgram(tui.NewMainModel())
	tui.Prog.EnterAltScreen()
	if err := tui.Prog.Start(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
}

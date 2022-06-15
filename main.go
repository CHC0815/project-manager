package main

import (
	"fmt"
	"os"

	"cophee.team/project-manager/config"
	"cophee.team/project-manager/project"
	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	project.Configuration = config.ReadConfig()
	if len(project.Configuration.ImportDirs) == 0 {
		fmt.Println("No import directories configured")
		os.Exit(1)
	}

	if err := tea.NewProgram(project.NewProjectModel()).Start(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
}

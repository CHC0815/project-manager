package styles

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var StatusMessageStyle = lipgloss.NewStyle().
						Foreground(lipgloss.AdaptiveColor{
							Light: "#04B575",
							Dark: "#04B575"}).Render

var AppStyle = lipgloss.NewStyle().Padding(1, 2)
var TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

var EntryBox = lipgloss.NewStyle().Padding(1, 2).
			Align(lipgloss.Center).Border(lipgloss.DoubleBorder())


func UpdateEntryBox(size tea.WindowSizeMsg) {
	EntryBox = lipgloss.NewStyle().Padding(1, 2).
			Align(lipgloss.Center).Border(lipgloss.DoubleBorder()).Width(size.Width - 6)
}
package project

import (
	"io/ioutil"
	"log"

	"cophee.team/project-manager/config"
	"cophee.team/project-manager/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var Configuration *config.Config

type SelectMsg struct {
	Path string
}
type Project struct {
	Name string
	Path string
}

func (p Project) FilterValue() string {
	return p.Name
}

func (p Project) Title() string {
	return p.Name
}

func (p Project) Description() string {
	return p.Path
}

type ProjectModel struct {
	list list.Model
	keys *constants.ListKeyMap
	delegateKeys *constants.KeyMap
}

func newListKeyMap() *constants.ListKeyMap {
	return &constants.ListKeyMap{
		InsertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		ToggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		ToggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		TogglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		ToggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}


func getProjects() []Project {
	var projects []Project
	
	items, _ := ioutil.ReadDir(Configuration.ImportDirs[0])
	for _, item := range items {
		if item.IsDir() {
			projects = append(projects, Project{
				Name: item.Name(),
				Path: Configuration.ImportDirs[0] + "/" + item.Name(),
			})
		}
	}
	
	return projects
}

func NewProjectModel() ProjectModel {

	var (
		delegateKeys = constants.DefaultKeyMap
		listKeys = newListKeyMap()
	)

	delegate := NewItemDelegate(&delegateKeys)
	projects := getProjects()
	items := make([]list.Item, len(projects))
	for i, project := range projects {
		items[i] = project
	}
	projectList := list.New(items, delegate, 0, 0)
	projectList.Title = "Projects"
	projectList.Styles.Title = constants.TitleStyle
	projectList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.InsertItem,
			listKeys.ToggleTitleBar,
			listKeys.ToggleStatusBar,
			listKeys.TogglePagination,
			listKeys.ToggleHelpMenu,
		}
	}
	return ProjectModel{
		list: projectList,
		keys: listKeys,
		delegateKeys: &delegateKeys,
	}
}

func (m ProjectModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m ProjectModel) View() string {
	return constants.AppStyle.Render(m.list.View())
}

func (m ProjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := constants.AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch {
		case key.Matches(msg, m.keys.ToggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil
		case key.Matches(msg, m.keys.ToggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil
		case key.Matches(msg, m.keys.TogglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil
		case key.Matches(msg, m.keys.ToggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func NewItemDelegate(keys *constants.KeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func (msg tea.Msg, m *list.Model) tea.Cmd {
		var path string

		if i, ok := m.SelectedItem().(Project); ok { 
			path = i.Path
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.Select):
				return selectProjectCmd(path)
			// case key.Matches(msg, keys.Delete):
			// 	index := m.Index()
			// 	m.RemoveItem(index)
			// 	if len(m.Items()) == 0 {
			// 		keys.Delete.SetEnabled(false)
			// 	}
			// 	return m.NewStatusMessage(constants.StatusMessageStyle("Deleted " + title))
			}
		}
		return nil
	}

	help := []key.Binding{keys.Select, keys.Delete, keys.Quit, keys.Up, keys.Down}

	d.ShortHelpFunc = func () []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

func selectProjectCmd(path string) tea.Cmd{
	return func() tea.Msg {
		log.Println("Selected project: " + path)
		return SelectMsg{Path: path}
	}
}
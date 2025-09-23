package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight = 14
	listWidth  = 50
)

var (
	hashStyle = lipgloss.NewStyle().Bold(true)

	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#444444")).
			Foreground(lipgloss.Color("#FFFFFF"))

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#808080"))
)

type commitItem struct {
	title string
	desc  string
}

func (ct *commitItem) Title() string       { return hashStyle.Render(ct.title) }
func (ct *commitItem) Description() string { return ct.desc }
func (ct *commitItem) FilterValue() string { return ct.desc }

type logModel struct {
	list list.Model
}

func NewLogList(commits []string) tea.Model {
	items := make([]list.Item, len(commits))

	for i, c := range commits {
		parts := strings.SplitN(c, " ", 2)
		if len(parts) == 2 {
			items[i] = &commitItem{
				title: parts[0],
				desc:  parts[1],
			}
		}
	}

	d := list.NewDefaultDelegate()
	d.Styles.NormalTitle = normalStyle
	d.Styles.NormalDesc = normalStyle
	d.Styles.SelectedTitle = selectedStyle
	d.Styles.SelectedDesc = selectedStyle

	l := list.New(items, d, listWidth, listHeight)
	l.Title = "Git Log"
	l.Styles.Title = selectedStyle.PaddingLeft(2).PaddingRight(2)
	l.SetShowStatusBar(false)

	return logModel{list: l}
}

func (m logModel) Init() tea.Cmd {
	return nil
}

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m logModel) View() string {
	return m.list.View()
}

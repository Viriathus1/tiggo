package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	listView     = lipgloss.NewStyle().Margin(1, 2)
	brutalBorder = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("15")). // pure white
			Padding(0, 1)
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("51")).Underline(true)
	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("15")). // white background
			Foreground(lipgloss.Color("0")).  // black text
			Bold(true)
	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15"))

	hashStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Bold(true) // gray hash
	messageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
)

type commitItem struct {
	title string
	desc  string
}

func (ct *commitItem) Title() string       { return hashStyle.Render(ct.title) }
func (ct *commitItem) Description() string { return ct.desc }
func (ct *commitItem) FilterValue() string { return messageStyle.Render(ct.desc) }

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

	l := list.New(items, d, 0, 0)
	l.Title = "Git Log"

	return logModel{list: l}
}

func (m logModel) Init() tea.Cmd {
	return nil
}

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		x, y := listView.GetFrameSize()
		m.list.SetSize(msg.Width-x, msg.Height-y)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m logModel) View() string {
	return brutalBorder.Render(m.list.View())
}

package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"sort"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

type ColumnData struct {
	Title string
	Width int
}

func ShowTruthtable(res []map[string]bool) {

	// Extrair as chaves únicas dos mapas
	var keys []string
	keyMap := make(map[string]bool)
	for _, d := range res {
		for k := range d {
			keyMap[k] = true
		}
		break
	}
	for k := range keyMap {
		keys = append(keys, k)
	}

	// ordena por tamanho + ordem alfabetica
	sort.Slice(keys, func(i, j int) bool {
		if len(keys[i]) == len(keys[j]) {
			return keys[i] < keys[j] // Ordenar em ordem alfabética em caso de empate
		}
		return len(keys[i]) < len(keys[j])
	})

	// Criar as colunas da tabela
	var columns []table.Column
	for _, k := range keys {
		columns = append(columns, table.Column{Title: k, Width: max(len(k), 8)})
	}

	// Preencher as linhas da tabela
	var rows []table.Row
	for _, d := range res {
		var row table.Row
		for _, k := range keys {
			row = append(row, fmt.Sprintf("%v", d[k]))
		}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}

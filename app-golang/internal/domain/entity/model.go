package entity

type Model struct {
	Name      string
	MaxTokens int
}

func NewModel(name string, maxToken int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxToken,
	}
}

func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}

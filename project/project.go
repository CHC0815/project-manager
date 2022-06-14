package project

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
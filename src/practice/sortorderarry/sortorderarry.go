package main

import (
	"fmt"
	"sort"
)

type model struct {
	id   int
	name string
}

func (e *model) String() string {
	return fmt.Sprintf("%+v", *e)
}

type arryModel []*model

func (m arryModel) Len() int {
	return len(m)
}

func (m arryModel) Less(i, j int) bool {
	if m[i].id < m[j].id {
		return true
	}
	if m[i].id > m[j].id {
		return false
	}
	return m[i].name < m[j].name
}

func (m arryModel) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	var models []*model
	models = append(models, &model{id: 50, name: "c"})
	models = append(models, &model{id: 10, name: "c"})
	models = append(models, &model{id: 30, name: "d"})
	models = append(models, &model{id: 30, name: "a"})
	models = append(models, &model{id: 30, name: "c"})
	models = append(models, &model{id: 30, name: "b"})
	models = append(models, &model{id: 20, name: "c"})
	models = append(models, &model{id: 70, name: "c"})
	fmt.Printf("%+v\n", models)
	sort.Sort(arryModel(models))
	fmt.Printf("%+v\n", models)
	sort.Sort(sort.Reverse(arryModel(models)))
	fmt.Printf("%+v\n", models)

}

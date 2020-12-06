package main

import (
	"github.com/awalterschulze/gographviz"
	"strconv"
	"strings"
)

const (
	graphName = "G"
)

var (
	nodeAttrs = map[string]string{
		"colorscheme": "rdylgn11",
		"style":       "\"solid,filled\"",
		"fontcolor":   "6",
		"fontname":    "\"Migu 1M\"",
		"color":       "7",
		"fillcolor":   "11",
		"shape":       "circle",
	}
	edgeAttrs = map[string]string{
		"color": "white",
	}
	graphAttrs = map[string]string{
		"bgcolor":   "#343434",
		"layout":    "dot",
		//"labelloc":  "t",
		//"labeljust": "c",
		"rankdir":   "TB",
	}
	ignoreModules = map[string]bool{
		"React":  true,
		"styled": true,
	}
	rankPackages = map[string]bool{
		"controller": true,
		"presenter":  true,
		"gateway":    true,
		"repository": true,
		"entity":     true,
		"interactor": true,
		"server":     true,
	}
)

// Dependency relation between modules
type Dependency struct {
	relation map[string]map[string]bool
	modules  []string
	graph    *gographviz.Graph
}

func newDependency() *Dependency {
	dependency := new(Dependency)
	dependency.relation = make(map[string]map[string]bool)

	g := gographviz.NewGraph()
	if err := g.SetName(graphName); err != nil {
		panic(err)
	}

	// 有向グラフか
	if err := g.SetDir(true); err != nil {
		panic(err)
	}

	for field, value := range graphAttrs {
		if err := g.AddAttr(graphName, field, strconv.Quote(value)); err != nil {
			panic(err)
		}
	}

	dependency.graph = g
	return dependency
}

func getNodeName(name string) string {
	return strconv.Quote(strings.ReplaceAll(name, ".tsx", ""))
}

func getRankNodeAttrs() map[string]string {
	rankNodeAttrs := make(map[string]string)
	for key, value := range nodeAttrs {
		rankNodeAttrs[key] = value
	}
	rankNodeAttrs["shape"] = "doublecircle"
	return rankNodeAttrs
}

func (d *Dependency) addModule(module string) {
	if isIgnore, ok := ignoreModules[module]; ok && isIgnore {
		return
	}
	if _, ok := d.relation[module]; ok {
		return
	}

	moduleNodeName := getNodeName(module)

	d.relation[module] = make(map[string]bool)
	d.modules = append(d.modules, module)

	if isSame, ok := rankPackages[module]; ok && isSame {
		if err := d.graph.AddNode(graphName, moduleNodeName, getRankNodeAttrs()); err != nil {
			panic(err)
		}
		return
	}
	if err := d.graph.AddNode(graphName, moduleNodeName, nodeAttrs); err != nil {
		panic(err)
	}
}

func (d *Dependency) addTo(module, to string) {
	if isIgnore, ok := ignoreModules[to]; ok && isIgnore {
		return
	}
	if _, ok := d.relation[module][to]; ok {
		return
	}

	moduleNodeName := getNodeName(module)
	toNodeName := getNodeName(to)

	d.addModule(to)

	if err := d.graph.AddEdge(moduleNodeName, toNodeName, true, edgeAttrs); err != nil {
		panic(err)
	}
	d.relation[module][to] = true

}

func (d *Dependency) add(module, to string) {
	d.addModule(module)
	d.addTo(module, to)
}

func (d *Dependency) concat(e *Dependency) {
	for _, module := range e.modules {
		for to := range e.relation[module] {
			d.add(module, to)
		}
	}
}

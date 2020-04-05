package main
//cygdrive
// /cygdrive/c/Users/rquon/OneDrive/Desktop/BevStuff/Research/DFXM_Go/DFXM_Go/visuals/go-echarts

//gitbash
// ~/OneDrive/Desktop/BevStuff/Research/DFXM_Go/DFXM_Go/visuals/go-echarts/charts
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"math/rand"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/gobuffalo/packr"
)

var graphNodes = []charts.GraphNode{
	{Name: "Task1",Symbol:"arrow"},
	{Name: "Task2"},
	{Name: "Task3"},
	{Name: "Task4"},
	{Name: "Task5"},
	{Name: "Task6"},
	{Name: "Task7"},
	{Name: "Task8"},
}

func genLinks() []charts.GraphLink {
	links := make([]charts.GraphLink, 0)
	for i := 0; i < len(graphNodes); i++ {
		for j := 0; j < len(graphNodes); j++ {
			if(rand.Intn(101) > 70){
			links = append(links,
				charts.GraphLink{Source: graphNodes[i].Name, Target: graphNodes[j].Name, })
			}
		}
	}
	return links
}

func graphBase() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-Tasks-No-Label"})
	graph.Add("graph", graphNodes, genLinks(),
		charts.GraphOpts{Force: charts.GraphForce{Repulsion: 8000}},
	)
	return graph
}
//ref: ../visuals/go-echarts/charts/graph.go
func graphCircle() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-Tasks-Labeled"})
	graph.Add("Task Dependencies", graphNodes, genLinks(),
		charts.GraphOpts{Layout: "circular", Force: charts.GraphForce{Repulsion: 8000}},
		charts.LabelTextOpts{Show: true, Position: "right"},
	)
	return graph
}

func graphNpmDep() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-npm package 依赖关系"})
	box := packr.NewBox(path.Join(".", "fixtures"))
	f, err := box.Find("npmdepgraph.json")
	if err != nil {
		log.Fatal(err)
	}
	type Data struct {
		Nodes []charts.GraphNode
		Links []charts.GraphLink
	}

	var data Data
	if err := json.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}
	graph.Add("graph", data.Nodes, data.Links,
		charts.GraphOpts{Layout: "none", Roam: true, FocusNodeAdjacency: true},
		charts.EmphasisOpts{Label: charts.LabelTextOpts{Show: true, Position: "left", Color: "black"}},
		charts.LineStyleOpts{Curveness: 0.3},
	)
	return graph
}

func graphHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("graph")...)
	page.Add(
		graphBase(),
		graphCircle(),
		graphNpmDep(),
	)
	f, err := os.Create(getRenderPath("graph.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}

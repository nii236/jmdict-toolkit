package cayley

import (
	"log"

	"github.com/google/cayley"
	"github.com/google/cayley/graph"
)

func writeToCayley() {
	// Initialize the database
	path := "./bolt.db"
	graph.InitQuadStore("bolt", path, nil)
	// Open and use the database
	store, err := cayley.NewGraph("bolt", path, nil)
	// store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}
	store.AddQuad(cayley.Quad("food", "is", "good", ""))
}

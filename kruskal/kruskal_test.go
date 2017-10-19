package main

import "testing"
// import "fmt"

func TestWikiKruskal(t *testing.T) {
  wikiGraph := []Edge{
    Edge{1,2,3},
    Edge{2,3,5},
    Edge{1,5,1},
    Edge{2,5,4},
    Edge{3,5,6},
    Edge{3,4,2},
    Edge{4,5,7},
  }
  actualMinimumSpan := Kruskal(wikiGraph)

  expectedMinimumSpan := []Edge{
    Edge{1,5,1},
    Edge{3,4,2},
    Edge{1,2,3},
    Edge{2,3,5},
  }

  if !areSpansEqual(actualMinimumSpan, expectedMinimumSpan) {
    t.Error("Actual and expected did not match")
  }
}

func areSpansEqual(actual, expected []Edge) bool {
  for i, edge := range expected {
    if(edge != actual[i]) {
      return false
    }
  }
  return true
}

func TestIsACycleFormed(t *testing.T) {
  existingGraph := []Edge{
    Edge{1,2,3},
    Edge{1,3,4},
  }
  additionalEdge := Edge{2,3,4}
  formed := IsACycleFormed(existingGraph, additionalEdge)

  if !formed {
    t.Error("Test failed failed to handle cycle")
  }

  additionalEdge1 := Edge{2,4,2}
  formed = IsACycleFormed(existingGraph, additionalEdge1)

  if formed {
    t.Error("Test failed to handle no cycle")
  }

  existingGraph = []Edge{}
  additionalEdge2 := Edge{4,5,3}
  formed = IsACycleFormed(existingGraph, additionalEdge2)

  if formed {
    t.Error("Test failed to handle empty graph")
  }
}


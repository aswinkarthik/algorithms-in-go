package main

import "testing"

func TestSum(t *testing.T) {
  minimumSpan := Kruskal([]Edge{
    Edge{1,2,3},
    Edge{1,3,2},
    Edge{1,4,4},
    Edge{2,4,5},
  })

  if validEdgeCount(minimumSpan) != 3 {
    t.Error("Test failed")
  }
}

func validEdgeCount(edges []Edge) int {
  count := 0
  for _, edge := range edges {
    if edge.Weight != 0 {
      count += 1
    }
  }
  return count
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


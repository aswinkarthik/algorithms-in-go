package main

import "sort"

func Kruskal(graph []Edge) []Edge {
  minimumSpan := make([]Edge, len(graph))
  msIndex := 0
  sort.Sort(ByWeight(graph))

  //A map of all nodes and if they are visited
  allNodesMap := make(map[int]bool)

  //Collect all nodes to visit and initialize to false
  for _, edge := range graph {
    allNodesMap[edge.Start] = false
    allNodesMap[edge.End] = false
  }

  for _, edge := range graph {
    allNodesMap[edge.Start] = true
    allNodesMap[edge.End] = true
    if(IsACycleFormed(minimumSpan, edge)) {
      continue
    }
    minimumSpan[msIndex] = edge
    msIndex += 1
    if(areAllNodesVisited(allNodesMap)) {
      break
    }
  }

  return minimumSpan
}

func areAllNodesVisited(allNodesMap map[int]bool) bool {
  for _, visited := range allNodesMap {
    if(!visited) {
      return false
    }
  }
  return true
}

func IsACycleFormed(existingGraph []Edge, additionalEdge Edge) bool {
  graphToTest := existingGraph

  allNodes := make(map[int]bool)
  for _, edge := range graphToTest {
    allNodes[edge.Start] = true
    allNodes[edge.End] = true
  }

  return allNodes[additionalEdge.End] && allNodes[additionalEdge.Start]
}

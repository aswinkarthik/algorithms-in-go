package main

import "sort"

func Kruskal(graph []Edge) []Edge {
  minimumSpan := make([]Edge, len(graph))
  edgeCount := 0
  sort.Sort(ByWeight(graph))

  nodeCount := countOfNodes(graph)

  for _, edge := range graph {
    if(IsACycleFormed(minimumSpan[:edgeCount], edge)) {
      continue
    }
    minimumSpan[edgeCount] = edge
    edgeCount += 1
    if done(edgeCount, nodeCount) {
      break
    }
  }
  return minimumSpan[:edgeCount]
}

func countOfNodes(graph []Edge) int {
  allNodesMap := make(map[int]bool)

  for _, edge := range graph {
    allNodesMap[edge.Start] = false
    allNodesMap[edge.End] = false
  }

  return len(allNodesMap)
}

func IsACycleFormed(existingGraph []Edge, additionalEdge Edge) bool {
  graphToTest := append(existingGraph, additionalEdge)

  max := -1
  for _, edge := range graphToTest {
    if edge.Start > max {
      max = edge.Start
    }
    if edge.End > max {
      max = edge.End
    }
  }
  parent := make([]int,max + 1)

  for i := 0; i < max + 1 ; i++ {
    parent[i] = -1
  }

  for _, edge := range graphToTest {
    startRoot := findParent(parent, edge.Start)
    endRoot := findParent(parent, edge.End)
    if startRoot == endRoot {
      return true
    }
    parent[endRoot] = startRoot
  }
  return false
}

func done(edgeCount, nodeCount int) bool {
  return edgeCount >= nodeCount
}

func findParent(parent []int, node int) int {
  if parent[node] == -1 {
    return node
  }
  return findParent(parent, parent[node])
}

package kruskal

//Structure of an Edge and is sortable by weight
type Edge struct {
  Start int
  End int
  Weight int
}

type ByWeight []Edge

func (e ByWeight) Len() int {
  return len(e);
}

func (e ByWeight) Swap(i, j int) {
  e[i], e[j] = e[j], e[i]
}

func (e ByWeight) Less(i, j int) bool {
  return e[i].Weight < e[j].Weight
}

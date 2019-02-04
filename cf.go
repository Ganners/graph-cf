package cf

import (
	"container/list"
	"errors"
	"log"
	"sort"
)

type Rating struct {
	User   string
	Item   string
	Rating int
}

type Ratings []Rating

func (r Ratings) Len() int           { return len(r) }
func (r Ratings) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Ratings) Less(i, j int) bool { return r[i].Rating < r[j].Rating }

type GraphCF struct {
	userToNode map[string]*Node
	itemToNode map[string]*Node
}

type Node struct {
	Type       nodeType
	Identifier string
	edges      []*Node
}

func (n *Node) AddEdge(child *Node) {
	n.edges = append(n.edges, child)
}

type nodeType uint8

const (
	nodeTypeUser nodeType = 1 << iota
	nodeTypeItem
)

func NewSimpleGraphCF() *GraphCF {
	return &GraphCF{
		userToNode: make(map[string]*Node, 1000),
		itemToNode: make(map[string]*Node, 1000),
	}
}

func (cf *GraphCF) AddRatings(ratings Ratings) {
	for _, r := range ratings {
		// don't store 0 ratings
		if r.Rating <= 0 {
			continue
		}

		// look for user node
		userNode, ok := cf.userToNode[r.User]
		if !ok {
			// if it doesn't exist already, create and link
			userNode = &Node{
				Type:       nodeTypeUser,
				Identifier: r.User,
				edges:      []*Node{},
			}
			cf.userToNode[r.User] = userNode
		}

		// look for item node
		itemNode, ok := cf.itemToNode[r.Item]
		if !ok {
			// if it doesn't exist already, create and link
			itemNode = &Node{
				Type:       nodeTypeItem,
				Identifier: r.Item,
				edges:      []*Node{},
			}
			cf.itemToNode[r.Item] = itemNode
		}

		// create bidirectional links
		userNode.AddEdge(itemNode)
		itemNode.AddEdge(userNode)
	}
}

func (cf *GraphCF) UserTopK(user string, k int) (Ratings, error) {
	root, ok := cf.userToNode[user]
	if !ok {
		return Ratings{}, errors.New("user does not exist in graph")
	}

	hist := cf.histBreadthFirst(root)
	log.Println(hist)

	// delete observed entries
	for _, child := range root.edges {
		delete(hist, child.Identifier)
	}

	ratings := Ratings{}
	for identifier, score := range hist {
		ratings = append(ratings, Rating{
			User:   root.Identifier,
			Item:   identifier,
			Rating: score,
		})
	}

	// sort reverse
	sort.Sort(sort.Reverse(ratings))

	k = min(len(ratings), k)
	return ratings[:k], nil
}

func (cf *GraphCF) histBreadthFirst(node *Node) map[string]int {
	maxVisits := 5
	seen := make(map[*Node]int)
	hist := make(map[string]int)

	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		node = queue.Remove(queue.Front()).(*Node)

		// if it's an already visited user node then skip
		if (node.Type & nodeTypeUser) > 0 {
			if _, ok := seen[node]; ok {
				continue
			} else {
				seen[node] = 1
			}
		}
		// if it's an item that has been visited more than the
		// maximum number of times then skip
		if (node.Type & nodeTypeItem) > 0 {
			if visitCount, _ := seen[node]; visitCount >= maxVisits {
				continue
			} else {
				seen[node] = visitCount + 1
			}
		}

		prefix := ""
		if (node.Type & nodeTypeItem) > 0 {
			prefix = "-"
		}

		log.Println(prefix + node.Identifier)

		// queue up children, we'll take care of seen and not in
		// this bit
		for _, child := range node.edges {
			// queue
			queue.PushBack(child)
		}

		if (node.Type & nodeTypeItem) > 0 {
			if _, ok := hist[node.Identifier]; !ok {
				hist[node.Identifier] = 0
			}
			hist[node.Identifier] += 1
		}
	}

	return hist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

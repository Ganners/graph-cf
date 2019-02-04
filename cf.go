package cf

import (
	"container/list"
	"errors"
	"sort"
)

// Rating represents a user -> item mapping with a float rating
type Rating struct {
	User   string
	Item   string
	Rating float64
}

// Ratings represents a set of rating triples
type Ratings []Rating

// By default we will sort by Rating
func (r Ratings) Len() int           { return len(r) }
func (r Ratings) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Ratings) Less(i, j int) bool { return r[i].Rating < r[j].Rating }

// nodeType is a bitwise flag which specifies if the node is
// either for a user or item. It is bitwise because there are
// alternatives to try out later with various latent features
type nodeType uint8

const (
	nodeTypeUser nodeType = 1 << iota
	nodeTypeItem
)

// Node is a graph node which can be an item or a user
type Node struct {
	Type       nodeType
	Identifier string
	edges      []*Node
}

// AddEdge adds an edge from this to another node
func (n *Node) AddEdge(child *Node) {
	n.edges = append(n.edges, child)
}

// GraphCF wraps the primary model. Links to the nodes are
// referenced through the two maps
type GraphCF struct {
	userToNode map[string]*Node
	itemToNode map[string]*Node
}

func NewSimpleGraphCF() *GraphCF {
	return &GraphCF{
		userToNode: make(map[string]*Node, 1000),
		itemToNode: make(map[string]*Node, 1000),
	}
}

func (cf *GraphCF) AddRatings(ratings Ratings) {
	for _, r := range ratings {
		// don't store 0 ratings
		if r.Rating <= 0. {
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

// UserTopK will, given a user's identifier, pull the top K
// ratings (up to that number), it excludes observed ratings
func (cf *GraphCF) UserTopK(user string, k int) (Ratings, error) {
	root, ok := cf.userToNode[user]
	if !ok {
		return Ratings{}, errors.New("user does not exist in graph")
	}

	// compute histogram of items and their scores
	hist := cf.histBreadthFirst(root)

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

// histBreadthFirstSearch will return a histogram of Indentifiers
// to their exponentially decreasing inverse difference
//
// This is currently an unbound traversal, limits could be placed
// to prevent it from walking each node
func (cf *GraphCF) histBreadthFirst(node *Node) map[string]float64 {
	seen := make(map[*Node]struct{})
	hist := make(map[string]float64)

	queue := list.New()
	queue.PushBack(node)
	generation := nodeTypeUser
	depth := 1.

	for queue.Len() > 0 {
		node = queue.Remove(queue.Front()).(*Node)

		if (node.Type & generation) == 0 {
			generation = node.Type
			depth /= 2
		}

		if (node.Type & nodeTypeItem) > 0 {
			if _, ok := hist[node.Identifier]; !ok {
				hist[node.Identifier] = 0
			}
			hist[node.Identifier] += depth
		}

		// if it's an already visited node then skip
		if _, ok := seen[node]; ok {
			continue
		}
		seen[node] = struct{}{}

		// queue up children, we'll take care of seen and not in
		// this bit
		for _, child := range node.edges {
			// queue
			queue.PushBack(child)
		}
	}

	return hist
}

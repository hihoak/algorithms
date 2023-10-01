package dijkstras_algorithm

type NodeID string
type EdgeValue int

type Graph struct {
	Nodes map[NodeID]interface{}
	Edges map[NodeID]map[NodeID]EdgeValue
}

type queue struct {
	head *queueNode
	tail *queueNode
}

func newQueue(head *queueNode) *queue {
	return &queue{
		head: head,
		tail: head,
	}
}

func (q *queue) pop() *queueNode {
	if q.head == nil {
		return nil
	}
	res := q.head
	q.head = res.next
	if res == q.tail {
		q.tail = res.next
	}
	return res
}

func (q *queue) add(node NodeID) {
	qNode := &queueNode{
		nodeID: node,
	}
	if q.tail == nil {
		q.head, q.tail = qNode, qNode
		return
	}
	q.tail.next = qNode
	q.tail = qNode
}

type queueNode struct {
	nodeID NodeID
	next   *queueNode
}

func newQueueNode(nodeID NodeID, next *queueNode) *queueNode {
	return &queueNode{
		nodeID: nodeID,
		next:   next,
	}
}

func DijkstrasAlgorithm(graph Graph, startNode NodeID) map[NodeID]EdgeValue {
	q := newQueue(newQueueNode(startNode, nil))

	resultDistances := make(map[NodeID]EdgeValue, len(graph.Nodes))
	resultDistances[startNode] = 0
	for currentNode := q.pop(); currentNode != nil; currentNode = q.pop() {
		currentDistance := resultDistances[currentNode.nodeID]
		neighbors := graph.Edges[currentNode.nodeID]
		for neighbor, distance := range neighbors {
			distanceToNeighbor := currentDistance + distance
			if knownDistance, ok := resultDistances[neighbor]; !ok || distanceToNeighbor < knownDistance {
				resultDistances[neighbor] = distanceToNeighbor
				q.add(neighbor)
			}
		}
	}

	return resultDistances
}

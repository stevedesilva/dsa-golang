package vertex

import (
	"errors"
	"fmt"
	"github.com/stevedesilva/dsa-golang.git/commonsense/queue"
)

type Vertex[T comparable] struct {
	value            T
	adjacentVertices []*Vertex[T]
}

func NewVertex[T comparable](value T) *Vertex[T] {
	return &Vertex[T]{
		value: value,
	}
}

func (v *Vertex[T]) AddAdjacentVertex(adjacentVertex *Vertex[T]) {
	// if slice contains adjacentVertex, return
	for _, vertex := range v.adjacentVertices {
		if vertex == adjacentVertex {
			return
		}
	}
	// add to this vertex slice
	v.adjacentVertices = append(v.adjacentVertices, adjacentVertex)

	// add to calling vertex slice
	adjacentVertex.AddAdjacentVertex(v)
}

func (v *Vertex[T]) DfsTraverse(vertex *Vertex[T]) {
	visited := make(map[*Vertex[T]]bool)
	depthFirstSearchTraverse(vertex, visited)
}

func depthFirstSearchTraverse[T comparable](vertex *Vertex[T], visited map[*Vertex[T]]bool) {
	// already visited
	visited[vertex] = true
	// print
	fmt.Printf("%v\n", vertex.value)
	// loop over adjacent vertx
	for _, v := range vertex.adjacentVertices {
		if _, ok := visited[v]; ok {
			continue
		} else {
			depthFirstSearchTraverse(v, visited)
		}
	}
}

func (v *Vertex[T]) DepthFirstSearch(value T) (*Vertex[T], error) {
	// get root value
	return dfs(v, make(map[*Vertex[T]]bool), value)
}

func dfs[T comparable](vertx *Vertex[T], visited map[*Vertex[T]]bool, value T) (*Vertex[T], error) {
	visited[vertx] = true
	if vertx.value == value {
		return vertx, nil
	}
	for _, v := range vertx.adjacentVertices {
		if _, ok := visited[v]; ok {
			continue
		} else {
			v2, _ := dfs(v, visited, value)
			if v2 != nil {
				// return if found
				return v2, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (v *Vertex[T]) BreadthFirstSearch(value T) (*Vertex[T], error) {
	q := queue.New[*Vertex[T]](v)
	q.Enqueue(v)
	visited := make(map[*Vertex[T]]bool)
	visited[v] = true
	for q.Size() > 0 {
		current, err := q.Dequeue()
		if err != nil {
			return nil, err
		}
		if current.value == value {
			return current, nil
		}
		for _, vertx := range current.adjacentVertices {
			// add if we haven't already processed
			if _, ok := visited[vertx]; !ok {
				q.Enqueue(vertx)
				visited[vertx] = true
			}
		}
	}
	return nil, errors.New("not found")
}

func (v *Vertex[T]) BreadthFirstSearchTraversal() error {
	q := queue.New[*Vertex[T]](v)
	visited := make(map[*Vertex[T]]bool)
	visited[v] = true
	for q.Size() > 0 {
		current, err := q.Dequeue()
		if err != nil {
			return err
		}
		fmt.Println(current.value)
		for _, vertx := range current.adjacentVertices {
			// add if we haven't already processed
			if _, ok := visited[vertx]; !ok {
				q.Enqueue(vertx)
				visited[vertx] = true
			}
		}
	}
	return nil
}

/*
public List<T> shortestPath(Vertex<T> start, Vertex<T> end) {
        // visited nodes
        Set<Vertex<T>> visitedVertices = new HashSet<>();
        visitedVertices.add(start);
        // queue
        Queue<Vertex<T>> queue = new LinkedList<>();
        queue.add(start);

        Map<Vertex<T>, Vertex<T>> previousVertex = new HashMap<>();
        // while queue is not empty
        while (!queue.isEmpty()) {
            Vertex<T> currentVertex = queue.remove();
            for (Vertex<T> adjacentVertx : currentVertex.edges) {
                if (!visitedVertices.contains(adjacentVertx)) {
                    visitedVertices.add(adjacentVertx);
                    queue.add(adjacentVertx);

                    previousVertex.put(adjacentVertx, currentVertex);
                }
            }
        }

        List<T> shortestPath = new ArrayList<>();
        Vertex<T> current = end;
        while (current != start) {
            shortestPath.add(current.value);
            current = previousVertex.get(current);
        }
        shortestPath.add(start.value);
        return shortestPath.reversed();
    }
*/

func (v *Vertex[T]) shortestPath(start, end *Vertex[T]) []string {
	// first to reach the end node wins - since it would be added to the visited node table
	// setup
	//	visited nodes
	//  queue
	//  previous node table
	visitedVertices := make(map[*Vertex[T]]bool)
	visitedVertices[start] = true

	queue := queue.New[*Vertex[T]](start)
	queue.Enqueue(start)

	//previousVertex := make(map[string]string)
	previousVertex := make(map[T]T)
	// while queue is not empty
	for queue.Size() > 0 {
		current, err := queue.Dequeue()
		if err != nil {
			return nil
		}

		for _, vertx := range current.adjacentVertices {
			if _, ok := visitedVertices[vertx]; !ok {
				queue.Enqueue(vertx)
				visitedVertices[vertx] = true
				previousVertex[vertx.value] = current.value
			}
		}
	}

	// reverse  cheapestPreviousStopoverCity and current
	//result := make([]string, 0, len(cheapestPreviousStopoverCity))
	//currentCityName := destinationCity.Name
	//for currentCityName != startCity.Name {
	//	result = append(result, currentCityName)
	//	currentCityName = cheapestPreviousStopoverCity[currentCityName]
	//}
	//result = append(result, startCity.Name)
	//// reverse result
	//for i := 0; i < len(result)/2; i++ {
	//	j := len(result) - i - 1
	//	result[i], result[j] = result[j], result[i]
	//}
	//
	//return result
	return nil
}

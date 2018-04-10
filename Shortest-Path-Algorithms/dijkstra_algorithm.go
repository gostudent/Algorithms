//Program in Go to implement Djikstra's Algorithm
package main

import "fmt"

type Edge struct {
	first_node_id int;
	second_node_id int;
	edge_length int;
}

type Graph struct {
	edges []Edge;
	nodes []Node;
}

type Node struct {
	id int;
	temp_val int;
	temp_path []int;
}

//Takes an input graph, starting node.id, ending node.id
func shortest_path(graph Graph, start int, end int) int {
	unvisited_nodes := graph.nodes
	var visited_nodes []Node

	start_node := Node{start, 0, []int{start}}
	//Start with input start node
	visited_nodes = append(visited_nodes, start_node)
	unvisited_nodes = remove(unvisited_nodes, start_node.id)
	end_flag := false
	new_neighbors := update_neighbors(graph,visited_nodes)
	minimum := -1
	for end_flag == false {
		minimum = -1
		new_neighbors = update_neighbors(graph,visited_nodes)
		unvisited_nodes = update_unvisited_nodes(unvisited_nodes, new_neighbors)
		//Choose New node to visit
		for _, node := range new_neighbors {
			if minimum == -1 {
				minimum = node.temp_val
				start_node = node
			}else if minimum > node.temp_val{
				minimum = node.temp_val
				start_node = node
			}
		}
		visited_nodes = append(visited_nodes, start_node)
		unvisited_nodes = remove(unvisited_nodes, start_node.id)
		if start_node.id == end {
			end_flag = true
		}
	}
	//And thus the start node has become the end node
	//fmt.Println(start_node.temp_path)
	return start_node.temp_val
}

func remove(slice []Node, s int) []Node {
	var my_index int
	for index, element := range slice {
		if element.id == s {
			my_index = index
		}
	}
    return append(slice[:my_index], slice[my_index+1:]...)
}

//Takes input Nodes and calculates all nieghbors and distances
func update_neighbors(graph Graph, start_nodes []Node) []Node {
	var neighbors []Node
	for _, start := range start_nodes{
		my_neighbors := get_neighbors(graph,start, start_nodes)
		//Update Neighbor Values
		for index, element := range my_neighbors {
			my_neighbors[index].temp_val = element.temp_val + start.temp_val
			my_neighbors[index].temp_path = append(my_neighbors[index].temp_path, my_neighbors[index].id)
		}
		//Compare to all neighbors
		for _, neigh := range my_neighbors {
			found := false
			for aindex, aneigh := range neighbors {
				if neigh.id == aneigh.id {
					found = true
					if neigh.temp_val < aneigh.temp_val {
						neighbors[aindex].temp_val = neigh.temp_val
						neighbors[aindex].temp_path = neigh.temp_path
					}
				}
			}
			if found == false {
				neighbors = append(neighbors, neigh)
			}
		}
	}
	return neighbors
}

//Takes input unvisited_nodes and new neighbors and updates their temp values
func update_unvisited_nodes(unvisited_nodes []Node, new_neighbors []Node) []Node {
	for index, element := range unvisited_nodes {
		for _, neighbor := range new_neighbors {
			if neighbor.id == element.id {
				if neighbor.temp_val < element.temp_val || element.temp_val == -1 {
					unvisited_nodes[index].temp_val = neighbor.temp_val
					unvisited_nodes[index].temp_path = neighbor.temp_path
				}
			}
		}
	}
	return unvisited_nodes
}

//Takes an input Node and calculates distance to immediate neighbors
func get_neighbors(graph Graph, start Node, visited_nodes []Node) []Node {
	var neighbors []Node
	for _, element := range graph.edges {
		my_first_node_id := element.first_node_id
		my_second_node_id := element.second_node_id
		first_node_visited := false
		second_node_visited := false
		for _, node := range visited_nodes {
			if my_first_node_id == node.id {
				first_node_visited = true
			}
			if my_second_node_id == node.id{
				second_node_visited = true
			}
		}
		if my_first_node_id == start.id && second_node_visited == false {
			found := false
			for _, element2 := range neighbors {
				if element2.id == my_second_node_id {
					found = true
				}
			}
			if found != true {
				neighbors = append(neighbors, Node{my_second_node_id,element.edge_length,start.temp_path})
			}
		}
		if my_second_node_id == start.id && first_node_visited == false{
			found := false
			for _, element2 := range neighbors {
				if element2.id == my_first_node_id {
					found = true
				}
			}
			if found != true {
				neighbors = append(neighbors, Node{my_first_node_id,element.edge_length,start.temp_path})
			}
		}
	}
	return neighbors
}

func main() {
	var edges_list []Edge
	edge_1_to_2 := Edge{1, 2, 1}
	edge_2_to_3 := Edge{2, 3, 5}
	edge_1_to_3 := Edge{1, 3, 5}
	edge_2_to_4 := Edge{2, 4, 1}
	edge_2_to_5 := Edge{2, 5, 5}
	edge_3_to_4 := Edge{3, 4, 5}
	edge_4_to_5 := Edge{4, 5, 5}
	edges_list = append(edges_list, edge_1_to_2)
	edges_list = append(edges_list, edge_2_to_3)
	edges_list = append(edges_list, edge_1_to_3)
	edges_list = append(edges_list, edge_2_to_4)
	edges_list = append(edges_list, edge_2_to_5)
	edges_list = append(edges_list, edge_3_to_4)
	edges_list = append(edges_list, edge_4_to_5)
	//fmt.Println(edges_list)
	var node_list []Node
	node_list = append(node_list, Node{1,-1,[]int{0}})
	node_list = append(node_list, Node{2,-1,[]int{0}})
	node_list = append(node_list, Node{3,-1,[]int{0}})
	node_list = append(node_list, Node{4,-1,[]int{0}})
	node_list = append(node_list, Node{5,-1,[]int{0}})
	graph := Graph{edges_list, node_list}
	//fmt.Println(graph)
	fmt.Println(shortest_path(graph,1,5))
}

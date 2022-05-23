package solver

var NUM_JOBS uint = 4

var Jobs = make(chan *Astar, NUM_JOBS)
var Results = make(chan *Node, NUM_JOBS)

func exWorker(id int, jobs <-chan *Astar, results chan<- *Node) {
	for j := range jobs {
		if Open_List.list == nil {
			results <- nil
			return
		}
		if Open_List.Size() == 0 {
			results <- nil
			return
		}
		node := Open_List.DeleteMin()

		state, ok := node.(*Node)
		if !ok {
			results <- nil
			return
		}

		uuid := state.State.CreateUuid()

		if node.(*Node).H == 0 {

			results <- node.(*Node)
			return
			// return node.(*Node), nil
		}
		node.(*Node).Execute(j)
		num := Open_List.Size()
		if num > int(j.MaxState) {
			j.MaxState = uint(num)
		}
		Closed_List.Insert(uuid)
		results <- nil
	}
}

func init() {

	// create workers

	for w := 0; w < int(NUM_JOBS); w++ {
		go exWorker(w, Jobs, Results)
	}
}

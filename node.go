package gribbon

type node struct {
	data *worker
	next *node
}

func newNode(d *worker, n *node) *node {
	return &node{data: d, next: n}
}

func (n *node) insertTail(nd *node) {
	n.next = nd
}

func (n *node) isWorking() bool {
	return n.data.isWorking()
}

func (n *node) submit(t *task) {
	n.data.submit(t)
}

func (n *node) run() {
	n.data.run()
}

func (n *node) runWithArg() {
	n.data.runWithArg()
}

func (n *node) close() {
	n.data.close()
}
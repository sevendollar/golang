package queue

type queue []string

func queueInit() queue {
        return []string{}
}

func (q *queue) push(item string) {
        (*q) = append((*q), "")
        copy((*q)[1:], (*q)[:])
        (*q)[0] = item

}

func (q *queue) pop() (o string) {
        if len((*q)) > 0 {
                o, (*q) = (*q)[len((*q))-1], (*q)[:len((*q))-1]
                return
        }
        return "nil"
}

type stack []string

func stackInit() stack {
        return []string{}
}

func (q *stack) push(item string) {
        (*q) = append((*q), item)
}

func (q *stack) pop() (o string) {
        if len(*q) > 0 {
                o, (*q) = (*q)[len(*q)-1], (*q)[:len(*q)-1]
                return
        }
        return "nil"
}


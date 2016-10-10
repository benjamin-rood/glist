package glist

func (l *Listy) Head() G {
	return l.headptr.g
}

func (l *Listy) Tail() G {
	return l.tailptr.g
}

func (l *Listy) Length() G {
	return l.length
}

func (l *Listy) Append(g G) {
	l.tailptr = &Node{
		g:    g,
		next: nil,
		prev: l.tailptr,
	}
	l.length++
}

func (l *Listy) AppendHead(g G) {
	l.headptr = &Node{
		g:    g,
		next: l.headptr,
		prev: nil,
	}
	l.length++
}

func (l *Listy) Pop() G {
	if l == nil {
		return nil
	}
	value := l.tailptr.g
	l.tailptr = l.tailptr.prev
	l.tailptr.next = nil
	l.length--
	return value
}

func (l *Listy) PopHead() G {
	if l == nil {
		return nil
	}
	value := l.headptr.g
	l.headptr = l.headptr.next
	l.tailptr.prev = nil
	l.length--
	return value
}

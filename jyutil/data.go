package jyutil

// ==============================================
// CList[T]
// ==============================================
type CList[T comparable] struct {
	list []T
}

func (c *CList[T]) GetData() []T {
	return c.list
}

func (c *CList[T]) Set(list []T) {
	c.list = make([]T, len(list))
	copy(c.list, list)
}

func (c *CList[T]) Add(data T) int {
	c.list = append(c.list, data)
	return len(c.list)
}

func (c *CList[T]) GetSize() int {
	return len(c.list)
}

func (c *CList[T]) GetAt(idx int) any {
	return c.list[idx]
}

func (c *CList[T]) SetAt(idx int, data T) {
	c.list[idx] = data
}

func (c *CList[T]) Del(idx int) bool {
	size := c.GetSize()

	if size <= 0 {
		Output("CList::Del - List is empty!")
		return false
	} else if idx < 0 || idx >= size {
		OutputDebugString("CList::Del - Invalid index! ( DelIdx: %d, Size: %d )\n", idx, size)
		return false
	}

	//
	c.list = append(c.list[:idx], c.list[idx+1:]...)

	return true
}

func (c *CList[T]) DelValue(dtDel T) bool {
	if c.GetSize() <= 0 {
		Output("CList::DelValue - List is empty!")
		return false
	}

	//
	listTmp := make([]T, len(c.list))
	copy(listTmp, c.list)
	c.list = make([]T, 0)

	for _, data := range listTmp {
		if data != dtDel {
			c.list = append(c.list, data)
		}
	}

	return true
}

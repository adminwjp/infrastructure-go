package utils

type Collection struct {
	set       map[interface{}]bool
	array     []interface{}
	isSet     bool
	Increment int
	Count     int
	size      int
}

func NewList() Collection {
	collection := Collection{isSet: false, Increment: 10,
		size: 10, Count: 0,
		array: make([]interface{}, 10)}
	return collection
}
func NewSet() Collection {
	collection := Collection{isSet: true, Increment: 10,
		Count: 0, size: 10, set: make(map[interface{}]bool, 10)}
	return collection
}
func (collection Collection) changeSize() {
	if collection.Count == collection.size {
		var size = collection.size + collection.Increment
		if !collection.isSet {
			var temp = make([]interface{}, collection.size)
			for i := 0; i < collection.size; i++ {
				temp[i] = collection.array[i]
			}
			collection.array = temp
			collection.size = size
		} else {
			var temp = make(map[interface{}]bool, collection.size)
			for k, v := range collection.set {
				temp[k] = v
			}
			collection.set = temp
			collection.size = size
		}
	}
}
func (collection Collection) Add(val interface{}) bool {
	collection.changeSize()
	if collection.isSet {
		_, ok := collection.set[val]
		if !ok {
			collection.set[val] = true
			collection.Count++
		}
		return !ok
	} else {
		collection.array[collection.Count] = val
		collection.Count++
		return true
	}
}
func (collection Collection) Remove(val interface{}) bool {
	if collection.isSet {
		_, ok := collection.set[val]
		if ok {
			collection.set[val] = false
			collection.Count--
			return true
		}
		return false
	} else {
		var remove = false
		for i := range collection.array {
			if collection.array[i] == val {
				remove = true
				if i != collection.Count {
					collection.array[i-1] = collection.array[i]
				}
				collection.Count--
			}
		}
		return remove
	}
}
func (collection Collection) ToArray() []interface{} {
	if !collection.isSet && collection.Count+1 == collection.size {
		return collection.array
	}
	var temp = make([]interface{}, collection.Count+1)
	if collection.isSet {
		var i = 0
		for k, v := range collection.set {
			if v {
				temp[i] = k
				i++
			}
		}
	} else {
		for i := 0; i <= collection.Count; i++ {
			temp[i] = collection.array[i]
		}
	}
	return temp
}

type HashMap struct {
	keys []interface{}
	vals []interface{}
}

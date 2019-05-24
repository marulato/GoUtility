package collection

type ArrayList struct {
	size     int
	capacity int
	content  []interface{}
}

func (array *ArrayList) init() *ArrayList {
	array.size = 0
	array.content = make([]interface{}, 0)
	return array
}

func New() *ArrayList {
	return new(ArrayList).init()
}

func (array *ArrayList) Size() int {
	return array.size
}

func (array *ArrayList) Add(obj interface{}) {
	array.content = append(array.content, obj)
	array.size++
}

func (array *ArrayList) Get(index int) interface{} {
	if index >= 0 && index < array.size {
		return array.content[index]
	}
	return nil
}

func (array *ArrayList) IsEmpty() bool {
	return array.size == 0
}

func (array *ArrayList) Remove(index int) {
	if index >= 0 && index < array.size {
		frontSlice := array.content[0:index]
		backSlice := array.content[index+1 : array.size]
		array.content = frontSlice
		array.content = append(array.content, backSlice...)
		array.size--
	}
}

func (array *ArrayList) LastIndexOf(obj interface{}) int {
	if array.size > 0 {
		lastIndex := -1
		for idx, ctn := range array.content {
			if obj == ctn {
				lastIndex = idx
			}
		}
		return lastIndex
	}
	return -1
}

func (array *ArrayList) Contains(obj interface{}) bool {
	return array.LastIndexOf(obj) >= 0
}

func (array *ArrayList) AddAll(objList *ArrayList) {
	if objList != nil && objList.Size() > 0 {
		array.content = append(array.content, objList.content...)
		array.size += objList.Size()
	}
}

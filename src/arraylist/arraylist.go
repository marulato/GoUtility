package arraylist

type ArrayList struct {
	Size int
	capacity int
	content []interface{}
}

func (array *ArrayList) init() *ArrayList {
	array.Size = 0
	array.content = make([]interface{}, 0)
	return array
}

func New() *ArrayList {
	return new(ArrayList).init()
}

func (array *ArrayList) Add(obj interface{}) {
	array.content = append(array.content, obj)
	array.Size ++
}

func (array *ArrayList) Get(index int) interface{} {
	if index >= 0 && index < array.Size {
		return array.content[index]
	}
	return nil
}

func (array *ArrayList) IsEmpty() bool  {
	return array.Size == 0
}


func (array *ArrayList) Remove(index int) {
	if index >= 0 && index < array.Size {
		frontSlice := array.content[0 : index]
		backSlice := array.content[index + 1 : array.Size]
		array.content = frontSlice
		array.content = append(array.content, backSlice...)
		array.Size --
	}
}

func (array *ArrayList) LastIndexOf(obj interface{}) int {
	if array.Size > 0 {
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

func (array *ArrayList) AddAll (objList []interface{}) {

}
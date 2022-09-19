package utils
type Table struct {
	Table string `json:"table"`
	Class string `json:"class"`
	Comment string `json:"comment"`
	Key string `json:"key"`
	RefKey string `json:"ref_key"`
	Columns []*Column `json:"columns"`

}


type Column struct {
	Column string `json:"column"`
	ProName string `json:"pro_name"`
	Length int `json:"length"`
	Type string `json:"type"`
	Comment string `json:"comment"`
	Default string `json:"default"`
	Orders int `json:"orders"`
}

type TplUtil struct {

}
type TableList struct {
	set       map[*Table]bool
	array     []*Table
	isSet     bool
	Increment int
	Count     int
	size      int
	orders bool
}

func NewTableList() TableList {
	collection := TableList{isSet: false, Increment: 10,
		size: 10, Count: 0,
		array: make([]*Table, 10)}
	return collection
}
func NewTableSet() TableList {
	collection := TableList{isSet: true, Increment: 10,
		Count: 0, size: 10, set: make(map[*Table]bool, 10)}
	return collection
}
func (tableList TableList) changeSize() {
	if tableList.Count == tableList.size {
		var size = tableList.size + tableList.Increment
		if !tableList.isSet {
			var temp = make([]*Table, tableList.size)
			for i := 0; i < tableList.size; i++ {
				temp[i] = tableList.array[i]
			}
			tableList.array = temp
			tableList.size = size
		} else {
			var temp = make(map[*Table]bool, tableList.size)
			for k, v := range tableList.set {
				temp[k] = v
			}
			tableList.set = temp
			tableList.size = size
		}
	}
}
func (tableList TableList) Sort(){
	if tableList.orders{
		return
	}
	if tableList.isSet {
		var i = 0
		for k, v := range tableList.set {
			if v {
				if k.Columns!=nil{
					for j := 0; j < len(k.Columns); j++ {
						if k.Columns[j].Orders==0{
							k.Columns[j].Orders=j
							continue
						}else if k.Columns[j].Orders>0{
							if(k.Columns[j].Orders>=len(k.Columns)){
								continue
							}
							 if k.Columns[k.Columns[j].Orders].Orders< k.Columns[j].Orders{
								var temp=k.Columns[j]
								k.Columns[k.Columns[j].Orders]=temp
								k.Columns[j]=k.Columns[k.Columns[j].Orders]
							}
							continue
						}else if k.Columns[j].Orders<0{
							continue
						}
					}
				}
				i++
			}
		}
	} else {
		for i := 0; i <= tableList.Count; i++ {
			var table=tableList.array[i]
			for j := 0; j < len(table.Columns); j++ {
				if table.Columns[j].Orders==0{
					continue
				}else if table.Columns[j].Orders>0{

					continue
				}else if table.Columns[j].Orders<0{
					continue
				}
			}
		}
	}
	tableList.orders=true
}
func (tableList TableList) getIndex(columns []*Column, orders int){
	for j := 0; j < len(columns); j++ {
		if columns[j].Orders==0{
			continue
		}else if columns[j].Orders>0{

			continue
		}else if columns[j].Orders<0{
			continue
		}
	}
}
func (tableList TableList) Add(val *Table) bool {
	tableList.changeSize()
	if tableList.isSet {
		_, ok := tableList.set[val]
		if !ok {
			tableList.set[val] = true
			tableList.Count++
		}
		return !ok
	} else {
		tableList.array[tableList.Count] = val
		tableList.Count++
		return true
	}
}
func (tableList TableList) Remove(val *Table) bool {
	if tableList.isSet {
		_, ok := tableList.set[val]
		if ok {
			tableList.set[val] = false
			tableList.Count--
			return true
		}
		return false
	} else {
		var remove = false
		for i := range tableList.array {
			if tableList.array[i] == val {
				remove = true
				if i != tableList.Count {
					tableList.array[i-1] = tableList.array[i]
				}
				tableList.Count--
			}
		}
		return remove
	}
}
func (tableList TableList) ToArray() []*Table {
	if !tableList.isSet && tableList.Count+1 == tableList.size {
		return tableList.array
	}
	var temp = make([]*Table, tableList.Count+1)
	if tableList.isSet {
		var i = 0
		for k, v := range tableList.set {
			if v {
				temp[i] = k
				i++
			}
		}
	} else {
		for i := 0; i <= tableList.Count; i++ {
			temp[i] = tableList.array[i]
		}
	}
	return temp
}


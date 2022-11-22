package counternumbers

type Allowed interface {
	int | []any
}

type ElementI[T Allowed] interface {
	get() T
	isArray() bool
}

type IntegerEl struct {
	Value int
}

func (a *IntegerEl) get() int {
	return a.Value
}

func (a *IntegerEl) isArray() bool {
	return false
}

type ArrayEl[T Allowed] struct {
	Elements []T
}

func (a *ArrayEl[Allowed]) isArray() bool {
	return true
}

func (a *ArrayEl[Allowed]) get() []Allowed {
	return a.Elements
}

//type Element struct {
//	IntValue int
//	ArrValue ArrayElement
//}

//func (e *Element) isArray() bool {
//	if e.ArrValue.Elements != nil && len(e.ArrValue.Elements) > 0 {
//		return true
//	} else {
//		return false
//	}
//}

func GetNumbersRestrictedInf(element ElementI[[]any]) []int {
	result := make([]int, 0)
	return getNumbersRestrictedRecInf(element, result)
}

func getNumbersRestrictedRecInf(t ElementI[[]any], res []int) []int {
	if t.isArray() {
		//for _, el := range t.ArrValue.Elements {
		//	if el.isArray() {
		//		res = getNumbersRestrictedRec(el, res)
		//	} else {
		//		res = append(res, el.IntValue)
		//	}
		//}
		//for _, el := range t.get(){
		//	if el.isArray() {
		//		res = getNumbersRestrictedRecInf(el, res)
		//	} else {
		//		res = append(res, el.IntValue)
		//	}
		//}
	}
	return res
}

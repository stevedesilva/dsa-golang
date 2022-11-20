package counternumbers

func GetNumbersFromSlice(array []any) []int {
	result := make([]int, 0)
	return getNumbersFromSliceRec(array, result)
}

func getNumbersFromSliceRec(array []any, res []int) []int {
	for _, v := range array {
		if num, ok := v.(int); ok {
			res = append(res, num)
		} else if arr, ok := v.([]any); ok {
			res = getNumbersFromSliceRec(arr, res)
		} else {
			// error
		}
	}
	return res
}

//
//type AllowedTypes interface {
//	IntegerElement | ArrayElement
//}
//
//
//func (a *IntegerElement) isArray() bool {
//	return false
//}

//type IntegerElement struct {
//	Value int
//}

type ArrayElement struct {
	Elements []Element
}

type Element struct {
	IntValue int
	ArrValue ArrayElement
}

func (e *Element) isArray() bool {
	if e.ArrValue.Elements != nil && len(e.ArrValue.Elements) > 0 {
		return true
	} else {
		return false
	}
}

func GetNumbersRestricted(element Element) []int {
	result := make([]int, 0)
	return getNumbersRestrictedRec(element, result)
}

func getNumbersRestrictedRec(t Element, res []int) []int {
	if t.isArray() {
		for _, el := range t.ArrValue.Elements {
			if el.isArray() {
				res = getNumbersRestrictedRec(el, res)
			} else {
				res = append(res, el.IntValue)
			}
		}
	}
	return res
}

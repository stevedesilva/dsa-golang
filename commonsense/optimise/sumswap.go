package optimise

import "fmt"

func SumSwap(a, b []int) (int, int) {
	sumA, sumB := 0, 0
	for _, v := range a {
		sumA += v
	}
	for _, v := range b {
		sumB += v
	}
	diff := sumA - sumB
	if diff%2 != 0 {
		return 0, 0
	}
	diff /= 2
	for _, v := range a {
		for _, w := range b {
			if v-w == diff {
				return v, w
			}
		}
	}
	return 0, 0
}

/*
    public int[] swapToMakeEqual(int[] a1, int[] a2) {
        validate(a1, a2);
        // key
        Map<Integer, Integer> a1Map = new HashMap<>();
        // get a sum of a1 while storing it's values in hashmap
        int a1Total = 0;
        for (int index = 0; index < a1.length; index++) {
            final int number = a1[index];
            a1Total += number;
            a1Map.put(number, index);
        }

        final int a2Total = Arrays.stream(a2).sum();

        // calculate the difference between the two arrays
        final int difference = (a1Total - a2Total) / 2;

        for (int i=0; i < a2.length; i++) {
            // check map for the numbers counterpart in the 1st array,
            // which is calculated as the current number
            // plus the amount it has to shift by:
            final int num = a2[i];
            final int numToFind = num + difference;
            try {
                if (a1Map.containsKey(numToFind)) {
//                    return new int[]{a1[numToFind],i};
                    return new int[]{a1Map.get(numToFind),i};
                }
            } catch (Exception e) {
                // continue
            }

        }
        // throw item not found exception
        throw new IllegalArgumentException("no items found");
    }
*/

func SwapToMakeEqual(a1, a2 []int) ([]int, error) {
	// validate if a1 or a2 is empty
	if len(a1) == 0 || len(a2) == 0 {
		return nil, fmt.Errorf("empty array")
	}
	// get a sum of a1 while storing it's values in hashmap
	sumA := 0
	hashA := make(map[int]int)
	for i, v := range a1 {
		sumA += v
		hashA[v] = i
	}

	// get a sum of a2
	sumB := 0
	for _, v := range a2 {
		sumB += v
	}

	// calculate the difference
	diff := (sumA - sumB) / 2
	//
	//for (int i=0; i < a2.length; i++) {
	//	// check map for the numbers counterpart in the 1st array,
	//	// which is calculated as the current number
	//	// plus the amount it has to shift by:
	//	final int num = a2[i];
	//	final int numToFind = num + difference;
	//	try {
	//		if (a1Map.containsKey(numToFind)) {
	//		//                    return new int[]{a1[numToFind],i};
	//		return new int[]{a1Map.get(numToFind),i};
	//	}
	//	} catch (Exception e) {
	//		// continue
	//	}
	//
	//}
	// check map for the numbers counterpart in the 1st array,
	// which is calculated as the current number
	// plus the amount it has to shift by:
	for i, v := range a2 {
		numToFind := diff + v
		if _, ok := hashA[numToFind]; ok {
			return []int{hashA[numToFind], i}, nil
		}
	}

	return nil, fmt.Errorf("empty array")
}

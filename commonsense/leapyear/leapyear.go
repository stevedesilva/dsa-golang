package leapyear

/*

yes	Leap Years are any year that can be exactly divided by 4 (such as 2016, 2020, 2024, etc)
 	not	except if it can be exactly divided by 100, then it isn't (such as 2100, 2200, etc)
 	 	yes	except if it can be exactly divided by 400, then it is (such as 2000, 2400)

*/
func IsLeap(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else {
		return true
	}
}

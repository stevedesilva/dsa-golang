package array

import "errors"

var (
	ErrMinimumValue = errors.New("minimum of 1 value is required")
)

func PickResume(resumes []string) (*string, error) {
	if resumes == nil || len(resumes) < 1 {
		return nil, ErrMinimumValue
	}
	top := true
	for len(resumes) > 1 {
		half := len(resumes) / 2
		if top {
			resumes = resumes[half:]
		} else {
			resumes = resumes[:half]
		}
		top = !top
	}
	return &resumes[0], ErrMinimumValue
}

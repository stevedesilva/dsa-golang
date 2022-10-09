package array

import "errors"

var (
	ErrMinimumValue = errors.New("minimum of 1 value is required")
)

func PickResume(resumes []string) (*string, error) {
	if resumes == nil || len(resumes) < 1 {
		return nil, ErrMinimumValue
	}
	return nil, ErrMinimumValue
}

// Code generated by "stringer -type=Status"; DO NOT EDIT.

package core

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusUp-0]
	_ = x[StatusTransitioningUp-1]
	_ = x[StatusTransitioningDown-2]
	_ = x[StatusDown-3]
	_ = x[StatusUnknown-4]
}

const _Status_name = "StatusUpStatusTransitioningUpStatusTransitioningDownStatusDownStatusUnknown"

var _Status_index = [...]uint8{0, 8, 29, 52, 62, 75}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
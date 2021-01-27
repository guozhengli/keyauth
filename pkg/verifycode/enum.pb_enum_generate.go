// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package verifycode

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseNotifyTypeFromString Parse NotifyType from string
func ParseNotifyTypeFromString(str string) (NotifyType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := NotifyType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown NotifyType: %s", str)
	}

	return NotifyType(v), nil
}

// Equal type compare
func (t NotifyType) Equal(target NotifyType) bool {
	return t == target
}

// IsIn todo
func (t NotifyType) IsIn(targets ...NotifyType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t NotifyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(t.String())
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *NotifyType) UnmarshalJSON(b []byte) error {
	ins, err := ParseNotifyTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseIssueTypeFromString Parse IssueType from string
func ParseIssueTypeFromString(str string) (IssueType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := IssueType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown IssueType: %s", str)
	}

	return IssueType(v), nil
}

// Equal type compare
func (t IssueType) Equal(target IssueType) bool {
	return t == target
}

// IsIn todo
func (t IssueType) IsIn(targets ...IssueType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t IssueType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(t.String())
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *IssueType) UnmarshalJSON(b []byte) error {
	ins, err := ParseIssueTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

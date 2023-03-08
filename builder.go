package dateformat

import (
	"fmt"
	"strconv"
	"strings"
)

type strBuilder struct {
	strings.Builder
}

func newStrBuilder() *strBuilder {
	return &strBuilder{}
}

func (sb *strBuilder) WriteInt(i int) error {
	_, err := sb.WriteString(strconv.Itoa(i))
	return err
}

func (sb *strBuilder) WriteIntPrefix(i int, size int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%0*d", size, i))
}

func (sb *strBuilder) WriteStringSlice(s []string) (int, error) {
	return sb.WriteString(strings.Join(s, ""))
}

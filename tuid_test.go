package tuid

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	is := assert.New(t)
	id := New()
	idStr := id.String()

	is.Equal(strings.ToUpper(id.prefix), idStr[:len(id.prefix)])

	re := regexp.MustCompile(`^\d{6}`)
	is.True(re.MatchString(idStr))

	is.Equal(strings.ToUpper(id.randomCode), idStr[len(idStr)-len(id.randomCode):])
}

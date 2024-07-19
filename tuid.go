package tuid

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TUID struct {
	prefix     string
	timeLayout string
	time       time.Time
	randomCode string
	seq        int64
}

func (id TUID) TimeString() string {
	ts := id.time.Format(id.timeLayout)

	re := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	ts = re.ReplaceAllString(ts, "")

	return ts
}

func (id TUID) EncodeString() string {
	weekday := strconv.Itoa(int(id.time.Weekday()) + 1) // avoid zero case
	toEncodeStr := weekday + id.time.Format("150405")

	if id.seq > 0 {
		toEncodeStr += fmt.Sprintf("%04d", id.seq)
	}

	toEncodeInt64, _ := strconv.ParseInt(toEncodeStr, 10, 64)
	return strconv.FormatInt(toEncodeInt64, 36)
}

func (id TUID) String() string {
	sb := strings.Builder{}
	sb.WriteString(id.prefix)
	sb.WriteString(id.TimeString())
	sb.WriteString(id.EncodeString())
	sb.WriteString(id.randomCode)
	return strings.ToUpper(sb.String())
}

func NewTUID(opts ...GenOption) TUID {
	id := TUID{
		timeLayout: TimeLayoutDefualt,
		time:       time.Now(),
		randomCode: RandomString(3),
	}

	for _, opt := range opts {
		opt.apply(&id)
	}

	return id
}

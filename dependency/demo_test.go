package dependency

import (
	"fmt"
	"strings"
	"testing"
)


type FakeLogger struct {
	s strings.Builder
}

func (fl *FakeLogger) Println(v ...interface{}) {
	fmt.Fprintln(&fl.s, v...)
}
func (fl *FakeLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&fl.s, format, v...)
}

func TestDemoV3(t *testing.T) {

	var fl FakeLogger
	DemoV3(&fl) // demoV3 is accepting FakeLogger because it implements Logger interface in demo.go file
	want := "error opening in file"
	got := fl.s.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs %s; want substring %s", got, want)
	}

}

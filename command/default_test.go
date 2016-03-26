package command

import (
	"testing"
	"time"
)

func TestTimeOfDay(t *testing.T) {
	output := timeOfDay(time.Date(2016, time.March, 26, 10, 0, 0, 0, time.UTC))
	if output != "morning" {
		t.Errorf("expected text to equal morning not: %s", output)
	}

	output = timeOfDay(time.Date(2016, time.March, 26, 17, 0, 0, 0, time.UTC))
	if output != "afternoon" {
		t.Errorf("expected text to equal afternoon not: %s", output)
	}

	output = timeOfDay(time.Date(2016, time.March, 26, 23, 0, 0, 0, time.UTC))
	if output != "evening" {
		t.Errorf("expected text to equal morning not: %s", output)
	}
}

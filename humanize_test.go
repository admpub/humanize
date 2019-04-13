package humanize

import "testing"

func TestNew_Correct(t *testing.T) {
	var humanizer *Humanizer
	var err error
	humanizer, err = New("en")
	if humanizer == nil {
		t.Error("Humanizer creation failed.")
	}
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
}

func TestNew_Wrong(t *testing.T) {
	var humanizer *Humanizer
	var err error
	humanizer, err = New("xyz")
	if err != nil {
		t.Fatal(err)
	}
	expected := `en`
	language := humanizer.Language()
	if expected != language {
		t.Errorf(`Expected '%s', got '%s'.`, expected, language)
	}
}

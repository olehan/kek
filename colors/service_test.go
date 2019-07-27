package colors

import "testing"

func TestColor_Add(t *testing.T) {
    t.Log(Green.Add("this text is green"))
    t.Log(Green.Add("this text is green and bold", Bold), "this is normal")
    t.Log(Blue.Add("this text is blue and underlined", Underline))
}

func TestColor_String(t *testing.T) {
    t.Log(Blue.String(), "blue text", Reset.String())
    t.Log(Red.String(Bold, Underline), "underlined red bold", Reset.String())
}

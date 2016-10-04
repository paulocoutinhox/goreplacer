package goreplace

import "testing"
import "os"

func TestReplaceByMapOfString(t *testing.T) {
	text := "The $[book] is on the table"
	expected := "The Big Book is on the table"
	data := map[string]string{
		"book": "Big Book",
	}

	result := ReplaceByMapOfString(text, "$[", "]", data)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

func TestReplaceAllByMapOfString(t *testing.T) {
	text := "The $[book] is on the table, but the $[book] is too heavy!"
	expected := "The Big Book is on the table, but the Big Book is too heavy!"
	data := map[string]string{
		"book": "Big Book",
	}

	result := ReplaceByMapOfString(text, "$[", "]", data)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

func TestReplaceNotFound(t *testing.T) {
	text := "The $[book] is on the table, but the $[book] is too heavy!"
	expected := "The $[book] is on the table, but the $[book] is too heavy!"
	data := map[string]string{
		"book": "Big Book",
	}

	result := ReplaceByMapOfString(text, "{", "}", data)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

func TestReplaceWihoutStartText(t *testing.T) {
	text := "The $[book] is on the table, but the $[book] is too heavy!"
	expected := "The $[book] is on the table, but the $[book] is too heavy!"
	data := map[string]string{
		"book": "Big Book",
	}

	result := ReplaceByMapOfString(text, "", "]", data)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

func TestReplaceWihoutEndText(t *testing.T) {
	text := "The $[book] is on the table, but the $[book] is too heavy!"
	expected := "The $[book] is on the table, but the $[book] is too heavy!"
	data := map[string]string{
		"book": "Big Book",
	}

	result := ReplaceByMapOfString(text, "$[", "", data)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

func TestReplaceByEnvironmentVars(t *testing.T) {
	text := "The home var is: ENV{MY_HOME}"
	expected := "The home var is: /tmp"
	os.Setenv("MY_HOME", "/tmp")

	result, _ := ReplaceFunc(text, "ENV{", "}", os.Getenv)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

func TestReplaceWithThreeValues(t *testing.T) {
	text := "My name is @|name1|@, your name is @|name2|@ and we are in @|country|@"
	expected := "My name is Paulo Coutinho, your name is Guest and we are in Brasil"
	data := map[string]string{
		"name1":   "Paulo Coutinho",
		"name2":   "Guest",
		"country": "Brasil",
	}

	result := ReplaceByMapOfString(text, "@|", "|@", data)

	if result != expected {
		t.Errorf("Expected: %v, But it was: %v", expected, result)
	}
}

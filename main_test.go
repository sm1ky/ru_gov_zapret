package main

import "testing"

func TestParseEntryIgnoresExtraWhitespace(t *testing.T) {
	entry, err := parseEntry("example.com   @blocked")
	if err != nil {
		t.Fatalf("parseEntry returned error: %v", err)
	}

	if entry.Type != "domain" {
		t.Fatalf("unexpected type: %q", entry.Type)
	}

	if entry.Value != "example.com" {
		t.Fatalf("unexpected value: %q", entry.Value)
	}

	if len(entry.Attrs) != 1 {
		t.Fatalf("expected 1 attr, got %d", len(entry.Attrs))
	}

	if entry.Attrs[0].Key != "blocked" {
		t.Fatalf("unexpected attr key: %q", entry.Attrs[0].Key)
	}
}

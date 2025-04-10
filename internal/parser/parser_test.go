package parser

import (
	"strings"
	"testing"
)

func TestParser_ParseComponent(t *testing.T) {
	input := `component Position {
		x: number;
		y: number;
	}`

	p := New(input)
	program, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(program.Components) != 1 {
		t.Fatalf("program.Components does not contain 1 component. got=%d", len(program.Components))
	}

	comp := program.Components[0]
	if comp.Name != "Position" {
		t.Errorf("component.Name not 'Position'. got=%q", comp.Name)
	}

	if len(comp.Fields) != 2 {
		t.Fatalf("component.Fields does not contain 2 fields. got=%d", len(comp.Fields))
	}

	tests := []struct {
		expectedName string
		expectedType string
	}{
		{"x", "number"},
		{"y", "number"},
	}

	for i, tt := range tests {
		field := comp.Fields[i]
		if field.Name != tt.expectedName {
			t.Errorf("field[%d].Name not '%s'. got=%s", i, tt.expectedName, field.Name)
		}
		if field.Type != tt.expectedType {
			t.Errorf("field[%d].Type not '%s'. got=%s", i, tt.expectedType, field.Type)
		}
	}
}

func TestParser_ParseSystem(t *testing.T) {
	input := `system Movement {
		query: (Position, Velocity);
		run(pos, vel) {
			pos.x = pos.x + vel.x;
			pos.y = pos.y + vel.y;
		}
	}`

	p := New(input)
	program, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(program.Systems) != 1 {
		t.Fatalf("program.Systems does not contain 1 system. got=%d", len(program.Systems))
	}

	sys := program.Systems[0]
	if sys.Name != "Movement" {
		t.Errorf("system.Name not 'Movement'. got=%q", sys.Name)
	}

	if sys.Query == nil {
		t.Fatalf("system.Query is nil")
	}

	if len(sys.Query.Components) != 2 {
		t.Fatalf("system.Query.Components does not contain 2 components. got=%d", len(sys.Query.Components))
	}

	expectedComponents := []string{"Position", "Velocity"}
	for i, expected := range expectedComponents {
		if sys.Query.Components[i] != expected {
			t.Errorf("system.Query.Components[%d] not '%s'. got=%s", i, expected, sys.Query.Components[i])
		}
	}

	// Compare code ignoring whitespace differences
	expectedCode := strings.Join([]string{
		"pos.x = pos.x + vel.x;",
		"pos.y = pos.y + vel.y;",
	}, "\n")
	gotCode := strings.TrimSpace(sys.Code)
	if gotCode != expectedCode {
		t.Errorf("system.Code not '%s'. got='%s'", expectedCode, gotCode)
	}
}

func TestParser_ParseRelationship(t *testing.T) {
	input := `relationship ChildOf {}`

	p := New(input)
	program, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(program.Relationships) != 1 {
		t.Fatalf("program.Relationships does not contain 1 relationship. got=%d", len(program.Relationships))
	}

	rel := program.Relationships[0]
	if rel.Name != "ChildOf" {
		t.Errorf("relationship.Name not 'ChildOf'. got=%q", rel.Name)
	}
}

func TestParser_ParseQuery(t *testing.T) {
	input := `system Test {
		query: (Position, pair(ChildOf, *), Velocity);
	}`

	p := New(input)
	program, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(program.Systems) != 1 {
		t.Fatalf("program.Systems does not contain 1 system. got=%d", len(program.Systems))
	}

	sys := program.Systems[0]
	if sys.Query == nil {
		t.Fatalf("system.Query is nil")
	}

	if len(sys.Query.Components) != 2 {
		t.Fatalf("system.Query.Components does not contain 2 components. got=%d", len(sys.Query.Components))
	}

	if len(sys.Query.Relations) != 1 {
		t.Fatalf("system.Query.Relations does not contain 1 relation. got=%d", len(sys.Query.Relations))
	}

	rel := sys.Query.Relations[0]
	if rel.Type != "ChildOf" {
		t.Errorf("relation.Type not 'ChildOf'. got=%q", rel.Type)
	}
	if rel.Target != "*" {
		t.Errorf("relation.Target not '*'. got=%q", rel.Target)
	}
}

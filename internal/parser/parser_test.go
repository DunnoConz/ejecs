package parser

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/ejecs/ejecs/internal/ast"
	"github.com/stretchr/testify/assert"
	// Keep token import if needed for error checks
)

// Helper function to compare strings ignoring whitespace differences AND spaces around operators
func assertEqualIgnoringWhitespace(t *testing.T, expected, actual string) {
	t.Helper()
	// Regex to remove spaces around common operators and punctuation
	// Adjust the character set [=+\-*/.,;:(){}\[\]] as needed (Note: hyphen is escaped)
	spaceAroundOpsRegex := regexp.MustCompile(`\s*([=+\-*/.,;:(){}\[\]])\s*`)

	normalize := func(s string) string {
		s = spaceAroundOpsRegex.ReplaceAllString(s, "$1") // Remove space around ops
		return strings.Join(strings.Fields(s), " ")       // Collapse other whitespace
	}

	normalizedExpected := normalize(expected)
	normalizedActual := normalize(actual)
	assert.Equal(t, normalizedExpected, normalizedActual, "Normalized code strings do not match")
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestParser_ParseComponent(t *testing.T) {
	input := `component Position {
		number x;
		number y = 10.5;
	}`

	p := New(input)
	program, err := p.ParseProgram() // Changed from p.Parse()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.Component)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.Component. got=%T", program.Statements[0])
	}

	comp := stmt
	if comp.Name != "Position" {
		t.Errorf("component.Name not 'Position'. got=%q", comp.Name)
	}

	if len(comp.Fields) != 2 {
		t.Fatalf("component.Fields does not contain 2 fields. got=%d", len(comp.Fields))
	}

	tests := []struct {
		expectedName         string
		expectedType         string
		expectedDefaultValue ast.Expression // Check expression type
	}{
		{"x", "number", nil}, // No default
		{"y", "number", &ast.NumberLiteral{Value: "10.5"}}, // Expect NumberLiteral
	}

	for i, tt := range tests {
		field := comp.Fields[i]
		if field.Name != tt.expectedName {
			t.Errorf("field[%d].Name not '%s'. got=%s", i, tt.expectedName, field.Name)
		}
		if field.Type != tt.expectedType {
			t.Errorf("field[%d].Type not '%s'. got=%s", i, tt.expectedType, field.Type)
		}
		// Compare DefaultValue (might need a helper function for deep comparison)
		if fmt.Sprintf("%v", field.DefaultValue) != fmt.Sprintf("%v", tt.expectedDefaultValue) {
			t.Errorf("field[%d].DefaultValue wrong. expected=%v, got=%v", i, tt.expectedDefaultValue, field.DefaultValue)
		}
	}
}

func TestParser_ParseSystem(t *testing.T) {
	input := `system Movement {
		query(Position, Velocity)
		priority: 10
		{
			pos.x = pos.x + vel.x;
			pos.y = pos.y + vel.y;
		}
	}`

	p := New(input)
	program, err := p.ParseProgram() // Changed from p.Parse()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.System)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.System. got=%T", program.Statements[0])
	}

	sys := stmt
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

	// Check priority
	expectedPriority := &ast.NumberLiteral{Value: "10"}
	if fmt.Sprintf("%v", sys.Priority) != fmt.Sprintf("%v", expectedPriority) {
		t.Errorf("system.Priority wrong. expected=%v, got=%v", expectedPriority, sys.Priority)
	}

	// Compare code ignoring whitespace differences
	expectedCode := "pos.x = pos.x + vel.x; pos.y = pos.y + vel.y;"
	// The current parseCodeBlock joins with spaces - adjust test or parser
	// gotCode := strings.TrimSpace(sys.Code)
	gotCode := sys.Code // Check raw parsed code for now
	assertEqualIgnoringWhitespace(t, expectedCode, gotCode)
}

func TestParser_ParseRelationship(t *testing.T) {
	input := `@parent relationship ChildOf {
		child: A
		parent: B
	}`

	p := New(input)
	program, err := p.ParseProgram() // Changed from p.Parse()
	if err != nil {
		t.Fatalf("ParseProgram() error: %v", err)
	}
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.Relationship)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.Relationship. got=%T", program.Statements[0])
	}

	rel := stmt
	if rel.Name != "ChildOf" {
		t.Errorf("relationship.Name not 'ChildOf'. got=%q", rel.Name)
	}
	if rel.Type != "parent" {
		t.Errorf("relationship.Type not 'parent'. got=%q", rel.Type)
	}
	if rel.Child != "A" {
		t.Errorf("relationship.Child not 'A'. got=%q", rel.Child)
	}
	if rel.Parent != "B" {
		t.Errorf("relationship.Parent not 'B'. got=%q", rel.Parent)
	}
}

// Add Test for Table Type Parsing in Field
func TestParseField_TableType(t *testing.T) {
	input := `table<string, boolean> flags;`
	p := New(input)
	// Directly parse a field for testing (adjust if needed)
	// This requires exposing parseField or testing via parseComponent
	// Let's test via parseComponent
	compInput := fmt.Sprintf("component Test { %s }", input)
	p = New(compInput)
	program, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("ParseProgram error: %v", err)
	}
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement")
	}
	comp, ok := program.Statements[0].(*ast.Component)
	if !ok {
		t.Fatalf("Expected *ast.Component")
	}
	if len(comp.Fields) != 1 {
		t.Fatalf("Expected 1 field")
	}

	field := comp.Fields[0]
	if field.Type != "table" {
		t.Errorf("field.Type not 'table'. got=%q", field.Type)
	}
	if field.Name != "flags" {
		t.Errorf("field.Name not 'flags'. got=%q", field.Name)
	}
	if field.MapKeyType != "string" {
		t.Errorf("field.MapKeyType not 'string'. got=%q", field.MapKeyType)
	}
	if field.MapValueType != "boolean" {
		t.Errorf("field.MapValueType not 'boolean'. got=%q", field.MapValueType)
	}
}

// Add Test for Default Value Expression Parsing
func TestParseField_DefaultValueExpr(t *testing.T) {
	input := `CFrame camera = CFrame.new(0, 1, -5);`
	compInput := fmt.Sprintf("component Test { %s }", input)
	p := New(compInput)
	program, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("ParseProgram error: %v", err)
	}
	checkParserErrors(t, p)
	// ... (checks for component and field) ...
	comp, _ := program.Statements[0].(*ast.Component)
	field := comp.Fields[0]

	if field.DefaultValue == nil {
		t.Fatalf("field.DefaultValue is nil, expected expression")
	}
	// Basic check using String() representation
	expected := "CFrame.new(0, 1, (-10))" // Note: String() adds parens for prefix
	// Let's re-run with actual input to get correct expected string
	expected = "CFrame.new(0, 1, (-5))"
	got := field.DefaultValue.String()
	if got != expected {
		t.Errorf("DefaultValue String() wrong.\nexpected=%q\ngot=%q", expected, got)
	}
}

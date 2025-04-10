package ast

import (
	"strings"
	"testing"
)

func TestComponent_String(t *testing.T) {
	tests := []struct {
		name     string
		comp     *Component
		expected string
	}{
		{
			name: "basic component",
			comp: &Component{
				Name: "Position",
				Fields: []*Field{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
			},
			expected: "component Position {\n    x: number\n    y: number\n}",
		},
		{
			name: "component with attributes",
			comp: &Component{
				Name:       "Player",
				Attributes: []string{"replicated", "networked"},
				Fields: []*Field{
					{Name: "name", Type: "string"},
				},
			},
			expected: "@replicated @networked\ncomponent Player {\n    name: string\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.comp.String()
			if got != tt.expected {
				t.Errorf("Component.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestField_String(t *testing.T) {
	tests := []struct {
		name     string
		field    *Field
		expected string
	}{
		{
			name:     "basic field",
			field:    &Field{Name: "x", Type: "number"},
			expected: "x: number",
		},
		{
			name:     "optional field",
			field:    &Field{Name: "name", Type: "string", Optional: true},
			expected: "name: string?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.String()
			if got != tt.expected {
				t.Errorf("Field.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSystem_String(t *testing.T) {
	tests := []struct {
		name     string
		sys      *System
		expected string
	}{
		{
			name: "basic system",
			sys: &System{
				Name:       "Movement",
				Components: []string{"Position", "Velocity"},
			},
			expected: "system Movement {\n    using Position, Velocity\n}",
		},
		{
			name: "system with frequency and priority",
			sys: &System{
				Name:       "Physics",
				Components: []string{"RigidBody"},
				Frequency:  "60hz",
				Priority:   "1",
			},
			expected: "system Physics {\n    using RigidBody\n    frequency: 60hz\n    priority: 1\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sys.String()
			if got != tt.expected {
				t.Errorf("System.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRelationship_String(t *testing.T) {
	tests := []struct {
		name     string
		rel      *Relationship
		expected string
	}{
		{
			name: "basic relationship",
			rel: &Relationship{
				Name:   "ChildOf",
				Child:  "child",
				Parent: "parent",
			},
			expected: "relationship ChildOf {\n    child: child\n    parent: parent\n}",
		},
		{
			name: "relationship with type",
			rel: &Relationship{
				Name:   "Inventory",
				Child:  "item",
				Parent: "container",
				Type:   "many_to_one",
			},
			expected: "@many_to_one\nrelationship Inventory {\n    child: item\n    parent: container\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rel.String()
			if got != tt.expected {
				t.Errorf("Relationship.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestProgram_String(t *testing.T) {
	prog := &Program{
		Statements: []Node{
			&Component{
				Name: "Position",
				Fields: []*Field{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
			},
			&Relationship{
				Name:   "ChildOf",
				Child:  "child",
				Parent: "parent",
			},
			&System{
				Name:       "Movement",
				Components: []string{"Position", "Velocity"},
			},
		},
	}

	expected := strings.Join([]string{
		"component Position {\n    x: number\n    y: number\n}",
		"relationship ChildOf {\n    child: child\n    parent: parent\n}",
		"system Movement {\n    using Position, Velocity\n}",
	}, "\n")

	got := prog.String()
	if got != expected {
		t.Errorf("Program.String() wrong.\nexpected=%q\ngot=%q", expected, got)
	}
}

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Node{
			&Component{
				Name: "Player",
				Fields: []*Field{
					{Name: "health", Type: "int"},
					{Name: "position", Type: "Vector3"},
				},
			},
			&Relationship{
				Type:   "parent",
				Name:   "PlayerMovement",
				Child:  "Player",
				Parent: "Movement",
			},
			&System{
				Name: "MovementSystem",
				Parameters: []*Parameter{
					{Name: "speed", Type: "float", DefaultValue: "1.0"},
					{Name: "maxSpeed", Type: "float", DefaultValue: "10.0"},
				},
				Query: &Query{
					Components: []string{"Position", "Velocity"},
					Relations: []*Relation{
						{Type: "parent", Component: "Movement"},
					},
				},
				Frequency: "60",
				Priority:  "1",
				Code:      "position.x += velocity.x * speed",
			},
		},
	}

	expected := `component Player {
    health: int
    position: Vector3
}

@parent
relationship PlayerMovement {
    child: Player
    parent: Movement
}

system MovementSystem {
    parameters: {
        speed: float = 1.0
        maxSpeed: float = 10.0
    }
    query: {
        components: [Position, Velocity]
        relations: [pair(parent, Movement)]
    }
    frequency: 60
    priority: 1
    code: {
        position.x += velocity.x * speed
    }
}`

	if program.String() != expected {
		t.Errorf("program.String() wrong.\nexpected=%q\ngot=%q",
			expected, program.String())
	}
}

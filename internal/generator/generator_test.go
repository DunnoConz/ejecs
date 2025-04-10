package generator

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/ejecs/ejecs/internal/ast"
)

// normalizeWhitespace removes insignificant whitespace differences
func normalizeWhitespace(s string) string {
	// Remove empty lines
	s = regexp.MustCompile(`(?m)^\s*$`).ReplaceAllString(s, "")

	// Normalize indentation to 4 spaces
	s = regexp.MustCompile(`(?m)^[\t ]+`).ReplaceAllString(s, "    ")

	// Remove trailing commas in arrays/objects
	s = regexp.MustCompile(`,(\s*[}\]])`).ReplaceAllString(s, "$1")

	// Normalize spaces around operators and commas
	s = regexp.MustCompile(`\s*([-+*/=,])\s*`).ReplaceAllString(s, "$1")

	// Normalize spaces in function declarations
	s = regexp.MustCompile(`function\s*\(`).ReplaceAllString(s, "function(")

	// Remove extra spaces in parameter lists
	s = regexp.MustCompile(`\(\s+`).ReplaceAllString(s, "(")
	s = regexp.MustCompile(`\s+\)`).ReplaceAllString(s, ")")

	// Normalize spaces after commas in parameter lists
	s = regexp.MustCompile(`,\s*([a-zA-Z])`).ReplaceAllString(s, ", $1")

	// Normalize spaces around braces
	s = regexp.MustCompile(`{\s+`).ReplaceAllString(s, "{")
	s = regexp.MustCompile(`\s+}`).ReplaceAllString(s, "}")

	// Normalize spaces around brackets
	s = regexp.MustCompile(`\[\s+`).ReplaceAllString(s, "[")
	s = regexp.MustCompile(`\s+\]`).ReplaceAllString(s, "]")

	// Normalize spaces around parentheses
	s = regexp.MustCompile(`\(\s+`).ReplaceAllString(s, "(")
	s = regexp.MustCompile(`\s+\)`).ReplaceAllString(s, ")")

	// Normalize spaces around semicolons
	s = regexp.MustCompile(`\s*;\s*`).ReplaceAllString(s, ";")

	// Normalize spaces around dots
	s = regexp.MustCompile(`\s*\.\s*`).ReplaceAllString(s, ".")

	// Normalize spaces around colons
	s = regexp.MustCompile(`\s*:\s*`).ReplaceAllString(s, ":")

	// Normalize newlines
	s = regexp.MustCompile(`\r\n|\r|\n`).ReplaceAllString(s, "\n")

	// Normalize multiple spaces
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	// Trim trailing whitespace
	s = regexp.MustCompile(`(?m)[ \t]+$`).ReplaceAllString(s, "")

	return strings.TrimSpace(s)
}

func TestGenerator_Component(t *testing.T) {
	tests := []struct {
		name     string
		comp     *ast.Component
		expected string
	}{
		{
			name: "basic component",
			comp: &ast.Component{
				Name: "Position",
				Fields: []*ast.Field{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
			},
			expected: `Types.Position = {
    x = 0,
    y = 0
}`,
		},
		{
			name: "component with attributes",
			comp: &ast.Component{
				Name:       "Player",
				Attributes: []string{"replicated", "networked"},
				Fields: []*ast.Field{
					{Name: "name", Type: "string"},
					{Name: "health", Type: "number"},
				},
			},
			expected: `Types.Player = {
    name = "",
    health = 0
}`,
		},
	}

	for _, library := range []string{"ecr", "jecs"} {
		g := New(Config{Library: library})

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s_%s", library, tt.name), func(t *testing.T) {
				program := &ast.Program{
					Components: []*ast.Component{tt.comp},
				}
				got, err := g.Generate(program)
				if err != nil {
					t.Fatalf("Generate() error = %v", err)
				}

				// Extract just the component definition
				lines := strings.Split(got, "\n")
				var componentLines []string
				inComponent := false
				for _, line := range lines {
					if strings.Contains(line, tt.comp.Name) {
						inComponent = true
					}
					if inComponent {
						componentLines = append(componentLines, line)
						if strings.Contains(line, "}") {
							break
						}
					}
				}

				// Compare just the component definition
				got = strings.Join(componentLines, "\n")
				expected := strings.TrimSpace(tt.expected)

				if got != expected {
					t.Errorf("Generate() = %v, want %v", got, expected)
				}
			})
		}
	}
}

func TestGenerator_System(t *testing.T) {
	tests := []struct {
		name         string
		sys          *ast.System
		ecrExpected  string
		jecsExpected string
	}{
		{
			name: "basic system",
			sys: &ast.System{
				Name:       "Movement",
				Components: []string{"Position", "Velocity"},
				Code:       "pos.x = pos.x + vel.x;\npos.y = pos.y + vel.y;",
			},
			ecrExpected: `world:system({
    name = "Movement",
    query = {
        all = {
            Position,
            Velocity
        }
    },
    callback = function(entity, components)
        pos.x = pos.x + vel.x;
        pos.y = pos.y + vel.y;
    end
})`,
			jecsExpected: `world:system({
    name = "Movement",
    query = {
        with = {
            Position,
            Velocity
        }
    },
    callback = function(entity, components)
        pos.x = pos.x + vel.x;
        pos.y = pos.y + vel.y;
    end
})`,
		},
		{
			name: "system with frequency and priority",
			sys: &ast.System{
				Name:       "Physics",
				Components: []string{"RigidBody"},
				Frequency:  "60hz",
				Priority:   "1",
				Code:       "body.simulate();",
			},
			ecrExpected: `world:system({
    name = "Physics",
    query = {
        all = {
            RigidBody
        }
    },
    frequency = 60hz,
    priority = 1,
    callback = function(entity, components)
        body.simulate();
    end
})`,
			jecsExpected: `world:system({
    name = "Physics",
    query = {
        with = {
            RigidBody
        }
    },
    frequency = 60hz,
    priority = 1,
    callback = function(entity, components)
        body.simulate();
    end
})`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test ECR
			g := New(Config{Library: "ecr"})
			program := &ast.Program{
				Systems: []*ast.System{tt.sys},
			}
			got, err := g.Generate(program)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			// Extract just the system definition
			lines := strings.Split(got, "\n")
			var systemLines []string
			inSystem := false
			for _, line := range lines {
				if strings.Contains(line, tt.sys.Name) {
					inSystem = true
				}
				if inSystem {
					systemLines = append(systemLines, line)
					if strings.Contains(line, "})") {
						break
					}
				}
			}

			// Compare just the system definition
			got = strings.Join(systemLines, "\n")
			expected := strings.TrimSpace(tt.ecrExpected)

			if got != expected {
				t.Errorf("ECR Generate() = %v, want %v", got, expected)
			}

			// Test jecs
			g = New(Config{Library: "jecs"})
			got, err = g.Generate(program)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			// Extract just the system definition
			lines = strings.Split(got, "\n")
			systemLines = nil
			inSystem = false
			for _, line := range lines {
				if strings.Contains(line, tt.sys.Name) {
					inSystem = true
				}
				if inSystem {
					systemLines = append(systemLines, line)
					if strings.Contains(line, "})") {
						break
					}
				}
			}

			// Compare just the system definition
			got = strings.Join(systemLines, "\n")
			expected = strings.TrimSpace(tt.jecsExpected)

			if got != expected {
				t.Errorf("jecs Generate() = %v, want %v", got, expected)
			}
		})
	}
}

func TestGenerator_Relationship(t *testing.T) {
	tests := []struct {
		name         string
		rel          *ast.Relationship
		ecrExpected  string
		jecsExpected string
	}{
		{
			name: "basic relationship",
			rel: &ast.Relationship{
				Name: "ChildOf",
				From: "child",
				To:   "parent",
			},
			ecrExpected:  `local ChildOf = world:component()`,
			jecsExpected: `local ChildOf = world:component()`,
		},
		{
			name: "relationship with type",
			rel: &ast.Relationship{
				Name: "Inventory",
				From: "item",
				To:   "container",
				Type: "many_to_one",
			},
			ecrExpected:  `local Inventory = world:component()`,
			jecsExpected: `local Inventory = world:component()`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test ECR
			g := New(Config{Library: "ecr"})
			program := &ast.Program{
				Relationships: []*ast.Relationship{tt.rel},
			}
			got, err := g.Generate(program)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			// Extract just the relationship definition
			lines := strings.Split(got, "\n")
			var relationshipLines []string
			for _, line := range lines {
				if strings.Contains(line, tt.rel.Name) {
					relationshipLines = append(relationshipLines, line)
					break
				}
			}

			// Compare just the relationship definition
			got = strings.Join(relationshipLines, "\n")
			expected := strings.TrimSpace(tt.ecrExpected)

			if got != expected {
				t.Errorf("ECR Generate() = %v, want %v", got, expected)
			}

			// Test jecs
			g = New(Config{Library: "jecs"})
			got, err = g.Generate(program)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			// Extract just the relationship definition
			lines = strings.Split(got, "\n")
			relationshipLines = nil
			for _, line := range lines {
				if strings.Contains(line, tt.rel.Name) {
					relationshipLines = append(relationshipLines, line)
					break
				}
			}

			// Compare just the relationship definition
			got = strings.Join(relationshipLines, "\n")
			expected = strings.TrimSpace(tt.jecsExpected)

			if got != expected {
				t.Errorf("jecs Generate() = %v, want %v", got, expected)
			}
		})
	}
}

func TestGenerator_Complete(t *testing.T) {
	program := &ast.Program{
		Components: []*ast.Component{
			{
				Name: "Position",
				Fields: []*ast.Field{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
			},
		},
		Systems: []*ast.System{
			{
				Name:       "Movement",
				Components: []string{"Position", "Velocity"},
				Code:       "pos.x = pos.x + vel.x;\npos.y = pos.y + vel.y;",
			},
		},
		Relationships: []*ast.Relationship{
			{
				Name: "ChildOf",
				From: "child",
				To:   "parent",
			},
		},
	}

	ecrExpected := `-- Generated by EJECS IDL Compiler
local ECR = require(game.ReplicatedStorage.ECR)
local world = ECR.World.new()

Types.Position = {
    x = 0,
    y = 0
}

local ChildOf = world:component()

world:system({
    name = "Movement",
    query = {
        all = {
            Position,
            Velocity
        }
    },
    callback = function(entity, components)
        pos.x = pos.x + vel.x;
        pos.y = pos.y + vel.y;
    end
})`

	jecsExpected := `-- Generated by EJECS IDL Compiler
local jecs = require(game.ReplicatedStorage.jecs)
local world = jecs.World.new()

Types.Position = {
    x = 0,
    y = 0
}

local ChildOf = world:component()

world:system({
    name = "Movement",
    query = {
        with = {
            Position,
            Velocity
        }
    },
    callback = function(entity, components)
        pos.x = pos.x + vel.x;
        pos.y = pos.y + vel.y;
    end
})`

	// Test ECR
	g := New(Config{Library: "ecr"})
	got, err := g.Generate(program)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Normalize whitespace for comparison
	got = normalizeWhitespace(got)
	expected := normalizeWhitespace(ecrExpected)

	if got != expected {
		t.Errorf("ECR Generate() = %v, want %v", got, expected)
	}

	// Test jecs
	g = New(Config{Library: "jecs"})
	got, err = g.Generate(program)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Normalize whitespace for comparison
	got = normalizeWhitespace(got)
	expected = normalizeWhitespace(jecsExpected)

	if got != expected {
		t.Errorf("jecs Generate() = %v, want %v", got, expected)
	}
}

func TestGenerator_QuerySyntax(t *testing.T) {
	tests := []struct {
		name         string
		sys          *ast.System
		ecrExpected  string
		jecsExpected string
	}{
		{
			name: "complex query with relationships",
			sys: &ast.System{
				Name:       "ParentSystem",
				Components: []string{"Transform", "RigidBody"},
				Code:       "local transform = components.Transform",
			},
			ecrExpected: `world:system({
    name = "ParentSystem",
    query = {
        all = {
            Transform,
            RigidBody,
        },
    },
    callback = function(entity, components)
        local transform = components.Transform
    end
})`,
			jecsExpected: `world:system({
    name = "ParentSystem",
    query = {
        with = {
            Transform,
            RigidBody,
        },
    },
    callback = function(entity, components)
        local transform = components.Transform
    end
})`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test ECR
			g := New(Config{Library: "ecr"})
			program := &ast.Program{
				Systems: []*ast.System{tt.sys},
			}
			got, err := g.Generate(program)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			// Extract just the system definition
			lines := strings.Split(got, "\n")
			var systemLines []string
			inSystem := false
			for _, line := range lines {
				if strings.Contains(line, tt.sys.Name) {
					inSystem = true
				}
				if inSystem {
					systemLines = append(systemLines, line)
					if strings.Contains(line, "})") {
						break
					}
				}
			}

			// Compare just the system definition
			got = strings.Join(systemLines, "\n")
			expected := strings.TrimSpace(tt.ecrExpected)

			if got != expected {
				t.Errorf("ECR Generate() = %v, want %v", got, expected)
			}

			// Test jecs
			g = New(Config{Library: "jecs"})
			got, err = g.Generate(program)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			// Extract just the system definition
			lines = strings.Split(got, "\n")
			systemLines = nil
			inSystem = false
			for _, line := range lines {
				if strings.Contains(line, tt.sys.Name) {
					inSystem = true
				}
				if inSystem {
					systemLines = append(systemLines, line)
					if strings.Contains(line, "})") {
						break
					}
				}
			}

			// Compare just the system definition
			got = strings.Join(systemLines, "\n")
			expected = strings.TrimSpace(tt.jecsExpected)

			if got != expected {
				t.Errorf("jecs Generate() = %v, want %v", got, expected)
			}
		})
	}
}

func TestGenerateComponent(t *testing.T) {
	tests := []struct {
		name     string
		library  string
		comp     *ast.Component
		expected string
	}{
		{
			name:    "ECR component",
			library: "ecr",
			comp: &ast.Component{
				Name: "Position",
				Fields: []*ast.Field{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
			},
			expected: `Types.Position = {
    x = 0,
    y = 0
}`,
		},
		{
			name:    "Jecs component",
			library: "jecs",
			comp: &ast.Component{
				Name: "Position",
				Fields: []*ast.Field{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
			},
			expected: `Types.Position = {
    x = 0,
    y = 0
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(Config{Library: tt.library})
			err := g.generateComponent(tt.comp)
			if err != nil {
				t.Fatalf("generateComponent() error = %v", err)
			}
			got := normalizeWhitespace(g.buffer.String())
			expected := normalizeWhitespace(tt.expected)
			if got != expected {
				t.Errorf("generateComponent() = %v, want %v", got, expected)
			}
		})
	}
}

func TestGenerateSystem(t *testing.T) {
	tests := []struct {
		name     string
		library  string
		sys      *ast.System
		expected string
	}{
		{
			name:    "ECR system with parameters",
			library: "ecr",
			sys: &ast.System{
				Name: "MoveSystem",
				Parameters: []*ast.Field{
					{Name: "speed", Type: "number"},
				},
				Components: []string{"Position", "Velocity"},
				Code:       "Position.x = Position.x + Velocity.x * speed",
			},
			expected: `world:system({
    name = "MoveSystem",
    parameters = {
        speed = 0
    },
    query = {
        all = {
            Position,
            Velocity
        }
    },
    callback = function(entity, components, speed)
        Position.x = Position.x + Velocity.x * speed
    end
})`,
		},
		{
			name:    "Jecs system with parameters",
			library: "jecs",
			sys: &ast.System{
				Name: "MoveSystem",
				Parameters: []*ast.Field{
					{Name: "speed", Type: "number"},
				},
				Components: []string{"Position", "Velocity"},
				Code:       "Position.x = Position.x + Velocity.x * speed",
			},
			expected: `world:system({
    name = "MoveSystem",
    parameters = {
        speed = 0
    },
    query = {
        with = {
            Position,
            Velocity
        }
    },
    callback = function(entity, components, speed)
        Position.x = Position.x + Velocity.x * speed
    end
})`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(Config{Library: tt.library})
			got, err := g.generateSystem(tt.sys)
			if err != nil {
				t.Fatalf("generateSystem() error = %v", err)
			}
			got = normalizeWhitespace(got)
			expected := normalizeWhitespace(tt.expected)
			if got != expected {
				t.Errorf("generateSystem() = %v, want %v", got, expected)
			}
		})
	}
}

package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ejecs/ejecs/internal/ast"
)

// Config holds the configuration for the generator
// type Config struct { // Removed Config struct
// 	Library string // Target library (ecr, jecs)
// }

// Generator handles the code generation process
type Generator struct {
	buffer bytes.Buffer
	indent int
}

// New creates a new Generator instance
// func New(config Config) *Generator { // Old New function signature
func New() *Generator { // Simplified New function signature
	return &Generator{
		// library: config.Library, // Removed library assignment
	}
}

// Generate generates code from an AST
func (g *Generator) Generate(node ast.Node) (string, error) {
	switch n := node.(type) {
	case *ast.Program:
		var out strings.Builder
		for i, stmt := range n.Statements {
			if i > 0 {
				out.WriteString("\n\n")
			}
			result, err := g.Generate(stmt)
			if err != nil {
				return "", err
			}
			out.WriteString(result)
		}
		return out.String(), nil

	case *ast.Component:
		return g.generateComponent(n), nil

	case *ast.System:
		return g.generateSystem(n), nil

	case *ast.Relationship:
		return g.generateRelationship(n), nil

	default:
		return "", fmt.Errorf("unknown node type: %T", n)
	}
}

func (g *Generator) writeHeader() {
	g.writeLine("-- Generated by EJECS IDL Compiler")
	g.writeLine("local ECR = require(game.ReplicatedStorage.ECR)")
	g.writeLine("local world = ECR.World.new()")
	g.writeLine("")
}

func (g *Generator) generateComponent(comp *ast.Component) string {
	var sb strings.Builder
	sb.WriteString("local ")
	sb.WriteString(comp.Name)
	sb.WriteString(" = {\n")
	for i, field := range comp.Fields {
		if i > 0 {
			sb.WriteString(",\n")
		}
		sb.WriteString("    ")
		sb.WriteString(field.Name)
		sb.WriteString(" = ")
		sb.WriteString(field.Type)
	}
	sb.WriteString("\n}")
	return sb.String()
}

// getDefaultValue returns the default value for a given type
func (g *Generator) getDefaultValue(typeName string) string {
	switch typeName {
	case "int", "float", "number":
		return "0"
	case "string":
		return "\"\""
	case "boolean":
		return "false"
	case "Vector2":
		return "Vector2.new(0, 0)"
	case "Vector3":
		return "Vector3.new(0, 0, 0)"
	case "CFrame":
		return "CFrame.new()"
	case "Color3":
		return "Color3.new(1, 1, 1)"
	case "UDim2":
		return "UDim2.new(0, 0, 0, 0)"
	case "UDim":
		return "UDim.new(0, 0)"
	default:
		return "nil"
	}
}

func (g *Generator) generateRelationship(rel *ast.Relationship) string {
	var sb strings.Builder
	sb.WriteString("local ")
	sb.WriteString(rel.Name)
	sb.WriteString(" = {\n")
	sb.WriteString("    child = \"")
	sb.WriteString(rel.Child)
	sb.WriteString("\",\n")
	sb.WriteString("    parent = \"")
	sb.WriteString(rel.Parent)
	sb.WriteString("\"\n}")
	return sb.String()
}

func (g *Generator) generateSystem(system *ast.System) string {
	var sb strings.Builder
	sb.WriteString("world:system({")
	sb.WriteString("\n    name = \"")
	sb.WriteString(system.Name)
	sb.WriteString("\",")

	if len(system.Parameters) > 0 {
		sb.WriteString("\n    parameters = {")
		for i, param := range system.Parameters {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(param.Name)
			sb.WriteString(" = ")
			if param.DefaultValue != "" {
				sb.WriteString(param.DefaultValue)
			} else {
				sb.WriteString("0") // Default to 0 for numeric parameters
			}
		}
		sb.WriteString("},")
	}

	if system.Query != nil {
		sb.WriteString("\n    query = {")
		if len(system.Query.Components) > 0 {
			sb.WriteString("\n        all = {")
			for i, comp := range system.Query.Components {
				if i > 0 {
					sb.WriteString(",")
				}
				sb.WriteString("\n            ")
				sb.WriteString(comp)
			}
			sb.WriteString("\n        },")
		}
		sb.WriteString("\n    },")
	}

	if system.Frequency != "" {
		sb.WriteString("\n    frequency = ")
		sb.WriteString(system.Frequency)
		sb.WriteString(",")
	}

	if system.Priority != "" {
		sb.WriteString("\n    priority = ")
		sb.WriteString(system.Priority)
		sb.WriteString(",")
	}

	if system.Code != "" {
		sb.WriteString("\n    callback = function(entity, components")
		for _, param := range system.Parameters {
			sb.WriteString(", ")
			sb.WriteString(param.Name)
		}
		sb.WriteString(")\n        ")
		sb.WriteString(system.Code)
		sb.WriteString("\n    end")
	}

	sb.WriteString("\n})")
	return sb.String()
}

func (g *Generator) generateSystemWithIndent(system *ast.System, useIndent bool) (string, error) {
	var sb strings.Builder
	indent := "    "

	// Write system name
	if useIndent {
		sb.WriteString(fmt.Sprintf("name = %q,\n", system.Name))
	} else {
		sb.WriteString(fmt.Sprintf("name=%q,", system.Name))
	}

	// Write parameters if present
	if len(system.Parameters) > 0 {
		if useIndent {
			sb.WriteString("parameters = {\n")
			for i, param := range system.Parameters {
				if i > 0 {
					sb.WriteString(",\n")
				}
				sb.WriteString(fmt.Sprintf("%s%s = %s", indent, param.Name, g.getDefaultValue(param.Type)))
			}
			sb.WriteString("\n},\n")
		} else {
			sb.WriteString("parameters={")
			for i, param := range system.Parameters {
				if i > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(fmt.Sprintf("%s=%s", param.Name, g.getDefaultValue(param.Type)))
			}
			sb.WriteString("},")
		}
	}

	// Write query
	if useIndent {
		sb.WriteString("query = {\n")
		sb.WriteString(fmt.Sprintf("%sall = {\n", indent))
		for i, comp := range system.Components {
			if i > 0 {
				sb.WriteString(",\n")
			}
			sb.WriteString(fmt.Sprintf("%s%s%s", indent, indent, comp))
		}
		sb.WriteString(fmt.Sprintf("\n%s}", indent))
		if system.Query != nil && len(system.Query.Relations) > 0 {
			sb.WriteString(",\n")
			for i, rel := range system.Query.Relations {
				if i > 0 {
					sb.WriteString(",\n")
				}
				sb.WriteString(fmt.Sprintf("%s%spair(%s, %s)", indent, indent, rel.Type, rel.Component))
			}
		}
		sb.WriteString("\n},\n")
	} else {
		sb.WriteString("query={")
		sb.WriteString("all={")
		for i, comp := range system.Components {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(comp)
		}
		sb.WriteString("}")
		if system.Query != nil && len(system.Query.Relations) > 0 {
			sb.WriteString(",")
			for i, rel := range system.Query.Relations {
				if i > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(fmt.Sprintf("pair(%s,%s)", rel.Type, rel.Component))
			}
		}
		sb.WriteString("},")
	}

	// Write frequency if present
	if system.Frequency != "" {
		if useIndent {
			sb.WriteString(fmt.Sprintf("frequency = %s,\n", system.Frequency))
		} else {
			sb.WriteString(fmt.Sprintf("frequency=%s,", system.Frequency))
		}
	}

	// Write priority if present
	if system.Priority != "" {
		if useIndent {
			sb.WriteString(fmt.Sprintf("priority = %s,\n", system.Priority))
		} else {
			sb.WriteString(fmt.Sprintf("priority=%s,", system.Priority))
		}
	}

	// Write callback
	if useIndent {
		sb.WriteString("callback = function(entity, components")
		if len(system.Parameters) > 0 {
			for _, param := range system.Parameters {
				sb.WriteString(fmt.Sprintf(", %s", param.Name))
			}
		}
		sb.WriteString(")\n")
		lines := strings.Split(system.Code, "\n")
		for _, line := range lines {
			sb.WriteString(fmt.Sprintf("%s%s\n", indent, strings.TrimSpace(line)))
		}
		sb.WriteString("end")
	} else {
		sb.WriteString("callback=function(entity,components")
		if len(system.Parameters) > 0 {
			for _, param := range system.Parameters {
				sb.WriteString(fmt.Sprintf(",%s", param.Name))
			}
		}
		sb.WriteString(") ")
		sb.WriteString(strings.TrimSpace(system.Code))
		sb.WriteString(" end")
	}

	return sb.String(), nil
}

func (g *Generator) writeLine(line string) {
	if line == "" {
		g.buffer.WriteString("\n")
		return
	}
	indent := strings.Repeat("    ", g.indent)
	g.buffer.WriteString(indent + line + "\n")
}

func (g *Generator) writeString(str string) {
	if str == "" {
		return
	}
	indent := strings.Repeat("    ", g.indent)
	g.buffer.WriteString(indent + str)
}

func luauType(t string) string {
	switch t {
	case "number":
		return "number"
	case "string":
		return "string"
	case "boolean":
		return "boolean"
	default:
		return "any"
	}
}

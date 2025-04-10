package ast

import (
	"fmt"
	"strings"
)

// Node represents a node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Program represents the root node of every AST
type Program struct {
	Statements []Node
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out strings.Builder
	for i, s := range p.Statements {
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString(s.String())
	}
	return out.String()
}

// Component represents a component declaration
type Component struct {
	Name       string
	Fields     []*Field
	Attributes []string
}

func (c *Component) TokenLiteral() string { return "component" }
func (c *Component) String() string {
	var out strings.Builder
	out.WriteString("component ")
	out.WriteString(c.Name)
	out.WriteString(" {\n")
	for _, field := range c.Fields {
		out.WriteString("    ")
		out.WriteString(field.Name)
		out.WriteString(": ")
		out.WriteString(field.Type)
		out.WriteString("\n")
	}
	out.WriteString("}")
	return out.String()
}

// Field represents a field in a component
type Field struct {
	Name     string
	Type     string // Base type (e.g., "int", "Vector3", "table")
	Optional bool
	// IsMap    bool   // Removed flag
	MapKeyType   string // Used if Type is "table"
	MapValueType string // Used if Type is "table"
	DefaultValue string
}

func (f *Field) TokenLiteral() string { return "field" }
func (f *Field) String() string {
	opt := ""
	if f.Optional {
		opt = "?"
	}
	return fmt.Sprintf("%s: %s%s", f.Name, f.Type, opt)
}

// Relationship represents a relationship declaration
type Relationship struct {
	Type   string
	Name   string
	Child  string
	Parent string
}

func (r *Relationship) TokenLiteral() string { return "relationship" }
func (r *Relationship) String() string {
	var out strings.Builder
	if r.Type != "" {
		out.WriteString("@")
		out.WriteString(r.Type)
		out.WriteString("\n")
	}
	out.WriteString("relationship ")
	out.WriteString(r.Name)
	out.WriteString(" {\n")
	out.WriteString("    child: ")
	out.WriteString(r.Child)
	out.WriteString("\n")
	out.WriteString("    parent: ")
	out.WriteString(r.Parent)
	out.WriteString("\n}")
	return out.String()
}

type SystemParameter struct {
	// Token lexer.Token // Removed Token field
	Name string
	Type string
	// Add Line/Column ints here if needed for errors
	Line   int
	Column int
}

// System represents a system declaration
type System struct {
	// Token      lexer.Token // Removed Token field
	Name       string
	Parameters []*Parameter
	Components []string
	Query      *Query
	Frequency  string
	Priority   string
	Code       string
	// Add Line/Column ints here if needed for errors
	Line   int
	Column int
}

func (s *System) TokenLiteral() string { return "system" }
func (s *System) String() string {
	var out strings.Builder
	out.WriteString("system ")
	out.WriteString(s.Name)
	out.WriteString(" {\n")
	if len(s.Parameters) > 0 {
		out.WriteString("    parameters: {\n")
		for _, param := range s.Parameters {
			out.WriteString("        ")
			out.WriteString(param.Name)
			out.WriteString(": ")
			out.WriteString(param.Type)
			if param.DefaultValue != "" {
				out.WriteString(" = ")
				out.WriteString(param.DefaultValue)
			}
			out.WriteString("\n")
		}
		out.WriteString("    }\n")
	}
	if s.Query != nil {
		out.WriteString("    query: {\n")
		if len(s.Query.Components) > 0 {
			out.WriteString("        components: [")
			for i, comp := range s.Query.Components {
				if i > 0 {
					out.WriteString(", ")
				}
				out.WriteString(comp)
			}
			out.WriteString("]\n")
		}
		if len(s.Query.Relations) > 0 {
			out.WriteString("        relations: [")
			for i, rel := range s.Query.Relations {
				if i > 0 {
					out.WriteString(", ")
				}
				out.WriteString(rel.String())
			}
			out.WriteString("]\n")
		}
		out.WriteString("    }\n")
	}
	if s.Frequency != "" {
		out.WriteString("    frequency: ")
		out.WriteString(s.Frequency)
		out.WriteString("\n")
	}
	if s.Priority != "" {
		out.WriteString("    priority: ")
		out.WriteString(s.Priority)
		out.WriteString("\n")
	}
	if s.Code != "" {
		out.WriteString("    code: {\n")
		out.WriteString("        ")
		out.WriteString(s.Code)
		out.WriteString("\n    }\n")
	}
	out.WriteString("}")
	return out.String()
}

// Query represents a system's query
type Query struct {
	Components []string
	Relations  []*Relation
}

func (q *Query) TokenLiteral() string { return "query" }
func (q *Query) String() string {
	var parts []string
	parts = append(parts, strings.Join(q.Components, ", "))
	for _, r := range q.Relations {
		parts = append(parts, r.String())
	}
	return fmt.Sprintf("query: (%s);", strings.Join(parts, ", "))
}

// Relation represents a relationship query
type Relation struct {
	Type      string
	Component string
}

func (r *Relation) TokenLiteral() string { return "relation" }
func (r *Relation) String() string {
	return fmt.Sprintf("pair(%s, %s)", r.Type, r.Component)
}

// Type represents a type in the EJECS language
type Type struct {
	Name string
}

type Parameter struct {
	Name         string
	Type         string
	DefaultValue string
}

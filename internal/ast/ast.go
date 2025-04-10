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
	Components    []*Component
	Relationships []*Relationship
	Systems       []*System
}

func (p *Program) TokenLiteral() string { return "program" }
func (p *Program) String() string {
	var out strings.Builder
	for _, c := range p.Components {
		out.WriteString(c.String() + "\n")
	}
	for _, r := range p.Relationships {
		out.WriteString(r.String() + "\n")
	}
	for _, s := range p.Systems {
		out.WriteString(s.String() + "\n")
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
	var fields []string
	for _, f := range c.Fields {
		fields = append(fields, f.String())
	}
	attrs := ""
	if len(c.Attributes) > 0 {
		attrs = "@" + strings.Join(c.Attributes, " @") + "\n"
	}
	return fmt.Sprintf("%scomponent %s {\n    %s\n}", attrs, c.Name, strings.Join(fields, "\n    "))
}

// Field represents a field in a component
type Field struct {
	Name     string
	Type     string
	Optional bool
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
	Name string
	From string
	To   string
	Type string // one_to_one, one_to_many, many_to_many
}

func (r *Relationship) TokenLiteral() string { return "relationship" }
func (r *Relationship) String() string {
	typeStr := ""
	if r.Type != "" {
		typeStr = fmt.Sprintf("@%s\n", r.Type)
	}
	return fmt.Sprintf("%srelationship %s {\n    %s: Entity\n    %s: Entity\n}", typeStr, r.Name, r.From, r.To)
}

// System represents a system declaration
type System struct {
	Name       string
	Query      *Query
	Code       string
	Components []string
	Frequency  string
	Priority   int
}

func (s *System) TokenLiteral() string { return "system" }
func (s *System) String() string {
	var parts []string
	if len(s.Components) > 0 {
		parts = append(parts, fmt.Sprintf("    using %s", strings.Join(s.Components, ", ")))
	}
	if s.Frequency != "" {
		parts = append(parts, fmt.Sprintf("    frequency: %s", s.Frequency))
	}
	if s.Priority != 0 {
		parts = append(parts, fmt.Sprintf("    priority: %d", s.Priority))
	}
	return fmt.Sprintf("system %s {\n%s\n}", s.Name, strings.Join(parts, "\n"))
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
	Type   string
	Target string
}

func (r *Relation) TokenLiteral() string { return "relation" }
func (r *Relation) String() string {
	return fmt.Sprintf("pair(%s, %s)", r.Type, r.Target)
}

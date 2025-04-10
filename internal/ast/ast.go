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
	// Add attributes if present
	if len(c.Attributes) > 0 {
		for i, attr := range c.Attributes {
			out.WriteString("@")
			out.WriteString(attr)
			if i < len(c.Attributes)-1 {
				out.WriteString(" ")
			}
		}
		out.WriteString("\n")
	}
	out.WriteString("component ")
	out.WriteString(c.Name)
	out.WriteString(" {\n")
	for _, field := range c.Fields {
		out.WriteString("    ")
		out.WriteString(field.Name)
		out.WriteString(": ")
		out.WriteString(field.Type)
		if field.Optional {
			out.WriteString("?")
		}
		if field.DefaultValue != nil {
			out.WriteString(" = ")
			out.WriteString(field.DefaultValue.String())
		}
		out.WriteString("\n")
	}
	out.WriteString("}")
	return out.String()
}

// --- Expression Nodes ---

type Expression interface {
	Node // Inherit TokenLiteral() and String() if needed for debugging
	expressionNode()
}

// Identifier represents an identifier used as an expression (e.g., variable name, function name)
type Identifier struct {
	Value string
	// Add token info if needed
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Value } // Placeholder
func (i *Identifier) String() string       { return i.Value }

// Basic Literal types (can reuse existing token literals or define specific nodes)
type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return "STRING" }
func (sl *StringLiteral) String() string       { return fmt.Sprintf("%q", sl.Value) }

type NumberLiteral struct { // Can represent int or float
	Value string // Store as string initially
}

func (nl *NumberLiteral) expressionNode()      {}
func (nl *NumberLiteral) TokenLiteral() string { return "NUMBER" } // Generic
func (nl *NumberLiteral) String() string       { return nl.Value }

type BooleanLiteral struct {
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return "BOOLEAN" }
func (bl *BooleanLiteral) String() string       { return fmt.Sprintf("%t", bl.Value) }

// CallExpression represents a function call like CFrame.new(...)
type CallExpression struct {
	Function  Expression   // The expression being called (e.g., Identifier "CFrame.new")
	Arguments []Expression // List of argument expressions
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return "CALL" }
func (ce *CallExpression) String() string {
	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	return fmt.Sprintf("%s(%s)", ce.Function.String(), strings.Join(args, ", "))
}

// TableConstructor represents a table literal like { key = value, ... }
type TableConstructor struct {
	Fields []*TableField
}

func (tc *TableConstructor) expressionNode()      {}
func (tc *TableConstructor) TokenLiteral() string { return "TABLE_LITERAL" }
func (tc *TableConstructor) String() string {
	var fields []string
	for _, f := range tc.Fields {
		fields = append(fields, f.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
}

// TableField represents a field within a table constructor
type TableField struct {
	Key   Expression // Can be nil for array-like tables, IDENT, or STRING
	Value Expression
}

func (tf *TableField) String() string {
	if tf.Key != nil {
		// TODO: Handle different key types correctly (e.g., ["key"] vs key)
		return fmt.Sprintf("%s = %s", tf.Key.String(), tf.Value.String())
	} else {
		return tf.Value.String()
	}
}

// --- Add PrefixExpression Node ---
type PrefixExpression struct {
	Operator string // e.g., "-", "!"
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Operator }
func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String())
}

// --- Add MemberAccessExpression Node ---
type MemberAccessExpression struct {
	Object     Expression  // The expression on the left of the dot (e.g., Identifier "CFrame")
	MemberName *Identifier // The identifier on the right of the dot (e.g., Identifier "new")
}

func (ma *MemberAccessExpression) expressionNode()      {}
func (ma *MemberAccessExpression) TokenLiteral() string { return "." } // Or "MEMBER_ACCESS"?
func (ma *MemberAccessExpression) String() string {
	return fmt.Sprintf("%s.%s", ma.Object.String(), ma.MemberName.String())
}

// Field represents a field in a component
type Field struct {
	Name         string
	Type         string // Base type (e.g., "int", "Vector3", "table")
	Optional     bool
	MapKeyType   string     // Used if Type is "table"
	MapValueType string     // Used if Type is "table"
	DefaultValue Expression // Changed from string to Expression node
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
	Name       string
	Parameters []*Parameter
	Components []string // DEPRECATED: Use Query field
	Query      *Query
	Frequency  Expression // Changed from string
	Priority   Expression // Changed from string
	Code       string
	Line       int
	Column     int
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
			if param.DefaultValue != nil {
				out.WriteString(" = ")
				out.WriteString(param.DefaultValue.String())
			}
			out.WriteString("\n")
		}
		out.WriteString("    }\n")
	}
	if s.Query != nil {
		out.WriteString("    query: {\n")
		// Always include components and relations lines, even if empty
		out.WriteString("        components: [")
		for i, comp := range s.Query.Components {
			if i > 0 {
				out.WriteString(", ")
			}
			out.WriteString(comp)
		}
		out.WriteString("]\n")
		out.WriteString("        relations: [")
		for i, rel := range s.Query.Relations {
			if i > 0 {
				out.WriteString(", ")
			}
			out.WriteString(rel.String())
		}
		out.WriteString("]\n")
		out.WriteString("    }\n")
	} else if len(s.Components) > 0 { // Fallback for old Components field if Query is nil (for old tests)
		out.WriteString("    using ") // Re-add "using" for compatibility if needed
		out.WriteString(strings.Join(s.Components, ", "))
		out.WriteString("\n")
	}
	if s.Frequency != nil {
		out.WriteString("    frequency: ")
		out.WriteString(s.Frequency.String())
		out.WriteString("\n")
	}
	if s.Priority != nil {
		out.WriteString("    priority: ")
		out.WriteString(s.Priority.String())
		out.WriteString("\n")
	}
	if s.Code != "" {
		out.WriteString("    code: {\n")
		// Basic code indentation
		lines := strings.Split(s.Code, "\n")
		for _, line := range lines {
			out.WriteString("        ")
			out.WriteString(strings.TrimSpace(line))
			out.WriteString("\n")
		}
		out.WriteString("    }\n")
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
	DefaultValue Expression // Changed from string
}

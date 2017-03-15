// Copyright 2017 Pulumi, Inc. All rights reserved.

// Package graph defines CocoGL graphs.  Each graph is directed and acyclic, and the nodes have been topologically
// sorted based on dependencies (edges) between them.  Each node in the graph has a type and a set of properties.
//
// There are two forms of graph: complete and incomplete.  A complete graph is one in which all nodes and their property
// values are known.  An incomplete graph is one where two uncertainties may arise: (1) an edge might be "conditional",
// indicating that its presence or absence is dependent on a piece of information not yet available (like an output
// property from a resource), and/or (2) a property may either be similarly conditional or computed as an output value.
//
// In general, CocoPacks may be evaluated to produce graphs.  These may then be compared to other graphs to produce
// and/or carry out deployment plans.  This package therefore also exposes operations necessary for diffing graphs.
package graph

// Graph is an instance of a CocoGL digraph.  Each is associated with a single CocoPack as its input, along
// with a set of optional arguments used to evaluate it, along with the output DAG with node types and properties.
type Graph interface {
	Roots() []Edge // the root edges.
}

// Vertex is a single vertex within an overall CocoGL graph.
type Vertex interface {
	Data() interface{} // arbitrary data associated with this vertex.
	Label() string     // the vertex's label.
	Ins() []Edge       // incoming edges from other vertices within the graph to this vertex.
	Outs() []Edge      // outgoing edges from this vertex to other vertices within the graph.
}

// Edge is a directed edge from one vertex to another.
type Edge interface {
	Data() interface{} // arbitrary data associated with this edge.
	Label() string     // this edge's label.
	To() Vertex        // the vertex this edge connects to.
	From() Vertex      // the vertex this edge connects from.
}

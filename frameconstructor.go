package main

import "sync"

// similar to what is from package dataframe

// DataFrame - is a table where columns are variables and rows are measurements.
// Each row contains an instance. Each variable can have a different type.
type DataFrame struct {
	// Describes the data.
	Description string `json:"description"`
	// Identifies the batch or data. For example: a session, a file, etc.
	BatchID string `json:"batchid"`
	// Ordered list of variable names.
	VarNames []string `json:"var_names"`
	// Ordered list of variables.
	Data [][]interface{} `json:"data"`
	// Can be used to store custom properties related to the data frame.
	Properties map[string]string `json:"properties"`
	// maps var name to var index for faster access.
	varMap map[string]int
}

// DataGrid implementations represent data addressable by rows and columns.
type DataGrid interface {
	// Retrieves a given Attribute's specification
	GetAttribute(Attribute) (AttributeSpec, error)
	// Retrieves details of every Attribute
	AllAttributes() []Attribute
	// Marks an Attribute as a class Attribute
	AddClassAttribute(Attribute) error
	// Unmarks an Attribute as a class Attribute
	RemoveClassAttribute(Attribute) error
	// Returns details of all class Attributes
	AllClassAttributes() []Attribute
	// Gets the bytes at a given position or nil
	Get(AttributeSpec, int) []byte
	// Convenience function for iteration.
	MapOverRows([]AttributeSpec, func([][]byte, int) (bool, error)) error
}

// similar to what is used in Go Learn
// see gophernotes example notebook for more

// DenseInstances stores each Attribute value explicitly
// in a large grid.
type DenseInstances struct {
	agMap        map[string]int
	agRevMap     map[int]string
	ags          []AttributeGroup
	lock         sync.Mutex
	fixed        bool
	classAttrs   map[AttributeSpec]bool
	maxRow       int
	attributes   []Attribute
	tmpAttrAgMap map[Attribute]string
	// Counters for each AttributeGroup type
	floatRowSizeBytes int
	catRowSizeBytes   int
	binRowSizeBits    int
}

func Dataframe() {

	// meat

}

func ReadCSV() {

	// meat

}

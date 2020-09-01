package model

type Patch struct{
	OP string
	Path string
	Values []interface{}
}
package attr

import (
	"fmt"
)

// Int create int attributer with generated function
func Int(name string, genFunc func() int, options ...string) Attributer {
	return &intAttribute{
		name:    name,
		colName: getColName(options),
		genFunc: genFunc,
	}
}

type intAttribute struct {
	name    string
	colName string
	val     int
	genFunc func() int
	process Processor
	object  interface{}
}

func (attr *intAttribute) GetObject() interface{} {
	return attr.object
}

func (attr *intAttribute) Process(procFunc Processor) Attributer {
	attr.process = procFunc
	return attr
}

func (attr intAttribute) GetVal() interface{} {
	return attr.val
}

func (attr *intAttribute) SetVal(val interface{}) error {
	realVal, ok := val.(int)
	if !ok {
		return fmt.Errorf("set attribute val: val %+v is not int", val)
	}

	attr.val = realVal
	return nil
}

func (attr intAttribute) ColName() string {
	return attr.colName
}

func (attr intAttribute) Name() string {
	return attr.name
}

func (intAttribute) Kind() Type {
	return IntAttr
}

func (attr *intAttribute) Gen(data interface{}) (interface{}, error) {
	attr.val = attr.genFunc()
	attr.object = data
	if attr.process != nil {
		if err := attr.process(attr); err != nil {
			return nil, err
		}
	}

	return attr.val, nil
}

// SeqInt create int attributer with a slice of int
func SeqInt(name string, intSet []int, options ...string) Attributer {
	return &seqIntAttr{
		intSet:  intSet,
		index:   0,
		name:    name,
		colName: getColName(options),
	}
}

type seqIntAttr struct {
	intSet  []int
	index   int
	val     int
	name    string
	colName string
	process Processor
	object  interface{}
}

func (attr *seqIntAttr) Process(procFunc Processor) Attributer {
	attr.process = procFunc
	return attr
}

func (attr *seqIntAttr) GetObject() interface{} {
	return attr.object
}

func (attr seqIntAttr) GetVal() interface{} {
	return attr.val
}

func (attr *seqIntAttr) SetVal(val interface{}) error {
	realVal, ok := val.(int)
	if !ok {
		return fmt.Errorf("set attribute val: val %+v is not int", val)
	}
	attr.val = realVal
	return nil
}

func (attr seqIntAttr) ColName() string {
	return attr.colName
}

func (attr *seqIntAttr) Gen(data interface{}) (interface{}, error) {
	attr.val = attr.intSet[attr.index]
	attr.object = data
	if attr.process != nil {
		if err := attr.process(attr); err != nil {
			return nil, err
		}
	}

	attr.index++
	if attr.index >= len(attr.intSet) {
		attr.index = 0
	}
	return attr.val, nil
}

func (seqIntAttr) Kind() Type {
	return IntAttr
}

func (attr seqIntAttr) Name() string {
	return attr.name
}

// Seq create int attributer with sequential number from start parameter
func Seq(name string, start int, options ...string) Attributer {

	return &seqAttr{
		name:    name,
		colName: getColName(options),
		seq:     start,
	}
}

type seqAttr struct {
	seq     int
	val     int
	name    string
	colName string
	process Processor
	object  interface{}
}

func (attr *seqAttr) Process(procFunc Processor) Attributer {
	attr.process = procFunc
	return attr
}

func (attr *seqAttr) GetObject() interface{} {
	return attr.object
}

func (attr seqAttr) GetVal() interface{} {
	return attr.val
}

func (attr *seqAttr) SetVal(val interface{}) error {
	realVal, ok := val.(int)
	if !ok {
		return fmt.Errorf("set attribute val: val %+v is not int", val)
	}
	attr.val = realVal
	return nil
}

func (attr seqAttr) ColName() string {
	return attr.colName
}

func (attr *seqAttr) Gen(data interface{}) (interface{}, error) {
	attr.val = attr.seq
	attr.object = data
	if attr.process != nil {
		if err := attr.process(attr); err != nil {
			return nil, err
		}
	}
	attr.seq++

	return attr.val, nil
}

func (seqAttr) Kind() Type {
	return IntAttr
}

func (attr seqAttr) Name() string {
	return attr.name
}

// Float create float attributer with generated function
func Float(name string, genFunc func() float64, options ...string) Attributer {
	return &floatAttr{
		name:    name,
		colName: getColName(options),
		genFunc: genFunc,
	}
}

type floatAttr struct {
	name    string
	colName string
	val     float64
	genFunc func() float64
	process Processor
	object  interface{}
}

func (attr *floatAttr) Process(procFunc Processor) Attributer {
	attr.process = procFunc
	return attr
}

func (attr *floatAttr) GetObject() interface{} {
	return attr.object
}

func (attr floatAttr) GetVal() interface{} {
	return attr.val
}

func (attr *floatAttr) SetVal(val interface{}) error {
	realVal, ok := val.(float64)
	if !ok {
		return fmt.Errorf("set attribute val: val %+v is not float64", val)
	}

	attr.val = realVal
	return nil
}

func (attr floatAttr) ColName() string {
	return attr.colName
}

func (attr *floatAttr) Gen(data interface{}) (interface{}, error) {
	attr.val = attr.genFunc()
	attr.object = data
	if attr.process != nil {
		if err := attr.process(attr); err != nil {
			return nil, err
		}
	}

	return attr.val, nil
}

func (floatAttr) Kind() Type {
	return FloatAttr
}

func (attr floatAttr) Name() string {
	return attr.name
}

// Uint create uint attributer with generated function
func Uint(name string, genFunc func() uint, options ...string) Attributer {
	return &uintAttr{
		name:    name,
		colName: getColName(options),
		genFunc: genFunc,
	}
}

type uintAttr struct {
	name    string
	colName string
	val     uint
	genFunc func() uint
	process Processor
	object  interface{}
}

func (attr *uintAttr) Process(procFunc Processor) Attributer {
	attr.process = procFunc
	return attr
}

func (attr *uintAttr) GetObject() interface{} {
	return attr.object
}

func (attr uintAttr) GetVal() interface{} {
	return attr.val
}

func (attr *uintAttr) SetVal(val interface{}) error {
	realVal, ok := val.(uint)
	if !ok {
		return fmt.Errorf("set attribute val: val %+v is not uint", val)
	}

	attr.val = realVal
	return nil
}

func (attr uintAttr) ColName() string {
	return attr.colName
}

func (attr *uintAttr) Gen(data interface{}) (interface{}, error) {
	attr.val = attr.genFunc()
	attr.object = data
	if attr.process != nil {
		if err := attr.process(attr); err != nil {
			return nil, err
		}
	}
	return attr.val, nil
}

func (uintAttr) Kind() Type {
	return UintAttr
}

func (attr uintAttr) Name() string {
	return attr.name
}

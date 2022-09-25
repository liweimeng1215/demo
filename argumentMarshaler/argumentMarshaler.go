package argumentMarshaler

import "polymorphism/arglist"

type ArgumentMarshaler interface {
	Set(arg *arglist.ArgList)
}

type BooleanArgumentMarshaler struct {
	boolValue bool
}

func (bam *BooleanArgumentMarshaler) Set(arg *arglist.ArgList) {
	bam.boolValue = true
}

type StringArgumentMarsheler struct {
	stringValue string
}

func (sam *StringArgumentMarsheler) Set(arg *arglist.ArgList) {
	sam.stringValue = arg.Next()
}

func GetBoolValue(obj interface{}) bool {
	if boolObj, ok := obj.(*BooleanArgumentMarshaler); ok {
		return boolObj.boolValue
	}
	return false
}

func GetStringValue(obj interface{}) string {
	if stringObj, ok := obj.(*StringArgumentMarsheler); ok {
		return stringObj.stringValue
	}
	return ""
}

package args

import (
	"polymorphism/arglist"
	"polymorphism/argumentMarshaler"
	"polymorphism/log"
	"strings"
)

type Args struct {
	schema               string
	args                 *arglist.ArgList
	currentArgumentIndex int
	marshalers           map[string]argumentMarshaler.ArgumentMarshaler
}

func New() Args {
	return Args{
		marshalers: make(map[string]argumentMarshaler.ArgumentMarshaler),
	}
}

// Parse 解析参数规则和用户输入 schema: "l,d*"  args: [-l -d /a/b/c]
func (a *Args) Parse(schema string, args []string) {
	a.schema = schema
	a.args = arglist.New(args)
	a.parseSchema()
	a.parseArgs()
}

func (a *Args) parseSchema() {
	for _, element := range strings.Split(a.schema, ",") {
		a.parseSchemaElement(strings.TrimSpace(element))
	}
}

func (a *Args) parseSchemaElement(element string) {
	elementId := element[:1]
	elementTail := element[1:]
	if len(elementTail) == 0 {
		a.marshalers[elementId] = &argumentMarshaler.BooleanArgumentMarshaler{}
	} else if elementTail == "*" {
		a.marshalers[elementId] = &argumentMarshaler.StringArgumentMarsheler{}
	} else {
		log.Fetal("elementTail not support. elementTail: %s", elementTail)
	}
}

func (a *Args) parseArgs() {
	for a.args.HasNext() {
		success := a.parseArg(a.args.Next())
		if !success {
			break
		}
	}
}

func (a *Args) parseArg(arg string) bool {
	success := true
	if strings.HasPrefix(arg, "-") {
		elementId := arg[1:]
		if a.marshalers[elementId] == nil {
			log.Error("elementId not set in marshalers. elementId: %s", elementId)
			success = false
		}
		a.marshalers[elementId].Set(a.args)
	}
	return success
}

func (a *Args) GetBool(elementId string) bool {
	return argumentMarshaler.GetBoolValue(a.marshalers[elementId])
}

func (a *Args) GetString(elementId string) string {
	return argumentMarshaler.GetStringValue(a.marshalers[elementId])
}

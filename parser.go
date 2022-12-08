package Parser

import (
	"flag"
	"reflect"
	"strconv"
)

const TAG_IGNORE string = "ignore"
const TAG_DEFAUTL_VALUE string = "default"
const TAG_USAGE_VALUE string = "usage"
const TAG_ALIAS string = "alias"

type argParser struct {
	argStruct      interface{}
	pointers       []reflect.Value
	reflectedValue reflect.Value
}

func (ap *argParser) parseConsoleArgs() {
	for i := 0; i < len(ap.pointers); i++ {
		if ap.getTag(i, TAG_IGNORE) == "yes" {
			continue
		}
		ap.setFlag(i)
	}
	flag.Parse()
}
func (ap *argParser) setFlag(fIndex int) {
	switch ap.getKind(fIndex) {
	case reflect.Int:
		flag.IntVar(
			ap.getInterfaceOnValue(fIndex).(*int),
			ap.getFieldName(fIndex),
			ap.getTagAsInt(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.IntVar(
				ap.getInterfaceOnValue(fIndex).(*int),
				alias,
				ap.getTagAsInt(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Int8:
		flag.Var(
			newIntXBitsValue(
				ap.getTagAsInt8(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*int8),
				8,
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newIntXBitsValue(
					ap.getTagAsInt8(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*int8),
					8,
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Int16:
		flag.Var(
			newIntXBitsValue(
				ap.getTagAsInt16(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*int16),
				16,
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newIntXBitsValue(
					ap.getTagAsInt16(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*int16),
					16,
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Int32:
		flag.Var(
			newIntXBitsValue(
				ap.getTagAsInt32(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*int32),
				32,
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newIntXBitsValue(
					ap.getTagAsInt32(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*int32),
					32,
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Int64:
		flag.Int64Var(
			ap.getInterfaceOnValue(fIndex).(*int64),
			ap.getFieldName(fIndex),
			ap.getTagAsInt64(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Int64Var(
				ap.getInterfaceOnValue(fIndex).(*int64),
				alias,
				ap.getTagAsInt64(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Uint:
		flag.UintVar(
			ap.getInterfaceOnValue(fIndex).(*uint),
			ap.getFieldName(fIndex),
			ap.getTagAsUint(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.UintVar(
				ap.getInterfaceOnValue(fIndex).(*uint),
				alias,
				ap.getTagAsUint(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Uint8:
		flag.Var(
			newUintXBitsValue(
				ap.getTagAsUint8(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*uint8),
				8,
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newUintXBitsValue(
					ap.getTagAsUint8(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*uint8),
					8,
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Uint16:
		flag.Var(
			newUintXBitsValue(
				ap.getTagAsUint16(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*uint16),
				16,
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newUintXBitsValue(
					ap.getTagAsUint16(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*uint16),
					16,
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Uint32:
		flag.Var(
			newUintXBitsValue(
				ap.getTagAsUint32(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*uint32),
				32,
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newUintXBitsValue(
					ap.getTagAsUint32(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*uint32),
					32,
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Uint64:
		flag.Uint64Var(
			ap.getInterfaceOnValue(fIndex).(*uint64),
			ap.getFieldName(fIndex),
			ap.getTagAsUint64(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Uint64Var(
				ap.getInterfaceOnValue(fIndex).(*uint64),
				alias,
				ap.getTagAsUint64(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Float32:
		flag.Var(
			newFloat32Value(
				ap.getTagAsFloat32(fIndex, TAG_DEFAUTL_VALUE),
				ap.getInterfaceOnValue(fIndex).(*float32),
			),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Var(
				newFloat32Value(
					ap.getTagAsFloat32(fIndex, TAG_DEFAUTL_VALUE),
					ap.getInterfaceOnValue(fIndex).(*float32),
				),
				alias,
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Float64:
		flag.Float64Var(
			ap.getInterfaceOnValue(fIndex).(*float64),
			ap.getFieldName(fIndex),
			ap.getTagAsFloat64(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.Float64Var(
				ap.getInterfaceOnValue(fIndex).(*float64),
				alias,
				ap.getTagAsFloat64(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.String:
		flag.StringVar(
			ap.getInterfaceOnValue(fIndex).(*string),
			ap.getFieldName(fIndex),
			ap.getTagAsString(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.StringVar(
				ap.getInterfaceOnValue(fIndex).(*string),
				alias,
				ap.getTagAsString(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	case reflect.Bool:
		flag.BoolVar(
			ap.getInterfaceOnValue(fIndex).(*bool),
			ap.getFieldName(fIndex),
			ap.getTagAsBool(fIndex, TAG_DEFAUTL_VALUE),
			ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
		)
		var alias string
		if alias = ap.getTagAsString(fIndex, TAG_ALIAS); alias != "" {
			flag.BoolVar(
				ap.getInterfaceOnValue(fIndex).(*bool),
				alias,
				ap.getTagAsBool(fIndex, TAG_DEFAUTL_VALUE),
				ap.getTagAsString(fIndex, TAG_USAGE_VALUE),
			)
		}
	}
}
func (ap *argParser) getInterfaceOnValue(fIndex int) interface{} {
	if ap.pointers[fIndex].Type().Kind() == reflect.Ptr {
		return ap.pointers[fIndex].Interface()
	} else {
		return ap.pointers[fIndex].Addr().Interface()
	}
}
func (ap *argParser) getKind(fIndex int) (kind reflect.Kind) {
	kind = ap.pointers[fIndex].Type().Kind()
	if kind == reflect.Ptr {
		//Getting the field kind behind a pointer
		kind = ap.pointers[fIndex].Type().Elem().Kind()
		//Initializing the point when it is nil
		ap.initializePtr(fIndex)
	}
	return
}
func (ap *argParser) analyzeStructObject() {
	ap.reflectedValue = reflect.ValueOf(ap.argStruct)
	nv := ap.reflectedValue.Type().Elem().NumField()
	ap.pointers = make([]reflect.Value, nv)
	for i := 0; i < nv; i++ {
		ap.pointers[i] = ap.reflectedValue.Elem().Field(i)
	}
}
func (ap *argParser) initializePtr(fIndex int) {
	var newValue reflect.Value
	if ap.pointers[fIndex].Elem().Kind() == reflect.Invalid {
		newValue = reflect.New(ap.pointers[fIndex].Type().Elem())
		ap.pointers[fIndex].Set(newValue)
	}
}

func (ap *argParser) getTag(fIndex int, tagName string) string {
	var tagValue string
	tagValue, _ = ap.reflectedValue.Elem().Type().Field(fIndex).Tag.Lookup(tagName)
	return tagValue
}

func (ap *argParser) getTagAsString(fIndex int, tagName string) string {
	return ap.getTag(fIndex, tagName)
}

func (ap *argParser) getTagAsInt(fIndex int, tagName string) int {
	var err error
	var tag string
	var tagInt int
	tag = ap.getTag(fIndex, tagName)
	if tag == "" {
		return 0
	} else {
		if tagInt, err = strconv.Atoi(tag); err != nil {
			return 0
		}
		return tagInt
	}
}
func (ap *argParser) getTagAsInt8(fIndex int, tagName string) int8 {
	return int8(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsInt16(fIndex int, tagName string) int16 {
	return int16(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsInt32(fIndex int, tagName string) int32 {
	return int32(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsInt64(fIndex int, tagName string) int64 {
	return int64(ap.getTagAsInt(fIndex, tagName))
}

func (ap *argParser) getTagAsUint(fIndex int, tagName string) uint {
	return uint(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsUint8(fIndex int, tagName string) uint8 {
	return uint8(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsUint16(fIndex int, tagName string) uint16 {
	return uint16(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsUint32(fIndex int, tagName string) uint32 {
	return uint32(ap.getTagAsInt(fIndex, tagName))
}
func (ap *argParser) getTagAsUint64(fIndex int, tagName string) uint64 {
	return uint64(ap.getTagAsInt(fIndex, tagName))
}

func (ap *argParser) getTagAsFloat(fIndex int, tagName string, bitSize int) interface{} {
	var err error
	var tag string
	var tagFloat float64
	tag = ap.getTag(fIndex, tagName)
	if tag == "" {
		return 0
	} else {
		if tagFloat, err = strconv.ParseFloat(tag, bitSize); err != nil {
			return 0
		}
		return tagFloat
	}
}
func (ap *argParser) getTagAsFloat32(fIndex int, tagName string) float32 {
	return float32(ap.getTagAsFloat(fIndex, tagName, 32).(float64))
}
func (ap *argParser) getTagAsFloat64(fIndex int, tagName string) float64 {
	return ap.getTagAsFloat(fIndex, tagName, 64).(float64)
}

func (ap *argParser) getTagAsBool(fIndex int, tagName string) bool {
	var err error
	var tag string
	var tagBool bool
	tag = ap.getTag(fIndex, tagName)
	if tag == "" {
		return false
	} else {
		if tagBool, err = strconv.ParseBool(tag); err != nil {
			return false
		}
		return tagBool
	}
}

func (ap *argParser) tagExists(fIndex int, tagName string) bool {
	var tagExists bool
	_, tagExists = ap.reflectedValue.Elem().Type().Field(fIndex).Tag.Lookup(tagName)
	return tagExists
}
func (ap *argParser) getFieldName(fIndex int) string {
	return ap.reflectedValue.Elem().Type().Field(fIndex).Name
}

func Parse(ArgStructPtr interface{}) {
	var ap argParser
	ap.argStruct = ArgStructPtr
	ap.analyzeStructObject()
	ap.parseConsoleArgs()
}

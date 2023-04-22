package utils

import (
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUUID() string {
	return uuid.New().String()
}

type InputMap map[string]interface{}

func (i InputMap) Create() {
	i["_id"] = GenerateUUID()
	i["createdAt"] = time.Now()
	i["updatedAt"] = time.Now()
}

func (i InputMap) Update() {
	i["updatedAt"] = time.Now()
}

type Method string

const (
	MethodCreate Method = "create"
	MethodUpdate Method = "update"
	MethodFilter Method = "filter"
)

// ToMap converts a struct or an object to a map.
func ToMap(from interface{}, method ...Method) InputMap {
	if len(method) > 1 {
		log.Panic("ToMap only accepts one method")
	}

	jsonFields := make(map[string]interface{})

	switch from.(type) {
	// if the object is a bson primitive object then we can use the bson.Mapper
	case primitive.M, primitive.D, primitive.E, primitive.A, primitive.Regex:
		marshaled, err := bson.Marshal(from)
		if err != nil {
			log.Panic("Could not marshal primitives toMap")
		}

		err = bson.Unmarshal(marshaled, &jsonFields)
		if err != nil {
			log.Panic("Could not Unmarshal primitives toMap")
		}

		return jsonFields
	default:
		resultFields := make(map[string]interface{})
		// if the object is a struct then we can use the structs.Mapper
		// we extract the fields from the struct, then we use the structs.Mapper to convert the fields to a map
		fields := structs.Fields(from)

		// if update flag is passed we use flattenNestMap, then we call Update
		if len(method) > 0 {
			switch method[0] {
			case MethodCreate:
				nestMap(fields, resultFields)
				InputMap(resultFields).Create()
			case MethodUpdate:
				flattenNestMap("", fields, resultFields)
				InputMap(resultFields).Update()
			case MethodFilter:
				flattenNestMap("", fields, resultFields)

				if val, ok := resultFields["id"]; ok {
					resultFields["_id"] = val
					delete(resultFields, "id")
				}
			}
		} else {
			nestMap(fields, resultFields)
			InputMap(resultFields).Create()
		}

		return resultFields
	}
}

// nestMap converts a struct to a nested map.
func nestMap(src []*structs.Field, dest map[string]interface{}) {
	for _, v := range src {
		bsonTag := v.Tag("bson")

		if bsonTag == "-" || v.Name() == "ID" || v.Name() == "_id" {
			continue
		}

		// switch based on the field kind
		switch v.Kind() { //nolint:exhaustive//reason: the use of default here is enough to cover all cases
		case reflect.Struct:
			// if object is a struct we get the object fields
			// we call nestMap on the fields
			processSpecialType(v, dest)
		case reflect.Ptr:
			// if object is a pointer we check if tests's empty
			if !v.IsZero() {
				val := v.Value()
				// we get the element of the pointer
				ptVal := reflect.ValueOf(val).Elem()

				// we check if the element is a struct
				if ptVal.Kind() == reflect.Struct {
					// if tests is we call nestMap on the fields
					processSpecialType(v, dest)
				} else {
					// if tests's not we just set the value
					dest[strcase.ToLowerCamel(v.Name())] = val
				}

				break
			}

		default:
			// by default, we set the value
			dest[strcase.ToLowerCamel(v.Name())] = v.Value()
		}
	}
}

// flattenNestMap converts a struct to a flat map.
func flattenNestMap(prefix string, src []*structs.Field, dest map[string]interface{}) {
	if len(prefix) > 0 {
		prefix += "."
	}

	for _, v := range src {
		// ignore the id field
		if v.Value() == "id" || v.Value() == "_id" {
			continue
		}

		// switch based on the field kind
		switch v.Kind() { //nolint:exhaustive//reason: the use of default here is enough to cover all cases
		case reflect.Struct:
			// if object is a struct we get the object fields
			// we call flattenNestMap on the fields
			processSpecialType(v, dest, prefix)
		case reflect.Ptr:
			// if object is a pointer we check if tests's empty
			if !v.IsZero() {
				val := v.Value()
				// we get the element of the pointer
				ptVal := reflect.ValueOf(val).Elem()

				// we check if the element is a struct
				if ptVal.Kind() == reflect.Struct {
					processSpecialType(v, dest, prefix)
				} else {
					// if tests's not we just set the value
					dest[prefix+strcase.ToLowerCamel(v.Name())] = val
				}
				// if the pointer is empty we ignore the value
				// to avoid updating values to nil
				break
			}
		default:
			if v.IsZero() {
				continue
			}
			// by default, we set the value
			dest[prefix+strcase.ToLowerCamel(v.Name())] = v.Value()
		}
	}
}

func processSpecialType(v *structs.Field, dest map[string]interface{}, prefix ...string) {
	flat := false
	sPrefix := ""
	inline := false

	if bsonTag := v.Tag("bson"); strings.Contains(bsonTag, "inline") {
		inline = true
	}

	if !inline && len(prefix) > 0 {
		sPrefix = prefix[0]
		flat = true
	}

	switch v.Value().(type) {
	case time.Time:
		if flat {
			dest[sPrefix+strcase.ToLowerCamel(v.Name())] = v.Value()
		} else {
			dest[strcase.ToLowerCamel(v.Name())] = v.Value()
		}
	default:
		if flat {
			// if tests is we call flattenNestMap on the fields
			flattenNestMap(
				sPrefix+strcase.ToLowerCamel(v.Name()),
				v.Fields(),
				dest,
			)
		} else {
			if inline {
				nestMap(v.Fields(), dest)
			} else {
				dest[strcase.ToLowerCamel(v.Name())] = map[string]interface{}{}
				nestMap(v.Fields(), dest[strcase.ToLowerCamel(v.Name())].(map[string]interface{}))
			}
		}
	}
}

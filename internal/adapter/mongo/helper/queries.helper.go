package helper

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuildQuery(cond map[string]interface{}) bson.M {
	if cond == nil {
		return bson.M{}
	}

	out := bson.M{}
	for k, v := range cond {
		if isZero(v) {
			continue
		}

		if (k == "_id" || strings.HasSuffix(k, "_id")) && isHex24String(v) {
			if oid, err := primitive.ObjectIDFromHex(v.(string)); err == nil {
				out[k] = oid
				continue
			}
		}

		out[k] = v
	}

	return out
}

func isHex24String(v any) bool {
	s, ok := v.(string)
	if !ok || len(s) != 24 {
		return false
	}
	for i := 0; i < 24; i++ {
		c := s[i]
		if !((c >= '0' && c <= '9') ||
			(c >= 'a' && c <= 'f') ||
			(c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

func isZero(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return rv.IsNil()
	}
	// ค่าอื่น ๆ ใช้ zero ของ type เทียบ
	z := reflect.Zero(rv.Type())
	return reflect.DeepEqual(v, z.Interface())
}

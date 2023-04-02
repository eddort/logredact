package logredact

import (
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

type LogRedact struct {
	secrets  []string
	replacer string
}

func New(secrets []string, replacer string) *LogRedact {
	return &LogRedact{secrets: secrets, replacer: replacer}
}

func (h *LogRedact) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogRedact) Fire(entry *logrus.Entry) error {
	entry.Message = h.replaceSecrets(entry.Message)

	for key, value := range entry.Data {
		entry.Data[key] = h.processValue(reflect.ValueOf(value))
	}
	return nil
}

func (h *LogRedact) processValue(v reflect.Value) interface{} {
	if !v.IsValid() {
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		return h.replaceSecrets(v.String())

	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		elem := v.Elem()
		newElem := reflect.New(elem.Type())
		h.processValueRecursively(elem, newElem.Elem())
		return newElem.Interface()

	case reflect.Struct:
		newStruct := reflect.New(v.Type()).Elem()
		h.processValueRecursively(v, newStruct)
		return newStruct.Interface()

	case reflect.Slice:
		newSlice := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
		for i := 0; i < v.Len(); i++ {
			newSlice.Index(i).Set(reflect.ValueOf(h.processValue(v.Index(i))))
		}
		return newSlice.Interface()
	}

	return v.Interface()
}

func (h *LogRedact) processValueRecursively(src, dest reflect.Value) {
	for i := 0; i < src.NumField(); i++ {
		dest.Field(i).Set(reflect.ValueOf(h.processValue(src.Field(i))))
	}
}

func (h *LogRedact) replaceSecrets(s string) string {
	for _, secret := range h.secrets {
		s = strings.Replace(s, secret, h.replacer, -1)
	}
	return s
}

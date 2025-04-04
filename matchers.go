package matchers

import (
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"google.golang.org/protobuf/proto"
)

// ProtoEqualMatcher is a custom Gomega matcher for protocol buffers
type ProtoEqualMatcher struct {
	Expected proto.Message
}

func (matcher *ProtoEqualMatcher) Match(actual interface{}) (success bool, err error) {
	actualMessage, ok := actual.(proto.Message)
	if !ok {
		return false, nil
	}
	return proto.Equal(actualMessage, matcher.Expected), nil
}

func (matcher *ProtoEqualMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to equal", matcher.Expected)
}

func (matcher *ProtoEqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to equal", matcher.Expected)
}

// ProtoEqual returns a Gomega matcher that checks if a Protobuf message is equal to the expected message.
// It verifies that the input implements proto.Message and matches the expected message using proto.Equal.
// If the input is not a proto.Message or does not match the expected message, the matcher fails.
//
// This matcher is useful for testing scenarios where a single Protobuf message should exactly match an expected
// template. It performs a deep comparison of all fields, including nested messages, as defined by proto.Equal.
//
// Example usage:
//
//	expected := &v1.Foo{Bar: "test", Baz: "baz"}
//	item := &v1.Foo{Bar: "test", Baz: "baz"}
//	Expect(item).To(ProtoEqual(expected)) // Passes
//	mismatched := &v1.Foo{Bar: "different", Baz: "baz"}
//	Expect(mismatched).ToNot(ProtoEqual(expected)) // Passes
func ProtoEqual(expected proto.Message) types.GomegaMatcher {
	return &ProtoEqualMatcher{
		Expected: expected,
	}
}

// ProtoContainsMatcher is a custom Gomega matcher to check if a slice of protocol buffers contains specific elements
type ProtoContainsMatcher struct {
	Elements []proto.Message
}

func (matcher *ProtoContainsMatcher) Match(actual interface{}) (success bool, err error) {
	v := reflect.ValueOf(actual)
	if v.Kind() != reflect.Slice {
		return false, nil
	}

	// Convert actual slice to []proto.Message
	actualSlice := make([]proto.Message, v.Len())
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		actualMessage, ok := elem.(proto.Message)
		if !ok {
			return false, nil
		}
		actualSlice[i] = actualMessage
	}

	// Check if all expected elements are present in the actual slice
	for _, expectedElem := range matcher.Elements {
		found := false
		for _, actualElem := range actualSlice {
			if proto.Equal(actualElem, expectedElem) {
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}
	return true, nil
}

func (matcher *ProtoContainsMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to contain elements", matcher.Elements)
}

func (matcher *ProtoContainsMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to contain elements", matcher.Elements)
}

// ProtoContains returns a Gomega matcher that checks if a slice of Protobuf messages contains all the specified elements.
// It verifies that the input is a slice of proto.Message implementations and that each specified element is present in the
// actual slice (in any order), using proto.Equal for comparisons. The actual slice may contain additional elements beyond
// those specified.
//
// This matcher is useful for testing scenarios where you want to ensure certain Protobuf messages are present in a
// collection, regardless of order or additional elements. It accepts variadic arguments, allowing individual elements to
// be passed directly.
//
// Example usage:
//
//	Expect(items).To(ProtoContains(
//	    &v1.Foo{Bar: "test1"},
//	    &v1.Foo{Bar: "test2"},
//	)) // Passes if items contains both elements (and possibly more)
//	Expect(missing).ToNot(ProtoContains(
//	    &v1.Foo{Bar: "test1"},
//	    &v1.Foo{Bar: "test2"},
//	)) // Passes if missing lacks at least one of these elements
func ProtoContains(elements ...proto.Message) types.GomegaMatcher {
	return &ProtoContainsMatcher{
		Elements: elements,
	}
}

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

// AllProtoEqualMatcher is a custom Gomega matcher to check if all elements in a slice match a protocol buffer
type AllProtoEqualMatcher struct {
	Expected proto.Message
}

func (matcher *AllProtoEqualMatcher) Match(actual interface{}) (success bool, err error) {
	v := reflect.ValueOf(actual)
	if v.Kind() != reflect.Slice {
		return false, nil // Return false, nil for non-slice types, consistent with ProtoEqualMatcher
	}

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		actualMessage, ok := elem.(proto.Message)
		if !ok {
			return false, nil // Return false if an element isn’t a proto.Message
		}
		if !proto.Equal(actualMessage, matcher.Expected) {
			return false, nil
		}
	}
	return true, nil
}

func (matcher *AllProtoEqualMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have all elements equal", matcher.Expected)
}

func (matcher *AllProtoEqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have all elements equal", matcher.Expected)
}

// AllProtoEqual returns a Gomega matcher that checks if all elements in a slice are equal to the expected Protobuf message.
// It verifies that the input is a slice and that each element implements proto.Message and matches the expected message
// using proto.Equal. If the input is not a slice or any element does not match, the matcher fails.
//
// This matcher is useful for testing scenarios where a collection of Protobuf messages should all conform to a single
// expected template, ignoring order. An empty slice is considered a match (vacuously true), and non-slice inputs
// result in a mismatch without error.
//
// Example usage:
//
//	expected := &v1.Foo{Bar: "test", Baz: "baz"}
//	items := []*v1.Foo{{Bar: "test", Baz: "baz"}, {Bar: "test", Baz: "baz"}}
//	Expect(items).To(AllProtoEqual(expected)) // Passes
//	mismatched := []*v1.Foo{{Bar: "test", Baz: "baz"}, {Bar: "different", Baz: "baz"}}
//	Expect(mismatched).ToNot(AllProtoEqual(expected)) // Passes
func AllProtoEqual(expected proto.Message) types.GomegaMatcher {
	return &AllProtoEqualMatcher{
		Expected: expected,
	}
}

// ProtoSliceEqualMatcher is a custom Gomega matcher to check if a slice of protocol buffers matches an expected slice
type ProtoSliceEqualMatcher struct {
	Expected []proto.Message
}

func (matcher *ProtoSliceEqualMatcher) Match(actual interface{}) (success bool, err error) {
	v := reflect.ValueOf(actual)
	if v.Kind() != reflect.Slice {
		return false, nil
	}

	if v.Len() != len(matcher.Expected) {
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

	// Check if slices match in any order
	return slicesMatch(actualSlice, matcher.Expected), nil
}

// slicesMatch checks if two slices contain the same elements in any order
func slicesMatch(actual, expected []proto.Message) bool {
	if len(actual) != len(expected) {
		return false
	}

	// Create a map to track matched expected elements
	matched := make([]bool, len(expected))

	// For each actual element, try to find a matching expected element
	for _, actualMsg := range actual {
		foundMatch := false
		for j, expectedMsg := range expected {
			if !matched[j] && proto.Equal(actualMsg, expectedMsg) {
				matched[j] = true
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			return false
		}
	}
	return true
}

func (matcher *ProtoSliceEqualMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to equal slice", matcher.Expected)
}

func (matcher *ProtoSliceEqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to equal slice", matcher.Expected)
}

// ProtoSliceEqual returns a Gomega matcher that checks if a slice of Protobuf messages matches the expected slice.
// It verifies that the input is a slice of proto.Message implementations and contains the same elements as the
// expected slice in any order, using proto.Equal for comparisons. The slices must be the same length and contain
// equivalent messages, but order doesn't matter.
//
// This matcher is useful for testing scenarios where a collection of Protobuf messages should match an expected
// set of messages regardless of their order.
//
// Example usage:
//
//	expected := []proto.Message{&v1.Foo{Bar: "test"}, &v1.Foo{Bar: "baz"}}
//	items := []proto.Message{&v1.Foo{Bar: "baz"}, &v1.Foo{Bar: "test"}}
//	Expect(items).To(ProtoSliceEqual(expected)) // Passes
//	mismatched := []proto.Message{&v1.Foo{Bar: "test"}, &v1.Foo{Bar: "different"}}
//	Expect(mismatched).ToNot(ProtoSliceEqual(expected)) // Passes
func ProtoSliceEqual(expected []proto.Message) types.GomegaMatcher {
	return &ProtoSliceEqualMatcher{
		Expected: expected,
	}
}

// ProtoSliceContainsMatcher is a custom Gomega matcher to check if a slice of protocol buffers contains specific elements
type ProtoSliceContainsMatcher struct {
	Elements []proto.Message
}

func (matcher *ProtoSliceContainsMatcher) Match(actual interface{}) (success bool, err error) {
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

func (matcher *ProtoSliceContainsMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to contain elements", matcher.Elements)
}

func (matcher *ProtoSliceContainsMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to contain elements", matcher.Elements)
}

// ProtoSliceContains returns a Gomega matcher that checks if a slice of Protobuf messages contains all the specified elements.
// It verifies that the input is a slice of proto.Message implementations and that each element in the expected slice is
// present in the actual slice (in any order), using proto.Equal for comparisons. The actual slice may contain additional
// elements beyond those specified.
//
// This matcher is useful for testing scenarios where you want to ensure certain Protobuf messages are present in a
// collection, regardless of order or additional elements.
//
// Example usage:
//
//	expected := []proto.Message{&v1.Foo{Bar: "test1"}, &v1.Foo{Bar: "test2"}}
//	items := []proto.Message{&v1.Foo{Bar: "test1"}, &v1.Foo{Bar: "test2"}, &v1.Foo{Bar: "test3"}}
//	Expect(items).To(ProtoSliceContains(expected)) // Passes
//	missing := []proto.Message{&v1.Foo{Bar: "test1"}}
//	Expect(missing).ToNot(ProtoSliceContains(expected)) // Passes
func ProtoSliceContains(elements []proto.Message) types.GomegaMatcher {
	return &ProtoSliceContainsMatcher{
		Elements: elements,
	}
}

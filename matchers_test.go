package matchers_test

import (
	"testing"

	v1 "github.com/afritzler/protoequal/test/api/v1"
	"github.com/onsi/gomega"
	"google.golang.org/protobuf/proto"

	matchers "github.com/afritzler/protoequal"
)

func TestProtoEqualMatcher(t *testing.T) {
	g := gomega.NewWithT(t)

	// Define test cases
	testCases := []struct {
		name        string
		actual      proto.Message
		expected    proto.Message
		shouldMatch bool
	}{
		{
			name: "Should match identical messages",
			actual: &v1.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &v1.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			expected: &v1.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &v1.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			shouldMatch: true,
		},
		{
			name: "Should not match different messages",
			actual: &v1.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &v1.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			expected: &v1.Foo{
				Bar: "different-bar",
				Baz: "test-baz",
				Qux: &v1.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			shouldMatch: false,
		},
		{
			name: "Should not match messages with different nested fields",
			actual: &v1.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &v1.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			expected: &v1.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &v1.Qux{
					Driver: "different-driver",
					Handle: "foo-handle",
				},
			},
			shouldMatch: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldMatch {
				g.Expect(tc.actual).To(matchers.ProtoEqual(tc.expected))
			} else {
				g.Expect(tc.actual).ToNot(matchers.ProtoEqual(tc.expected))
			}
		})
	}
}

func TestProtoContainsMatcher(t *testing.T) {
	g := gomega.NewWithT(t)

	// Define test cases for ProtoContains
	testCases := []struct {
		name        string
		actual      interface{}
		elements    []proto.Message
		shouldMatch bool
	}{
		{
			name: "Should match when all elements are present",
			actual: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test2"},
				&v1.Foo{Bar: "test3"},
			},
			elements: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test2"},
			},
			shouldMatch: true,
		},
		{
			name: "Should match when all elements are present in different order",
			actual: []proto.Message{
				&v1.Foo{Bar: "test2"},
				&v1.Foo{Bar: "test3"},
				&v1.Foo{Bar: "test1"},
			},
			elements: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test2"},
			},
			shouldMatch: true,
		},
		{
			name: "Should not match when an element is missing",
			actual: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test3"},
			},
			elements: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test2"},
			},
			shouldMatch: false,
		},
		{
			name: "Should match with exact same elements",
			actual: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test2"},
			},
			elements: []proto.Message{
				&v1.Foo{Bar: "test1"},
				&v1.Foo{Bar: "test2"},
			},
			shouldMatch: true,
		},
		{
			name: "Should match with complex messages",
			actual: []proto.Message{
				&v1.Foo{
					Bar: "test1",
					Baz: "baz1",
					Qux: &v1.Qux{
						Driver: "driver1",
						Handle: "handle1",
					},
				},
				&v1.Foo{
					Bar: "test2",
					Baz: "baz2",
					Qux: &v1.Qux{
						Driver: "driver2",
						Handle: "handle2",
					},
				},
				&v1.Foo{Bar: "test3"},
			},
			elements: []proto.Message{
				&v1.Foo{
					Bar: "test1",
					Baz: "baz1",
					Qux: &v1.Qux{
						Driver: "driver1",
						Handle: "handle1",
					},
				},
				&v1.Foo{
					Bar: "test2",
					Baz: "baz2",
					Qux: &v1.Qux{
						Driver: "driver2",
						Handle: "handle2",
					},
				},
			},
			shouldMatch: true,
		},
		{
			name:   "Should not match non-slice input",
			actual: "not-a-slice",
			elements: []proto.Message{
				&v1.Foo{Bar: "test1"},
			},
			shouldMatch: false,
		},
		{
			name:        "Should match when checking empty elements",
			actual:      []proto.Message{&v1.Foo{Bar: "test1"}},
			elements:    []proto.Message{},
			shouldMatch: true, // Vacuously true: all (no) elements are present
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldMatch {
				g.Expect(tc.actual).To(matchers.ProtoContains(tc.elements...))
			} else {
				g.Expect(tc.actual).ToNot(matchers.ProtoContains(tc.elements...))
			}
		})
	}
}

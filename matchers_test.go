package matchers_test

import (
	"testing"

	"github.com/onsi/gomega"
	"google.golang.org/protobuf/proto"

	matchers "github.com/afritzler/protoequal"
	"github.com/afritzler/protoequal/test"
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
			actual: &test.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &test.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			expected: &test.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &test.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			shouldMatch: true,
		},
		{
			name: "Should not match different messages",
			actual: &test.Foo{
				Bar: "test-bar",
				Baz: "test-baz",
				Qux: &test.Qux{
					Driver: "foo-driver",
					Handle: "foo-handle",
				},
			},
			expected: &test.Foo{
				Bar: "different-bar",
				Baz: "test-baz",
				Qux: &test.Qux{
					Driver: "foo-driver",
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

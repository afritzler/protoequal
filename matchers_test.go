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

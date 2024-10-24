package matchers

import (
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

// ProtoEqual is a convenience function to create the matcher
func ProtoEqual(expected proto.Message) types.GomegaMatcher {
	return &ProtoEqualMatcher{
		Expected: expected,
	}
}

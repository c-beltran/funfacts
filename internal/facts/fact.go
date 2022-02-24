package facts

type (
	// Topic defines the variety of facts flavors.
	Topic struct {
		Dog           string
		Cat           string
		Entertainment string
		Trivial       string
	}

	// TopicType defines the fact topic type.
	TopicType uint
)

const (
	// TopicTypeUnknown defines the value when topic type is unknown.
	TopicTypeUnknown TopicType = iota

	// TopicTypeDog defines the value when the topic is a Dog.
	TopicTypeDog

	// TopicTypeCat defines the value when the topic is a Cat.
	TopicTypeCat

	// TopicTypeEntertainment defines the value when the topic is an Entertainment.
	TopicTypeEntertainment

	// TopicTypeTrivial defines the value when the topic is Trivial.
	TopicTypeTrivial
)

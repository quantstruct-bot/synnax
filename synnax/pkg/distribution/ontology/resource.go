package ontology

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/synnaxlabs/synnax/pkg/distribution/ontology/schema"
	"strings"
)

// ID is a unique identifier for a Resource. An example:
//
//	userKey := ID{
//	    ID:  "748d31e2-5732-4cb5-8bc9-64d4ad51efe8",
//	    Type: "user",
//	}
//
// They key has two elements so for two reasons. First, by storing the Type we know which
// Service to query for additional info on the Resource. Second, while a [ID.Key] may be
// unique for a particular resource (e.g. channel), it might not be unique across ALL
// resources. We need something universally unique across the entire delta cluster.
type ID struct {
	// Key is a string that uniquely identifies a Resource within its Type.
	Key string
	// Type defines the type of Resource the Key refers to :). For example,
	// a channel is a Resource of type "channel". Key user is a Resource of type
	// "user".
	Type Type
}

func (k ID) Validate() error {
	if k.Key == "" {
		return errors.Newf("[resource] - key is required")
	}
	if k.Type == "" {
		return errors.Newf("[resource] - type is required")
	}
	return nil
}

func (k ID) String() string {
	return fmt.Sprintf("%s:%s", k.Type, k.Key)
}

func ParseID(s string) (ID, error) {
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return ID{}, errors.Errorf("[ontology] - failed to parse id: %s", s)
	}
	return ID{Type: schema.Type(split[0]), Key: split[1]}, nil
}

func ParseIds(s []string) ([]ID, error) {
	var ids []ID
	for _, id := range s {
		parsed, err := ParseID(id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, parsed)
	}
	return ids, nil
}

// Resource represents a resource with a unique [ID] in the ontology along with its
// [Entity] information.
type Resource struct {
	// ID is the unique identifier for the resource.
	ID ID `json:"id" msgpack:"id"`
	// Entity is the entity information for the resource.
	Entity Entity `json:"entity" msgpack:"entity"`
}

// GorpKey implements the gorp.Entry interface.
func (r Resource) GorpKey() ID { return r.ID }

// SetOptions implements the gorp.Entry interface.
func (r Resource) SetOptions() []interface{} { return nil }

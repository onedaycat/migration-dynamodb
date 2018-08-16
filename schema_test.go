package migration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gopkg.in/yaml.v2"
)

var schema = `
- table: playg_user
  attr:
    - name: id
      type: S
  key:
    - name: id
      type: HASH
    - name: createdAt
      type: RANGE
  provisioned:
    read: 5
    write: 5
  globalIndex:
    - name: emailIndex
      key:
        - name: email
          type: HASH
      projection:
        type: ALL
      provisioned:
        read: 5
        write: 5
  localIndex:
    - name: emailIndex
      key:
        - name: email
          type: HASH
      projection:
        type: INCLUDE
        key: [id, email]
`

func TestSchema(t *testing.T) {
	s := []*Schema{}
	expS := []*Schema{
		{
			Table: "playg_user",
			Attr: []NameType{
				{"S", "id"},
			},
			Key: []NameType{
				{"HASH", "id"},
				{"RANGE", "createdAt"},
			},
			Provisioned: Provisioned{
				Read:  5,
				Write: 5,
			},
			GlobalIndex: []GlobalIndex{
				{
					Name: "emailIndex",
					Key: []NameType{
						{"HASH", "email"},
					},
					Projection: Projection{
						Type: "ALL",
					},
					Provisioned: Provisioned{
						Read:  5,
						Write: 5,
					},
				},
			},
			LocalIndex: []LocalIndex{
				{
					Name: "emailIndex",
					Key: []NameType{
						{"HASH", "email"},
					},
					Projection: Projection{
						Type: "INCLUDE",
						Key:  []string{"id", "email"},
					},
				},
			},
		},
	}
	err := yaml.Unmarshal([]byte(schema), &s)

	require.NoError(t, err)
	require.Equal(t, expS, s)
}

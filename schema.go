package migration

type Schema struct {
	Table       string        `yaml:"table"`
	Attr        []NameType    `yaml:"attr"`
	Key         []NameType    `yaml:"key"`
	Provisioned Provisioned   `yaml:"provisioned"`
	GlobalIndex []GlobalIndex `yaml:"globalIndex"`
	LocalIndex  []LocalIndex  `yaml:"localIndex"`
}

type NameType struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

type Provisioned struct {
	Read  int64 `yaml:"read"`
	Write int64 `yaml:"write"`
}

type GlobalIndex struct {
	Name        string      `yaml:"name"`
	Key         []NameType  `yaml:"key"`
	Projection  Projection  `yaml:"projection"`
	Provisioned Provisioned `yaml:"provisioned"`
}

type LocalIndex struct {
	Name       string     `yaml:"name"`
	Key        []NameType `yaml:"key"`
	Projection Projection `yaml:"projection"`
}

type Projection struct {
	Type string   `yaml:"type"`
	Key  []string `yaml:"key"`
}

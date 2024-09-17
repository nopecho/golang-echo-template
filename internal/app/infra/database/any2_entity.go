package database

type Jsonb map[string]interface{}

type Any2Entity struct {
	BaseModel
	Payload Jsonb `json:"payload" gorm:"type:jsonb;serializer:json"`
}

func (Any2Entity) TableName() string {
	return "domain2"
}

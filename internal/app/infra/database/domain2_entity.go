package database

type Jsonb map[string]interface{}

type Domain2Entity struct {
	BaseModel
	Payload Jsonb `json:"payload" gorm:"type:jsonb;serializer:json"`
}

func (Domain2Entity) TableName() string {
	return "domain2"
}

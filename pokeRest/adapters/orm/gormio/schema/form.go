package schema

type Form struct {
	FormSchema
	// relations
	Pokemon1 Pokemon `gorm:"foreignKey:FormNo;references:id;"` // has one
	Pokemon2 Pokemon `gorm:"foreignKey:FormName;references:name"`
}

type FormSchema struct {
	ID                uint16 `gorm:"primaryKey;autoIncrement:true"`
	Name              string `gorm:"unique"`
	IsRegionalVariant bool   `gorm:"default:false"`
	RegionName        string
}

func (Form) TableName() string {
	return "forms"
}

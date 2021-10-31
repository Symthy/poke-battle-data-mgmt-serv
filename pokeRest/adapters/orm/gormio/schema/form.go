package schema

type Form struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	Name              string `gorm:"unique"`
	IsRegionalVariant bool   `gorm:"default:false"`
	RegionName        string

	Pokemon Pokemon `gorm:"foreignKey:FormNo;references:id;foreignKey:FormName;references:name"` // has one
}

func (Form) TableName() string {
	return "forms"
}

package schema

type Form struct {
	ID                uint   `gorm:"primaryKey; autoIncrement"`
	Name              string `gorm:"unique"`
	IsRegionalVariant bool   `gorm:"default:false"`
	RegionName        string
}

func (Form) TableName() string {
	return "form"
}

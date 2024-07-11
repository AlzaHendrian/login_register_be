package authdto

type LoginResponse struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	LastName string `gorm:"type: varchar(255)" json:"last_name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	IsAdmin  bool   `gorm:"type: boolean" json:"isAdmin"`
}

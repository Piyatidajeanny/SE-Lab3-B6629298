package models
import "gorm.io/gorm"
type Student struct {
 gorm.Model
 FirstName string `json:"first_name"`
 LastName string `json:"last_name"`
 Age int `json:"age"`
 Email string `json:"email"`

 EnrollmentNumber string `json:"enrollment_number" gorm:"unique"`
 Major string `json:"major"`
}  
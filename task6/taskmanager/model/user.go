package model
import(
 "go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
	Name  string  `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role`
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	
}
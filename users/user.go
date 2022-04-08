package users

// User ...
// Custom object which can be stored in the claims
type User struct {
	Guid interface{} `json:"guid"`
}

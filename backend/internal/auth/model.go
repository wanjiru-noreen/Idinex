package auth

import ( 
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID  uuid.UUID 'json:"id"'
	Username string 'json:"username"'
	Email  string 'json:"email"'
	PassWordHash string 'json:"-"'
	CreatedAt  time.Time 'json:"created_at"'
	UpdatedAt time.Time 'json:"updated_at"'
}

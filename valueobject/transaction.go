package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a valueobject because it has no identifier and is immutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}

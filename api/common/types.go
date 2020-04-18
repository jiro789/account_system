package common

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AmountType float64
type DateType struct {
	time.Time
}

type IdType struct {
	uuid.UUID
}

func NewId() IdType {
	return IdType{uuid.New()}
}

func NewIdFromString(id string) IdType {
	return IdType{uuid.MustParse(id)}
}

func NewDate() DateType {
	return DateType{time.Now()}
}

func (d DateType) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateType) String() string {
	return fmt.Sprintf("%q", d.Format(time.RFC3339))
}

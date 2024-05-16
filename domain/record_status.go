package domain

import "database/sql/driver"

type RecordStatus string

var (
	RecordStatusActive   RecordStatus = "app.status.Active"
	RecordStatusDeleted  RecordStatus = "app.status.Deleted"
	RecordStatusInactive RecordStatus = "app.status.Inactive"
)

func (ct *RecordStatus) Scan(value interface{}) error {
	*ct = RecordStatus(value.(string))
	return nil
}

func (ct RecordStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

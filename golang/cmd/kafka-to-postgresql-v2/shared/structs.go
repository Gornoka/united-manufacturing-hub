package shared

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const (
	DbTagSeparator = "$"
)

type TopicDetails struct {
	Enterprise     string
	Site           string
	Area           string
	ProductionLine string
	WorkCell       string
	OriginId       string
	Schema         string
	Tag            string
}

type HistorianValue struct {
	NumericValue *float32
	StringValue  *string
	Name         string
	IsNumeric    bool
}

type Status int

const (
	Planned Status = iota
	InProgress
	Completed
)

type WorkOrderCreateMessageProduct struct {
	ExternalProductId string `json:"external_product_id"`
	CycleTimeMs       uint64 `json:"cycle_time_ms,omitempty"` //Note: omitempty is not checked when unmarshalling from JSON, and only used as a note for the reader
}

type WorkOrderCreateMessage struct {
	ExternalWorkOrderId string                        `json:"external_work_order_id"`
	Product             WorkOrderCreateMessageProduct `json:"product"`
	Quantity            uint64                        `json:"quantity"`
	Status              Status                        `json:"status"`
	StartTimeUnixMs     uint64                        `json:"start_time_unix_ms,omitempty"`
	EndTimeUnixMs       uint64                        `json:"end_time_unix_ms,omitempty"`
}

type WorkOrderStartMessage struct {
	ExternalWorkOrderId string `json:"external_work_order_id"`
	StartTimeUnixMs     uint64 `json:"start_time_unix_ms"`
}

type WorkOrderStopMessage struct {
	ExternalWorkOrderId string `json:"external_work_order_id"`
	EndTimeUnixMs       uint64 `json:"end_time_unix_ms"`
}

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

package storage

type PackType map[string]interface{}

// Simple storage interface for saving data,
// and uses for save FSM data
type Storage interface {
	SetData(ChatID int64, UserID int64, data *PackType)
	GetData(ChatID int64, UserID int64) *PackType
	SetState(int64, int64, string)
	GetState(ChatID int64, UserID int64) string
	Clear()
}
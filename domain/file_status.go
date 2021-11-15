package domain

type FileStatus string

const (
	FileStatusPending   FileStatus = "PENDING"
	FileStatusProcessed FileStatus = "PROCESSED"
	FileStatusError     FileStatus = "ERROR"
)

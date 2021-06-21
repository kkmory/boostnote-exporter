package entity

import "time"

type BoostNoteImportConfig struct {
	Token             string
	TargetFolderNames []string
	FloorUpdatedAt    time.Time
}

type BoostNoteFolder struct {
	ID              string    `json:"id"`
	Emoji           string    `json:"emoji"`
	Name            string    `json:"name"`
	Pathname        string    `json:"pathname"`
	Description     string    `json:"description"`
	ParentFolderID  string    `json:"parentFolderId"`
	TeamID          string    `json:"teamId"`
	Generated       bool      `json:"generated"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Version         int64     `json:"version"`
	WorkspaceID     string    `json:"workspaceId"`
	ChildDocsIDs    []string  `json:"childDocsIds"`
	ChildFoldersIDs []string  `json:"childFoldersIds"`
}

type BoostNoteDoc struct {
	ID             string           `json:"id"`
	Emoji          string           `json:"emoji"`
	Title          string           `json:"title"`
	Content        string           `json:"content"`
	HeadID         int64            `json:"headId"`
	ArchivedAt     time.Time        `json:"archivedAt"`
	FolderPathname string           `json:"folderPathname"`
	ParentFolderID string           `json:"parentFolderId"`
	DueDate        string           `json:"dueDate"`
	Status         string           `json:"status"`
	TeamID         string           `json:"teamId"`
	Generated      bool             `json:"generated"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	Version        int64            `json:"version"`
	UserID         string           `json:"userId"`
	WorkspaceID    string           `json:"workspaceId"`
	Tags           []BoostNoteTag   `json:"tags"`
	Head           BoostNoteDocHead `json:"head"`
	ShareLink      string           `json:"shareLink"`
	Assignees      []string         `json:"assignees"`
}

type BoostNoteDocHead struct {
	ID       int64              `json:"id"`
	DocID    string             `json:"docId"`
	Content  string             `json:"content"`
	Message  string             `json:"message"`
	Created  time.Time          `json:"created"`
	Creators []BoostNoteCreator `json:"creators"`
}

type BoostNoteCreator struct {
	ID          string    `json:"id"`
	UniqueName  string    `json:"uniqueName"`
	DisplayName string    `json:"displayName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type BoostNoteTag struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	TeamID    string    `json:"teamId"`
	CreatedAt time.Time `json:"createdAt"`
}

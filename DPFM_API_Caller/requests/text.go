package requests

type Text struct {
	MessageType				string  `json:"MessageType"`
	Language          		string  `json:"Language"`
	MessageTypeName			string  `json:"MessageTypeName"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}

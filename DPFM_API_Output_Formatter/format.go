package dpfm_api_output_formatter

import (
	"data-platform-api-message-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToMessageType(rows *sql.Rows) (*[]MessageType, error) {
	defer rows.Close()
	messageType := make([]MessageType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.MessageType{}

		err := rows.Scan(
			&pm.MessageType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &messageType, nil
		}

		data := pm
		messageType = append(messageType, MessageType{
			MessageType:			data.MessageType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &messageType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.MessageType,
			&pm.Language,
			&pm.MessageTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			MessageType:			data.MessageType,
			Language:          		data.Language,
			MessageTypeName:		data.MessageTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}

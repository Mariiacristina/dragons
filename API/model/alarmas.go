package model

import (
	"Reptile/API/connection"
	"Reptile/API/schema"
	"strconv"
 )


// GetAlarmas is a fucking function
func GetAlarmas(id string) ([]schema.Alarma, error) {
	newID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	db := connection.Connect()
	defer connection.Disconnect(db)

	rows, err := db.Query("SELECT hora, razon FROM alarmas WHERE alarmas.id_cliente = ?", newID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []schema.Alarma
	for rows.Next() {
		var hora string
		var razon string
		err = rows.Scan(&hora, &razon)
		res = append(res, schema.Alarma{Id_cliente: newID, Hora: hora, Razon: razon})
	}

	return res, nil
 }

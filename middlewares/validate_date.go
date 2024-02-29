package middlewares

import (
	"time" )

func ValidateAndConvertDate(date string) (string, error) {
    // Intenta parsear la cadena de fecha en formato "DD-MM-YYYY"
    t, err := time.Parse("02-01-2006", date)
    if err != nil {
        // Si no se puede parsear como "DD-MM-YYYY", intenta parsear como "YYYY-MM-DD"
        t, err = time.Parse("2006-01-02", date)
        if err != nil {
            // Si no se puede parsear en ninguno de los formatos, devuelve un error
            return "", err
        }
    }

    // Formatea el objeto time.Time en el formato "YYYY-MM-DD"
    formattedDate := t.Format("2006-01-02")

    return formattedDate, nil
}

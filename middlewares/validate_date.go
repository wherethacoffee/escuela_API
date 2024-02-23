package middlewares

import ( "time" )

func ValidateAndConvertDate(date string) (string, error) {
    t, err := time.Parse("02-01-2006", date) 
    if err != nil {
	return "", err
    }

    date = t.Format("02/01/2006")

    return date, nil
}

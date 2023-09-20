package email

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"

	"github.com/ashishjaiswal123/bookings/internal/models"
)

type data struct {
	GuestName          string
	HotelName          string
	CheckInDate        string
	CheckOutDate       string
	RoomType           string
	ConfirmationNumber string
}

// SendRoomConfirmationEmail returns html message of room confirmation email
func SendRoomConfirmationEmail(res models.Reservation) (string, error) {
	templateFile := "internal/email/room_confirmation.html"

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Printf("Error reading template file: %v\n", err)
		return "", err
	}

	guestData := data{
		GuestName:          res.FirstName + " " + res.LastName,
		HotelName:          "ABC Hotel",
		CheckInDate:        res.StartDate.Format("2000-01-02"),
		CheckOutDate:       res.EndDate.Format("2000-01-02"),
		RoomType:           res.Room.RoomName,
		ConfirmationNumber: strconv.Itoa(res.ID),
	}

	// Execute the template with the data
	var htmlMessageBuffer bytes.Buffer
	err = tmpl.Execute(&htmlMessageBuffer, guestData)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return "", err
	}

	htmlMessage := htmlMessageBuffer.String()
	return htmlMessage, nil
}

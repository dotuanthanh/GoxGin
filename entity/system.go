package entity

type WorkoutBooking struct {
	Id          int    `json:"id"`
	UserId      int    `json:"UserId"`
	BookingTime string `json:"bookingTime"`
	Status      bool   `json:"status"`
}

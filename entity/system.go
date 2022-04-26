package entity

type WorkoutBooking struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	BookingTime string `json:"bookingTime"`
	Status      bool   `json:"status"`
}

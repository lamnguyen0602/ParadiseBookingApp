package constant

const (
	BookingStatusPending = iota + 1
	BookingStatusConfirmed
	BookingStatusCheckIn
	BookingStatusCheckOut
	BookingStatusCompleted
	BookingStatusCancel
)

type TypeBooking int

const (
	TypeBookingForMySelf TypeBooking = iota + 1
	TypeBookingForFriend
)

const (
	BookingGuiderStatusPending = iota + 1
	BookingGuiderStatusConfirmed
	BookingGuiderStatusCompleted
	BookingGuiderStatusCancel
)

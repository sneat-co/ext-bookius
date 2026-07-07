package dto4bookius

const ExtensionID = "bookius"

type TargetKind string

const (
	TargetKindPerson       TargetKind = "person"
	TargetKindAppointment  TargetKind = "appointment"
	TargetKindConsultation TargetKind = "consultation"
	TargetKindMeetingRoom  TargetKind = "meeting-room"
	TargetKindFacility     TargetKind = "facility"
	TargetKindAsset        TargetKind = "asset"
	TargetKindEquipment    TargetKind = "equipment"
	TargetKindService      TargetKind = "service"
	TargetKindEventSession TargetKind = "event-session"
	TargetKindCustom       TargetKind = "custom"
)

type BookingState string

const (
	BookingStateDraft       BookingState = "draft"
	BookingStateHeld        BookingState = "held"
	BookingStateRequested   BookingState = "requested"
	BookingStateConfirmed   BookingState = "confirmed"
	BookingStateRescheduled BookingState = "rescheduled"
	BookingStateCancelled   BookingState = "cancelled"
	BookingStateExpired     BookingState = "expired"
	BookingStateCompleted   BookingState = "completed"
	BookingStateNoShow      BookingState = "no-show"
)

type ExtensionRef struct {
	Ext        string `json:"ext"`
	Collection string `json:"collection,omitempty"`
	ID         string `json:"id"`
}

type Slot struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Timezone string `json:"timezone,omitempty"`
}

type BookingType struct {
	ID                    string        `json:"id,omitempty"`
	Title                 string        `json:"title"`
	Slug                  string        `json:"slug"`
	Description           string        `json:"description,omitempty"`
	DurationMinutes       int           `json:"durationMinutes"`
	TargetKind            TargetKind    `json:"targetKind"`
	TargetRef             *ExtensionRef `json:"targetRef,omitempty"`
	AvailabilitySourceRef *ExtensionRef `json:"availabilitySourceRef,omitempty"`
	ConfirmationMode      string        `json:"confirmationMode"`
}

type CreateBookingRequest struct {
	SpaceID       string `json:"spaceID,omitempty"`
	BookingTypeID string `json:"bookingTypeID"`
	BookingPageID string `json:"bookingPageID,omitempty"`
	RequestedSlot Slot   `json:"requestedSlot"`
	VisitorName   string `json:"visitorName"`
	VisitorEmail  string `json:"visitorEmail"`
	VisitorPhone  string `json:"visitorPhone,omitempty"`
	Subject       string `json:"subject,omitempty"`
	Message       string `json:"message,omitempty"`
}

package dto4bookius

import (
	"context"
	"errors"
	"strings"
	"time"
)

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
	// TargetKindCompetitionEntry is a paid entry into an externally-owned
	// competition. The provider-specific target identity remains opaque.
	TargetKindCompetitionEntry TargetKind = "competition-entry"
	TargetKindCustom           TargetKind = "custom"
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

var ErrInvalidCompetitionEntry = errors.New("bookius: invalid competition entry value")

type CompetitionEntryReservationID string
type CheckoutID string
type SettlementID string

// CompetitionEntryTarget is an opaque external target. Bookius does not
// import Competios or assume a sport/game schema, but a Competition Entry must
// identify both an umbrella Event and a Tournament beneath it.
type CompetitionEntryTarget struct {
	ExtensionID   string `json:"extensionID"`
	EventID       string `json:"eventID"`
	TournamentID  string `json:"tournamentID"`
	CompetitionID string `json:"competitionID"`
}

type ReservationState string

const (
	ReservationRequested  ReservationState = "requested"
	ReservationWaitlisted ReservationState = "waitlisted"
	ReservationHeld       ReservationState = "held"
	ReservationCheckout   ReservationState = "checkout-ready"
	ReservationConfirmed  ReservationState = "confirmed"
	ReservationFailed     ReservationState = "failed"
	ReservationExpired    ReservationState = "expired"
	ReservationCancelled  ReservationState = "cancelled"
	ReservationRefunded   ReservationState = "refunded"
)

// CompetitionEntryPaymentState is independent of booking state: for example,
// an active hold can be not_started while a confirmed booking is paid.
type CompetitionEntryPaymentState string

const (
	CompetitionEntryPaymentFree          CompetitionEntryPaymentState = "free"
	CompetitionEntryPaymentNotStarted    CompetitionEntryPaymentState = "not_started"
	CompetitionEntryPaymentCheckoutOpen  CompetitionEntryPaymentState = "checkout_open"
	CompetitionEntryPaymentPaid          CompetitionEntryPaymentState = "paid"
	CompetitionEntryPaymentRefundPending CompetitionEntryPaymentState = "refund_pending"
	CompetitionEntryPaymentRefunded      CompetitionEntryPaymentState = "refunded"
	CompetitionEntryPaymentFailed        CompetitionEntryPaymentState = "failed"
)

// CompetitionEntryReservationRequest intentionally contains no amount or
// currency. The target provider resolves the immutable offer server-to-server;
// browser input can only name the already-authorised participant/entry.
type CompetitionEntryReservationRequest struct {
	RequestID            string                 `json:"requestID"`
	Target               CompetitionEntryTarget `json:"target"`
	ParticipantReference string                 `json:"participantReference"`
	EntryReference       string                 `json:"entryReference"`
}

// CompetitionEntryReservation is Bookius' safe projection. Amount and
// currency are copied only from the server-authoritative offer/hold; callers
// must never construct this from a browser request.
type CompetitionEntryReservation struct {
	ID                   CompetitionEntryReservationID `json:"id"`
	RequestID            string                        `json:"requestID"`
	Target               CompetitionEntryTarget        `json:"target"`
	ParticipantReference string                        `json:"participantReference"`
	EntryReference       string                        `json:"entryReference"`
	State                ReservationState              `json:"state"`
	AmountMinor          int64                         `json:"amountMinor"`
	Currency             string                        `json:"currency"`
	OfferReference       string                        `json:"offerReference"`
	OfferVersion         uint32                        `json:"offerVersion"`
	OfferChecksum        string                        `json:"offerChecksum"`
	PaymentState         CompetitionEntryPaymentState  `json:"paymentState"`
	ExpiresAt            *time.Time                    `json:"expiresAt,omitempty"`
}

type CheckoutProjection struct {
	CheckoutID    CheckoutID                    `json:"checkoutID"`
	ReservationID CompetitionEntryReservationID `json:"reservationID"`
	State         ReservationState              `json:"state"`
	CheckoutURL   string                        `json:"checkoutURL,omitempty"`
	ExpiresAt     time.Time                     `json:"expiresAt"`
}

type SettlementStatus string

const (
	SettlementPaid     SettlementStatus = "paid"
	SettlementFailed   SettlementStatus = "failed"
	SettlementRefunded SettlementStatus = "refunded"
)

// SettlementNotification is a server-to-server provider result. It repeats
// the held amount/currency so Bookius can fail closed on a mismatched payment;
// it is never browser-supplied input.
type SettlementNotification struct {
	SettlementID  SettlementID                  `json:"settlementID"`
	ReservationID CompetitionEntryReservationID `json:"reservationID"`
	Status        SettlementStatus              `json:"status"`
	AmountMinor   int64                         `json:"amountMinor"`
	Currency      string                        `json:"currency"`
	OccurredAt    time.Time                     `json:"occurredAt"`
}

func ValidateCompetitionEntryReservationRequest(value CompetitionEntryReservationRequest) error {
	if strings.TrimSpace(value.RequestID) == "" || !validCompetitionEntryTarget(value.Target) || strings.TrimSpace(value.ParticipantReference) == "" || strings.TrimSpace(value.EntryReference) == "" {
		return ErrInvalidCompetitionEntry
	}
	return nil
}

func ValidateCompetitionEntryReservation(value CompetitionEntryReservation) error {
	if value.ID == "" || strings.TrimSpace(value.RequestID) == "" || !validCompetitionEntryTarget(value.Target) || strings.TrimSpace(value.ParticipantReference) == "" || strings.TrimSpace(value.EntryReference) == "" || value.AmountMinor < 0 || !validISOCurrency(value.Currency) || strings.TrimSpace(value.OfferReference) == "" || value.OfferVersion == 0 || !validSHA256Checksum(value.OfferChecksum) {
		return ErrInvalidCompetitionEntry
	}
	switch value.State {
	case ReservationRequested, ReservationWaitlisted, ReservationHeld, ReservationCheckout, ReservationConfirmed, ReservationFailed, ReservationExpired, ReservationCancelled, ReservationRefunded:
	default:
		return ErrInvalidCompetitionEntry
	}
	if (value.State == ReservationHeld || value.State == ReservationCheckout) && (value.ExpiresAt == nil || value.ExpiresAt.IsZero()) {
		return ErrInvalidCompetitionEntry
	}
	switch value.PaymentState {
	case CompetitionEntryPaymentFree, CompetitionEntryPaymentNotStarted, CompetitionEntryPaymentCheckoutOpen, CompetitionEntryPaymentPaid, CompetitionEntryPaymentRefundPending, CompetitionEntryPaymentRefunded, CompetitionEntryPaymentFailed:
	default:
		return ErrInvalidCompetitionEntry
	}
	return nil
}

func ValidateSettlementNotification(value SettlementNotification) error {
	if value.SettlementID == "" || value.ReservationID == "" || value.AmountMinor < 0 || !validISOCurrency(value.Currency) || value.OccurredAt.IsZero() {
		return ErrInvalidCompetitionEntry
	}
	switch value.Status {
	case SettlementPaid, SettlementFailed, SettlementRefunded:
		return nil
	default:
		return ErrInvalidCompetitionEntry
	}
}

func validCompetitionEntryTarget(value CompetitionEntryTarget) bool {
	return value.ExtensionID != "" && value.EventID != "" && value.TournamentID != "" && value.CompetitionID != ""
}

func validISOCurrency(value string) bool {
	if len(value) != 3 {
		return false
	}
	for _, character := range value {
		if character < 'A' || character > 'Z' {
			return false
		}
	}
	return true
}

func validSHA256Checksum(value string) bool {
	if len(value) != len("sha256:")+64 || !strings.HasPrefix(value, "sha256:") {
		return false
	}
	for _, character := range value[len("sha256:"):] {
		if (character < '0' || character > '9') && (character < 'a' || character > 'f') {
			return false
		}
	}
	return true
}

// CompetitionEntryReservations is the narrow Bookius server port for paid
// competition participation. A provider resolves price and capacity when a
// hold is created; the browser cannot send amount/currency or settlement facts.
// Implementations must deduplicate request IDs and settlement IDs.
type CompetitionEntryReservations interface {
	ReserveCompetitionEntry(context.Context, CompetitionEntryReservationRequest) (CompetitionEntryReservation, error)
	BeginCompetitionEntryCheckout(context.Context, CompetitionEntryReservationID) (CheckoutProjection, error)
	RecordCompetitionEntrySettlement(context.Context, SettlementNotification) (CompetitionEntryReservation, error)
	CancelCompetitionEntryReservation(context.Context, CompetitionEntryReservationID, string) (CompetitionEntryReservation, error)
}

package dto4bookius

import (
	"errors"
	"testing"
	"time"
)

func validCompetitionReservation() CompetitionEntryReservation {
	expiresAt := time.Now().UTC().Add(time.Minute)
	return CompetitionEntryReservation{
		ID: "reservation-1", RequestID: "request-1", Target: CompetitionEntryTarget{ExtensionID: "competios", EventID: "event-1", TournamentID: "tournament-1", CompetitionID: "competition-1"},
		ParticipantReference: "participant-1", EntryReference: "entry-1", State: ReservationHeld,
		AmountMinor: 2500, Currency: "EUR", OfferReference: "offer-1", OfferVersion: 1, OfferChecksum: "sha256:1111111111111111111111111111111111111111111111111111111111111111", PaymentState: CompetitionEntryPaymentNotStarted, ExpiresAt: &expiresAt,
	}
}

func TestCompetitionEntryReservationNeverAcceptsBrowserPriceInput(t *testing.T) {
	request := CompetitionEntryReservationRequest{RequestID: "request-1", Target: CompetitionEntryTarget{ExtensionID: "competios", EventID: "event-1", TournamentID: "tournament-1", CompetitionID: "competition-1"}, ParticipantReference: "participant-1", EntryReference: "entry-1"}
	if err := ValidateCompetitionEntryReservationRequest(request); err != nil {
		t.Fatalf("valid price-free request: %v", err)
	}
	if err := ValidateCompetitionEntryReservation(validCompetitionReservation()); err != nil {
		t.Fatalf("valid server reservation: %v", err)
	}
	for name, reservation := range map[string]CompetitionEntryReservation{
		"negative amount": func() CompetitionEntryReservation {
			value := validCompetitionReservation()
			value.AmountMinor = -1
			return value
		}(),
		"lowercase currency": func() CompetitionEntryReservation {
			value := validCompetitionReservation()
			value.Currency = "eur"
			return value
		}(),
		"hold without expiry": func() CompetitionEntryReservation {
			value := validCompetitionReservation()
			value.ExpiresAt = nil
			return value
		}(),
	} {
		if err := ValidateCompetitionEntryReservation(reservation); !errors.Is(err, ErrInvalidCompetitionEntry) {
			t.Errorf("%s error = %v, want ErrInvalidCompetitionEntry", name, err)
		}
	}
}

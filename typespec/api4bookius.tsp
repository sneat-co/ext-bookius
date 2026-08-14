namespace Sneat.Bookius;

alias BookingState =
  "draft" | "held" | "requested" | "confirmed" | "rescheduled" |
  "cancelled" | "expired" | "completed" | "no-show";

alias TargetKind =
  "person" | "appointment" | "consultation" | "meeting-room" | "facility" |
  "asset" | "equipment" | "service" | "event-session" | "competition-entry" | "custom";

model ExtensionRef {
  ext: string;
  collection?: string;
  id: string;
}

model Slot {
  start: string;
  end: string;
  timezone?: string;
}

model BookingType {
  id?: string;
  title: string;
  slug: string;
  description?: string;
  durationMinutes: int32;
  targetKind: TargetKind;
  targetRef?: ExtensionRef;
  availabilitySourceRef?: ExtensionRef;
  confirmationMode: "automatic" | "manual" | "request";
}

model CreateBookingRequest {
  spaceID?: string;
  bookingTypeID: string;
  bookingPageID?: string;
  requestedSlot: Slot;
  visitorName: string;
  visitorEmail: string;
  visitorPhone?: string;
  subject?: string;
  message?: string;
}

model Booking {
  id?: string;
  bookingTypeID: string;
  bookingPageID?: string;
  state: BookingState;
  requestedSlot?: Slot;
  confirmedSlot?: Slot;
}

model CompetitionEntryTarget {
  extensionID: string;
  eventID: string;
  tournamentID: string;
  competitionID: string;
}

model CompetitionEntryReservationRequest {
  requestID: string;
  target: CompetitionEntryTarget;
  participantReference: string;
  entryReference: string;
}

model CompetitionEntryReservation {
  id: string;
  requestID: string;
  target: CompetitionEntryTarget;
  participantReference: string;
  entryReference: string;
  state: "requested" | "waitlisted" | "held" | "checkout-ready" | "confirmed" | "failed" | "expired" | "cancelled" | "refunded";
  // Server-authored hold terms, never fields on the create request.
  amountMinor: int64;
  currency: string;
  offerReference: string;
  offerVersion: uint32;
  offerChecksum: string;
  paymentState: "free" | "not_started" | "checkout_open" | "paid" | "refund_pending" | "refunded" | "failed";
  expiresAt?: utcDateTime;
}

model SettlementNotification {
  settlementID: string;
  reservationID: string;
  status: "paid" | "failed" | "refunded";
  amountMinor: int64;
  currency: string;
  occurredAt: utcDateTime;
}

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
  targetVersion: uint32;
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
  bookingRevision: uint32;
  state: "requested" | "waitlisted" | "held" | "checkout-ready" | "confirmed" | "failed" | "expired" | "cancelled" | "refunded";
  // Server-authored hold terms, never fields on the create request.
  amountMinor: int64;
  currency: string;
  offerReference: string;
  offerVersion: uint32;
  offerChecksum: string;
  paymentState: "free" | "not_started" | "checkout_open" | "paid" | "refund_pending" | "refunded" | "failed";
  confirmationEvidence?: CompetitionEntryConfirmationEvidence;
  checkoutOperation: "none" | "pending" | "ready" | "failed";
  confirmationDelivery: "none" | "pending" | "delivered" | "failed";
  refundDelivery: "none" | "pending" | "delivered" | "failed";
  expiresAt?: utcDateTime;
}

model SettlementNotification {
  settlementID: string;
  settlementReference: string;
  refundReference: string;
  reservationID: string;
  status: "paid" | "failed" | "refunded";
  amountMinor: int64;
  currency: string;
  occurredAt: utcDateTime;
}

model CompetitionEntryConfirmationEvidence {
  kind: "free" | "settled";
  settlementReference?: string;
}

model ApproveCompetitionEntryReservationRequest {
  commandID: string;
  reservationID: string;
}

model PromoteCompetitionEntryWaitlistRequest {
  commandID: string;
  reservationID: string;
}

model ExpireCompetitionEntryReservationRequest {
  commandID: string;
  reservationID: string;
}

model CancelCompetitionEntryReservationRequest {
  commandID: string;
  reservationID: string;
  origin: "participant" | "organiser";
  actorReference: string;
  authorityEvidence: string;
  reason: string;
}

model CompetitionEntryCancellationValidation {
  authorized: boolean;
  /** False only for a locked participant cancellation without an organiser refund override. */
  refundAuthorized: boolean;
  currentTournamentVersion: uint32;
  registrationLocked: boolean;
  authoriserReference?: string;
  authorityEvidence: string;
  validatedAt: utcDateTime;
}

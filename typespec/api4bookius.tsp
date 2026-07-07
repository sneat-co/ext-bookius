namespace Sneat.Bookius;

alias BookingState =
  "draft" | "held" | "requested" | "confirmed" | "rescheduled" |
  "cancelled" | "expired" | "completed" | "no-show";

alias TargetKind =
  "person" | "appointment" | "consultation" | "meeting-room" | "facility" |
  "asset" | "equipment" | "service" | "event-session" | "custom";

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

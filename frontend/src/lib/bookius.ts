export const bookiusExtensionID = "bookius" as const;

export type BookiusTargetKind =
  | "person"
  | "appointment"
  | "consultation"
  | "meeting-room"
  | "facility"
  | "asset"
  | "equipment"
  | "service"
  | "event-session"
  | "custom";

export type BookiusConfirmationMode = "automatic" | "manual" | "request";

export type BookiusBookingState =
  | "draft"
  | "held"
  | "requested"
  | "confirmed"
  | "rescheduled"
  | "cancelled"
  | "expired"
  | "completed"
  | "no-show";

export interface BookiusExtensionRef {
  readonly ext: string;
  readonly collection?: string;
  readonly id: string;
}

export interface BookiusAvailabilitySourceRef extends BookiusExtensionRef {
  readonly ext: "calendarius";
}

export interface BookiusContactRef extends BookiusExtensionRef {
  readonly ext: "contactus";
}

export interface BookiusLocation {
  readonly mode: "online" | "phone" | "physical" | "ask-host" | "ask-visitor";
  readonly label?: string;
  readonly url?: string;
  readonly address?: string;
}

export interface BookiusSlot {
  readonly start: string;
  readonly end: string;
  readonly timezone?: string;
}

export interface BookiusBookingType {
  readonly id?: string;
  readonly title: string;
  readonly slug: string;
  readonly description?: string;
  readonly durationMinutes: number;
  readonly targetKind: BookiusTargetKind;
  readonly targetRef?: BookiusExtensionRef;
  readonly availabilitySourceRef?: BookiusAvailabilitySourceRef;
  readonly location?: BookiusLocation;
  readonly confirmationMode: BookiusConfirmationMode;
  readonly visibility?: "public" | "hidden" | "disabled";
}

export interface BookiusBookingPage {
  readonly id?: string;
  readonly title: string;
  readonly slug: string;
  readonly handle?: string;
  readonly businessSlug?: string;
  readonly category?: string;
  readonly resourceSlug?: string;
  readonly bookingTypeIDs: readonly string[];
  readonly visibility?: "public" | "hidden" | "disabled";
}

export interface BookiusBooking {
  readonly id?: string;
  readonly bookingTypeID: string;
  readonly bookingPageID?: string;
  readonly state: BookiusBookingState;
  readonly requestedSlot?: BookiusSlot;
  readonly confirmedSlot?: BookiusSlot;
  readonly visitorName?: string;
  readonly visitorEmail?: string;
  readonly visitorPhone?: string;
  readonly subject?: string;
  readonly message?: string;
  readonly hostRef?: BookiusContactRef;
  readonly inviteeContactRef?: BookiusContactRef;
  readonly resourceRef?: BookiusExtensionRef;
  readonly calendarCommitmentRef?: BookiusAvailabilitySourceRef;
}

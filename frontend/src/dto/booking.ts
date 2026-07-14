import { IWithCreated } from '@sneat/dto';

export type BookiusTargetKind =
  | 'person'
  | 'appointment'
  | 'consultation'
  | 'meeting-room'
  | 'facility'
  | 'asset'
  | 'equipment'
  | 'service'
  | 'event-session'
  | 'custom';

export type BookiusConfirmationMode = 'automatic' | 'manual' | 'request';

export type BookiusBookingState =
  | 'draft'
  | 'held'
  | 'requested'
  | 'confirmed'
  | 'rescheduled'
  | 'cancelled'
  | 'expired'
  | 'completed'
  | 'no-show';

export interface BookiusExtensionRef {
  readonly ext: string;
  readonly collection?: string;
  readonly id: string;
}

export interface BookiusPublicPath {
  readonly handle?: string;
  readonly businessSlug?: string;
  readonly category?: string;
  readonly resourceSlug?: string;
}

export interface BookiusAvailabilitySourceRef extends BookiusExtensionRef {
  readonly ext: 'calendarius';
}

export interface BookiusContactRef extends BookiusExtensionRef {
  readonly ext: 'contactus';
}

export interface BookiusLocation {
  readonly mode: 'online' | 'phone' | 'physical' | 'ask-host' | 'ask-visitor';
  readonly label?: string;
  readonly url?: string;
  readonly address?: string;
}

export interface BookiusSlot {
  readonly start: string;
  readonly end: string;
  readonly timezone?: string;
}

export interface BookiusBookingTypeBrief {
  readonly id: string;
  readonly title: string;
  readonly slug: string;
  readonly durationMinutes: number;
  readonly targetKind: BookiusTargetKind;
  readonly confirmationMode: BookiusConfirmationMode;
}

export interface BookiusBookingTypeDbo
  extends Omit<BookiusBookingTypeBrief, 'id'>,
    IWithCreated {
  readonly description?: string;
  readonly visibility?: 'public' | 'hidden' | 'disabled';
  readonly availabilitySourceRef?: BookiusAvailabilitySourceRef;
  readonly targetRef?: BookiusExtensionRef;
  readonly location?: BookiusLocation;
  readonly intakeFields?: readonly string[];
}

export interface BookiusBookingPageDbo extends IWithCreated {
  readonly title: string;
  readonly slug: string;
  readonly path: BookiusPublicPath;
  readonly bookingTypeIDs: readonly string[];
  readonly visibility?: 'public' | 'hidden' | 'disabled';
  readonly defaultContext?: Record<string, string>;
}

export interface BookiusBookingTransition {
  readonly at: string;
  readonly by?: string;
  readonly from?: BookiusBookingState;
  readonly to: BookiusBookingState;
  readonly reason?: string;
}

export interface BookiusBookingDbo extends IWithCreated {
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
  readonly transitionLog?: readonly BookiusBookingTransition[];
  readonly confirmedAt?: string;
  readonly cancelledAt?: string;
}

export interface BookiusBookingBrief {
  readonly id: string;
  readonly bookingTypeID: string;
  readonly state: BookiusBookingState;
  readonly requestedSlot?: BookiusSlot;
  readonly confirmedSlot?: BookiusSlot;
  readonly visitorName?: string;
  readonly subject?: string;
}

export interface BookiusCreateBookingTypeRequest {
  readonly spaceID: string;
  readonly bookingType: BookiusBookingTypeDbo;
}

export interface BookiusCreateBookingRequest {
  readonly spaceID?: string;
  readonly bookingTypeID: string;
  readonly bookingPageID?: string;
  readonly requestedSlot: BookiusSlot;
  readonly visitorName: string;
  readonly visitorEmail: string;
  readonly visitorPhone?: string;
  readonly subject?: string;
  readonly message?: string;
}

/**
 * Public booking-type shape retained from the initial standalone contract.
 * The richer DBO above is used by the implementation; this remains useful to
 * lightweight consumers that do not need persistence metadata.
 */
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
  readonly visibility?: 'public' | 'hidden' | 'disabled';
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
  readonly visibility?: 'public' | 'hidden' | 'disabled';
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

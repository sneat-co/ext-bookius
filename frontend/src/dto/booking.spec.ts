import {
  BookiusBookingTypeBrief,
  BookiusCompetitionEntryReservationRequest,
  BookiusCreateBookingRequest,
} from './booking';
import { describe, expect, it } from 'vitest';

describe('Bookius booking DTOs', () => {
  it('supports generic booking type targets', () => {
    const bookingType: BookiusBookingTypeBrief = {
      id: 'office-meeting',
      title: 'Office Meeting',
      slug: 'office-meeting',
      durationMinutes: 60,
      targetKind: 'meeting-room',
      confirmationMode: 'request',
    };

    expect(bookingType.targetKind).toBe('meeting-room');
  });

  it('captures anonymous public booking requests', () => {
    const request: BookiusCreateBookingRequest = {
      bookingTypeID: 'investor-call',
      requestedSlot: {
        start: '2026-07-07T10:00:00Z',
        end: '2026-07-07T10:30:00Z',
        timezone: 'Europe/Dublin',
      },
      visitorName: 'Alex',
      visitorEmail: 'alex@example.com',
      subject: 'Sneat.co investment opportunity',
    };

    expect(request.visitorEmail).toBe('alex@example.com');
  });

  it('keeps price and currency out of a competition-entry browser request', () => {
    const request: BookiusCompetitionEntryReservationRequest = {
      requestID: 'request-1',
      target: {
        extensionID: 'competios',
        eventID: 'event-1',
        tournamentID: 'tournament-1',
        competitionID: 'competition-1',
      },
      participantReference: 'participant-1',
      entryReference: 'entry-1',
    };

    expect(request.target.tournamentID).toBe('tournament-1');
    expect('amountMinor' in request).toBe(false);
    expect('currency' in request).toBe(false);
  });
});

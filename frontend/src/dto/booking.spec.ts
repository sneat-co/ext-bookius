import {
  BookiusBookingTypeBrief,
  BookiusCreateBookingRequest,
} from './booking';

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
});

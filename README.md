# bookius-contract

Public contract surface for **Bookius**, the Sneat ecosystem booking platform.

Bookius owns booking types, public booking pages, bookings/reservations,
booking lifecycle, confirmations and the public booking experience. It stays
thin by composing existing Sneat owners:

- Calendarius owns calendars, availability, schedules and calendar commitments.
- Contactius owns contacts, invitees and the relationship graph.
- Resource-owning extensions own their resources; Bookius stores references.

Implementation repo: `sneat-co/bookius` (private by default).

## Packages

- `src/` — TypeScript DTOs and constants for frontend contracts.
- `typespec/` — API shape sketch for public and space-scoped endpoints.
- `backend/` — Go DTOs/constants for backend consumers.


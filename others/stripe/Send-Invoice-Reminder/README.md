# Send Invoice Reminder

Schedule invoice emails relative to invoice time (e.g., upcoming, new,
reminder, due). Also process payments and avoid sending reminders for invoices
that are already fully paid.

Each scheduled message is offset by a number of days relative to the invoice
time.

## Input

- `send_schedule`: map of offset -> label.
- `customer_invoices`: list of invoices with `invoice_time`, `name`, `amount`.
- `customer_payments`: list of payments with `payment_time`, `name`, `amount`.

## Output

- Printed reminder messages in chronological order.
- Final list of customers with remaining balance.

## Example

If a customer pays before the due date, only the earlier reminders may be sent.

## Constraints

- Offsets can be negative or positive.
- Timestamps are integers.

from collections import defaultdict
import heapq


class Invoicer:
    def __init__(self, send_schedule):
        self.send_schedule = send_schedule

    def send_emails(self, customer_invoices, customer_payments):
        res = []  # (t, type, person, amount)
        total_owed = {}
        ###------------
        pay_type = "pay"
        ###------------
        for customer in customer_invoices:
            invoice_time, name, amount = (
                customer["invoice_time"],
                customer["name"],
                customer["amount"],
            )
            total_owed[name] = amount
            for k, v in self.send_schedule.items():
                heapq.heappush(res, (k + invoice_time, v, name, amount))
        ###------------
        for payments in customer_payments:
            pay_time, name, amount = (
                payments["payment_time"],
                payments["name"],
                payments["amount"],
            )
            heapq.heappush(res, (pay_time, pay_type, name, amount))
        ###------------
        while res:
            t, type, person, amount = heapq.heappop(res)
            ###------------
            if type == pay_type:
                total_owed[person] -= amount
            elif total_owed[person] > 0:
                ###------------
                print(f"{t}: [{type}] Invoice for {person} for {total_owed[person]}")

        for k, v in total_owed.items():
            if v > 0:
                print(f"{k} owes {v} dollars")
        return res

    # def send_emails(self, customer_invoices, customer_payments):
    # 	return delinquent_customers


send_schedule = {-10: "Upcoming", 0: "New", 20: "Reminder", 30: "Due"}

invoicer = Invoicer(send_schedule)

customer_invoices = [
    {"invoice_time": 0, "name": "Alice", "amount": 200},
    {"invoice_time": 1, "name": "Bob", "amount": 100},
]

# invoicer.send_emails(customer_invoices)


customer_payments = [
    {"payment_time": -9, "name": "Alice", "amount": 100},
    {"payment_time": 1, "name": "Alice", "amount": 50},
    {"payment_time": 0, "name": "Bob", "amount": 100},
]

invoicer.send_emails(customer_invoices, customer_payments)

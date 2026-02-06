# Store accepted transactions to check them later
accepted_transactions = []

def process_fraud_rules(rules_input):
    fraud_rules = []
    for rule in rules_input:
        timestamp, field, value = rule.split(',')
        fraud_rules.append((int(timestamp), field, value))
    return sorted(fraud_rules, key=lambda rule: rule[0])

# Step 2: Process Authorization Requests and keep track of accepted ones
def process_authorization_request(requests, fraud_rules):
    # Sort requests by the timestamp (first element)
    sorted_requests = sorted(requests, key=lambda req: int(req.split(',')[0]))
    
    for request in sorted_requests:
        timestamp, unique_id, amount, card_number, merchant = request.split(',')
        timestamp = int(timestamp)
        # Check if the request matches any active fraud rules
        for rule_time, rule_field, rule_value in fraud_rules:
            if rule_time <= timestamp:  # Rule is active for this request
                if rule_field == 'merchant' and rule_value == merchant:
                    print(f"{timestamp} {unique_id} {amount} REJECT")
                    break
                if rule_field == 'card_number' and rule_value == card_number:
                    print(f"{timestamp} {unique_id} {amount} REJECT")
                    break
        else:
            # If no fraud rule matches, approve the request and track it
            print(f"{timestamp} {unique_id} {amount} APPROVE")
            accepted_transactions.append((timestamp, unique_id, amount, card_number, merchant))

# Example input
fraud_rule_input = [
    "1,merchant,bobs_burgers",
    "20,card_number,4242111111111111"
]

authorization_requests = [
    "0,R1,9.99,4242424242424242,bobs_burgers",
    "5,R2,5.60,4242424242424242,bobs_burgers",
    # "10,R3,500.00,4242111111111111,a_corp"
]

fraud_rules = process_fraud_rules(fraud_rule_input)
# Process and sort Authorization Requests
process_authorization_request(authorization_requests, fraud_rules)

# Step 3: Initialize the sum of money lost to fraud
def get_sum_lost_to_fraud(accepted_transactions, fraud_rules):
    total = 0
    for timestamp, field, value in fraud_rules:
        # Check if any accepted transactions become fraudulent after this rule is introduced
        for trans_time, trans_id, trans_amount, trans_card, trans_merchant in accepted_transactions:
            # Only check transactions before the fraud rule's timestamp
            if trans_time < int(timestamp):
                # If the fraud rule is about merchants and the transaction matches the merchant
                if field == 'merchant' and value == trans_merchant:
                    # Mark the transaction as lost to fraud
                    print(f"Transaction {trans_id} with amount {trans_amount} is lost to fraud.")
                    total += float(trans_amount)
                elif field == 'card_number' and value == trans_card:
                    print(f"Transaction {trans_id} with amount {trans_amount} is lost to fraud.")
                    total += float(trans_amount)
    return total


# Output the total sum of money lost to fraud
print(f"Total money lost to fraud: {get_sum_lost_to_fraud(accepted_transactions, fraud_rules)}")


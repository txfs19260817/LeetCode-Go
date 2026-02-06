# https://www.1point3acres.com/bbs/thread-1089011-1-1.html
from io import StringIO
import csv
from collections import defaultdict
import difflib

csv_data = """customer_id,merchant_id,payout_date,card_type,amount
cust1,merchantA,2021-12-30,Visa,150
cust2,merchantA,2021-12-30,Visa,200
cust3,merchantB,2021-12-31,MasterCard,300
cust4,merchantA,2021-12-30,Visa,50
"""

correct = """merchant_id,card_type,payout_date,amount
merchantA,Visa,2021-12-30,400
merchantB,MasterCard,2021-12-31,300"""

def register_receivables(csv_data):
	# Initialize a defaultdict to hold the aggregated amounts
	aggregated_data = defaultdict(int)

	# Read the input data
	reader = csv.reader(csv_data.strip().splitlines())
	next(reader)  # Skip the header line

	# Process each row in the CSV
	for row in reader:
		customer_id, merchant_id, payout_date, card_type, amount = row
		amount = int(amount)  # Convert amount to an integer
		key = (merchant_id, card_type, payout_date)
		aggregated_data[key] += amount

	# Prepare the output with the header
	output = ["merchant_id,card_type,payout_date,amount"]

	# Append each aggregated entry to the output
	for (merchant_id, card_type, payout_date), amount in aggregated_data.items():
		output.append(f"{merchant_id},{card_type},{payout_date},{amount}")

	return "\n".join(output)
assert register_receivables(csv_data) == correct

# def parse_amount(amount_str, total_amount):
#     """Parse the amount string which can be either a fixed number or a percentage"""
#     if amount_str.endswith('%'):
#         percentage = float(amount_str.rstrip('%'))
#         return int(total_amount * percentage / 100)
#     return int(amount_str)

def update_receivables(registered_receivables, contracts_csv):
	receivables = list(csv.reader(registered_receivables.strip().splitlines()))
	receivables_dict = {(row[0], row[1], row[2]): int(row[3]) for row in receivables[1:]}  # key: (merchant_id, card_type, payout_date)

	updated_receivables = []
	reader = csv.reader(contracts_csv.strip().splitlines())
	header = next(reader)  # Skip the header line
	# Process each contract
	for contract in reader:
		contract_id, merchant_id, payout_date, card_type, amount = contract
		amount = int(amount)  # Convert amount to an integer
		key = (merchant_id, card_type, payout_date)

		if key in receivables_dict:
			# Get the existing receivable
			existing_amount = receivables_dict[key]

			# Calculate actual amount based on whether it's fixed or percentage
			# amount = parse_amount(amount_str, existing_amount)

			if existing_amount > amount:
				# Partial payment, update the existing receivable
				receivables_dict[key] -= amount

			# Remove the existing receivable from the dictionary
			elif existing_amount == amount:
				del receivables_dict[key]	

			# Add the contract receivable
			updated_receivables.append([contract_id, card_type, payout_date, str(amount)])
			
		# else:
			# If there's no existing receivable, just add the contract
			# updated_receivables.append([contract_id, card_type, payout_date, str(amount)])

	for k,v in receivables_dict.items():
		receiveable = list(k)
		receiveable.append(str(v))
		updated_receivables.append(receiveable)

	# Prepare output with the correct header
	header[0] = 'id'
	output = [header]
	output.extend(updated_receivables)

	# Convert back to CSV format for output
	output_csv = "\n".join([",".join(row) for row in output])
	return output_csv


transactions = """customer_id,merchant_id,payout_date,card_type,amount
cust1,merchantA,2022-01-05,Visa,300
cust2,merchantA,2022-01-05,Visa,200
cust3,merchantB,2022-01-06,MasterCard,1000
"""
contracts = """contract_id,merchant_id,payout_date,card_type,amount
contract1,merchantA,2022-01-05,Visa,500
"""

correct = """id,merchant_id,payout_date,card_type,amount
contract1,Visa,2022-01-05,500
merchantB,MasterCard,2022-01-06,1000"""

assert update_receivables(register_receivables(transactions), contracts) == correct

transactions_t1 = """customer_id,merchant_id,payout_date,card_type,amount
cust1,merchantA,2022-01-07,Visa,500
cust2,merchantA,2022-01-07,Visa,250
cust3,merchantB,2022-01-08,MasterCard,1250
cust4,merchantC,2022-01-09,Visa,1500
"""
contracts_t1 = """contract_id,merchant_id,payout_date,card_type,amount
contract1,merchantA,2022-01-07,Visa,750
contract2,merchantC,2022-01-09,Visa,1500
"""
# print(update_receivables(register_receivables(transactions_t1), contracts_t1))


# Test functions
def test_update_receivables():
    # Test Case 1: Partial Payment with Existing Receivable
	registered_receivables = """customer_id,merchant_id,payout_date,card_type,amount
cust1,merchantA,2022-01-07,Visa,500
cust2,merchantA,2022-01-07,Visa,250
cust3,merchantB,2022-01-08,MasterCard,1250
cust4,merchantC,2022-01-09,Visa,1500
"""

	contracts = """contract_id,merchant_id,payout_date,card_type,amount
contract1,merchantA,2022-01-07,Visa,500
contract2,merchantC,2022-01-09,Visa,1500
"""
	expected_output = """id,merchant_id,payout_date,card_type,amount
contract1,Visa,2022-01-07,500
contract2,Visa,2022-01-09,1500
merchantA,Visa,2022-01-07,250
merchantB,MasterCard,2022-01-08,1250"""
	actual_output = update_receivables(register_receivables(registered_receivables), contracts)
	assert actual_output == expected_output

    # Test Case 2: Exact Match Payment
	registered_receivables = """customer_id,merchant_id,payout_date,card_type,amount
cust1,merchantA,2022-01-07,Visa,500
cust2,merchantA,2022-01-07,Visa,250
cust3,merchantB,2022-01-08,MasterCard,1250
cust4,merchantC,2022-01-09,Visa,1500
"""

	contracts = """contract_id,merchant_id,payout_date,card_type,amount
contract1,merchantA,2022-01-07,Visa,750
contract2,merchantC,2022-01-09,Visa,1500
"""
	expected_output = """id,merchant_id,payout_date,card_type,amount
contract1,Visa,2022-01-07,750
contract2,Visa,2022-01-09,1500
merchantB,MasterCard,2022-01-08,1250"""
	assert update_receivables(register_receivables(registered_receivables), contracts) == expected_output

    # Test Case 3: Contract Exceeds Existing Amount (Invalid Case)
    
    # Test Case 4: Multiple Contracts for One Merchant
    
    # Test Case 5: No Existing Receivables
    
    # Test Case 6: Complex Scenario
    
# Run tests
test_update_receivables()
print("All tests passed!")

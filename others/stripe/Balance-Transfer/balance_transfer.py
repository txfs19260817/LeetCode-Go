def redistribute_balances(accounts):
    # Step 1: Separate accounts into surplus and deficit lists
    surplus = []
    deficit = []
    
    for acc, balance in accounts.items():
        if balance > 100:
            surplus.append((acc, balance - 100))  # Store the surplus amount
        elif balance < 100:
            deficit.append((acc, 100 - balance))  # Store the deficit amount
    
    # Step 2: Redistribute money
    transfers = []
    for def_acc, def_amt in deficit:
        # While there is still a deficit to cover
        while def_amt > 0 and surplus:
            sur_acc, sur_amt = surplus[0]  # Get the first surplus account
            transfer_amt = min(def_amt, sur_amt)  # Transfer as much as needed or available
            
            # Add the transfer record
            transfers.append({'from': sur_acc, 'to': def_acc, 'amount': transfer_amt})
            
            # Update the surplus and deficit amounts
            def_amt -= transfer_amt
            sur_amt -= transfer_amt
            
            # If the surplus account is exhausted, remove it
            if sur_amt == 0:
                surplus.pop(0)
            else:
                surplus[0] = (sur_acc, sur_amt)  # Update the surplus list
    
    # Return the list of transfers
    return transfers

accounts = {
    'AU': 80,
    'US': 140,
    'MX': 110,
    'SG': 120,
    'FR': 70
}
# Execute redistribution
transfers = redistribute_balances(accounts)
# Print the transfers
for transfer in transfers:
    # print(f"from: {transfer['from']}, to: {transfer['to']}, amount: {transfer['amount']}")
    pass

# Final Balance Check: We apply the transfers to the initial balances and check whether all accounts have at least 100 after the transfers.
# Best Effort Check:
# We calculate the total surplus (amount of money above 100 in surplus accounts) and the total deficit
# (amount needed to bring deficit accounts to 100).
# If there’s still a surplus while some accounts are below 100, the transfers are not "best effort". 
# Otherwise, if the transfers fully utilized the surplus to reduce deficits, they are considered "best effort".

def evaluate_transfers_best_effort(accounts, transfers):
    # Calculate total balance
    total_balance = sum(accounts.values())
    num_accounts = len(accounts)
    
    # Check if it's even possible to meet the condition
    if total_balance < 100 * num_accounts:
        print("It's impossible for all accounts to reach at least 100.")
    
    # Apply transfers
    for transfer in transfers:
        from_account = transfer['from']
        to_account = transfer['to']
        amount = transfer['amount']
        accounts[from_account] -= amount
        accounts[to_account] += amount
    
    # Check final balances
    final_balances = {account: balance for account, balance in accounts.items()}
    
    # Check for accounts above 100
    below_hundred = [account for account, balance in final_balances.items() if balance < 100]
    
    if below_hundred:
        print("Accounts below 100 after transfers:", below_hundred)
        return "Not a best effort since some accounts have balances below 100."
    else:
        print("All accounts are at or above 100 after transfers.")
        return "Best effort since all accounts are at or above 100."

# Example usage
accounts = {
    'AU': 80,
    'US': 140,
    'MX': 110,
    'SG': 120,
    'FR': 70
}

transfers = [
    {'from': 'US', 'to': 'AU', 'amount': 20},
    {'from': 'US', 'to': 'FR', 'amount': 20},
    {'from': 'MX', 'to': 'FR', 'amount': 10}
]

result = evaluate_transfers_best_effort(accounts, transfers)
print(accounts, result)



def minimize_transfers(accounts):
    # Step 1: Separate accounts into surplus and deficit lists
    surplus = []
    deficit = []
    
    for acc, balance in accounts.items():
        if balance > 100:
            surplus.append((acc, balance - 100))  # Store the surplus amount
        elif balance < 100:
            deficit.append((acc, 100 - balance))  # Store the deficit amount
    
    # Sort surplus and deficit lists (not necessary but helps in understanding)
    surplus.sort(key=lambda x: -x[1])  # Sort surplus descending
    deficit.sort(key=lambda x: -x[1])  # Sort deficit descending
    
    # Step 2: Greedily match surplus to deficit
    transfers = []
    i, j = 0, 0  # Pointers to surplus and deficit
    
    while i < len(surplus) and j < len(deficit):
        sur_acc, sur_amt = surplus[i]  # Current surplus account
        def_acc, def_amt = deficit[j]  # Current deficit account
        
        # Transfer the minimum of surplus amount and deficit amount
        transfer_amt = min(sur_amt, def_amt)
        # Record the transfer
        transfers.append({'from': sur_acc, 'to': def_acc, 'amount': transfer_amt})
        
        # Update surplus and deficit amounts
        surplus[i] = (sur_acc, sur_amt - transfer_amt)
        deficit[j] = (def_acc, def_amt - transfer_amt)
        
        # If surplus is exhausted, move to the next surplus account
        if surplus[i][1] == 0:
            i += 1
        # If deficit is satisfied, move to the next deficit account
        if deficit[j][1] == 0:
            j += 1
    return transfers

# Example input
accounts = {
    'AU': 80,
    'US': 140,
    'MX': 110,
    'SG': 120,
    'FR': 70
}

# Execute the optimal redistribution
optimal_transfers = minimize_transfers(accounts)

# Print the transfers
for transfer in optimal_transfers:
    print(f"from: {transfer['from']}, to: {transfer['to']}, amount: {transfer['amount']}")

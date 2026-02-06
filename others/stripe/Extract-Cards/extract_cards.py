def extract_cards(card_string):
    index = 0
    cards = []
    
    while index < len(card_string):
        length = int(card_string[index:index + 2])  # Extract the length (first 2 digits)
        index += 2
        card = card_string[index:index + length]  # Extract the card based on the length
        cards.append(card)
        index += length
    return cards

def filter_supported_cards(cards, supported_cards):
    # Only keep cards that are in the system's supported cards
    return [card for card in cards if card in supported_cards]

def filter_generic_cards(cards, supported_cards):
    # Check if any system-supported card starts with the generic card
    unique_cards = []
    seen = set()
    
    for card in cards:
        if card in supported_cards:
            unique_cards.append(card)
            seen.add(card)
        else:
            for sup_card in supported_cards:
                if sup_card.startswith(card) and sup_card not in seen:
                    unique_cards.append(sup_card)
                    seen.add(sup_card)
                    
    return unique_cards
# Example usage:

# Part 1: Extract cards
card_string = "03abc02de04wxyz"  # Example: 03 for 'abc', 02 for 'de', 04 for 'wxyz'
cards = extract_cards(card_string)
print("Extracted cards:", cards)

# Part 2: Filter system-supported cards
supported_cards = ['abc', 'wxyz']  # Example of system supported cards
filtered_cards = filter_supported_cards(cards, supported_cards)
print("Filtered supported cards:", filtered_cards)

# Part 3: Include cards that start with a generic card
generic_filtered_cards = filter_generic_cards(cards, supported_cards)
print("Generic filtered cards:", generic_filtered_cards)

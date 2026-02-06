def redact_card_numbers(s):
    def redact_token(token):
        # Initialize variables to store the processed token
        result = []
        current_digits = []

        for char in token:
            if char.isdigit():
                current_digits.append(char)  # Collect digits
            else:
                # Process the digits if we have between 13 and 16
                if 13 <= len(current_digits) <= 16:
                    result.append(
                        "x" * (len(current_digits) - 4) + "".join(current_digits[-4:])
                    )
                else:
                    result.append("".join(current_digits))  # If not, keep them as is
                result.append(char)  # Add the non-digit character
                current_digits = []  # Reset digit collection

        # Check for any remaining digits at the end of the token
        if 13 <= len(current_digits) <= 16:
            result.append(
                "x" * (len(current_digits) - 4) + "".join(current_digits[-4:])
            )
        else:
            result.append("".join(current_digits))

        return "".join(result)

    # Split the input string into tokens and apply the redaction logic to each token
    tokens = s.split()
    redacted_tokens = [redact_token(token) for token in tokens]

    # Join the tokens back into a string with spaces
    return " ".join(redacted_tokens)


# Example test cases
# print(redact_card_numbers("a1234567890123456b is a number"))  # "axxxxxxxxxxxx3456b is a number"
# print(redact_card_numbers("basic_string 12345 no redaction"))  # "basic_string 12345 no redaction"
# print(redact_card_numbers("an embedded number 1234567890123456 in the string"))  # "an embedded number xxxxxxxxxxxx3456 in the string"
# print(redact_card_numbers("a123123123123b"))  # "a123123123123b"


def is_valid_card_number(number: str) -> bool:
    """Helper function to validate the card number based on the rules for Visa, Mastercard, and American Express."""
    # Visa: 13 or 16 digits, starts with '4'
    if (len(number) == 13 or len(number) == 16) and number.startswith("4"):
        return True
    # American Express: 15 digits, starts with '34' or '37'
    elif len(number) == 15 and (number.startswith("34") or number.startswith("37")):
        return True
    # Mastercard: 16 digits, starts with '51'-'55' or '2221'-'2720'
    elif len(number) == 16 and (
        51 <= int(number[:2]) <= 55 or 2221 <= int(number[:4]) <= 2720
    ):
        return True
    return False


def redact_card_numbers(s: str) -> str:
    def redact_token(token: str) -> str:
        result = []
        current_digits = []

        for char in token:
            if char.isdigit():
                current_digits.append(char)
            else:
                # Check the digits before the non-digit character and redact if it's a valid card number
                if is_valid_card_number("".join(current_digits)):
                    result.append(
                        "x" * (len(current_digits) - 4) + "".join(current_digits[-4:])
                    )
                else:
                    result.append(
                        "".join(current_digits)
                    )  # Add unchanged digits if not valid
                result.append(char)  # Add the non-digit character
                current_digits = []  # Reset digit collection

        # Process any remaining digits at the end
        if is_valid_card_number("".join(current_digits)):
            result.append(
                "x" * (len(current_digits) - 4) + "".join(current_digits[-4:])
            )
        else:
            result.append("".join(current_digits))

        return "".join(result)

    # Split the string into tokens and apply redaction
    tokens = s.split()
    redacted_tokens = [redact_token(token) for token in tokens]

    return " ".join(redacted_tokens)


# Example test cases
# print(redact_card_numbers("a1234567890123456b is a number"))  # "a1234567890123456b is a number" (not a valid card)
# print(redact_card_numbers("4111111111111111 is a Visa card"))  # "xxxxxxxxxxxx1111 is a Visa card"
# print(redact_card_numbers("341234567890123 is an Amex card"))  # "xxxxxxxxxxx0123 is an Amex card"
# print(redact_card_numbers("5512345678901234 is a Mastercard"))  # "xxxxxxxxxxxx1234 is a Mastercard"
# print(redact_card_numbers("2221123412341234 is another Mastercard"))  # "xxxxxxxxxxxx1234 is another Mastercard"

# print(redact_card_numbers("basic_string 12345 no redaction"))
# print(redact_card_numbers("1234567890123456 is not a card"))
# print(redact_card_numbers("a4234567890123456b is a valid visa"))


def luhn_checksum(card_number: str) -> bool:
    """Helper function to validate the card number using the Luhn algorithm."""
    total = 0
    reverse_digits = card_number[::-1]

    for i, digit in enumerate(reverse_digits):
        n = int(digit)
        if i % 2 == 1:  # Every second digit (from the rightmost digit)
            n *= 2
            if n > 9:
                n -= 9
        total += n

    # The card is valid if the total modulo 10 is 0
    return total % 10 == 0


def is_valid_card_number(number: str) -> bool:
    """Helper function to validate the card number based on the rules for Visa, Mastercard, and American Express, and the Luhn checksum."""
    # Visa: 13 or 16 digits, starts with '4'
    if (len(number) == 13 or len(number) == 16) and number.startswith("4"):
        return luhn_checksum(number)
    # American Express: 15 digits, starts with '34' or '37'
    elif len(number) == 15 and (number.startswith("34") or number.startswith("37")):
        return luhn_checksum(number)
    # Mastercard: 16 digits, starts with '51'-'55' or '2221'-'2720'
    elif len(number) == 16 and (
        51 <= int(number[:2]) <= 55 or 2221 <= int(number[:4]) <= 2720
    ):
        return luhn_checksum(number)
    return False


def redact_card_numbers(s: str) -> str:
    def redact_token(token: str) -> str:
        result = []
        current_digits = []

        for char in token:
            if char.isdigit():
                current_digits.append(char)
            else:
                # Check the digits before the non-digit character and redact if it's a valid card number
                if is_valid_card_number("".join(current_digits)):
                    result.append(
                        "x" * (len(current_digits) - 4) + "".join(current_digits[-4:])
                    )
                else:
                    result.append(
                        "".join(current_digits)
                    )  # Add unchanged digits if not valid
                result.append(char)  # Add the non-digit character
                current_digits = []  # Reset digit collection

        # Process any remaining digits at the end
        if is_valid_card_number("".join(current_digits)):
            result.append(
                "x" * (len(current_digits) - 4) + "".join(current_digits[-4:])
            )
        else:
            result.append("".join(current_digits))

        return "".join(result)

    # Split the string into tokens and apply redaction
    tokens = s.split()
    redacted_tokens = [redact_token(token) for token in tokens]

    return " ".join(redacted_tokens)


# Example test cases
print(
    redact_card_numbers("4111111111111111 is a Visa card")
)  # "xxxxxxxxxxxx1111 is a Visa card"
print(
    redact_card_numbers("341234567890123 is an Amex card")
)  # "xxxxxxxxxxx0123 is an Amex card"
print(
    redact_card_numbers("5112345678901234 is a Mastercard")
)  # "xxxxxxxxxxxx1234 is a Mastercard"
print(
    redact_card_numbers("2221123412341234 is another Mastercard")
)  # "xxxxxxxxxxxx1234 is another Mastercard"
print(
    redact_card_numbers("1234567890123456 is an invalid card")
)  # "1234567890123456 is an invalid card"

print(redact_card_numbers("421111111111111111 is not valid"))
print(redact_card_numbers("4234567890123456 is valid"))

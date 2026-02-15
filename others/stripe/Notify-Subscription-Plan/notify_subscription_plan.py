from notify_subscription_plan_part3 import (
    Part3Notifier,
    format_notifications,
    generate_notifications,
)

__all__ = ["Part3Notifier", "generate_notifications", "format_notifications"]


if __name__ == "__main__":
    input_data = {
        "schedule": {
            "start": "Welcome",
            "-15": "Upcoming expiration",
            "end": "Expired",
        },
        "user_accounts": [
            {"name": "A", "plan": "X", "begin_date": 0, "duration": 30},
            {"name": "B", "plan": "Y", "begin_date": 1, "duration": 15},
        ],
        "subscription_changes": [
            {"name": "B", "day": 5, "new_plan": "Z"},
            {"name": "A", "day": 10, "extension": 10},
            {"name": "A", "day": 20, "new_plan": "Gold", "extension": 5},
        ],
    }

    actual = generate_notifications(input_data)
    expected = [
        "0: [Welcome] A, subscribe in plan X",
        "1: [Welcome] B, subscribe in plan Y",
        "1: [Upcoming expiration] B, subscribe in plan Y",
        "5: [Updated] B, subscribe in plan Z",
        "10: [Renewed] A, subscribe in plan X",
        "16: [Expired] B, subscribe in plan Z",
        "20: [Updated] A, subscribe in plan Gold",
        "20: [Renewed] A, subscribe in plan Gold",
        "30: [Upcoming expiration] A, subscribe in plan Gold",
        "45: [Expired] A, subscribe in plan Gold",
    ]
    assert actual == expected, f"Unexpected output:\n{actual}"

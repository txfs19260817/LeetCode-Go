from collections import defaultdict

users = [
    {"name": "A", "plan": "X", "begin_date": 0, "duration": 30},
    {"name": "B", "plan": "Y", "begin_date": 1, "duration": 15},
]
changes = [
    {"name": "A", "new_plan": "Y", "change_date": 5},
    {"name": "B", "extension": 15, "change_date": 3},
]

# Email notifications
events = defaultdict(list)

# User information tracking
user_info = {}


def schedule_event(date, message):
    events[date].append(message)


def add_user_events(user):
    name, plan, begin_date, duration = (
        user["name"],
        user["plan"],
        user["begin_date"],
        user["duration"],
    )
    user_info[name] = {
        "current_plan": plan,
        "begin_date": begin_date,
        "duration": duration,
    }

    schedule_event(begin_date, f"[Welcome] {name}, subscribe in plan {plan}")

    upcoming_expiration_date = begin_date + duration - 15
    if upcoming_expiration_date >= begin_date:
        schedule_event(
            upcoming_expiration_date,
            f"[Upcoming expiration] {name}, subscribe in plan {plan}",
        )

    expiration_date = begin_date + duration
    schedule_event(expiration_date, f"[Expired] {name}, subscribe in plan {plan}")


def process_change(change):
    name = change["name"]
    current_plan, begin_date, duration = (
        user_info[name]["current_plan"],
        user_info[name]["begin_date"],
        user_info[name]["duration"],
    )
    if "new_plan" in change:
        new_plan, change_date = change["new_plan"], change["change_date"]

        schedule_event(change_date, f"[Changed] {name}, subscribe in plan {new_plan}")
        user_info[name]["current_plan"] = new_plan

        expiration_date = begin_date + duration
        upcoming_expiration_date = expiration_date - 15

        if upcoming_expiration_date >= change_date:
            # Update upcoming expiration event
            events[upcoming_expiration_date] = [
                msg for msg in events[upcoming_expiration_date] if name not in msg
            ]
            schedule_event(
                upcoming_expiration_date,
                f"[Upcoming expiration] {name}, subscribe in plan {new_plan}",
            )

        # Update expiration event
        events[expiration_date] = [
            msg for msg in events[expiration_date] if name not in msg
        ]
        schedule_event(
            expiration_date, f"[Expired] {name}, subscribe in plan {new_plan}"
        )

    elif "extension" in change:
        extension, change_date = change["extension"], change["change_date"]

        duration += extension
        user_info[name]["duration"] = duration

        schedule_event(
            change_date, f"[Renewed] {name}, subscribe in plan {current_plan}"
        )

        expiration_date = begin_date + duration
        upcoming_expiration_date = expiration_date - 15

        # Update upcoming expiration and expiration events
        events[upcoming_expiration_date] = [
            msg for msg in events[upcoming_expiration_date] if name not in msg
        ]
        events[expiration_date] = [
            msg for msg in events[expiration_date] if name not in msg
        ]

        if upcoming_expiration_date >= change_date:
            schedule_event(
                upcoming_expiration_date,
                f"[Upcoming expiration] {name}, subscribe in plan {current_plan}",
            )
        schedule_event(
            expiration_date, f"[Expired] {name}, subscribe in plan {current_plan}"
        )


# Process users and their events
for user in users:
    add_user_events(user)

# Process changes
for change in changes:
    process_change(change)

# Output email notifications in order
for day in sorted(events.keys()):
    for message in events[day]:
        print(f"{day}: {message}")

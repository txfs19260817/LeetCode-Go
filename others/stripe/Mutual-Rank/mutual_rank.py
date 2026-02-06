# Define the wishlist data structure as a dictionary
wishlists = {
    "a": ["c", "d"],
    "b": ["d", "a", "c"],
    "c": ["a", "b"],
    "d": ["c", "a", "b"],
}


def get_wishlist_rank(user, other_user):
    """Return the rank (0-based index) of other_user in user's wishlist"""
    if user not in wishlists or other_user not in wishlists[user]:
        return -1
    return wishlists[user].index(other_user)


def has_mutual_pair_for_rank(username, rank):
    """Check if the user has a mutual pair at the specified rank"""
    if username not in wishlists or rank >= len(wishlists[username]):
        return False

    # Get the user at the specified rank in username's wishlist
    other_user = wishlists[username][rank]

    # Check if the original user appears at the same rank in other_user's wishlist
    return (
        other_user in wishlists
        and username in wishlists[other_user]
        and get_wishlist_rank(other_user, username) == rank
    )


def has_mutual_first_choice(username):
    """Check if the user has a mutual first choice (rank 0)"""
    # Get the user at the specified rank in the source user's wishlist
    # target_user = wishlists[username][0]

    # Check if the source user appears at the same rank in the target's wishlist
    # return get_wishlist_rank(target_user, user) == rank
    return has_mutual_pair_for_rank(username, 0)


assert has_mutual_first_choice("a") == True  # a and c are mutual first choices
assert has_mutual_first_choice("b") == False  # b's first choice (d) doesn't reciprocate

assert has_mutual_pair_for_rank("a", 0) == True  # a and c
assert has_mutual_pair_for_rank("a", 1) == True  # a and d are mutual second choices


def changed_pairings(username, rank):
    """
    Return list of users whose mutual ranking status would change if the entry
    at the given rank was bumped up one position
    """
    if (
        username not in wishlists or rank >= len(wishlists[username]) or rank == 0
    ):  # Can't bump up rank 0
        return []

    affected_users = []
    current_user = wishlists[username][rank]
    previous_user = wishlists[username][rank - 1]

    # Check if the current pairing would be affected
    if (
        has_mutual_pair_for_rank(username, rank)
        or get_wishlist_rank(current_user, username) == rank - 1
    ):
        affected_users.append(current_user)

    # Check if the previous pairing would be affected
    if (
        has_mutual_pair_for_rank(username, rank - 1)
        or get_wishlist_rank(previous_user, username) == rank
    ):
        affected_users.append(previous_user)

    return list(set(affected_users))  # Remove any duplicates


assert changed_pairings("d", 1) == [
    "a"
]  # a and d would lose mutual second-choice status
assert changed_pairings("b", 2) == ["c"]  # b and c would become mutual second choices
assert changed_pairings("b", 1) == []  # no mutual pairings affected


def changed_antipairings(username, rank):
    """
    Return list of users whose mutual anti-ranking status would change if the entry
    at the given rank was bumped up one position
    """
    if username not in wishlists or rank >= len(wishlists[username]) or rank == 0:
        return []

    affected_users = []
    list_length = len(wishlists[username])

    # Convert rank from start to rank from end for the users being swapped
    current_from_end = list_length - 1 - rank
    previous_from_end = list_length - rank

    current_user = wishlists[username][rank]
    previous_user = wishlists[username][rank - 1]

    # Check if current user's anti-pairing would be affected
    if current_user in wishlists:
        other_list_length = len(wishlists[current_user])
        current_user_rank = get_wishlist_rank(current_user, username)
        if current_user_rank != -1:
            current_user_from_end = other_list_length - 1 - current_user_rank
            if (
                current_from_end == current_user_from_end
                or previous_from_end == current_user_from_end
            ):
                affected_users.append(current_user)

    # Check if previous user's anti-pairing would be affected
    if previous_user in wishlists:
        other_list_length = len(wishlists[previous_user])
        previous_user_rank = get_wishlist_rank(previous_user, username)
        if previous_user_rank != -1:
            previous_user_from_end = other_list_length - 1 - previous_user_rank
            if (
                current_from_end == previous_user_from_end
                or previous_from_end == previous_user_from_end
            ):
                affected_users.append(previous_user)

    return list(set(affected_users))  # Remove any duplicates


print(changed_antipairings("b", 2))  # ['c']
print(changed_antipairings("a", 1))  # ['c', 'd']

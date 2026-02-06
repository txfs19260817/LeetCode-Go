def get_user_features(user, features):
	res = []
	user_id, location = user['id'], user['location']
	for feature in features:
		location_effect = True if 'locations' in feature else False
		ab_test_effect = True if 'abTest' in feature else False
		if location_effect and location not in feature['locations']:
			continue
		if ab_test_effect and user_id % 2 != 0
			continue
		res.append(feature['id'])
	return res

def get_user_features(user, features):
	res = []
	user_id, location, opt_in, opt_out = user['id'], user['location'], user.get('optIn', []), user.get('optOut', [])
	for feature in features:
		location_effect = True if 'locations' in feature else False
		ab_test_effect = True if 'abTest' in feature else False
		if feature['id'] in opt_out:
			continue
		if location_effect and location not in feature['locations']:
			continue
		if ab_test_effect and user_id % 2 != 0 and feature['id'] not in opt_in:
			continue
		res.append(feature['id'])
	return res


users = [
{ "id": 0, "name": "eva",    "location": "US", "optIn": ["annual_sale"] },
{ "id": 1, "name": "tess",   "location": "US", "optIn": ["annual_sale"] },
{ "id": 2, "name": "rahool", "location": "CA", "optOut": ["enhanced_comments", "canada_promotion"] },
{ "id": 3, "name": "amanda", "location": "CA", "optIn": ["annual_sale"] }
]

features = [
	{
		"id": "annual_sale",
		"locations": ["US"],
		"abTest": True,
	},
	{
		"id": "enhanced_comments",
		"abTest": True,
	},
	{
		"id": "canada_promotion",
		"locations": ["CA"],
	}
]

# print(get_user_features(users[3], features))

for user in users:
	print(get_user_features(user, features))	


def get_user_features(user, features):
    """
    Determine which features apply to a given user based on location, A/B testing rules,
    and opt-in/out preferences.
    
    Args:
        user: Dict containing user information with keys 'id', 'name', 'location',
              and optional 'optIn' and 'optOut' lists
        features: List of feature dicts, each containing 'id' and optionally 'locations' and 'abTest'
    
    Returns:
        set: Set of feature IDs that apply to the given user
    """
    applicable_features = set()
    
    # Create quick lookup for features
    feature_dict = {f["id"]: f for f in features}
    
    # Get user preferences
    opt_ins = set(user.get("optIn", []))
    opt_outs = set(user.get("optOut", []))
    
    for feature in features:
        feature_id = feature["id"]
        
        # Skip if explicitly opted out
        if feature_id in opt_outs:
            continue
            
        # Check if feature should be included based on rules
        should_include = True
        
        # Always check location restriction, even for opt-ins
        if "locations" in feature:
            if user["location"] not in feature["locations"]:
                should_include = False
                # Location restrictions cannot be overridden by opt-in
                continue
        
        # Check A/B test restriction if it exists
        # Skip A/B test check if user has opted in
        if "abTest" in feature and feature["abTest"]:
            if feature_id not in opt_ins and user["id"] % 2 != 0:  # Odd IDs are excluded
                should_include = False
        
        # Include if either:
        # 1. All rules pass naturally
        # 2. User has opted in and location allows it
        if should_include or (feature_id in opt_ins):
            applicable_features.add(feature_id)
            
    return applicable_features

for user in users:
	print(get_user_features(user, features))

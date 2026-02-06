def parse_accept_language(accept_language, supported_languages):
    # Split the accept_language string into a list of language tags
    requested_languages = [lang.strip() for lang in accept_language.split(',')]
    # Filter and sort the supported languages based on the requested order
    result = []
    for lang in requested_languages:
        if lang in supported_languages:
            result.append(lang)
    return result

def test_parse_accept_language():
    assert parse_accept_language("en-US, fr-CA, fr-FR", ["fr-FR", "en-US"]) == ["en-US", "fr-FR"]
    assert parse_accept_language("fr-CA, fr-FR", ["en-US", "fr-FR"]) == ["fr-FR"]
    assert parse_accept_language("en-US", ["en-US", "fr-CA"]) == ["en-US"]  
    # Test case 4: No matching languages
    assert parse_accept_language("es-ES, de-DE", ["fr-FR", "it-IT"]) == []
    # Test case 5: Partial matches
    assert parse_accept_language("en-US, en-GB, fr-FR", ["en-GB", "fr-CA", "de-DE"]) == ["en-GB"]
    print("All tests passed!")

test_parse_accept_language()

def parse_accept_language(accept_language, supported_languages):
    # Split the accept_language string into a list of language tags
    requested_languages = [lang.strip() for lang in accept_language.split(',')]
    
    # Filter and sort the supported languages based on the requested order
    result = []
    remaining_languages = supported_languages.copy()
    for lang in requested_languages:
        if lang in remaining_languages:
            # Find the original case version from remaining_languages
            # original_case = next(supported_lang for supported_lang in remaining_languages 
                                 # if supported_lang.lower() == lang.lower())
            result.append(lang)
            remaining_languages.remove(lang)
        elif lang == '*':
            # Add all remaining supported languages when encountering a wildcard
            result.extend(remaining_languages)
            break
        else:
            # Check for partial matches (e.g., "fr" matches "fr-CA" and "fr-FR")
            partial_matches = [supported_lang for supported_lang in remaining_languages 
                               if supported_lang.startswith(lang)]
            result.extend(partial_matches)
            for match in partial_matches:
                remaining_languages.remove(match)
    return result

def test_parse_accept_language():
    assert parse_accept_language("fr-FR, fr, *", ["en-US", "fr-CA", "fr-FR"]) == ['fr-FR', 'fr-CA', 'en-US']
    assert parse_accept_language("en-US, *", ["en-US", "fr-CA", "fr-FR"]) == ['en-US', 'fr-CA', 'fr-FR']
    print("All tests passed!")

test_parse_accept_language()


def parse_accept_language(accept_language, supported_languages):
    # Parse the accept_language string into (language, q-factor) pairs
    def parse_language_tag(tag):
        parts = tag.strip().split(';q=')
        lang = parts[0].strip()
        q = float(parts[1]) if len(parts) > 1 else 1.0
        return (lang, q)
    
    # Split and parse each language tag
    requested_languages = [parse_language_tag(lang) for lang in accept_language.split(',')]
    
    # Separate zero and non-zero priority languages
    zero_priority = [lang for lang, q in requested_languages if q == 0]
    non_zero = [(lang, q) for lang, q in requested_languages if q > 0]
    
    # Sort by q-factor in descending order
    non_zero.sort(key=lambda x: -x[1])
    
    # Initialize remaining languages as a list to maintain input order
    remaining_languages = list(supported_languages)
    
    # Keep track of excluded languages while maintaining order
    excluded_languages = []
    for lang in zero_priority:
        if lang == '*':
            # If * has q=0, all remaining languages should be excluded
            excluded_languages.extend([l for l in remaining_languages if l not in excluded_languages])
            remaining_languages = []
        else:
            # Exclude exact and partial matches
            to_exclude = [
                rem for rem in remaining_languages 
                if lang == rem or rem.startswith(lang+'-')
            ]
            for exc in to_exclude:
                if exc in remaining_languages and exc not in excluded_languages:
                    remaining_languages.remove(exc)
                    excluded_languages.append(exc)
    
    result = []
    # Process non-zero priority languages
    for lang, q in non_zero:
        if lang == '*':
            # Add all remaining supported languages when encountering a wildcard
            result.extend(remaining_languages)
            remaining_languages = []
        elif lang in remaining_languages:
            # Exact match
            result.append(lang)
            remaining_languages.remove(lang)
        else:
            # Check for partial matches (e.g., "fr" matches "fr-CA" and "fr-FR")
            partial_matches = [supported_lang for supported_lang in remaining_languages 
                             if supported_lang.startswith(lang + '-')]
            partial_matches.sort()  # Sort partial matches for consistency
            result.extend(partial_matches)
            for match in partial_matches:
                remaining_languages.remove(match)
    
    # Add excluded languages at the end
    result.extend(excluded_languages)
    return result

def test_parse_accept_language():
    assert parse_accept_language("fr-FR;q=1, fr-CA;q=0, fr;q=0.5", ["fr-FR", "fr-CA", "fr-BG"]) == ["fr-FR", "fr-BG", "fr-CA"]
    assert parse_accept_language("fr-FR;q=1, fr-CA;q=0, *;q=0.5", ["fr-FR", "fr-CA", "fr-BG", "en-US"]) == ["fr-FR", "fr-BG", "en-US", "fr-CA"]
    assert parse_accept_language("fr-FR;q=1, fr-CA;q=0.8, *;q=0.5", ["fr-FR", "fr-CA", "fr-BG", "en-US"]) == ["fr-FR", "fr-CA", "fr-BG", "en-US"]
test_parse_accept_language()

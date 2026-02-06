def expand_braces(input_str: str) -> list[str]:
    # Handle None or empty input
    if not input_str:
        return []
    
    # Find positions of braces
    open_brace = input_str.find('{')
    close_brace = input_str.find('}')
    
    # Check for invalid cases:
    # - Missing either brace
    # - Closing brace before opening brace
    # - Empty or single token between braces
    if (open_brace == -1 or close_brace == -1 or 
        close_brace < open_brace or 
        open_brace + 1 >= close_brace):
        return [input_str]
    
    # Extract parts
    prefix = input_str[:open_brace]
    suffix = input_str[close_brace + 1:]
    middle = input_str[open_brace + 1:close_brace]
    
    # Split tokens
    tokens = middle.split(',')
    
    # If less than 2 tokens, return original input
    if len(tokens) < 2:
        return [input_str]
    
    # Generate expanded strings
    return [f"{prefix}{token}{suffix}" for token in tokens]

def run_tests():
    # Part 1 - Valid cases
    print("Part 1 - Valid cases:")
    
    input1 = "/2022/{jan,feb,march}/report"
    print(f"Input: {input1}")
    print(f"Output: {expand_braces(input1)}\n")
    
    input2 = "prefix{a,c,d}suffix"
    print(f"Input: {input2}")
    print(f"Output: {expand_braces(input2)}\n")
    
    # Part 2 - Invalid cases
    print("Part 2 - Invalid cases:")
    
    input3 = "pre{mid}suf"
    print(f"Input: {input3}")
    print(f"Output: {expand_braces(input3)}\n")
    
    input4 = "pre{}suf"
    print(f"Input: {input4}")
    print(f"Output: {expand_braces(input4)}\n")
    
    input5 = "pre}mid{suf"
    print(f"Input: {input5}")
    print(f"Output: {expand_braces(input5)}\n")
    
    input6 = "premidsuf"
    print(f"Input: {input6}")
    print(f"Output: {expand_braces(input6)}")


if __name__ == "__main__":
    run_tests()


def expand_braces(input_str: str) -> list[str]:
    if not input_str:
        return []

    # Start with the initial string as our only result
    results = [input_str]
    
    # Keep expanding until no more valid brace patterns are found
    while True:
        new_results = []
        any_expansion = False
        
        # Process each string in our current results
        for s in results:
            # Find first set of braces
            open_brace = s.find('{')
            close_brace = s.find('}')
            
            # If no valid brace pattern, keep string as is
            if (open_brace == -1 or close_brace == -1 or 
                close_brace < open_brace or 
                open_brace + 1 >= close_brace):
                new_results.append(s)
                continue
                
            # Extract parts
            prefix = s[:open_brace]
            suffix = s[close_brace + 1:]
            middle = s[open_brace + 1:close_brace]
            
            # Split tokens
            tokens = middle.split(',')
            
            # If less than 2 tokens, keep string as is
            if len(tokens) < 2:
                new_results.append(s)
                continue
            
            # Expand this set of braces
            for token in tokens:
                new_results.append(f"{prefix}{token}{suffix}")
            any_expansion = True
        
        # If no expansions were performed, we're done
        if not any_expansion:
            break
            
        results = new_results
    
    return results


def run_tests():
    print("Part 1 - Single brace:")
    input1 = "/2022/{jan,feb,march}/report"
    print(f"Input: {input1}")
    print(f"Output: {expand_braces(input1)}\n")
    
    print("Part 2 - Invalid cases:")
    input2 = "pre{mid}suf"
    print(f"Input: {input2}")
    print(f"Output: {expand_braces(input2)}\n")
    
    input3 = "pre{}suf"
    print(f"Input: {input3}")
    print(f"Output: {expand_braces(input3)}\n")
    
    print("Part 3 - Multiple sequential braces:")
    input4 = "/{2021,2022}/{jan,feb}/report"
    print(f"Input: {input4}")
    print(f"Output: {expand_braces(input4)}\n")
    
    input5 = "/test/{a,b}/{x,y,z}/end"
    print(f"Input: {input5}")
    print(f"Output: {expand_braces(input5)}")


if __name__ == "__main__":
    run_tests()

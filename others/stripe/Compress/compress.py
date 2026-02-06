# https://www.1point3acres.com/bbs/thread-808688-1-1.html
def compress_minor(arr) :
	words = ''.join(arr)
	count = 0
	for i in range(1, len(words)-1):
		if words[i].isnumeric():
			count += int(words[i])
		else:
			count += 1

	return words[0] + str(count) + words[-1]

def compress_string(string, minor_parts):
	parts = string.split('/')
	compressed_parts = []

	for part in parts:
	    minor_split = part.split('.')
	    compressed_part = ''
	    for minor in minor_split:
	        compressed_part += minor[0] + str(len(minor) - 2) + minor[-1] + '.'

	    # Remove the extra '.' at the end
	    compressed_part = compressed_part[:-1]

	    minor_split = compressed_part.split(".")
	    if len(minor_split) <= minor_parts :
	    	compressed_parts.append(compressed_part)
	    else :
	    	tmp = minor_split[:minor_parts-1]
	    	tmp.append(compress_minor(minor_split[minor_parts-1:]))
	    	compressed_parts.append(".".join(tmp))

	return "/".join(compressed_parts)

str1 = "stripe.com/payments/checkout/customer.john.doe"
minor_parts1 = 2
compressed_str1 = compress_string(str1, minor_parts1)
print("Compressed String 1:", compressed_str1)

str2 = "www.api.stripe.com/checkout"
minor_parts2 = 3
compressed_str2 = compress_string(str2, minor_parts2)
print("Compressed String 2:", compressed_str2)


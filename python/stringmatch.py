def string_match(a, b):
 	shorter = min(len(a), len(b))
 	count = 0
 	for i in range(shorter-1):
 		a_str = a[i:i+2]
 		b_str = b[i:i+2]
 		if a_str == b_str:
 			count = count + 1
 	return count
def longestword(str):
	split = str.word()
	largest = split[0]
	for i in range(len(split)):
		if len[i]>largest:
			largest = len[i]
	return largest

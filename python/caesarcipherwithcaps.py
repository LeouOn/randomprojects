alphabet = {' ': ' ', 'a':'b', 'b':'c', 'c':'d', 'd':'E', 'e':'f', 'f':'g', 'g':'h', 'h':'I', 'i':'j', 'j':'k', 'k':'l', 'l':'m', 'm':'n', 'n':'O', 'o':'p', 'p':'q', 'q':'r', 'r':'s', 's':'t', 't':'U', 'u':'v', 'v':'w', 'w':'x', 'x':'y', 'y':'z', 'z':'A'}
def cipher(str):
	output = ''
	str = str.lower()
	for i in range(len(str)):
		if str[i].isalpha == True:
			output = output + alphabet.get(str[i+1])
		else:
			output = output + str[i]
	return output
print cipher("It is a good day to die he said")
print cipher("a confusing /:sentence:/[ this is not!!!!!!!~")
print cipher("g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj.")
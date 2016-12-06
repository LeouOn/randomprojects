def FirstReverse(str): 
    length = len(str)
    curpos = 0
    result = ""
    if curpos < length:
            result[curpos] = str[-curpos-1]
            curpos = curpos + 1
    else:
         return result
    # code goes here 
   
    
# keep this function call here  
print FirstReverse("I am so blue I'm greener than purple.")  

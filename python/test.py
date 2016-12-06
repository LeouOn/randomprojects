def sum67(nums):
  passedsix = False
  for i in range(len(nums)):
    if nums[i] == 7 and passedsix == True:
      passedsix = False
      nums[i]=0
    elif passedsix == True:
      nums[i]=0
    elif nums[i] == 6:
      passedsix = True
      nums[i]=0
  print nums
  return sum(nums)


sum67([1, 2, 2])
sum67([1, 2, 2, 6, 99, 99, 7])
sum67([1, 1, 6, 7, 2])
sum67([1, 6, 2, 2, 7, 1, 6, 99, 99, 7]) 
sum67([1, 6, 2, 6, 2, 7, 1, 6, 99, 99, 7])
sum67([2, 7, 6, 2, 6, 7, 2, 7]) 
sum67([2, 7, 6, 2, 6, 2, 7]) 
sum67([1, 6, 7, 7]) 
sum67([6, 7, 1, 6, 7, 7]) 
sum67([6, 8, 1, 6, 7]) 
sum67([]) 
sum67([6, 7, 11]) 
sum67([11, 6, 7, 11]) 
sum67([2, 2, 6, 7, 7]) 
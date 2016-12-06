def make_bricks(small, big, goal):
  smaller_goal = goal - (big*5) 
  if (goal/5)<=big and goal%5<=small:
    return True
  elif smaller_goal <= small and goal>(big*5):
    return True
  else:
    return False
        
name = 'Yune Leou-On'
age = 25
height = 73 # inches
height_cm = height * 2.54
weight = 230 # lbs
weight_kg = weight * 0.45
eyes = 'Brown'
teeth = 'Cream'
hair = 'Black'

print "Let's talk about %s" % name
print "He's %d inches tall." % height
print "He's %d pounds heavy" % weight
print "That is %d cm and %d kg for non-americans" % (height_cm, weight_kg)
print "That is a little heavy"
print "He's got %s eyes and %s hair." % (eyes, hair)
print "His teeth are usually %s depending on the coffee" % teeth

print "If I add %d, %d, and %d I get %d." % (age, height, weight, age + height + weight)
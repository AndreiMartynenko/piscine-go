
#cat the-final-cl-test/mystery/memberships/AAA the-final-cl-test/mystery/memberships/Delta_SkyMiles the-final-cl-test/mystery/memberships/Terminal_City_Library the-final-cl-test/mystery/memberships/Museum_of_Bash_History | sort | uniq -cd | sort | grep 4 | cut -f 2- -d 4
#grep -A4 "L337" the-final-cl-test/mystery/vehicles | grep -A3 "Honda" | grep -A1 "Blue" | cat | grep "Owner" | cut -f 2- -d ":" | cut -f 2- -d " "
#interview-699607:However, she reports seeing the car that fled the scene.  Describes it as a blue Honda, with a license plate that starts with "L337" and ends with "9"

#Get all clues
#grep -i clue  the-final-cl-test/mystery/crimescene

#Run through interviews
#grep -i clue  -C5 the-final-cl-test/mystery/interviews/*

# Get everyone that have 4 memberships
#cat the-final-cl-test/mystery/memberships/AAA the-final-cl-test/mystery/memberships/Delta_SkyMiles the-final-cl-test/mystery/memberships/Terminal_City_Library the-final-cl-test/mystery/memberships/Museum_of_Bash_History | sort | uniq -cd | sort | grep 4 | cut -f 2- -d 4 | cut -f 2- -d " " > membership

#Get everyone with Blue Honda L337
#grep -A4 "L337" the-final-cl-test/mystery/vehicles | grep -A3 "Honda" | grep -A1 "Blue" | cat | grep "Owner" | cut -f 2- -d ":" | cut -f 2- -d " "  > car

#Match those with Hondas and have 4 memberships and get thier names
#cat car membership | sort | uniq -d > suspects
##Result 
##Dartey Henv - not real name
##Hellen Maher - not real name

#grep -i "tall" -C5 the-final-cl-test/mystery/interviews/*

#Match those with Hondas and have 4 memberships and get thier addresses
#grep -A4 "L337*9" the-final-cl-test/mystery/vehicles | grep -C5 "Honda" | grep -A1 "Blue" | cat | grep "Owner" | cut -f 2- -d ":" | cut -f 2- -d " "  > car

#grep "Annabel" -i -d skip the-final-cl-test/mystery/people | grep -w "F" | cut -f 1 -d$'\t' | cut -d " " -f 2 > annabels
#grep "Annabel" -i -d skip the-final-cl-test/mystery/people | grep -w "F" | cut -f 1 -d$'\t' | cut -d " " -f 2 | grep -l the-final-cl-test/mystery/interviews/*
echo "Dartey Henv"




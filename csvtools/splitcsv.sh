#!/bin/bash


# grab the filename and sizeargs from user input
while getopts f:s: flag
do
	case "${flag}" in
		f) csvfile=${OPTARG};;
		s) size=${OPTARG};;
	esac
done

# create the output directory, if not exists
mkdir -p results

# csv file might be in different directory, strip to prefix just in case
filenameonly=${csvfile##*/}
prefix=${filenameonly%.*}



HDR=$(head -1 $csvfile)		# Pick up CSV header line to apply to each file

split -l $size $csvfile xyz	# Split the file into chunks of 20 lines each

n=0				# set counter = 0

for f in xyz*			# Go through all newly created chunks
do
   echo $HDR > results/${prefix}_${n}	# Write out header to new file called "Part(n)"
   cat $f >> results/${prefix}_${n}	# Append the chunk from the "split" command
   rm $f			# Remove temporary file
   ((n++))			# Increment name of output part
done

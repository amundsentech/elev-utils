#! python

from csv import reader,writer
import os

cwd = os.getcwd()

def main():
	for entry in os.scandir(cwd):
		if entry.path.endswith(".csv"):
			transpose(entry.path)

def transpose(infile):
	base=os.path.basename(infile)
	base=os.path.splitext(base)[0]
	outfile=base + "_transposed" + ".csv"

	print(infile)
	print(outfile)

	with open(infile) as f, open(outfile, 'w+', newline='') as fw: 
    		writer(fw, delimiter=',').writerows(zip(*reader(f, delimiter=',')))


	#a = izip(*csv.reader(open(filename, "rb")))
	#csv.writer(open(output, "wb")).writerows(a)

main()

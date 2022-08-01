import sys
def files():
    n = 0
    while True:
        n += 1
        yield open('./%d.part' % n, 'w')
    

pat = '!'
fs = files()
outfile = next(fs) 
count = 1
with open(sys.argv[1]) as infile:
    for line in infile:
        if pat not in line[0]:
            outfile.write(line)
        else:
            items = line.split(pat)
            outfile.write(items[0])
            for item in items[1:]:
                outfile = next(fs)
                outfile.write(pat + item)
                count += 1

print(outfile.name)
from math import dist
import sys
import subprocess
file = sys.argv[1]
dist_file = sys.argv[2]
subprocess.call(['python3', './csv-to-yaml.py', file, dist_file])
subprocess.call(['python3', 'replace.py', '--regex', "--in-place", '--from', ".*: ''", '--to', '', dist_file])
subprocess.call(['python3', 'replace.py', '--regex', "--in-place", '--from', ": ", '--to', '=', dist_file])
subprocess.call(['python3', 'replace.py', '--regex', "--in-place", '--from', "\n- ", '--to', '\n!\n', dist_file])
#get file object reference to the file
# file_to_count = open(dist_file, "r")

# #read content of file to string
# data = file_to_count.read()

# #get number of occurrences of the substring in the string
# occurrences = data.count("---")
subprocess.call(['python3', './split.py', dist_file])

for x in range(1, int(sys.argv[3]) + 2 ):
    working_file = str(x) + ".part"
    subprocess.call(['python3', 'replace.py', '--regex', "--in-place", '--from', "\n  ", '--to', '\n', working_file])
    subprocess.call(['python3', 'replace.py', '--regex', "--in-place", '--from', "!\n", '--to', '', working_file])
    subprocess.call(['/bin/zsh', './remove-relations.sh', working_file])
    
# subprocess.call(['python3', './properties_to_yaml.py', dist_file])
# dist_file = dist_file + ".yaml"  
subprocess.call(['/bin/zsh', './clean.sh', dist_file])

import os
import sys
import getopt

username = ""
project = ""

try:
    opts, args = getopt.getopt(sys.argv[1:],"hu:p:",["username=","project="])
except getopt.GetoptError:
    print("setup.py -u[--username] <your_username> -p[--project] <your_project_name>")
    sys.exit(2)

for opt, arg in opts:
    if opt == '-h':
        print("setup.py -u[--username] <your_username> -p[--project] <your_project_name>")
        sys.exit()
    elif opt in ("-u", "--username"):
        username = arg
    elif opt in ("-p", "--project"):
        project = arg

if len(username)*len(project) == 0:
    print("Missing <your_username> and <your_project_name>")
    print("setup.py -u[--username] <your_username> -p[--project] <your_project_name>")
    sys.exit(2)

path_replacement = "dee-ex/dx-golang-rest"
database_replacement = "dx-golang-rest"

print("Prepare your setup...")
for dname, dirs, files in os.walk(os.getcwd()):
    if ".git" in dname:
        continue
    for fname in files:
        fpath = os.path.join(dname, fname)
        with open(fpath) as f:
            s = f.read()
        s = s.replace(path_replacement, username + "/" + project)
        s = s.replace(database_replacement, project)
        with open(fpath, "w") as f:
            f.write(s)

print("Setup Successfully!")

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

path_replacement = "some_username/some_project"
database_replacement = "some_project"

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

#!/usr/bin/env python

# For invoking system commands:
import subprocess
import shlex

# For timestamps:
import time
from datetime import datetime

# For info on the system configuration:
# import platform

# For verbose:
import argparse

# For system utilities
import pathlib
import sys
import os

# For coloured text in stdout
import colorama


##################################################
## Author: NovusEdge
## Copyright: Copyright 2021, gohack
## License: MIT License
## Version: 0.1
##################################################


# Changing pwd to parent-directory (of script) for the CLI.
PATH = pathlib.Path(__file__).parent.absolute()
os.chdir(PATH)


# Getting current date and time
DATE = datetime.now().strftime("%d-%m-%Y")
TIME = datetime.now().strftime("%H:%M:%S")
__ERROR_LOGS_PATH = f"logs/setup_errors_{DATE}.log"


# Creating a log file for the current run of the setup
if pathlib.Path(__ERROR_LOGS_PATH).exists():
    __ERROR_LOGS = open(__ERROR_LOGS_PATH, "a")
    __ERROR_LOGS.write(f"\n\nLog__{TIME}\n\n")

else:
    __ERROR_LOGS = open(__ERROR_LOGS_PATH, "w")
    __ERROR_LOGS.write(f"\n\nLog__{TIME}\n\n")


# Creating a directory for binaries
if not pathlib.Path("bin/").exists():
    os.mkdir("bin")


# Fetching sys env
ENV = os.environ


# Cleaning old binaries (if present)
# get_proc = subprocess.call("go get", stderr=__ERROR_LOGS, stdout=sys.stdout, env=ENV)
os.system("go get")



# Closing opened files...
__ERROR_LOGS.close()


#go get
#cd bin/
#go clean
#go build ../commands/*

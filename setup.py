#!/usr/bin/env python

# For invoking system commands:
import subprocess
import shlex

# For timestamps:
import time
from datetime import datetime

# For info on the system configuration:
import platform

# For verbose:
import argparse

# For system utilities
import pathlib
import sys
import os

from colorama import Fore, Style

# Colors!
RED     = Fore.RED
YELLOW  = Fore.YELLOW
GREEN   = Fore.GREEN
MAGENTA = Fore.MAGENTA
CYAN    = Fore.CYAN
WHITE   = Fore.WHITE
CLEAR   = Style.RESET_ALL


##################################################
## Author: NovusEdge
## Copyright: Copyright 2021, gohack
## License: MIT License
## Version: 0.1
##################################################


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

    # Changing pwd to parent-directory (of script) for the CLI.
else:
    __ERROR_LOGS = open(__ERROR_LOGS_PATH, "w")
    __ERROR_LOGS.write(f"\n\nLog__{TIME}\n\n")


# Creating a directory for binaries
if not pathlib.Path("bin/").exists():
    os.mkdir("bin")

if not pathlib.Path("tool_bin/").exists():
    os.mkdir("tool_bin")


# Fetching sys env
ENV       = os.environ
PLATFORM  = platform.system()

# TODO: Figure out a way to do this quietly...
# get_proc = subprocess.call("go get", stderr=__ERROR_LOGS, stdout=sys.stdout, env=ENV)


if PLATFORM in 'linux darwin':
    proc = subprocess.Popen(["python3", "setup_scripts/setup_linux.py"], stderr=__ERROR_LOGS)
    proc.communicate()
    proc.close()

elif PLATFORM == 'win32':
    proc = subprocess.Popen(["python3", "setup_scripts/setup_windows.py"], stderr=__ERROR_LOGS)
    proc.communicate()
    proc.close()

else:
    print(f"{RED}[-] Platform not Supported :({CLEAR}")

# Closing opened files...
__ERROR_LOGS.close()

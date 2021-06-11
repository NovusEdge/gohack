#!/usr/bin/env python

# For invoking system commands:
import subprocess
import shlex

# For timestamps:
import time

# For info on the system configuration:
import platform

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


PATH = pathlib.Path(__file__).parent().absolute()
os.chdir(PATH)


_os = platform.system()
if _os == "Windows":
    command = shlex.split("powershell .config/setup_windows.ps")
    subprocess.run(command)

elif _os in ["Linux"]:
    pass
else:
    pass


#go get
#cd bin/
#go clean
#go build ../commands/*




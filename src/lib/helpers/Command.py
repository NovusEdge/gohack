#!/usr/bin/env python

#############################################
## For some reason, the os/exec package
## fails to execute the binaries.
##
## Even if it does execute them, there's
## no output, so using this as a workaround.
#############################################

# For invoking system commands:
import subprocess
import sys

if __name__ == '__main__':
    args = sys.argv[1:]
    subprocess.run(' '.join(args), shell=True)

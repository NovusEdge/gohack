#!/usr/bin/env python

# For invoking system commands:
import subprocess

# For system utilities
import pathlib
import os

from colorama import Fore, Style

# Colors!                       # Color codes:
RED     = Fore.RED              # Red          0;31     Light Red     1;31
YELLOW  = Fore.YELLOW           # Brown/Orange 0;33     Yellow        1;33
GREEN   = Fore.GREEN            # Green        0;32     Light Green   1;32
MAGENTA = Fore.MAGENTA          # Purple       0;35     Light Purple  1;35
CYAN    = Fore.CYAN             # Cyan         0;36     Light Cyan    1;36
WHITE   = Fore.WHITE            # Light Gray   0;37     White         1;37
CLEAR   = Style.RESET_ALL
                                # Use: $'\033[<code>'

# TODO: add a check for presence of go and python installations on the system
# System Environment:
# ENV = os.environ

# Checking go version:
# proc = subprocess.run("go version", shell=True, capture_output=True)
# out = str(proc.stdout)
#
# GOVERSION = out.split(' ')[2][2:]

PATH = pathlib.Path(__file__).parent.parent.absolute()
os.chdir(PATH)

# Installing dependencies
print(f"\n{YELLOW}[*] Installing golang dependencies ...{CLEAR}")
command = '''go clean
go get
'''
os.chdir("src")
subprocess.run(command, shell=True)
print(f"{CYAN}[~] Done!{CLEAR}"); os.chdir(PATH)


# Removing all binaries in tool_bin
print(f"\n{YELLOW}[*] Cleaning old binaries ...{CLEAR}")
command = '''
rm src/tool_bin/*
rm src/bin/*
'''
subprocess.run(command, shell=True, capture_output=True)
print(f"{CYAN}[~] Done!{CLEAR}")


# Building all tools/commands in tool_bin as binaries.
print(f"\n{YELLOW}[*] Building binaries for tools ...")
TOOLS = os.listdir(f"{PATH}/src/commands")
os.chdir("src/tool_bin")
for tool in TOOLS:
    command = f"go build ../commands/{tool}"
    subprocess.run(command, shell=True)
print(f"{CYAN}[~] Done!{CLEAR}"); os.chdir(PATH)


# Building the main binary:
print(f"\n{YELLOW}[*] Building the main binary ...")
os.chdir("src/bin")
command = "go build ../gohack.go"
subprocess.run(command, shell=True)
print(f"{CYAN}[~] Done!{CLEAR}\n"); os.chdir(PATH)

print("\033[1;30m Setting Gohack Environment...\033[0m")
home = os.environ["HOME"]


if not pathlib.Path(f"{home}/.config").exists():
    os.mkdir(f"{home}/.config")

with open(f"{home}/.config/gohack", "w+") as f:
    f.write(f"GOHACKPATH={PATH}\n")
    f.write(f"BINARIES={PATH}/src/bin\n")
    f.write(f"TOOLBINARIES={PATH}/src/tool_bin")

print("\033[36m[~] Done!\033[0m\n")


print(f"{CYAN}[~] To check for any errors during setting up, please check logs...{CLEAR}")
print(f"{YELLOW}[*] You can find the logs at: {PATH}/logs/{CLEAR}\n")

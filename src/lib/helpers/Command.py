#!/usr/bin/env python

# For invoking system commands:
import subprocess
import shlex

# For system utilities
import pathlib
import sys
import os

# For reading env.yaml
import yaml

from colorama import Fore, Style

# Colors!
RED     = Fore.RED
YELLOW  = Fore.YELLOW
GREEN   = Fore.GREEN
MAGENTA = Fore.MAGENTA
CYAN    = Fore.CYAN
WHITE   = Fore.WHITE
CLEAR   = Style.RESET_ALL



class CommandTemplate:
    def __init__(self, binary_name, aliases, is_functional):
        self.binary_name = binary_name
        self.aliases = aliases
        self.is_functional = is_functional


COMMANDS = {
    "portScanner": CommandTemplate(
        binary_name="portScanner",
        aliases=["ps", "pscanner", "PORTSCANNER", "portscanner", "PortScanner"],
        is_functional=True,
    ),
    "bannerGrabber": CommandTemplate(
        binary_name="bannerGrabber",
        aliases=["bg", "bgrabber", "BANNERGRABBER", "bannergrabber", "BannerGrabber"],
        is_functional=True,
    ),
}


class Command:
    def __init__(self, template, args):
        self.template = template
        self.args = args

    def run_command(self):
        if not (self.__check_template() or self.template.is_functional):
            raise Exception(f"{RED}[-] E: Incorrect args/command-template {CLEAR}")

        bin_path = self.__get_env()["TOOLBINARIES"]
        command = f"{bin_path}/{self.template.binary_name} {self.__make_args_string()}"
        subprocess.run(command, shell=True)


    def __make_args_string(self):
        if "-h" in self.args:
            return "-h"

        return ' -'.join(self.args)


    def __check_template(self):
        return self.template in list(COMMANDS.keys())

    def __get_env(self):
        path = self.__get_project_path()
        efile = open(f"{path}/.config/env.yaml", "r")
        return yaml.load(efile, Loader=yaml.FullLoader)

    def __get_project_path(self):
        PATH = pathlib.Path(__file__).parent.parent.parent.parent.absolute()
        return PATH

if __name__ == '__main__':
    temp = COMMANDS["bannerGrabber"]
    c = Command(template=temp, args=["-url=\"https://github.com/\""])
    c.run_command()

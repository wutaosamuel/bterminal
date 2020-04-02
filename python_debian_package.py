# Aims:
# 1. build .deb in terms of architecture

# requirement:
# TODO: deb package dir tree

import os
import sys
import glob
import hashlib
import argparse
import platform

class python_debian_package:

  def __init__(self, deb_dir=None):
    self.deb_dir = deb_dir
    self.debian = None
    self.name=None
    self.system = None
    self.architecture = self.detectArch()
    self.version = "1.0.0"
    self.flag = None
    self.make = None

  def exec_arg(self):
    # TODO: provide go build command
    parser = argparse.ArgumentParser()
    parser.add_argument("-i", "--input",
                        help="parent directory for making dep package",
                        type=str, nargs="?", required=True)
    parser.add_argument("-n", "--name",
                        help="name of debian package",
                        type=str, nargs="?", required=True)
    parser.add_argument("-a", "--architecture",
                        help="target architecture for go building",
                        type=str, nargs="+", required=True)
    parser.add_argument("-s", "--system",
                        help="target system for go building",
                        type=str, nargs="+", required=False)
    parser.add_argument("-m", "--make",
                        help="make command for build program, Makefile required",
                        type=str, nargs="+", required=False)
    parser.add_argument("-V", "--Version",
                        help="deb package version",
                        type=str, nargs="+", required=False)
    parser.add_argument("-f", "--flag",
                        help="flag for go build",
                        type=str, nargs="+", required=False)
    
    args = parser.parse_args()
    self.deb_dir = args.input
    self.name = args.name
    self.architecture = args.architecture
    if args.Version != None:
      self.version = args.Version
    if args.system != None:
      self.system = args.system
    if args.flag != None:
      self.flag = args.flag
    if args.Version != None:
      self.version = args.Version
    if args.make != None:
      self.make = args.make
    
    # check all parameter
    if not os.path.isdir(self.deb_dir):
      parser.print_help()
      sys.exit(1)
    if not self.system == None:
      if len(self.system) != len(self.architecture):
        print("No. of system should same as No. of architectures")
        parser.print_help()
        sys.exit(1)
    if not self.make == None:
      if len(self.make) != len(self.architecture):
        print("No. of make command should same as No. of architectures")
        parser.print_help()
        sys.exit(1)
    
    # set extra value
    self.system = self.detectSys()
  
  def detectSys(self):
    #sys = platform.system()
    #if sys == "Linux":
      #return "linux"
    #else:
      #print(sys+" is not support")
      #print("Linux system only")
      #self.print_help()
    return platform.system()

  def detectArch(self):
    arch = platform.machine()
    # TODO: support more machine
    if arch == "i386":
      return "386"
    if arch == "x86_64":
      return "amd64"
    if arch == "AMD64":
      return "amd64"
    if arch == "PowerPC":
      return "ppc"
    else:
      print(arch+" is not support")
      self.print_help()
  
  def replace_control(self, version=None, architecture=None):
    if version == None or architecture == None:
      print("Please enter version or architecture")
      self.print_help()
    deb = os.path.join(self.deb_dir, "debian")
    Deb = os.path.join(self.deb_dir, "DEBIAN")
    if os.path.exists(deb):
      self.debian = deb
    if os.path.exists(Deb):
      self.debian = Deb
    if self.debian == None:
      print("no such "+deb+" or "+Deb+" directory")
      self.print_help()
    control = os.path.join(self.debian, "control")
    if not os.path.exists(control):
      print("no such "+control+" file")
    # replace version and architecture
    file_data = []
    with open(control, "r") as f:
      for line in f:
        if "Version" in line:
          line = "Version:"+version+"\n"
        if "Architecture" in line:
          line = "Architecture:"+architecture+"\n"
        file_data.append(line)
    with open(control, "w") as f:
      f.writelines(file_data)
  
  def md5sum(self):
    usr = os.path.join(self.deb_dir, "usr")
    files = []
    md5s = []
    # recode files and each md5 sum
    for f in sorted(
      glob.glob(os.path.join(usr, "**"), recursive=True)):
      if os.path.isfile(f):
        with open(f, "rb") as content:
          #md5 = hashlib.md5(str(content).encode("utf-8")).hexdigest()
          file_hash = hashlib.md5()
          while True:
            chunk = content.read(8192)
            if not chunk:
              break
            file_hash.update(chunk)
          md5 = file_hash.hexdigest()
        # reformat: ./usr/*
        filesplit = os.path.normpath(f).split(os.path.sep)
        isUsr = False
        filepath = "."
        for index in range(len(filesplit)):
          if filesplit[index] == "usr":
            isUsr = True
          if isUsr:
            filepath += "/"+filesplit[index]
        files.append(filepath)
        md5s.append(md5)
    # write it into md5sum
    # format of md5sums:
    # 3q2...  ./usr/**
    md5file = os.path.join(self.debian, "md5sums")
    md5data = []
    for line in range(len(files)):
      md5data.append(md5s[line] + "  " + files[line] + "\n")
    with open(md5file, "w") as md5f:
      md5f.writelines(md5data)
  
  def exec_make(self):
    for index in range(len(self.make)):
      # replace version and architecture in control file
      self.replace_control(self.version, self.architecture[index])
      # make binary exectution
      os.system(self.make[index])
      # calculate md5sums of ./usr/*
      self.md5sum()
      # build debian package
      debname = self.name+"_"+self.version+"_"+self.architecture[index]+".deb"
      os.system("dpkg -b "+self.deb_dir+" "+debname)

  def exec_build(self):
    pass

  def exec_main(self):
    self.exec_arg()
    self.exec_make()
  
  ################## HELP ###################

  def print_more_help(self):
    print("python_debian_package.py \'-h\' | \'--help\' for more helps")

  def usage(self):
    print("python_debian_package.py [-i <inputfolder>] [option]")
    print("or")
    self.print_more_help()

  def print_help(self):
    self.usage()
    print()
    sys.exit(0)

if __name__ == "__main__":
  deb = python_debian_package()
  deb.exec_main()
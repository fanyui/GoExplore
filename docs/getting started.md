---
layout: page
title: Getting Started
permalink: /get-started/
nav_order: 2
---

## Getting started

Learning a new language for the first time is no easy feat! However, keep in mind that every Go developer, no matter how skilled, was once in your shoes: learning the syntax and inner workings of the language for the first time.


Be sure to make use of all the resources available to you, and maintain a constant rhythm. Throughout this book, you'll have many opportunities to practice the Go skills you'll pick up along the way (especially those that you'll need for your entire go career!). By leveraging your creativity and problem-solving skills, I know you'll find development with Go rewarding. Good luck!

## Step-by-Step Guide for Installing Go

Getting Go installed on your local machine is straightforward. If you're using Linux, Mac, or Windows, you can download and execute an installer to get up and running immediately.

### Installing Go on Windows

To install Go on Windows, follow these steps:
1. Download the Windows installer from the official Go website: https://golang.org/dl/
2. Double-click on the installer to start the installation process.
3. Follow the prompts to complete the installation. Choose the default options unless you have a specific reason to change them.
4. Once the installation is complete, open the Command Prompt or PowerShell and type go version to verify that Go is installed correctly.

### Installing Go on Mac
To install Go on a Mac, follow these steps:
1. Download the macOS installer from the official Go website: https://golang.org/dl/
2. Double-click on the installer to start the installation process.
3. Follow the prompts to complete the installation. Choose the default options unless you have a specific reason to change them.
4. Once the installation is complete, open the Terminal app and type go version to verify that Go is installed correctly.


### Installing Go on Linux
To install Go on Linux, follow these steps:
1. Open a terminal window.
2. Update the package list by typing `sudo apt update` (for Debian/Ubuntu) or `sudo yum update` (for CentOS/Fedora).
3. Install the Go package by typing `sudo apt install golang`(for Debian/Ubuntu) or `sudo yum install golang` (for CentOS/Fedora).
4. Once the installation is complete, type `go version` to verify that Go is installed correctly.

If you still have issues after following the instructions above, you can checkout the official website [Download and Install](https://go.dev/doc/install)

## Setting Up Environment Variables
After the installation steps covered above,it's time for us to set up environment variables necessary for running go programs and using go tools.

### Setting Up Environment Variables on Windows
1. Open the Control Panel and go to System and Security > System > Advanced system settings.
2. Click on the Environment Variables button.
3. Under System Variables, click on New.
4. Enter the following information:
   - Variable name: `GOROOT`
   - Variable value: `C:\Go` (or the directory where Go is installed)
5. Click on OK to save the variable.
6. Under System Variables, click on New again.
7. Enter the following information:
   - Variable name: `GOPATH`
   - Variable value:`%USERPROFILE%\go` (or the directory where you want to store your Go projects)
8. Click on OK to save the variable

### Setting Up Environment Variables on Mac/Linux
Here the steps are both same for Mac and linux

1. Open the Terminal app.
2. Type `nano ~/.bash_profile` to open the Bash profile file.
3. Add the following lines to the file:
   `export GOROOT=/usr/local/go`
   `export GOPATH=$HOME/go`
   `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`
4. Press `Ctrl+X`, then `Y`, then Enter to save the file and exit Nano.
5. Type source `~/.bash_profile` to reload the Bash profile.


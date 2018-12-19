# nscommand
A small tool to control multiple machines via ssh. Written in Go

# background
I was learning the GO language and wanted to make a useful app in the process.
I also happened to be playing around with AWS to learn how that works as well. Then I realized there were no good tools to maintain multiple computers via shell. Even some conmercial tools have their limitations: not cross platform, need x-server to run, cannot easily switch the control set.
So I wrote this tool for fun. I hope it will be helpful to others too.

# design features
1. Cross platform, i.e. can be run at least on Windows, Linux, Mac, and FreeBSD
2. No x-server needed, since it is slow and hard to setup across different platforms (There is a tool called terminator, which is very good, but it needs x-server and only exists in Linux.)
3. Easy to setup (Some linux tools need lots of dependencies and hard to install.)
4. Easy to switch control set, i.e. I may want to send some commands to remote host a,b,c,d, and send other commands to remote host a,c,f. I would like to be able to easily switch between two sets of remote hosts even when they have overlaps.


# milestones
1. implement the host config based on tree structure, so that switching between controlled sets becomes easy.
2. can send commands to a set of remote computers (all hosts in a branch) and get feedback. Works like a CLI tool
3. can interactively switch control set and works like a multiple windows shell. windows are generated based on the inputed start node.
4. can switch controlled set dynamically and regenerate windows
5. can enhance the focus windows (change size, etc.) so that it can show more information without scrolling
6. start to support corner cases, like using ssh agents, jumping through proxy hosts, etc.

# what makes this tool different from other tools
The best part of the design is using a tree structure to store the configuration. By doing this, we can put hosts with the same role in the same branch; similar branches under bigger branches so that we can control all similar braches together. By sending a command to different branch roots, we can control all hosts that belong directly or indirectly to this branch. This abstraction covers almost all realistic use cases. 

# curent release
The current release is just past the second milestone. It is already a very useful tool. More work will be done soon!

# usage
.\nscommand.exe --help
Usage of nscommand.exe:

  -cmd string
  
    command send to remote hosts (default "ls")
        
  -conf string
  
    name of the configuration (default "config.json")
        
  -p
  
    print the forest in the configuration file
  
  -r
  
    run the command on remote hosts
  
  -root string
  
    root node positionID (default "1")

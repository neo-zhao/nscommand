# nscommand
A small tool to control multiple machines via ssh. Written in Go

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
        

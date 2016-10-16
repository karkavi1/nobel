
GO Language Examples:
====================
   
    Install The Go Tools:
    
        The following url has binary distributions for all operating systems:  https://golang.org/dl/
        
        After downloading GO installer / distributions, follow the installation instructions mentioned below:


        Mac OS:
        =======

           Open your terminal and run the following commands

                Setting your golang "bin" path in PATH environment variable by issuing this command:  

                > by default, it will be installed in '/user/local/go'. if you want to confirm the installation location, issue command: which go (it gives the path where go has installed in your system) and copy the location and append '/bin' directory 

             	> export PATH=$PATH:/usr/local/go/bin


       Prerequisite: 
       =============

         1. set the GOPATH environment variable to point to your workspace location
               
                export GOPATH=<workspacelocation> >>  ${HOME}/.bash_profile

         2. Gorilla/mux package for URL router and dispatcher. Command to install Gorilla/mux package: 

          		go get -u 'github.com/gorilla/mux'


         Executing Go Code:
         ==================

                command to run go code: 

                  go run $GOPATH/src/main.go 

                  you can see port 8081 is listening for incoming requests.

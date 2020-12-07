# Jobber

Jobber transform a simple command to a Job. 

To become called Job a command must have :
- an history of last run, indicating time of run, duration and status : success or failure
- for previous run : stdout, stderr must be logged
- must be abble to start a command to process result (status, stdout, stderr) (those command are called sinker)


Jobber is basicly used with cron, which allow you receive an email when your command failed and see history of previous run. 

## Install 

```bash
go get https://github.com/canadadry/jobber
```

## Usage

To start a job:

```bash
# with no command on failure or on succes
jobber -j jobName1
# with no command on failure and one on succes
jobber -j jobName2 -s SinkerName1
# with no command on succes and one on failure
jobber -j jobName3 -f SinkerName2
# with one command on succes and one on failure
jobber -j jobName4 -s SinkerName3 -f SinkerName4
```

### Configuration

Just add `.sh` in `.jobber/job` for job and in  `.jobber/sinker`
`.sh` files must be executable

**sinker command must never failed**

    └── ~/.jobber                                 
        ├── job                           
        │   ├── JobName1.sh               
        │   ├── JobName2.sh               
        │   ├── JobName3.sh               
        │   └── JobName4.sh                       
        ├── sinker                           
        │   ├── SinkerName1.sh               
        │   ├── SinkerName1.sh               
        │   ├── SinkerName1.sh               
        │   └── SinkerName1.sh                      
        └── log                              
            ├── history.log                  
            ├── JobName1.log               
            ├── JobName2.log               
            ├── JobName3.log               
            └── JobName4.log                        

### Sinker Example 

Jobber will call sinker command with two arguments : a unique identifier and the job stdout and stderr

If you want to receive an email on a job failure you can try this : 
```bash 
#!/usr/bin/env bash

echo -e "Notification for" $1 "\n" $2  | mail -s "Jobber Notification" -aFrom:server@example.com admin@example.com
```

# Jobber

Jobber transform a simple command to a Job. 

To become called Job a command must have :
- an history of last run, indicating time of run, duration and status : success or failure
- for previous run : stdout, stderr must be logged
- must be abble to start a command to process result (status, stdout, stderr) (those command are called sinker)
Jobber is basicly used with cron, which allow you receive an email when your command failed and see history of previous run. 


## Install 

```bash
go get https://github.com/everycheck/jobber
```

## Usage

To start a job:

```bash
# with no command on failure or on succes
jobber jobName1
# with no command on failure and one on succes
jobber jobName2 -succes SinkerName1
# with no command on succes and one on failure
jobber jobName3 -failure SinkerName2
# with one command on succes and one on failure
jobber jobName4 -succes SinkerName3 -failure SinkerName4
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



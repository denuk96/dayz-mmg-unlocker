## DAYZ MMG LOCKER OPENER(brute-force )
Did you forget your pass to your MMG locker in DayZ? Don't worry! This simple script will help you open the locker!(Dont use it for cheating)

### Usage
1. Open dayz and join the server
2. Head to your door or crate with MMG locker
3. Open a password input window
4. Alt+Tab to the location of the compiled script `dayz_mmg_unlocker.exe`
5. Execute it and return to your dayZ(you have 5 second delay for this)
6. Wait until it unlocks

#### Optional
By default start from 0000, you may change this behavior by: 
```cmd
setx DAYZ_OPENER_START "4444"
```

### Build for windows
```bash
export GOOS=windows
export GOARCH=amd64
go build -o dayz_mmg_unlocker.exe main.go
file dayz_mmg_unlocker.exe
```
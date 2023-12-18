package main

import (
	"errors"
	"fmt"
	"strconv"
	"os"
	"bufio"
)

type RiceCooker struct {
	isOn           bool
	isCooking      bool
	isPlugged      bool
	isKeepWarm     bool
	isSteamCooking bool
}

func (rc *RiceCooker) TurnOn() error {
	if rc.isPlugged {
		rc.isOn = true
		fmt.Println("Rice cooker is now on.")
		return nil
	}
	return errors.New("Rice cooker is not plugged in. Cannot turn on.")
}

func (rc *RiceCooker) TurnOff() {
	if rc.isOn {
		rc.isOn = false
		rc.isCooking = false
		rc.isKeepWarm = false
		rc.isSteamCooking = false
		fmt.Println("Rice cooker is now off.")
	} else {
		fmt.Println("Rice cooker is already off.")
	}
}

func (rc *RiceCooker) CookRice() error {
	if rc.isOn {
		rc.isCooking = true
		fmt.Println("Cooking rice...")
		return nil
	}
	return errors.New("Cannot cook rice. Rice cooker is off.")
}

func (rc *RiceCooker) StopCooking() {
	if rc.isCooking {
		rc.isCooking = false
		rc.isKeepWarm = true
		fmt.Println("Rice cooking process stopped. Keep Warm mode activated.")
	} else {
		fmt.Println("No rice cooking process to stop.")
	}
}

func (rc *RiceCooker) KeepWarm() error {
	if rc.isCooking {
		return errors.New("Cannot activate Keep Warm while cooking.")
	} else if rc.isOn {
		rc.isKeepWarm = true
		fmt.Println("Keep Warm mode activated.")
		return nil
	}
	return errors.New("Rice cooker is off. Cannot activate Keep Warm.")
}

func (rc *RiceCooker) SteamCooking() error {
	if rc.isOn {
		rc.isSteamCooking = true
		fmt.Println("Steam cooking activated.")
		return nil
	}
	return errors.New("Rice cooker is off. Cannot activate Steam Cooking.")
}

func (rc *RiceCooker) StopSteamCooking() {
	if rc.isSteamCooking {
		rc.isSteamCooking = false
		fmt.Println("Steam cooking process stopped.")
	} else {
		fmt.Println("No steam cooking process to stop.")
	}
}

func (rc *RiceCooker) CheckActivity() {
	if rc.isOn {
		if rc.isCooking {
			fmt.Println("Rice cooking process is currently running.")
		} else if rc.isKeepWarm {
			fmt.Println("Keep Warm mode is currently active.")
		} else if rc.isSteamCooking {
			fmt.Println("Steam cooking process is currently running.")
		} else {
			fmt.Println("No activity is currently running.")
		}
	} else {
		fmt.Println("Rice cooker is off. No activity is currently running.")
	}
}

func (rc *RiceCooker) PlugIn() {
	if rc.isPlugged {
		fmt.Println("Rice cooker is already plugged in.")
	} else {
		rc.isPlugged = true
		fmt.Println("Rice cooker is plugged in.")
	}
}

func (rc *RiceCooker) PlugOut() {
	rc.isPlugged = false
	rc.isOn = false
	rc.isCooking = false
	rc.isKeepWarm = false
	rc.isSteamCooking = false
	fmt.Println("Rice cooker is unplugged and turned off.")
}

func main() {
	riceCooker := RiceCooker{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Rice Cooker Console App ---")
		fmt.Println("1. Turn On")
		fmt.Println("2. Turn Off")
		fmt.Println("3. Cook Rice")
		fmt.Println("4. Stop Cooking")
		fmt.Println("5. Keep Warm")
		fmt.Println("6. Steam Cooking")
		fmt.Println("7. Stop Steam Cooking")
		fmt.Println("8. Check Activity")
		fmt.Println("9. Plug In")
		fmt.Println("10. Plug Out")
		fmt.Println("11. Exit")
		fmt.Print("Choose an option (1-11): ")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			err := riceCooker.TurnOn()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			riceCooker.TurnOff()
		case 3:
			err := riceCooker.CookRice()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			riceCooker.StopCooking()
		case 5:
			err := riceCooker.KeepWarm()
			if err != nil {
				fmt.Println(err)
			}
		case 6:
			err := riceCooker.SteamCooking()
			if err != nil {
				fmt.Println(err)
			}
		case 7:
			riceCooker.StopSteamCooking()
		case 8:
			riceCooker.CheckActivity()
		case 9:
			riceCooker.PlugIn()
		case 10:
			riceCooker.PlugOut()
		case 11:
			fmt.Println("Exiting Rice Cooker App. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 11.")
		}
	}
}

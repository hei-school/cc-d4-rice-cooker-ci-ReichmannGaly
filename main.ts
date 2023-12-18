import * as readline from 'readline'

export class RiceCooker {
  isOn: boolean = false
  isCooking: boolean = false
  isKeepWarm: boolean = false
  isSteamCooking: boolean = false
  isPlugged: boolean = false

  turnOn (): void {
    if (this.isPlugged) {
      this.isOn = true
      console.log('Rice cooker is now on.')
    } else {
      throw new Error('Rice cooker is not plugged in. Cannot turn on.')
    }
  }

  turnOff (): void {
    if (this.isOn) {
      this.isOn = false
      this.isCooking = false
      this.isKeepWarm = false
      this.isSteamCooking = false
      console.log('Rice cooker is now off.')
    } else {
      console.log('Rice cooker is already off.')
    }
  }

  
  cookRice (): void {
    if (this.isOn) {
      this.isCooking = true
      console.log('Cooking rice...')
    } else {
      throw new Error('Cannot cook rice. Rice cooker is off.')
    }
  }

  stopCooking (): void {
    if (this.isCooking) {
      this.isCooking = false
      this.isKeepWarm = true
      console.log('Rice cooking process stopped. Keep Warm mode activated.')
    } else {
      console.log('No rice cooking process to stop.')
    }
  }

  keepWarm (): void {
    if (this.isCooking) {
      throw new Error('Cannot activate Keep Warm while cooking.')
    } else if (this.isOn) {
      this.isKeepWarm = true
      console.log('Keep Warm mode activated.')
    } else {
      throw new Error('Rice cooker is off. Cannot activate Keep Warm.')
    }
  }

  steamCooking (): void {
    if (this.isOn) {
      this.isSteamCooking = true
      console.log('Steam cooking activated.')
    } else {
      throw new Error('Rice cooker is off. Cannot activate Steam Cooking.')
    }
  }

  stopSteamCooking (): void {
    if (this.isSteamCooking) {
      this.isSteamCooking = false
      console.log('Steam cooking process stopped.')
    } else {
      console.log('No steam cooking process to stop.')
    }
  }

  checkActivity (): void {
    if (this.isOn) {
      if (this.isCooking) {
        console.log('Rice cooking process is currently running.')
      } else if (this.isKeepWarm) {
        console.log('Keep Warm mode is currently active.')
      } else if (this.isSteamCooking) {
        console.log('Steam cooking process is currently running.')
      } else {
        console.log('No activity is currently running.')
      }
    } else {
      console.log('Rice cooker is off. No activity is currently running.')
    }
  }

  plugIn (): void {
    if (this.isPlugged) {
      console.log('Rice cooker is already plugged in.')
    } else {
      this.isPlugged = true
      console.log('Rice cooker is plugged in.')
    }
  }

  plugOut (): void {
    this.isPlugged = false
    this.isOn = false
    this.isCooking = false
    this.isKeepWarm = false
    this.isSteamCooking = false
    console.log('Rice cooker is unplugged and turned off.')
  }
}

function main (): void {
  const riceCooker = new RiceCooker()

  const showMenu = (): void => {
    console.log('\n--- Rice Cooker Console App ---')
    console.log('1. Turn On')
    console.log('2. Turn Off')
    console.log('3. Cook Rice')
    console.log('4. Stop Cooking')
    console.log('5. Keep Warm')
    console.log('6. Steam Cooking')
    console.log('7. Stop Steam Cooking')
    console.log('8. Check Activity')
    console.log('9. Plug In')
    console.log('10. Plug Out')
    console.log('11. Exit')
    console.log('Choose an option (1-11): ')
  }

  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  })

  const processUserInput = (input: string): void => {
    const choice: number = parseInt(input)
    try {
      switch (choice) {
        case 1:
          riceCooker.turnOn()
          break
        case 2:
          riceCooker.turnOff()
          break
        case 3:
          riceCooker.cookRice()
          break
        case 4:
          riceCooker.stopCooking()
          break
        case 5:
          riceCooker.keepWarm()
          break
        case 6:
          riceCooker.steamCooking()
          break
        case 7:
          riceCooker.stopSteamCooking()
          break
        case 8:
          riceCooker.checkActivity()
          break
        case 9:
          riceCooker.plugIn()
          break
        case 10:
          riceCooker.plugOut()
          break
        case 11:
          console.log('Exiting Rice Cooker App. Goodbye!')
          rl.close()
          break
        default:
          console.log('Invalid choice. Please enter a number between 1 and 11.')
      }
    } catch (error) {
      if (error instanceof Error) {
        console.error(error.message)
      } else {
        console.error('An unknown error occurred')
      }
    } finally {
      showMenu()
    }
  }

  rl.on('line', processUserInput)

  showMenu()

  rl.on('close', () => {
    process.exit(0)
  })
}

main()

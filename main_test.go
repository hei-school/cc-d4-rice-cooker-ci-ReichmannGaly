package main

import (
	"testing"
)

func TestRiceCooker(t *testing.T) {
	var riceCooker RiceCooker

	t.Run("StartCookingAndStop", func(t *testing.T) {
		riceCooker.PlugIn()
		riceCooker.TurnOn()
		riceCooker.CookRice()
		if !riceCooker.isCooking {
			t.Error("Expected rice cooker to be cooking")
		}
		riceCooker.StopCooking()
		if riceCooker.isCooking || !riceCooker.isKeepWarm {
			t.Error("Expected rice cooker to stop cooking and enter Keep Warm mode")
		}
	})

	t.Run("CannotActivateKeepWarmWhileCooking", func(t *testing.T) {
		riceCooker.PlugIn()
		riceCooker.TurnOn()
		riceCooker.CookRice()
		err := riceCooker.KeepWarm()
		if err == nil || err.Error() != "Cannot activate Keep Warm while cooking." {
			t.Error("Expected error when trying to activate Keep Warm while cooking")
		}
	})

	t.Run("CannotActivateKeepWarmWhenTurnedOff", func(t *testing.T) {
		// Reset riceCooker state
		riceCooker = RiceCooker{}
		err := riceCooker.KeepWarm()
		if err == nil || err.Error() != "Rice cooker is off. Cannot activate Keep Warm." {
			t.Error("Expected error when trying to activate Keep Warm when turned off")
		}
	})

	t.Run("CannotActivateSteamCookingWhenTurnedOff", func(t *testing.T) {
		// Reset riceCooker state
		riceCooker = RiceCooker{}
		err := riceCooker.SteamCooking()
		if err == nil || err.Error() != "Rice cooker is off. Cannot activate Steam Cooking." {
			t.Error("Expected error when trying to activate Steam Cooking when turned off")
		}
	})

	t.Run("StartSteamCookingAndStop", func(t *testing.T) {
		riceCooker.PlugIn()
		riceCooker.TurnOn()
		riceCooker.SteamCooking()
		if !riceCooker.isSteamCooking {
			t.Error("Expected rice cooker to be steam cooking")
		}
		riceCooker.StopSteamCooking()
		if riceCooker.isSteamCooking {
			t.Error("Expected rice cooker to stop steam cooking")
		}
	})

	t.Run("UnplugAndTurnOff", func(t *testing.T) {
		riceCooker.PlugIn()
		riceCooker.TurnOn()
		riceCooker.PlugOut()
		if riceCooker.isPlugged || riceCooker.isOn || riceCooker.isCooking || riceCooker.isKeepWarm || riceCooker.isSteamCooking {
			t.Error("Expected rice cooker to be unplugged and turned off")
		}
	})
}

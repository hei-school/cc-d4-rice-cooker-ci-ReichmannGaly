import { RiceCooker } from '../main';
import { describe, expect, test, beforeEach, afterEach } from '@jest/globals';

describe('RiceCooker', () => {
  let riceCooker: RiceCooker;

  beforeEach(() => {
    riceCooker = new RiceCooker();
  });

  afterEach(() => {
    riceCooker = null as any;
  });

  test('should start cooking rice and then stop cooking', () => {
    riceCooker.plugIn();
    riceCooker.turnOn();
    riceCooker.cookRice();
    expect(riceCooker.isCooking).toBe(true);
    riceCooker.stopCooking();
    expect(riceCooker.isCooking).toBe(false);
    expect(riceCooker.isKeepWarm).toBe(true);
  });

  test('should not be able to activate Keep Warm while cooking', () => {
    riceCooker.plugIn();
    riceCooker.turnOn();
    riceCooker.cookRice();
    expect(() => riceCooker.keepWarm()).toThrow('Cannot activate Keep Warm while cooking.');
  });

  test('should not be able to activate Keep Warm when turned off', () => {
    expect(() => riceCooker.keepWarm()).toThrow('Rice cooker is off. Cannot activate Keep Warm.');
  });

  test('should not be able to activate Steam Cooking when turned off', () => {
    expect(() => riceCooker.steamCooking()).toThrow('Rice cooker is off. Cannot activate Steam Cooking.');
  });

  test('should start steam cooking and then stop steam cooking', () => {
    riceCooker.plugIn();
    riceCooker.turnOn();
    riceCooker.steamCooking();
    expect(riceCooker.isSteamCooking).toBe(true);
    riceCooker.stopSteamCooking();
    expect(riceCooker.isSteamCooking).toBe(false);
  });

  test('should unplug and turn off the rice cooker', () => {
    riceCooker.plugIn();
    riceCooker.turnOn();
    riceCooker.plugOut();
    expect(riceCooker.isPlugged).toBe(false);
    expect(riceCooker.isOn).toBe(false);
    expect(riceCooker.isCooking).toBe(false);
    expect(riceCooker.isKeepWarm).toBe(false);
    expect(riceCooker.isSteamCooking).toBe(false);
  });
});

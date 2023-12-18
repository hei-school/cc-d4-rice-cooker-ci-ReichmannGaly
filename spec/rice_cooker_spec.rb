require_relative '../main'

RSpec.describe RiceCooker do
  let(:rice_cooker) { RiceCooker.new }

  before(:each) do
    rice_cooker.instance_variable_set(:@is_plugged, false)
    rice_cooker.instance_variable_set(:@is_on, false)
    rice_cooker.instance_variable_set(:@is_cooking, false)
    rice_cooker.instance_variable_set(:@is_keep_warm, false)
    rice_cooker.instance_variable_set(:@is_steam_cooking, false)
  end

  describe '#start_cooking_and_stop' do
    it 'should start cooking rice and then stop cooking' do
      rice_cooker.plug_in
      rice_cooker.turn_on
      rice_cooker.cook_rice
      expect(rice_cooker.is_cooking).to be(true)
      rice_cooker.stop_cooking
      expect(rice_cooker.is_cooking).to be(false)
      expect(rice_cooker.is_keep_warm).to be(true)
    end
  end

  describe '#cannot_activate_keep_warm_while_cooking' do
    it 'should not be able to activate Keep Warm while cooking' do
      rice_cooker.plug_in
      rice_cooker.turn_on
      rice_cooker.cook_rice
      expect { rice_cooker.keep_warm }.to raise_error('Cannot activate Keep Warm while cooking.')
    end
  end

  describe '#cannot_activate_keep_warm_when_turned_off' do
    it 'should not be able to activate Keep Warm when turned off' do
      expect { rice_cooker.keep_warm }.to raise_error('Rice cooker is off. Cannot activate Keep Warm.')
    end
  end

  describe '#cannot_activate_steam_cooking_when_turned_off' do
    it 'should not be able to activate Steam Cooking when turned off' do
      expect { rice_cooker.steam_cooking }.to raise_error('Rice cooker is off. Cannot activate Steam Cooking.')
    end
  end

  describe '#start_steam_cooking_and_stop' do
    it 'should start steam cooking and then stop steam cooking' do
      rice_cooker.plug_in
      rice_cooker.turn_on
      rice_cooker.steam_cooking
      expect(rice_cooker.is_steam_cooking).to be(true)
      rice_cooker.stop_steam_cooking
      expect(rice_cooker.is_steam_cooking).to be(false)
    end
  end

  describe '#plug_out_and_turn_off' do
    it 'should unplug and turn off the rice cooker' do
      rice_cooker.plug_in
      rice_cooker.turn_on
      rice_cooker.plug_out
      expect(rice_cooker.is_plugged).to be(false)
      expect(rice_cooker.is_on).to be(false)
      expect(rice_cooker.is_cooking).to be(false)
      expect(rice_cooker.is_keep_warm).to be(false)
      expect(rice_cooker.is_steam_cooking).to be(false)
    end
  end
end

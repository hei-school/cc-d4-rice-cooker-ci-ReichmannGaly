class RiceCooker
  attr_reader :is_on, :is_cooking, :is_plugged, :is_keep_warm, :is_steam_cooking
  
  def initialize
    @is_on = false
    @is_cooking = false
    @is_plugged = false
    @is_keep_warm = false
    @is_steam_cooking = false
  end

  def turn_on
    if @is_plugged
      @is_on = true
      puts "Rice cooker is now on."
    else
      raise "Rice cooker is not plugged in. Cannot turn on."
    end
  end

  def turn_off
    if @is_on
      @is_on = false
      @is_cooking = false
      @is_keep_warm = false
      @is_steam_cooking = false
      puts "Rice cooker is now off."
    else
      puts "Rice cooker is already off."
    end
  end

  def cook_rice
    if @is_on
      @is_cooking = true
      puts "Cooking rice..."
    else
      raise "Cannot cook rice. Rice cooker is off."
    end
  end

  def stop_cooking
    if @is_cooking
      @is_cooking = false
      @is_keep_warm = true
      puts "Rice cooking process stopped. Keep Warm mode activated."
    else
      puts "No rice cooking process to stop."
    end
  end

  def keep_warm
    if @is_cooking
      raise "Cannot activate Keep Warm while cooking."
    elsif @is_on
      @is_keep_warm = true
      puts "Keep Warm mode activated."
    else
      raise "Rice cooker is off. Cannot activate Keep Warm."
    end
  end

  def steam_cooking
    if @is_on
      @is_steam_cooking = true
      puts "Steam cooking activated."
    else
      raise "Rice cooker is off. Cannot activate Steam Cooking."
    end
  end

  def stop_steam_cooking
    if @is_steam_cooking
      @is_steam_cooking = false
      puts "Steam cooking process stopped."
    else
      puts "No steam cooking process to stop."
    end
  end

  def check_activity
    if @is_cooking
      puts "Rice cooking process is currently running."
    elsif @is_keep_warm
      puts "Keep Warm mode is currently active."
    elsif @is_steam_cooking
      puts "Steam cooking process is currently running."
    else
      puts "No activity is currently running."
    end
  end

  def plug_in
    if @is_plugged
      puts "Rice cooker is already plugged in."
    else
      @is_plugged = true
      puts "Rice cooker is plugged in."
    end
  end

  def plug_out
    @is_plugged = false
    @is_on = false
    @is_cooking = false
    @is_keep_warm = false
    @is_steam_cooking = false
    puts "Rice cooker is unplugged and turned off."
  end
end

def main
  rice_cooker = RiceCooker.new

  loop do
    puts "\n--- Rice Cooker Console App ---"
    puts "1. Turn On"
    puts "2. Turn Off"
    puts "3. Cook Rice"
    puts "4. Stop Cooking"
    puts "5. Keep Warm"
    puts "6. Steam Cooking"
    puts "7. Stop Steam Cooking"
    puts "8. Check Activity"
    puts "9. Plug In"
    puts "10. Plug Out"
    puts "11. Exit"
    print "Choose an option (1-11): "

    choice = gets.chomp.to_i

    begin
      case choice
      when 1
        rice_cooker.turn_on
      when 2
        rice_cooker.turn_off
      when 3
        rice_cooker.cook_rice
      when 4
        rice_cooker.stop_cooking
      when 5
        rice_cooker.keep_warm
      when 6
        rice_cooker.steam_cooking
      when 7
        rice_cooker.stop_steam_cooking
      when 8
        rice_cooker.check_activity
      when 9
        rice_cooker.plug_in
      when 10
        rice_cooker.plug_out
      when 11
        puts "Exiting Rice Cooker App. Goodbye!"
        break
      else
        puts "Invalid choice. Please enter a number between 1 and 11."
      end
    rescue StandardError => e
      puts e.message
    end
  end
end

main if $PROGRAM_NAME == __FILE__
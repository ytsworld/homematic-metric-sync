import sys
import time
import homematicip
from homematicip.home import Home
from homematicip.device import TemperatureHumiditySensorOutdoor

config = homematicip.find_and_load_config_file()
if config == None:
    print("Cannot find config.ini!")
    sys.exit()
    
home = Home()
home.set_auth_token(config.auth_token)
home.init(config.access_point)
    
def main():
    
    global home

    while True:
        home.get_current_state()
        for g in home.groups:
            if g.groupType=="META":
                for device in g.devices:
                    if isinstance(device, TemperatureHumiditySensorOutdoor):
                        with open('./data/' + device.label + '.log','a') as f:
                            f.write('{}, temp: {}, hum: {}\n'.format(device.lastStatusUpdate, device.actualTemperature, device.humidity))
                            f.close()
                        print(device.label, " ", device.lastStatusUpdate, ", temp: ", device.actualTemperature, ", hum: ", device.humidity)
        time.sleep(300)

if __name__ == "__main__":
    main()

import sys
import time
import homematicip
from homematicip.home import Home
from homematicip.device import TemperatureHumiditySensorOutdoor, WallMountedThermostatPro
from datetime import date
import json

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
        today = date.today().strftime("%Y%m%d")
        home.get_current_state()
        for g in home.groups:
            print(g)
            if g.groupType=="META":
                for device in g.devices:
                    print("Device: ",device.label, " Type: ", type(device))
                    if isinstance(device, TemperatureHumiditySensorOutdoor):
                        data = {
                            "lastStatusUpdate" : device.lastStatusUpdate.strftime("%Y-%m-%d %H:%M:%S"),
                            "deviceName": device.label,
                            "actualTemperature": device.actualTemperature,
                            "humidity": device.humidity
                        }
                        with open('./data/' + today + '-devices.log','a', encoding ='utf8') as f:
                            json.dump(data, f, ensure_ascii=False)
                            f.write("\n")
                            f.close()
                        print(device.label, " ", device.lastStatusUpdate, ", temp: ", device.actualTemperature, ", hum: ", device.humidity)
                    if isinstance(device, WallMountedThermostatPro):
                        data = {
                            "lastStatusUpdate" : device.lastStatusUpdate.strftime("%Y-%m-%d %H:%M:%S"),
                            "deviceName": device.label,
                            "actualTemperature": device.actualTemperature,
                            "humidity": device.humidity
                        }
                        with open('./data/' + today + '-devices.log','a', encoding ='utf8') as f:
                            json.dump(data, f, ensure_ascii=False)
                            f.write("\n")
                            f.close()

        time.sleep(60)

if __name__ == "__main__":
    main()

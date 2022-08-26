# homematic-metric-sync

## Create or switch to virtual envrionment

Create:

```sh
python -m venv .venv
# Install requirements:
pip install -r requirements.txt
```

Switch:

```sh
./.venv/bin/activate
```

## Install homematicip-rest-api (first time only)

```sh
pip install git+https://github.com/coreGreenberet/homematicip-rest-api
```

## Freeze requirements after changing modules

f:/Entwicklung/git/homeatic-metric-sync/.venv/Scripts/python.exe -m pip freeze > requirements.txt

## Create and run docker image

```sh
docker build . -t "homeatic-metric-sync:local"

# Mount local config, script and data directory for easy testing
docker run -it --rm -v "$(pwd)/config.ini:/config.ini" -v "$(pwd)/get_data.py:/get_data.py" -v "$(pwd)/data:/data" homeatic-metric-sync:local
```

## Links

<https://homematicip-rest-api.readthedocs.io/en/latest/>

# Automate Application Status

I applied to be a adjunct professor. They don't email you when there is a application status change. So let's automate checking for status changes!

This program will login to the university website and check the application status. If there was an application status change it will print status change and open a screenshot of the webpage.

If there is no status change then it will print no status change.

## Dependencies:
```
sudo apt-get install -y openimageio-tools
sudo apt-get install -y install xdg-utils
```

## Running this program:
### When there's a status change:
```
> byui
Status change! Openining screenshot of application status...
```

### When there's no status change:
```
> byui
No status change
```
# Automate Application Status

I applied to be a adjunct professor. They don't email you when there is a application status change. So let's automate checking for status changes!

This program will login to the university website and check the application status. If there was an application status change it will print status change and open a screenshot of the webpage. 

Here is what the automation does when you run this `byui` command line tool:

1. Login to https://employment.byui.net/login using your username and password (make sure to update the main.go with your username and password)
2. Navigates to https://employment.byui.net/job_applications
3. Takes a screenshot of the webpage
4. Checks if the `.byui` folder in the users home directory exists, if not creates it.
5. Checks if `appstatus.png` exists. If so rename it to `old_appstatus.png`
6. Saves the screenshot from #3 as `appstatus.png`
7. Uses `idiff` to check if there is a difference between the images: `appstatus.png` and `old_appstatus.png`
    1.  If so then print to the console there is a status change and open `appstatus.png` using `xdg-open`
    2. If not then print to the console no status change

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
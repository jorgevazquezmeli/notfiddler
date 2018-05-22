# notfiddler
Fiddler clone as an exercise for the IT accelerator program, made in Go.

## Usage

Just execute the goproxy file from the 'src/github.com/jorgevazquezmeli/goproxy' folder. The proxy runs on 127.0.0.1:8080, so you'll need to set up your proxy settings in your favorite browser. The proxy outputs data for each request and response data only for HTML/JSon content. You can also see an excerpt from the body.


## Browser configuration

To use it, you'll need to set up your favorite browser.

### Firefox

In Firefox, press the main menu button on the top right corner and select the *Preferences* option. From there, use the top search form to look for the "proxy" button. Click on the button and in the new window select "Manual Proxy". In the corresponding fields, type 127.0.0.1 as the server and 8080 as the port. Below, check that the option "Proxy DNS when using SOCKS 5" is also selected.

![Firefox Setup](/imgs/firefox.png?raw=true "Firefox Setup")

### Chrome Setup

In Chrome, go to the main menu button on the top right corner of your browser and select *Settings*. In the search bar, type "proxy" and click on "Open proxy settings." In Mac OS, a new window will open with your system's proxy settings. In there, select "Web Proxy (HTTP)" and type 127.0.0.1 as the URL and 8080 as the port. Click OK and Apply.

![Chrome Setup in Mac OS](/imgs/chromemac.png?raw=true "Chrome setup in Mac OS")

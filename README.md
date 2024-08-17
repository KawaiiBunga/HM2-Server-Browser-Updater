# HM2-Server-Browser-Updater
Easily update your favorites server list in HM2 with the most recent servers with a simple double click.

# Reccomended Usage - 
- Technically, you can run this program anywhere and it will generate a file called ```favourites.json```, which you can then move into ```root_game_dir/players2```.

- However, what I personally prefer, is placing the ```servergrabmwr.exe``` INSIDE of the ```h2m_root_game_dir/players2``` folder.
- - You can then right click on ```servergrabmwr.exe``` inside of the ```h2m_root_game_dir/players2``` folder and press "Create Shortcut"
- - Now, you can move the shortcut wherever you have your HM2 shortcut to launch it. Whenever you feel like you servers need updated, simply run the shortcut and in less than a couple seconds, your servers are updated with the newest list from ```https://master.iw4.zip/servers#```

# Building
1. Make sure you have the latest version of GO installed (https://go.dev/doc/install)
2. Install the dependency: ```go get github.com/PuerkitoBio/goquery```
3. Simply run ```go build``` in the project directory
4. This will output ```servergrabmwr.exe``` that you can now use (See Reccomended Usage)

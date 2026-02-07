set -e 
while true; do
    read -p "Do you wish to uninstall terminal notes? " yn
    case $yn in
        [Yy]* ) break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y or n.";;
    esac
done

echo "Deleting Executable..."
rm -f /usr/local/bin/terminal-notes

while true; do
    read -p "Do you wish to delete your notes? " yn
    case $yn in
        [Yy]* ) 
            # Fetches user's home directory 
            # Source - https://stackoverflow.com/a/7359006
            # Posted by Michał Šrajer, modified by community. See post 'Timeline' for change history
            # Retrieved 2026-02-07, License - CC BY-SA 3.0
            USER_HOME=$(getent passwd $SUDO_USER | cut -d: -f6)
            NOTES_DIR="$USER_HOME/.local/share/terminal-notes"
            echo "Deleting notes from $NOTES_DIR... "
            rm -r $NOTES_DIR
            break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y or n.";;
    esac
done

echo "Done."
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

echo "Done."
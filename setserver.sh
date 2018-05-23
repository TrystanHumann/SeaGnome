setcap 'cap_net_bind_service=+ep' /home/momam/main &&
sudo systemctl start momam &&
sudo systemctl status momam
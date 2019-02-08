# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.define "web" do |web|
    web.vm.box = "centos/7"
    web.vm.network "private_network", ip: "192.168.33.10"
    web.vm.provision "shell", inline: <<-SHELL
      echo "webserver" > /etc/hostname
      hostname -b webserver
      echo "192.168.33.20          clockserver" >> /etc/hosts
    SHELL
  end
  config.vm.define "clock" do |clock|
    clock.vm.box = "centos/7"
    clock.vm.network "private_network", ip: "192.168.33.20"
    clock.vm.provision "shell", inline: <<-SHELL
      echo "clockserver" > /etc/hostname
      hostname -b clockserver
      echo "192.168.33.10          webserver" >> /etc/hosts
    SHELL
  end
end

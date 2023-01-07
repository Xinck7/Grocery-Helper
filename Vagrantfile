
Vagrant.configure("2") do |config|
  
  config.vm.box = "generic/rocky8"

  config.vm.synced_folder "./", "/opt"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = 2048
    vb.cpus = 1
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo yum clean all
    sudo yum update -y
    # sudo yum groupinstall "Server with GUI" -y
    curl -sL https://rpm.nodesource.com/setup_18.x | sudo -E bash -
    sudo dnf install nodejs -y
    npm install npm@latest -g
    npm install -g @angular/cli
  SHELL
end

=================
PlugDJ Test 
=================

A PlugDJ Test using Vagrant.

Requirements
===========

Please install Vagrant and VirtualBox in your local machine in order to follow this instructions


Instructions
============

Run the following commands in a terminal. Git, VirtualBox and Vagrant must
already be installed.

.. code-block:: bash

    git clone https://github.com/alejandrobernalcollazos/plugdj-test
    cd plugdj-test
    vagrant plugin install vagrant-vbguest
    vagrant up


This will download an Ubuntu  VirtualBox image and create four virtual
machines for you. 

One will be a Salt Master named `master` two will be Salt Minions named `minion1` and `minion2` that will contain the https server configurations (nginx and a golanguage based application) and minion3 that will contain a postgressql database. 

 The Salt Minions will point to the Salt Master and the Minion's keys will already be accepted. Because the keys are pre-generated and reside in the repo.

You can then run the following commands to log into the Salt Master and begin
using Salt.

.. code-block:: bash

    vagrant ssh master
    sudo su
    salt '*' state.apply

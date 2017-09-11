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

Then once the provisioning has finished, please get into the browser with this urls to test the http server working through an nginx proxy 

.. code-block:: bash

    http://192.168.50.11:8080
    http://192.168.50.12:8080

In order to test the access to the DB please execute this command within any of minion1 or minion2 machines

.. code-block:: bash

   psql -h 192.168.50.13 -U plugdj
   password : plugdj

Testing "Now" Endpoint
======================

In both virtual machines please try in the browser this two urls

.. code-block:: bash

   http://192.168.50.11:8080/now
   http://192.168.50.12:8080/now

Testing "Check" Endpoint
========================

In both virtual machines please try in the browser this two urls

.. code-block:: bash

   http://192.168.50.11:8080/check
   http://192.168.50.12:8080/check

Testing "Later" Endpoint
========================

In order to test the later endpoint in both virtual machines please try in the console the next commands

.. code-block:: bash

   curl -H "Content-Type: application/json" -d '{"name":"Andrew"}' http://192.168.50.11:8080/later -vvv
   curl -H "Content-Type: application/json" -d '{"name":"Akita"}' http://192.168.50.12:8080/later -vvv

It will be useful if we see the check endpoint again in both servers

.. code-block:: bash

   http://192.168.50.11:8080/check
   http://192.168.50.12:8080/check

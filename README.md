Nebula - An Augmented Reality Storefront
===================


What is Nebula?
------------
Nebula is a augmented reality storefront that allows shopping online to be a virtual experience.  Using cutting edge technology with the Microsoft Hololens, Nebula brings shopping into the 21st Century.


Hololens App
----------
The main component of Nebula is the Microsoft Hololens app.  The application runs on the Microsoft Hololens and is the key to the shopping experience.  Using the Hololens technology, Nebula turns your entire field of view into a virtual space.  You can then view a holographic, 3D to scale model of an item prior to purchasing it.  Imagine being able to look at an empty room and holographiclly furnish it entirely without ever leaving your home.  That's the magic of Nebula.

Project can be found here: https://github.com/Generalkidd/NebulaShop

NebulaShop Website
----------
In order to submit an item for sale, you have to simply log into https://nebulashop.net/, create a profile (all sellers are required to have on) and then you can start uploading your items for sale.  All you need is a 3D model of your item and it'll be ready to viewed in the Hololens App for purchase.

In order to get a 3D model, there are a couple options:
* 3rd Party smartphone apps.  These exist to allow you to create a 3D model of an image with your phone, but they are not perfect.
* Our recommended method is using the official Microsoft App for scanning an object with your smarthphone to get a 3D model.  This is not yet released, but is slated for 2017.
* For certain items, we will have pre-loaded stock image models that you can leverage.  Just add the dimensions of these pre-loaded models and you can use that.  Note, these don't exist of every possilbe item.

Nebula Backend
----------
The Nebula backend is a modular Go based API that interacts with a MongoDB database.  The API is a standard REST based API that provides all of the needed interaction for both the NebulaShop website and the Hololens App.

The Nebula backend also contains a Go program for rebuilding the database for regression testing purposes. The regression tests are much more effective and accurate when run on a clean DB. It is required you use Testing DBs and the Testing API for this so as not to delete the entire production DB when rebuilding it.

Also included is a regression test suite for every endpoint currently in the Nebula API.  The testing suite is written in Python and is to be executed after every addition of a new endpoint or any other functionality.  It will output the final results of what endpoints passed and what endpoints failed as well as whether or not the regression test was passable (currently 80% of endpoints passing is a passable test, however you'll need to examine what endpoints/functions failed to ensure 100% success in subsequent runs of the test suite).  This will ensure that all previous functionality is maintained.  After that the test set for the new endpoint/function will be added to the suite.

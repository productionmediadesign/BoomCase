Setup - Structure and Launching GAE Application
===============================================

Google App Engine Setup
-----------------------

Google App Engine Setup via web (use Google Chrome):
* x

PC Setup
--------

Install Python 2.7.x:
* https://www.python.org/downloads/
  * Locate the newest version of 2.7.x and download 64-bit version
  * Install for all users to: C:\Python27\

Install Golang:
* https://golang.org/dl/
* Test Installation:
  * Open command prompt:
  * > go version

Install Git:
* x

Install Google Cloud:
* https://cloud.google.com/appengine/downloads
  * Click to get the Standard Environment for Golang
  * https://cloud.google.com/appengine/docs/standard/go/download
  * Leads to next step below..
* Get the Cloud SDK for Go and install it:
  * https://cloud.google.com/sdk/docs/
    * Install via the Cloud SDK Installer
    * Once install is complete, command line opens and asks to login (do it)
    * Select the cloud project pmd-thumpcase (that you created earlier via web)
    * NOTE: Google Cloud SDK Folder:
      * > cd "C:\Users\%USERNAME%\AppData\Local\Google\Cloud SDK"
  * Enable command completion in shell:
    * > "C:\Users\%USERNAME%\AppData\Local\Google\Cloud SDK\google-cloud-sdk\install.sh"
  * Review Documentation: https://cloud.google.com/appengine/docs/standard/go/

Install App Engine extension for Go:
* > gcloud components install app-engine-go

Ensure you are up-to-date:
* > gcloud components update

Set CLOUDSDK_PYTHON Environment Variable:
* Start > "environment variable" > Click: "Edit the System Environment Variables"
* Click: "Environment Variables" button
* Click: "New..."
  * Variable name: CLOUDSDK_PYTHON
  * Variable value: C:\Python27\python.exe

Run Application Locally:
* > dev_appserver.py app.yaml

Deploy Application to Google App Engine:
* > gcloud app deploy --version pre-prod-5 --project my-sample-app

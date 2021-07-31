# ******************************************************************************
# This script uses the TestFuncs.py functions and enters the data from each file found
# in the 'test_inputs' directory as an entry on the website.
# 
# It is run by specifying the URL of the webpage to run the tests on
# ie. python3 RunAllUITests.py localhost:80
# ******************************************************************************
# *** REQUIRED DEPENDENCIES ***
# -- selenium python library
# -- Firefox web browser
# -- Geckodriver (allows interaction between selenium/firefox)
# 
# Selenium can be installed with pip/pip3
#
# Installation instructions for Selenium and Geckodriver can be found here:
# https://selenium-python.readthedocs.io/installation.html#downloading-python-bindings-for-selenium
#
# ******************************************************************************

from TestFuncs import *

headless = True

def test_from_file(filepath):
    global driver
    global infile 
    infile = open_datafile(filepath)
    driver = initialize_driver(webpage_path)

    try:
        input_team_names()
        input_section_1()
        input_section_2()
        input_section_3()
        input_section_4()
        input_section_5()
        submit()
    except Exception as e:
        print("An error occurred. Closing Browser")    
        print(e)

if __name__ == "__main__":
    global driver
    driver = initialize_driver()

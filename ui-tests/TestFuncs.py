# ******************************************************************************
# This script uses Selenium to open a physical window and interact directly with the submissions webpage.
# This allows the creation of automated tests which can verify user input and interact with the page
# as a normal user.
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

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions
from time import sleep

# path to file containing information to enter
# data_path = 'C:/Users/brian/OneDrive/Documents/2021 Spring Semester/CPTS 421/bioretentiongardeningapp/tests/test1.txt'
data_path = '/Users/bturley/Documents/capstone/gardenapp/ui-tests/test_inputs/test1.txt'
# URL of webpage to open
# webpage_path = "http://bioretentionapp2/"
webpage_path = "http://129.146.241.222/"

# environment variables controlling automated web browser
headless = False
fullscreen = False

# opens the given file for read mode
def open_datafile(path_to_data):
    return open(path_to_data, "r")


# reads, strips, and returns line read from infile
def getline():
    return infile.readline().strip()


# creates the automated web browser using the specified environment variables
def initialize_driver(url):
    # open the browser
    options = Options()
    options.headless = headless
    driver = webdriver.Firefox(options = options)

    if not headless and fullscreen:
        driver.maximize_window()

    # open given webpage
    driver.get(url)
    try:
        # wait for the browser to load
        WebDriverWait(driver, 10).until(
            expected_conditions.element_to_be_clickable((By.NAME, 'name1'))
        )
    except:
        driver.quit()
        print("ERROR: Failed to load webpage.")
        exit(1)
    return driver


# performes any cleanup tasks necessary
# ie. closing the webrowser or deleting any generated files
def cleanup():
    driver.quit()


# provided the css_selector of a given text-box page element,
# fills the text-box with the string read from the input file
def fill_textbox_from_input(css_selector_str):
    elem = driver.find_element(
        By.CSS_SELECTOR, 
        css_selector_str
    )
    elem.send_keys(getline())


# provided the css_selector of a given drop-down page element,
# selects the drop-down entry that exactly matches the string read from the input file
def fill_dropdown_from_input(css_selector_str):
    choice = getline()
    dropdown = driver.find_element(
        By.CSS_SELECTOR,
        css_selector_str
    )

    # retrieves a list of all response elements
    selection_options = dropdown.find_elements_by_css_selector('*')

    # click the selection that matches the user's desired choice
    for elem in selection_options:
        if elem.text == choice:
            elem.click()
            return


# fills the page form with team names read from input file
def input_team_names():
    row = 1
    col = 5

    # input values into table
    for _ in range(4):
        fill_textbox_from_input(
            'section.center:nth-child(1) > div:nth-child(1) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(%d) > td:nth-child(%d) > input:nth-child(1)' %(row, col)
        )
        fill_textbox_from_input(
            'section.center:nth-child(1) > div:nth-child(1) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(%d) > td:nth-child(%d) > input:nth-child(1)' %(row, col+2)
        )
            
        # update row & column
        row += 1
        col = 4

    getline()  # eat empty line in file
    print("Successfully entered all team names.")
    

# TODO: Survey Date/Time don't fill properly
# fills "Section 1: Background Information" with data read from input file
def input_section_1():
    # site name
    fill_textbox_from_input(
        '.col-lg-5 > input:nth-child(3)'
    )

    # survey date
    fill_textbox_from_input(
        'div.row:nth-child(5) > div:nth-child(1) > input:nth-child(2)'
    )

    # start time
    fill_textbox_from_input(
        'div.col-6:nth-child(2) > input:nth-child(2)'
    )

    # group code
    fill_textbox_from_input(
        'div.col-lg-3:nth-child(3) > input:nth-child(2)'
    )

    # address
    fill_textbox_from_input(
        'div.col-lg-6 > input:nth-child(3)'
    )
    
    # city
    fill_textbox_from_input(
        '.mt-1 > div:nth-child(2) > input:nth-child(2)'
    )

    # county
    fill_textbox_from_input(
        '.mt-1 > div:nth-child(3) > input:nth-child(2)'
    )

    # latitude, longitude, sound impacts ID
    for n in range(1,4):
        fill_textbox_from_input(
            'div.row:nth-child(8) > div:nth-child(%d) > input:nth-child(2)' %n
        )
        
    # rainfall today, yesterday, two days ago
    for n in range(1,4):
        fill_textbox_from_input(
            'div.row:nth-child(10) > div:nth-child(%d) > input:nth-child(2)' %n
        )

    getline()  # eat empty line in file
    print("Successfully entered all components of section 1.")


# fills "Section 2: Site Overview" with data read from input file
def input_section_2():
    # type of site
    fill_dropdown_from_input('#site-type')

    # age of site
    fill_dropdown_from_input('#site-age')

    # source of 'age'
    temp = getline()
    if temp == "Verifiable Source":
        driver.find_element_by_css_selector('#VS').click()
    elif temp == "Estimate":
        driver.find_element_by_css_selector('#Es').click()
    
    # description
    fill_textbox_from_input('div.col-lg-8:nth-child(1) > textarea:nth-child(2)')

    getline()    # eat empty line in file
    print("Successfully entered all components of section 2.")


# fills "Section 3: Contributing Area/Hydrology" with data read from input file
def input_section_3():  
    # Contributin Water Source:
    water_sources = getline().split(", ")
    # check each clickable item
    # if the label's text is in the given list of water_sources then click it
    for row in range(1,8):
        elem = driver.find_element(
            By.CSS_SELECTOR,
            'div.align-items-center:nth-child(%d) > label:nth-child(2)' %row
        )
        if elem.text in water_sources:
            elem.click()

    # Overflow 1
    given_choice = getline()
    if given_choice == "Yes":
        driver.find_element(
            By.CSS_SELECTOR,
            'div.ml-4:nth-child(2) > span:nth-child(2) > input:nth-child(1)'
        ).click()
    elif given_choice == "No":
        driver.find_element(
            By.CSS_SELECTOR,
            'div.ml-4:nth-child(2) > span:nth-child(3) > input:nth-child(1)'
        ).click()
    elif given_choice == "Unknown":
        driver.find_element(
            By.CSS_SELECTOR,
            'div.ml-4:nth-child(2) > span:nth-child(4) > input:nth-child(1)'
        ).click()

    # Overflow 2
    given_choice = getline()
    if given_choice == "Yes":
        driver.find_element(
            By.CSS_SELECTOR,
            'div.ml-4:nth-child(4) > span:nth-child(2) > input:nth-child(1)'
        ).click()
    elif given_choice == "No":
        driver.find_element(
            By.CSS_SELECTOR,
            'div.ml-4:nth-child(4) > span:nth-child(3) > input:nth-child(1)'
        ).click()
    elif given_choice == "Unknown":
        driver.find_element(
            By.CSS_SELECTOR,
            'div.ml-4:nth-child(4) > span:nth-child(4) > input:nth-child(1)'
        ).click()

    # Overflow 3
    given_choice = getline()
    if given_choice == "Yes":
        driver.find_element(
            By.CSS_SELECTOR,
            '.mb- > div:nth-child(6) > span:nth-child(2) > input:nth-child(1)'
        ).click()
    elif given_choice == "No":
        driver.find_element(
            By.CSS_SELECTOR,
            '.mb- > div:nth-child(6) > span:nth-child(3) > input:nth-child(1)'
        ).click()
    elif given_choice == "Unknown":
        driver.find_element(
            By.CSS_SELECTOR,
            '.mb- > div:nth-child(6) > span:nth-child(4) > input:nth-child(1)'
        ).click()

    # percent block
    for n in range(1, 4):
        fill_dropdown_from_input('#percent_blockage_I%d' %n)
    fill_dropdown_from_input('#percent_blockage_SF')
    for n in range(1, 4):
        fill_dropdown_from_input('#percent_blockage_O%d' %n)
                                 
    # blockage type
    for n in range(1, 4):
        fill_dropdown_from_input('#blockage_type_I%d' %n)
    fill_dropdown_from_input('#blockage_type_SF')
    for n in range(1, 4):
        fill_dropdown_from_input('#blockage_type_O%d' %n)
                                 
    # zone erosion
    for n in range(1, 4):
        fill_dropdown_from_input('#Erosion_Z%d' %n)
                                 
    fill_textbox_from_input('textarea.form-control:nth-child(1)')
    
    getline()    # eat empty line in file
    print("Successfully entered all components of section 3.")


# fills "Section 4: Zone 1 Conditions" with data read from input file
def input_section_4():
    # zone 1 length
    fill_textbox_from_input('body > form > div:nth-child(9) > div:nth-child(2) > div:nth-child(2) > div > input')
        
    # water depth
    for n in range(2, 5):
        fill_textbox_from_input('table.table:nth-child(4) > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(%d) > input:nth-child(1)' %n)
    
    # siltation depth
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#siltation-depth-%c' %n)
                                 
    # liner present
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#liner-present-%c' %n)
                                 
    # liner depth
    for n in range(2, 5):
        fill_textbox_from_input('table.table:nth-child(4) > tbody:nth-child(2) > tr:nth-child(4) > td:nth-child(%d) > input:nth-child(1)' %n)
    
    # filter fabric present
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#filter-present-%c' %n)
                                 
    # filter depth
    for n in range(2, 5):
        fill_textbox_from_input('table.table:nth-child(4) > tbody:nth-child(2) > tr:nth-child(6) > td:nth-child(%d) > input:nth-child(1)' %n)
    
    # native soil depth
    for n in range(2, 5):
        fill_textbox_from_input(
            'div.container-fluid:nth-child(9) > div:nth-child(4) > section:nth-child(2) > div:nth-child(1) > table:nth-child(1) > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(%d) > input:nth-child(1)' %n
        )
        
    # compacted surface soil
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#compacted-soil-%c' %n)
                                 
    # rain garden texture
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#garden-soil-%c' %n)
                                 
    # native soil texture
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#native-soil-%c' %n)
                                 
    # other observations
    fill_textbox_from_input('div.container-fluid:nth-child(9) > div:nth-child(5) > div:nth-child(2) > textarea:nth-child(2)')
    
    getline()    # eat empty line in file    
    print("Successfully entered all components of section 4.")


# TODO: add support for End-Time 
# fills "Section 5: Overall substrates, Vegetation, Conditions" 
# with data read from input file=
def input_section_5():
    # mulch type
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#mulch-type-1%c' %n)
    for n in range (2, 4):
        fill_dropdown_from_input('#mulch-type-%d' %n)
                                 
    # mulch depth
    for n in [ 'A', 'B', 'C' ]:
        fill_dropdown_from_input('#mulch-depth-1%c' %n)
    for n in range (2, 4):
        fill_dropdown_from_input('#mulch-depth-%d' %n)
                                 
    # zone coverages
    for n in [ 'mulch', 'ground', 'gravel', 'drain', 'mid', 'low' ]:
        for m in range(1, 4):
            fill_dropdown_from_input('#%s-coverage-%d' % (n, m))
    
    # vegetation
    for n in range(1, 4):
        fill_dropdown_from_input('#vegetation-coverage-%d' %n)
    for n in [ 'problem', 'weeds', 'deciduous', 'evergreen', 'herbaceous', 'soil' ]:
        for m in [ 'coverage', 'vigor' ]:
            for o in range(1, 4):
                fill_dropdown_from_input('#%s-%s-%d' % (n, m, o))
                                         
    # vegetation observations
    fill_textbox_from_input('div.pt-3:nth-child(9) > div:nth-child(2) > textarea:nth-child(2)')
    
    # visible
    fill_textbox_from_input('#visible-to-public')
                            
    # aesthetically pleasing
    fill_textbox_from_input('#aesthetically-pleasing')
                            
    # well maintained
    fill_textbox_from_input('#well-maintained')
                            
    # educational signage
    given_choice = getline()
    if given_choice == "Yes":
        driver.find_element_by_css_selector('#signage-yes').click()
    elif given_choice == "No":
        driver.find_element_by_css_selector('#signage-no').click()
    
    # other observations
    fill_textbox_from_input('.textarea-resize')
    
    print("Successfully entered all components of section 5.")


# submits the webform using the button at the bottom of the page
def submit():
    elem = driver.find_element(By.CSS_SELECTOR, 'input.btn')
    elem.click() # click once to get out of last box?
    elem.click() # click again to submit form
    print("Submitted the webform.")


if __name__ == "__main__":
    global driver
    global infile 
    infile = open_datafile(data_path)
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

    cleanup()

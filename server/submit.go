/*
This file details structs and functions which allow submitting queries to the MYSQL database.
This includes parsing input from the GET/POST requests, cleaning it, and submitting to the database.
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type sectionInfo struct {
	water_depth     string
	siltation_depth int
	liner_present   int
	liner_depth     string
	filter_present  int
	filter_depth    string
	native_depth    string
	compacted_soil  int
	garden_soil     int
	native_soil     int
}

type zoneInfo struct {
	mulch_coverage           int
	bare_ground_coverage     int
	pea_gravel_coverage      int
	drain_rock_coverage      int
	rock_lower_coverage      int
	rock_upper_coverage      int
	all_vegetation_coverage  int
	problem_plants_coverage  int
	non_target_weed_coverage int
	deciduous_coverage       int
	evergreen_coverage       int
	herbaceous_coverage      int
	ground_cover_coverage    int
	problem_plants_vigor     int
	non_target_weed_vigor    int
	deciduous_vigor          int
	evergreen_vigor          int
	herbaceous_vigor         int
	ground_cover_vigor       int
}

type formData struct {
	name1     string
	name2     string
	name3     string
	name4     string
	numEmail1 string
	numEmail2 string
	numEmail3 string
	numEmail4 string

	site_name      string
	survey_date    string
	start_time     string
	group_code     string
	address        string
	city           string
	county         string
	lat            string
	lon            string
	sounds_impacts string
	today_rain     string
	yesterday_rain string
	twodays_rain   string
	site_type      int
	site_age       int

	age_source_radio int
	age_source_desc  string

	water_sources    [8]int
	water_source_ids string

	overflow_1_radio int
	overflow_2_radio int
	overflow_3_radio int

	percent_blockage_I1 int
	percent_blockage_I2 int
	percent_blockage_I3 int
	percent_blockage_SF int
	percent_blockage_01 int
	percent_blockage_02 int
	percent_blockage_03 int

	blockage_type_I1 int
	blockage_type_I2 int
	blockage_type_I3 int
	blockage_type_SF int
	blockage_type_01 int
	blockage_type_02 int
	blockage_type_03 int

	erosion_Z1 int
	erosion_Z2 int
	erosion_Z3 int

	hydroconc   string
	zone1_len   string
	section_len string

	section_info_A sectionInfo
	section_info_B sectionInfo
	section_info_C sectionInfo

	substrate_observations string

	mulch_type_1A int
	mulch_type_1B int
	mulch_type_1C int
	mulch_type_2  int
	mulch_type_3  int

	mulch_depth_1A int
	mulch_depth_1B int
	mulch_depth_1C int
	mulch_depth_2  int
	mulch_depth_3  int

	zone_info_1 zoneInfo
	zone_info_2 zoneInfo
	zone_info_3 zoneInfo

	vegetation_observations string

	visible_to_public      int
	aesthetically_pleasing int
	well_maintained        int
	educational_signage    int

	other_observations string
	end_time           string
}

func parseAndSubmit(r *http.Request) error {
	r.ParseForm()

	f := formData{}
	// TODO: handle errors received from here
	retrieveFormData(&f, r)
	checkValidInputs(&f)
	err := submitToDB(&f)
	if err != nil {
		fmt.Println(err) // this will print any error from db stuff
		return err
	}
	return nil
}

// returns the first data entry from the form at the given name as a string
func getFormDataString(r *http.Request, name string) string {
	if len(r.Form[name]) > 0 {
		return string(r.Form[name][0])
	}
	return ""
}

// returns the first data entry from the form at the given name as an int
func getFormDataInt(r *http.Request, name string) int {
	if len(r.Form[name]) > 0 {
		if len(r.Form[name][0]) == 0 {
			return -1 // default value for empty inputs
		}
		intVal, error := strconv.Atoi(r.Form[name][0])
		if error == nil {
			return intVal
		}
	}
	return 0 // default value for unselected radio/dropdown
}

// returns the first data entry from the form at the given name as a float
func getFormDataFloat(r *http.Request, name string) string {
	if len(r.Form[name]) > 0 {
		floatVal, error := strconv.ParseFloat(r.Form[name][0], 32)
		if error == nil {
			return fmt.Sprintf("%.6f", floatVal)
		}
	}
	return "NULL"
}

// makes a SQL query to submit the database
func submitToDB(f *formData) error {
	// Insert sections into db
	sectIds, err := submitSectionsToDB(f)
	if err != nil {
		return err
	}
	// Insert zones into db
	zoneIds, err := submitZonesToDB(f)
	if err != nil {
		return err
	}

	start_time := fmt.Sprintf("%v:00", f.start_time)
	if len(f.start_time) < 1 {
		start_time = ""
	}
	end_time := fmt.Sprintf("%v:00", f.end_time)
	if len(f.end_time) < 1 {
		end_time = ""
	}
	query := "INSERT INTO Form VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, ?, ?, " + fmt.Sprintf("%v, %v, ", f.lat, f.lon) + "?, " +
		fmt.Sprintf("%v, %v, %v, ", f.today_rain, f.yesterday_rain, f.twodays_rain) +
		"?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, " +
		fmt.Sprintf("%v, ", f.zone1_len) + "?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, " +
		"?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	fmt.Println(query)
	result, err := db.Exec(query,
		f.name1, f.numEmail1, f.name2, f.numEmail2, f.name3, f.numEmail3,
		f.name4, f.numEmail4, f.site_name, f.survey_date, start_time, f.group_code,
		f.address, f.city, f.county, f.sounds_impacts, f.site_type, f.site_age,
		f.age_source_radio, f.age_source_desc, f.overflow_1_radio, f.overflow_2_radio,
		f.overflow_3_radio, f.percent_blockage_I1, f.percent_blockage_I2,
		f.percent_blockage_I3, f.percent_blockage_SF, f.percent_blockage_01,
		f.percent_blockage_02, f.percent_blockage_03, f.blockage_type_I1,
		f.blockage_type_I2, f.blockage_type_I3, f.blockage_type_SF,
		f.blockage_type_01, f.blockage_type_02, f.blockage_type_03, f.erosion_Z1,
		f.erosion_Z2, f.erosion_Z3, f.hydroconc, sectIds[0], sectIds[1],
		sectIds[2], f.substrate_observations, f.mulch_type_1A, f.mulch_type_1B,
		f.mulch_type_1C, f.mulch_type_2, f.mulch_type_3, f.mulch_depth_1A,
		f.mulch_depth_1B, f.mulch_depth_1C, f.mulch_depth_2, f.mulch_depth_3,
		zoneIds[0], zoneIds[1], zoneIds[2], f.vegetation_observations,
		f.visible_to_public, f.aesthetically_pleasing, f.well_maintained,
		f.educational_signage, f.other_observations, end_time)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	err = insertWaterSources(id, f.water_sources)
	if err != nil {
		return err
	}

	return nil
}

// inserts sections to DB
func submitSectionsToDB(f *formData) ([3]int64, error) {
	ids := [3]int64{-1, -1, -1}
	idA, err := submitSectionToDB(&f.section_info_A)
	if err != nil {
		return ids, err
	}
	ids[0] = idA
	idB, err := submitSectionToDB(&f.section_info_B)
	if err != nil {
		return ids, err
	}
	ids[1] = idB
	idC, err := submitSectionToDB(&f.section_info_C)
	if err != nil {
		return ids, err
	}
	ids[2] = idC
	return ids, err
}

// inserts zones to DB
func submitZonesToDB(f *formData) ([3]int64, error) {
	ids := [3]int64{-1, -1, -1}
	id1, err := submitZoneToDB(&f.zone_info_1)
	if err != nil {
		return ids, err
	}
	ids[0] = id1
	id2, err := submitZoneToDB(&f.zone_info_2)
	if err != nil {
		return ids, err
	}
	ids[1] = id2
	id3, err := submitZoneToDB(&f.zone_info_3)
	if err != nil {
		return ids, err
	}
	ids[2] = id3
	return ids, err
}

func submitSectionToDB(s *sectionInfo) (int64, error) {
	query := "INSERT INTO SectionInfo VALUES (NULL, "
	query += fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		s.water_depth, s.siltation_depth, s.liner_present, s.liner_depth,
		s.filter_present, s.filter_depth, s.native_depth, s.compacted_soil,
		s.garden_soil, s.native_soil)
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	result, err := stmt.Exec()
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, err
}

func submitZoneToDB(z *zoneInfo) (int64, error) {
	query := "INSERT INTO ZoneInfo VALUES (NULL, "
	query += fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v, %v, ",
		z.mulch_coverage, z.bare_ground_coverage, z.pea_gravel_coverage,
		z.drain_rock_coverage, z.rock_lower_coverage, z.rock_upper_coverage,
		z.all_vegetation_coverage, z.problem_plants_coverage, z.non_target_weed_coverage)
	query += fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		z.deciduous_coverage, z.evergreen_coverage, z.herbaceous_coverage, z.ground_cover_coverage,
		z.problem_plants_vigor, z.non_target_weed_vigor, z.deciduous_vigor, z.evergreen_vigor,
		z.herbaceous_vigor, z.ground_cover_vigor)
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	result, err := stmt.Exec()
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, err
}

func insertWaterSources(formId int64, water_sources [8]int) error {
	for i := 0; i < 8; i++ {
		if water_sources[i] == 1 {
			err := insertWaterSource(formId, i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func insertWaterSource(formId int64, index int) error {
	query := "INSERT INTO FormWaterSource VALUES (NULL, "
	query += fmt.Sprintf("%v, %v);", formId, index)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

// ensures that all entries are valid
func checkValidInputs(f *formData) error {
	return nil
}

func retrieveFormData(f *formData, r *http.Request) error {
	f.site_name = getFormDataString(r, "site_name")
	f.name1 = getFormDataString(r, "name1")
	f.name2 = getFormDataString(r, "name2")
	f.name3 = getFormDataString(r, "name3")
	f.name4 = getFormDataString(r, "name4")
	f.numEmail1 = getFormDataString(r, "numEmail1")
	f.numEmail2 = getFormDataString(r, "numEmail2")
	f.numEmail3 = getFormDataString(r, "numEmail3")
	f.numEmail4 = getFormDataString(r, "numEmail4")
	f.site_name = getFormDataString(r, "site_name")
	f.survey_date = getFormDataString(r, "survey_date")
	f.start_time = getFormDataString(r, "start_time")
	f.group_code = getFormDataString(r, "group_code")
	f.address = getFormDataString(r, "address")
	f.city = getFormDataString(r, "city")
	f.county = getFormDataString(r, "county")
	f.lat = getFormDataFloat(r, "lat")
	f.lon = getFormDataFloat(r, "lon")
	f.sounds_impacts = getFormDataString(r, "Sound_Impacts")
	f.today_rain = getFormDataFloat(r, "today_rain")
	f.yesterday_rain = getFormDataFloat(r, "yesterday_rain")
	f.twodays_rain = getFormDataFloat(r, "twodays_rain")
	f.site_type = getFormDataInt(r, "site-type")
	f.site_age = getFormDataInt(r, "site-age")
	f.age_source_radio = getFormDataInt(r, "age_source_radio")
	f.age_source_desc = getFormDataString(r, "age_source_desc")
	for i := 0; i < len(f.water_sources); i++ {
		f.water_sources[i] = getFormDataInt(r, "WS"+strconv.Itoa(i+1))
	}
	f.overflow_1_radio = getFormDataInt(r, "overflow_1_radio")
	f.overflow_2_radio = getFormDataInt(r, "overflow_2_radio")
	f.overflow_3_radio = getFormDataInt(r, "overflow_3_radio")
	f.percent_blockage_I1 = getFormDataInt(r, "percent_blockage_I1")
	f.percent_blockage_I2 = getFormDataInt(r, "percent_blockage_I2")
	f.percent_blockage_I3 = getFormDataInt(r, "percent_blockage_I3")
	f.percent_blockage_SF = getFormDataInt(r, "percent_blockage_SF")
	f.percent_blockage_01 = getFormDataInt(r, "percent_blockage_O1")
	f.percent_blockage_02 = getFormDataInt(r, "percent_blockage_O2")
	f.percent_blockage_03 = getFormDataInt(r, "percent_blockage_O3")
	f.blockage_type_I1 = getFormDataInt(r, "blockage_type_I1")
	f.blockage_type_I2 = getFormDataInt(r, "blockage_type_I2")
	f.blockage_type_I3 = getFormDataInt(r, "blockage_type_I3")
	f.blockage_type_SF = getFormDataInt(r, "blockage_type_SF")
	f.blockage_type_01 = getFormDataInt(r, "blockage_type_O1")
	f.blockage_type_02 = getFormDataInt(r, "blockage_type_O2")
	f.blockage_type_03 = getFormDataInt(r, "blockage_type_O3")
	f.erosion_Z1 = getFormDataInt(r, "Erosion_Z1")
	f.erosion_Z2 = getFormDataInt(r, "Erosion_Z2")
	f.erosion_Z3 = getFormDataInt(r, "Erosion_Z3")
	f.hydroconc = getFormDataString(r, "HydroConc")
	f.zone1_len = getFormDataFloat(r, "zone1-length")
	f.section_info_A = retrieveSectionData(r, "A")
	f.section_info_B = retrieveSectionData(r, "B")
	f.section_info_C = retrieveSectionData(r, "C")
	f.substrate_observations = getFormDataString(r, "substrate-observations")
	f.mulch_type_1A = getFormDataInt(r, "mulch-type-1A")
	f.mulch_type_1B = getFormDataInt(r, "mulch-type-1B")
	f.mulch_type_1C = getFormDataInt(r, "mulch-type-1C")
	f.mulch_type_2 = getFormDataInt(r, "mulch-type-2")
	f.mulch_type_3 = getFormDataInt(r, "mulch-type-3")
	f.mulch_depth_1A = getFormDataInt(r, "mulch-depth-1A")
	f.mulch_depth_1B = getFormDataInt(r, "mulch-depth-1B")
	f.mulch_depth_1C = getFormDataInt(r, "mulch-depth-1C")
	f.mulch_depth_2 = getFormDataInt(r, "mulch-depth-2")
	f.mulch_depth_3 = getFormDataInt(r, "mulch-depth-3")
	f.zone_info_1 = retrieveZoneData(r, "1")
	f.zone_info_2 = retrieveZoneData(r, "2")
	f.zone_info_3 = retrieveZoneData(r, "3")
	f.vegetation_observations = getFormDataString(r, "vegetation-observations")
	f.visible_to_public = getFormDataInt(r, "visible-to-public")
	f.aesthetically_pleasing = getFormDataInt(r, "aesthetically-pleasing")
	f.well_maintained = getFormDataInt(r, "well-maintained")
	f.educational_signage = getFormDataInt(r, "signage-radio")
	f.other_observations = getFormDataString(r, "other-observations")
	f.end_time = getFormDataString(r, "end-time")

	return nil
}

func retrieveSectionData(r *http.Request, sectionLetter string) sectionInfo {
	section := sectionInfo{
		water_depth:     getFormDataFloat(r, "water-depth-"+sectionLetter),
		siltation_depth: getFormDataInt(r, "siltation-depth-"+sectionLetter),
		liner_present:   getFormDataInt(r, "liner-present-"+sectionLetter),
		liner_depth:     getFormDataFloat(r, "liner-depth-"+sectionLetter),
		filter_present:  getFormDataInt(r, "filter-present-"+sectionLetter),
		filter_depth:    getFormDataFloat(r, "filter-depth-"+sectionLetter),
		native_depth:    getFormDataFloat(r, "native-depth-"+sectionLetter),
		compacted_soil:  getFormDataInt(r, "compacted-soil-"+sectionLetter),
		garden_soil:     getFormDataInt(r, "garden-soil-"+sectionLetter),
		native_soil:     getFormDataInt(r, "native-soil-"+sectionLetter),
	}
	return section
}

func retrieveZoneData(r *http.Request, zoneNumber string) zoneInfo {
	zone := zoneInfo{
		mulch_coverage:           getFormDataInt(r, "mulch-coverage-"+zoneNumber),
		bare_ground_coverage:     getFormDataInt(r, "ground-coverage-"+zoneNumber),
		pea_gravel_coverage:      getFormDataInt(r, "gravel-coverage-"+zoneNumber),
		drain_rock_coverage:      getFormDataInt(r, "drain-coverage-"+zoneNumber),
		rock_lower_coverage:      getFormDataInt(r, "mid-coverage-"+zoneNumber),
		rock_upper_coverage:      getFormDataInt(r, "low-coverage-"+zoneNumber),
		all_vegetation_coverage:  getFormDataInt(r, "vegetation-coverage-"+zoneNumber),
		problem_plants_coverage:  getFormDataInt(r, "problem-coverage-"+zoneNumber),
		non_target_weed_coverage: getFormDataInt(r, "weeds-coverage-"+zoneNumber),
		deciduous_coverage:       getFormDataInt(r, "deciduous-coverage-"+zoneNumber),
		evergreen_coverage:       getFormDataInt(r, "evergreen-coverage-"+zoneNumber),
		herbaceous_coverage:      getFormDataInt(r, "herbaceous-coverage-"+zoneNumber),
		ground_cover_coverage:    getFormDataInt(r, "soil-coverage-"+zoneNumber),
		problem_plants_vigor:     getFormDataInt(r, "problem-vigor-"+zoneNumber),
		non_target_weed_vigor:    getFormDataInt(r, "weeds-vigor-"+zoneNumber),
		deciduous_vigor:          getFormDataInt(r, "deciduous-vigor-"+zoneNumber),
		evergreen_vigor:          getFormDataInt(r, "evergreen-vigor-"+zoneNumber),
		herbaceous_vigor:         getFormDataInt(r, "herbaceous-vigor-"+zoneNumber),
		ground_cover_vigor:       getFormDataInt(r, "soil-vigor-"+zoneNumber),
	}
	return zone
}

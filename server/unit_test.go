/*
This file contains unit tests for all elements of the 'main' package.
This includes tests for functions in server.go, submit.go, and download.go.
*/
package main

import (
	"net/http"
	"os"
	"testing"
)

var (
	emptyForm = formData{}

	section_info_ABC = sectionInfo{
		water_depth:     "14",
		siltation_depth: 1,
		liner_present:   1,
		liner_depth:     "1.25",
		filter_present:  1,
		filter_depth:    "2",
		native_depth:    "10",
		compacted_soil:  1,
		garden_soil:     1,
		native_soil:     1,
	}

	zone_info_123 = zoneInfo{
		mulch_coverage:           1,
		bare_ground_coverage:     1,
		pea_gravel_coverage:      1,
		drain_rock_coverage:      1,
		rock_lower_coverage:      1,
		rock_upper_coverage:      1,
		all_vegetation_coverage:  1,
		problem_plants_coverage:  1,
		non_target_weed_coverage: 1,
		deciduous_coverage:       1,
		evergreen_coverage:       1,
		herbaceous_coverage:      1,
		ground_cover_coverage:    1,
		problem_plants_vigor:     1,
		non_target_weed_vigor:    1,
		deciduous_vigor:          1,
		evergreen_vigor:          1,
		herbaceous_vigor:         1,
		ground_cover_vigor:       1,
	}

	// TODO: fill out more valid data
	validForm = formData{
		name1:     "bryan smith",
		numEmail1: "bsmith@gmail.com",
		name2:     "Bob Brown",
		numEmail2: "Brownie2112@gmail.com",
		name3:     "Sam Smalls",
		numEmail3: "smallntiny@gmail.com",
		name4:     "Stanley Yelnats",
		numEmail4: "c4V3m4N@gmail.com",

		site_name:   "The Best Site",
		survey_date: "04/09/2021",
		start_time:  "4:22 PM",
		group_code:  "ALPHA99",
		address:     "3220 NE 141 Street",
		city:        "Pullman, WA",
		county:      "Whitman",
		lat:         "46.727",
		lon:         "-117.162",
		//Stings not floats
		today_rain:          "4.5",
		yesterday_rain:      "1.2",
		twodays_rain:        "2.4",
		site_type:           1,
		site_age:            1,
		age_source_radio:    1,
		age_source_desc:     "Test age description",
		water_sources:       [8]int{1, 0, 1, 0, 1, 0, 1, 0},
		overflow_1_radio:    1,
		overflow_2_radio:    1,
		overflow_3_radio:    1,
		percent_blockage_I1: 1,
		percent_blockage_I2: 1,
		percent_blockage_I3: 1,
		percent_blockage_SF: 1,
		percent_blockage_01: 1,
		percent_blockage_02: 1,
		percent_blockage_03: 1,
		blockage_type_I1:    1,
		blockage_type_I2:    1,
		blockage_type_I3:    1,
		blockage_type_SF:    1,
		blockage_type_01:    1,
		blockage_type_02:    1,
		blockage_type_03:    1,
		erosion_Z1:          1,
		erosion_Z2:          1,
		erosion_Z3:          1,
		hydroconc:           "Test Hydrology Concerns",
		zone1_len:           "10.5",

		section_info_A: section_info_ABC,
		section_info_B: section_info_ABC,
		section_info_C: section_info_ABC,

		substrate_observations: "Test Substrate Observations",
		mulch_type_1A:          1,
		mulch_type_1B:          1,
		mulch_type_1C:          1,
		mulch_type_2:           1,
		mulch_type_3:           1,
		mulch_depth_1A:         1,
		mulch_depth_1B:         1,
		mulch_depth_1C:         1,
		mulch_depth_2:          1,
		mulch_depth_3:          1,

		zone_info_1: zone_info_123,
		zone_info_2: zone_info_123,
		zone_info_3: zone_info_123,

		vegetation_observations: "Test vegetation observations",
		visible_to_public:       1,
		aesthetically_pleasing:  1,
		well_maintained:         1,
		educational_signage:     1,
		other_observations:      "Test other observations",
		end_time:                "05:11 PM",
	}

	// TODO: fill with invalid data
	invalidForm = formData{
		name1:     "bryan smith",
		numEmail1: "bsmith@gmail.com",
		name2:     "Bob Brown",
		numEmail2: "Brownie2112@gmail.com",
		name3:     "Sam Smalls",
		numEmail3: "smallntiny@gmail.com",
		name4:     "Stanley Yelnats",
		numEmail4: "c4V3m4N@gmail.com",

		site_name:   "The Best Site",
		survey_date: "04/09/2021",
		start_time:  "4:22 PM",
		group_code:  "ALPHA99",
		address:     "3220 NE 141 Street",
		city:        "Pullman, WA",
		county:      "Whitman",
		lat:         "46.727",
		lon:         "-117.162",
		//Stings not floats
		today_rain:     "4.5",
		yesterday_rain: "1.2",
		twodays_rain:   "2.4",
		site_type:      1,
		site_age:       3,
		//3 is an invalid choice for site_age

		age_source_radio:    1,
		age_source_desc:     "Test age description",
		water_sources:       [8]int{1, 0, 1, 0, 1, 0, 1, 0},
		overflow_1_radio:    1,
		overflow_2_radio:    1,
		overflow_3_radio:    1,
		percent_blockage_I1: 1,
		percent_blockage_I2: 1,
		percent_blockage_I3: 1,
		percent_blockage_SF: 1,
		percent_blockage_01: 1,
		percent_blockage_02: 1,
		percent_blockage_03: 1,
		blockage_type_I1:    1,
		blockage_type_I2:    1,
		blockage_type_I3:    1,
		blockage_type_SF:    1,
		blockage_type_01:    1,
		blockage_type_02:    1,
		blockage_type_03:    1,
		erosion_Z1:          1,
		erosion_Z2:          1,
		erosion_Z3:          1,
		hydroconc:           "Test Hydrology Concerns",
		zone1_len:           "10.5",

		section_info_A: section_info_ABC,
		section_info_B: section_info_ABC,
		section_info_C: section_info_ABC,

		substrate_observations: "Test Substrate Observations",
		mulch_type_1A:          1,
		mulch_type_1B:          1,
		mulch_type_1C:          1,
		mulch_type_2:           1,
		mulch_type_3:           1,
		mulch_depth_1A:         1,
		mulch_depth_1B:         1,
		mulch_depth_1C:         1,
		mulch_depth_2:          1,
		mulch_depth_3:          1,

		zone_info_1: zone_info_123,
		zone_info_2: zone_info_123,
		zone_info_3: zone_info_123,

		vegetation_observations: "Test vegetation observations",
		visible_to_public:       1,
		aesthetically_pleasing:  1,
		well_maintained:         1,
		educational_signage:     1,
		other_observations:      "Test other observations",
		end_time:                "05:11 PM",
	}
)

// Entry point for "go test", allows specifying setup and teardown.
// 'm.Run()' executes all other tests
func TestMain(m *testing.M) {
	// configure and start server
	scanEnvVariables()
	os.Setenv("MYSQL_DATABASE", "test")
	os.Setenv("MYSQL_USER", "user1")
	os.Setenv("MYSQL_PASSWORD", "usbw")
	os.Setenv("SERVER_PORT", "80")
	hostname = "localhost"
	dbConnect()
	configureHandlers()
	go startServer()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRetrieveFormData(t *testing.T) {

}

func TestValidFormData(t *testing.T) {
	if checkValidInputs(&validForm) != nil {
		t.Error("validFormData() improperly identified a valid form as invalid")
	}
	if checkValidInputs(&emptyForm) == nil {
		t.Error("validFormData() improperly identified an empty form as valid")
	}
	if checkValidInputs(&invalidForm) == nil {
		t.Error("validFormData() improperly identified an invalid form as valid")
	}
}

// TODO
func TestSubmitToDB(t *testing.T) {

}

// verify that we can connect to the main webpage using http
func TestHTTPConn(t *testing.T) {
	// make request to port 80 (HTTP)
	_, err := http.Get("http://localhost")
	if err != nil {
		t.Error("failed to connect to server using http with error:", err)
	}
}

/*
This file contains all elements required to allow users to download the MYSQL database as a CSV.
*/
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var (
	simpleAnswer = [4]string{"", "Yes", "No", "Unknown"}
	siteType     = [4]string{"", "Rain Garden", "Bioretention", "Unknown"}
	siteAge      = [6]string{"", "<1 year", "1-3 years", "3-5 years", ">5 years", "Unknown"}
	ageSource    = [3]string{"", "Verifiable source", "Estimate"}
	waterSource  = [8]string{"Rooftop", "Driveway", "Lawn", "Maintained pasture",
		"Residential street, low use parking lot",
		"Moderate use street, high use parking lot",
		"High use street, livestock confinement area",
		"Industrial or other high containment area"}
	percentClass = [9]string{"", "None", "Trace <.5%", ".5-5%", "6-25%", "26-50%",
		"51-75%", "76-95%", ">96%"}
	blockageType    = [7]string{"", "None", "Siltation", "Organic", "Rock", "Trash", "Vegetation"}
	erosionSeverity = [5]string{"", "None", "Minor", "Moderate", "Extensive"}
	siltationDepth  = [5]string{"", "None", "Trace", "<.25'", ">.25'"}
	soilTexture     = [5]string{"", "Sandy", "Silty", "Clayey", "N/A"}
	mulchType       = [5]string{"", "None", "Shredded mulch", "Fine mulch", "Coarse mulch"}
	mulchDepth      = [5]string{"", "None", "Trace <1\"", "1-3\"", ">3\""}
	vigorRank       = [4]string{"", "Poor", "Moderate", "Robust"}
	publicAmenity   = [4]string{"", "Low", "Moderate", "High"}
)

func downloadCSV(w http.ResponseWriter, r *http.Request) error {
	err := retrieveForms(w)
	if err != nil {
		fmt.Println(err) // this will print any error from db stuff
		return err
	}
	return nil
}

func retrieveForms(w http.ResponseWriter) error {
	forms, err := getFormRows()
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	csvWriter := csv.NewWriter(b)
	err = csvWriter.Write(getColumnHeaders())
	if err != nil {
		return err
	}

	for _, form := range forms {
		var water_sources []string
		ids := strings.Split(form.water_source_ids, ",")
		for _, i := range ids {
			if len(i) > 0 {
				index, err := strconv.Atoi(i)
				if err != nil {
					return err
				}
				water_sources = append(water_sources, waterSource[index])
			}
		}
		record := createRecord(form, water_sources)
		err := csvWriter.Write(record)
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Disposition", "attachment; filename=GardenAppData.csv")
	w.Header().Set("Content-Type", "application/CSV")
	_, err = w.Write(b.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func getFormRows() ([]formData, error) {
	query := getQueryString()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forms []formData
	for rows.Next() {
		f := formData{}
		err := rows.Scan(&f.name1, &f.numEmail1, &f.name2, &f.numEmail2,
			&f.name3, &f.numEmail3, &f.name4, &f.numEmail4, &f.site_name,
			&f.survey_date, &f.start_time, &f.group_code, &f.address,
			&f.city, &f.county, &f.lat, &f.lon, &f.sounds_impacts, &f.today_rain,
			&f.yesterday_rain, &f.twodays_rain, &f.site_type, &f.site_age,
			&f.age_source_radio, &f.age_source_desc, &f.water_source_ids, &f.overflow_1_radio,
			&f.overflow_2_radio, &f.overflow_3_radio, &f.percent_blockage_I1,
			&f.percent_blockage_I2, &f.percent_blockage_I3, &f.percent_blockage_SF,
			&f.percent_blockage_01, &f.percent_blockage_02, &f.percent_blockage_03,
			&f.blockage_type_I1, &f.blockage_type_I2, &f.blockage_type_I3,
			&f.blockage_type_SF, &f.blockage_type_01, &f.blockage_type_02,
			&f.blockage_type_03, &f.erosion_Z1, &f.erosion_Z2, &f.erosion_Z3,
			&f.hydroconc, &f.zone1_len, &f.section_len, &f.section_info_A.water_depth,
			&f.section_info_A.siltation_depth, &f.section_info_A.liner_present,
			&f.section_info_A.liner_depth, &f.section_info_A.filter_present,
			&f.section_info_A.filter_depth, &f.section_info_A.native_depth,
			&f.section_info_A.compacted_soil, &f.section_info_A.garden_soil,
			&f.section_info_A.native_soil, &f.section_info_B.water_depth,
			&f.section_info_B.siltation_depth, &f.section_info_B.liner_present,
			&f.section_info_B.liner_depth, &f.section_info_B.filter_present,
			&f.section_info_B.filter_depth, &f.section_info_B.native_depth,
			&f.section_info_B.compacted_soil, &f.section_info_B.garden_soil,
			&f.section_info_B.native_soil, &f.section_info_C.water_depth,
			&f.section_info_C.siltation_depth, &f.section_info_C.liner_present,
			&f.section_info_C.liner_depth, &f.section_info_C.filter_present,
			&f.section_info_C.filter_depth, &f.section_info_C.native_depth,
			&f.section_info_C.compacted_soil, &f.section_info_C.garden_soil,
			&f.section_info_C.native_soil, &f.substrate_observations,
			&f.mulch_type_1A, &f.mulch_type_1B, &f.mulch_type_1C,
			&f.mulch_type_2, &f.mulch_type_3, &f.mulch_depth_1A, &f.mulch_depth_1B,
			&f.mulch_depth_1C, &f.mulch_depth_2, &f.mulch_depth_3,
			&f.zone_info_1.mulch_coverage, &f.zone_info_1.bare_ground_coverage,
			&f.zone_info_1.pea_gravel_coverage, &f.zone_info_1.drain_rock_coverage,
			&f.zone_info_1.rock_lower_coverage, &f.zone_info_1.rock_upper_coverage,
			&f.zone_info_1.all_vegetation_coverage, &f.zone_info_1.problem_plants_coverage,
			&f.zone_info_1.non_target_weed_coverage, &f.zone_info_1.deciduous_coverage,
			&f.zone_info_1.evergreen_coverage, &f.zone_info_1.herbaceous_coverage,
			&f.zone_info_1.ground_cover_coverage, &f.zone_info_1.problem_plants_vigor,
			&f.zone_info_1.non_target_weed_vigor, &f.zone_info_1.deciduous_vigor,
			&f.zone_info_1.evergreen_vigor, &f.zone_info_1.herbaceous_vigor,
			&f.zone_info_1.ground_cover_vigor, &f.zone_info_2.mulch_coverage,
			&f.zone_info_2.bare_ground_coverage, &f.zone_info_2.pea_gravel_coverage,
			&f.zone_info_2.drain_rock_coverage, &f.zone_info_2.rock_lower_coverage,
			&f.zone_info_2.rock_upper_coverage, &f.zone_info_2.all_vegetation_coverage,
			&f.zone_info_2.problem_plants_coverage, &f.zone_info_2.non_target_weed_coverage,
			&f.zone_info_2.deciduous_coverage, &f.zone_info_2.evergreen_coverage,
			&f.zone_info_2.herbaceous_coverage, &f.zone_info_2.ground_cover_coverage,
			&f.zone_info_2.problem_plants_vigor, &f.zone_info_2.non_target_weed_vigor,
			&f.zone_info_2.deciduous_vigor, &f.zone_info_2.evergreen_vigor,
			&f.zone_info_2.herbaceous_vigor, &f.zone_info_2.ground_cover_vigor,
			&f.zone_info_3.mulch_coverage, &f.zone_info_3.bare_ground_coverage,
			&f.zone_info_3.pea_gravel_coverage, &f.zone_info_3.drain_rock_coverage,
			&f.zone_info_3.rock_lower_coverage, &f.zone_info_3.rock_upper_coverage,
			&f.zone_info_3.all_vegetation_coverage, &f.zone_info_3.problem_plants_coverage,
			&f.zone_info_3.non_target_weed_coverage, &f.zone_info_3.deciduous_coverage,
			&f.zone_info_3.evergreen_coverage, &f.zone_info_3.herbaceous_coverage,
			&f.zone_info_3.ground_cover_coverage, &f.zone_info_3.problem_plants_vigor,
			&f.zone_info_3.non_target_weed_vigor, &f.zone_info_3.deciduous_vigor,
			&f.zone_info_3.evergreen_vigor, &f.zone_info_3.herbaceous_vigor,
			&f.zone_info_3.ground_cover_vigor, &f.vegetation_observations,
			&f.visible_to_public, &f.aesthetically_pleasing, &f.well_maintained,
			&f.educational_signage, &f.other_observations, &f.end_time)
		if err != nil {
			return forms, err
		}
		forms = append(forms, f)
	}
	return forms, err
}

func getQueryString() string {
	return "SELECT Name1, Email1, Name2, Email2, Name3, Email3, Name4, Email4, " +
		"SiteName, SurveyDate, COALESCE(StartTime, '') AS StartTime, GroupCode, " +
		"Address, City, County, COALESCE(Latitude, '') AS Latitude, " +
		"COALESCE(Longitude, '') AS Longitude, SoundImpactsId, " +
		"COALESCE(RainfallToday, '') AS RainfallToday, " +
		"COALESCE(RainfallYesterday, '') AS RainfallYesterday, " +
		"COALESCE(RainfallTwoDaysAgo, '') AS RainfallTwoDaysAgo, " +
		"SiteTypeId, SiteAgeId, AgeSourceId, AgeSourceDescription, " +
		"COALESCE(GROUP_CONCAT(FormWaterSource.WaterSourceId), '') AS WaterSources, " +
		"Overflow1Id, Overflow2Id, Overflow3Id, Inflow1BlockagePercentId, " +
		"Inflow2BlockagePercentId, Inflow3BlockagePercentId, " +
		"SheetFlowBlockagePercentId, Overflow1BlockagePercentId, " +
		"Overflow2BlockagePercentId, Overflow3BlockagePercentId, " +
		"Inflow1BlockageTypeId, Inflow2BlockageTypeId, " +
		"Inflow3BlockageTypeId, SheetFlowBlockageTypeId, " +
		"Overflow1BlockageTypeId, Overflow2BlockageTypeId, " +
		"Overflow3BlockageTypeId, Zone1ErosionId, Zone2ErosionId, Zone3ErosionId, " +
		"HydrologyConcerns, COALESCE(Zone1Length, '') AS Zone1Length, " +
		"COALESCE(Zone1Length / 3.0, '') AS SectionLength, " +
		"COALESCE(sectioninfo1A.StandingWaterDepth, '') AS StandingWaterDepth1A, " +
		"sectioninfo1A.SiltationDepthId AS SiltationDepth1A, " +
		"sectioninfo1A.LinerPresent AS LinerPresent1A, " +
		"COALESCE(sectioninfo1A.LinerDepth, '') AS LinerDepth1A, " +
		"sectioninfo1A.FilterFabricPresent AS FilterFabricPresent1A, " +
		"COALESCE(sectioninfo1A.FilterFabricDepth, '') AS FilterFabricDepth1A, " +
		"COALESCE(sectioninfo1A.NativeSoilDepth, '') AS NativeSoilDepth1A, " +
		"COALESCE(sectioninfo1A.CompactedSurfaceSoils, '') AS CompactedSurfaceSoils1A, " +
		"sectioninfo1A.RainGardenMixSoilTextureId AS GardenMixSoilTexture1A, " +
		"sectioninfo1A.NativeTextureId AS NativeSoilTexture1A, " +
		"COALESCE(sectioninfo1B.StandingWaterDepth, '') AS StandingWaterDepth1B, " +
		"sectioninfo1B.SiltationDepthId AS SiltationDepth1B, " +
		"sectioninfo1B.LinerPresent AS LinerPresent1B, " +
		"COALESCE(sectioninfo1B.LinerDepth, '') AS LinerDepth1B, " +
		"sectioninfo1B.FilterFabricPresent AS FilterFabricPresent1B, " +
		"COALESCE(sectioninfo1B.FilterFabricDepth, '') AS FilterFabricDepth1B, " +
		"COALESCE(sectioninfo1B.NativeSoilDepth, '') AS NativeSoilDepth1B, " +
		"COALESCE(sectioninfo1B.CompactedSurfaceSoils, '') AS CompactedSurfaceSoils1B, " +
		"sectioninfo1B.RainGardenMixSoilTextureId AS GardenMixSoilTexture1B, " +
		"sectioninfo1B.NativeTextureId AS NativeSoilTexture1B, " +
		"COALESCE(sectioninfo1C.StandingWaterDepth, '') AS StandingWaterDepth1C, " +
		"sectioninfo1C.SiltationDepthId AS SiltationDepth1C, " +
		"sectioninfo1C.LinerPresent AS LinerPresent1C, " +
		"COALESCE(sectioninfo1C.LinerDepth, '') AS LinerDepth1C, " +
		"sectioninfo1C.FilterFabricPresent AS FilterFabricPresent1C, " +
		"COALESCE(sectioninfo1C.FilterFabricDepth, '') AS FilterFabricDepth1C, " +
		"COALESCE(sectioninfo1C.NativeSoilDepth, '') AS NativeSoilDepth1C, " +
		"COALESCE(sectioninfo1C.CompactedSurfaceSoils, '') AS CompactedSurfaceSoils1C, " +
		"sectioninfo1C.RainGardenMixSoilTextureId AS GardenMixSoilTexture1C, " +
		"sectioninfo1C.NativeTextureId AS NativeSoilTexture1C, " +
		"SubstrateObservations, MulchType1AId, MulchType1BId, MulchType1CId, " +
		"MulchType2Id, MulchType3Id, MulchDepth1AId, MulchDepth1BId, " +
		"MulchDepth1CId, MulchDepth2Id, MulchDepth3Id, " +
		"zoneinfo1.MulchCoverageId AS MulchCoverage1, " +
		"zoneinfo1.BareGroundCoverageId AS GroundCoverage1, " +
		"zoneinfo1.PeaGravelCoverageId AS GravelCoverage1, " +
		"zoneinfo1.DrainRockCoverageId AS DrainRockCoverage1, " +
		"zoneinfo1.Rock2To12InchCoverageId AS LowRockCoverage1, " +
		"zoneinfo1.RockGreaterThan12InchCoverageId AS HighRockCoverage1, " +
		"zoneinfo1.AllVegetationCoverageId AS AllVegetationCoverage1, " +
		"zoneinfo1.ProblemPlantsCoverageId AS ProblemPlantCoverage1, " +
		"zoneinfo1.NonTargetWeedCoverageId AS WeedCoverage1, " +
		"zoneinfo1.DeciduousCoverageId AS DeciduousCoverage1, " +
		"zoneinfo1.EvergreenCoverageId AS EvergreenCoverage1, " +
		"zoneinfo1.HerbaceousCoverageId AS HerbaceousCoverage1, " +
		"zoneinfo1.GroundCoverCoverageId AS GroundCoverCoverage1, " +
		"zoneinfo1.ProblemPlantsVigorId AS PlantVigor1, " +
		"zoneinfo1.NonTargetWeedVigorId AS WeedVigor1, " +
		"zoneinfo1.DeciduousVigorId AS DeciduousVigor1, " +
		"zoneinfo1.EvergreenVigorId AS EvergreenVigor1, " +
		"zoneinfo1.HerbaceousVigorId AS HerbaceousVigor1, " +
		"zoneinfo1.GroundCoverVigorId AS GroundVigor1, " +
		"zoneinfo2.MulchCoverageId AS MulchCoverage2, " +
		"zoneinfo2.BareGroundCoverageId AS GroundCoverage2, " +
		"zoneinfo2.PeaGravelCoverageId AS GravelCoverage2, " +
		"zoneinfo2.DrainRockCoverageId AS DrainRockCoverage2, " +
		"zoneinfo2.Rock2To12InchCoverageId AS LowRockCoverage2, " +
		"zoneinfo2.RockGreaterThan12InchCoverageId AS HighRockCoverage2, " +
		"zoneinfo2.AllVegetationCoverageId AS AllVegetationCoverage2, " +
		"zoneinfo2.ProblemPlantsCoverageId AS ProblemPlantCoverage2, " +
		"zoneinfo2.NonTargetWeedCoverageId AS WeedCoverage2, " +
		"zoneinfo2.DeciduousCoverageId AS DeciduousCoverage2, " +
		"zoneinfo2.EvergreenCoverageId AS EvergreenCoverage2, " +
		"zoneinfo2.HerbaceousCoverageId AS HerbaceousCoverage2, " +
		"zoneinfo2.GroundCoverCoverageId AS GroundCoverCoverage2, " +
		"zoneinfo2.ProblemPlantsVigorId AS PlantVigor2, " +
		"zoneinfo2.NonTargetWeedVigorId AS WeedVigor2, " +
		"zoneinfo2.DeciduousVigorId AS DeciduousVigor2, " +
		"zoneinfo2.EvergreenVigorId AS EvergreenVigor2, " +
		"zoneinfo2.HerbaceousVigorId AS HerbaceousVigor2, " +
		"zoneinfo2.GroundCoverVigorId AS GroundVigor2, " +
		"zoneinfo3.MulchCoverageId AS MulchCoverage3, " +
		"zoneinfo3.BareGroundCoverageId AS GroundCoverage3, " +
		"zoneinfo3.PeaGravelCoverageId AS GravelCoverage3, " +
		"zoneinfo3.DrainRockCoverageId AS DrainRockCoverage3, " +
		"zoneinfo3.Rock2To12InchCoverageId AS LowRockCoverage3, " +
		"zoneinfo3.RockGreaterThan12InchCoverageId AS HighRockCoverage3, " +
		"zoneinfo3.AllVegetationCoverageId AS AllVegetationCoverage3, " +
		"zoneinfo3.ProblemPlantsCoverageId AS ProblemPlantCoverage3, " +
		"zoneinfo3.NonTargetWeedCoverageId AS WeedCoverage3, " +
		"zoneinfo3.DeciduousCoverageId AS DeciduousCoverage3, " +
		"zoneinfo3.EvergreenCoverageId AS EvergreenCoverage3, " +
		"zoneinfo3.HerbaceousCoverageId AS HerbaceousCoverage3, " +
		"zoneinfo3.GroundCoverCoverageId AS GroundCoverCoverage3, " +
		"zoneinfo3.ProblemPlantsVigorId AS PlantVigor3, " +
		"zoneinfo3.NonTargetWeedVigorId AS WeedVigor3, " +
		"zoneinfo3.DeciduousVigorId AS DeciduousVigor3, " +
		"zoneinfo3.EvergreenVigorId AS EvergreenVigor3, " +
		"zoneinfo3.HerbaceousVigorId AS HerbaceousVigor3, " +
		"zoneinfo3.GroundCoverVigorId AS GroundVigor3, " +
		"VegetationObservations, VisibleToPublicId, " +
		"AestheticallyPleasingId, WellMaintainedId, " +
		"COALESCE(EducationalSignage, '') AS EducationalSignage, " +
		"OtherObservations, COALESCE(EndTime, '') AS EndTime " +
		"FROM Form " +
		"JOIN SectionInfo AS sectioninfo1A " +
		"ON Form.Section1AInfoId = sectioninfo1A.SectionInfoId " +
		"JOIN SectionInfo AS sectioninfo1B " +
		"ON Form.Section1BInfoId = sectioninfo1B.SectionInfoId " +
		"JOIN SectionInfo AS sectioninfo1C " +
		"ON Form.Section1CInfoId = sectioninfo1C.SectionInfoId " +
		"JOIN ZoneInfo AS zoneinfo1 " +
		"ON Form.Zone1InfoId = zoneinfo1.ZoneInfoId " +
		"JOIN ZoneInfo AS zoneinfo2 " +
		"ON Form.Zone2InfoId = zoneinfo2.ZoneInfoId " +
		"JOIN ZoneInfo AS zoneinfo3 " +
		"ON Form.Zone3InfoId = zoneinfo3.ZoneInfoId " +
		"LEFT JOIN FormWaterSource " +
		"ON Form.FormId = FormWaterSource.FormId " +
		"GROUP BY Form.FormId;"
}

func getColumnHeaders() []string {
	return []string{"Site Name", "Address", "City", "County", "Latitude", "Longitude",
		"Sound Impacts Id", "Team Name 1", "Email 1", "Team Name 2", "Email 2",
		"Team Name 3", "Email 3", "Team Name 4", "Email 4", "Survey Date",
		"Start Time", "Group Code", "Rainfall Today", "Rainfall Yesterday", "Rainfall Two Days Ago",
		"Type of Site", "Age of Site", "Source of Age", "Description for Source of Age",
		"Contributing Water Source(s)", "Overflow 1", "Overflow 2", "Overflow 3",
		"Inflow 1 Percent Blockage", "Inflow 1 Blockage Type", "Inflow 2 Percent Blockage",
		"Inflow 2 Blockage Type", "Inflow 3 Percent Blockage", "Inflow 3 Blockage Type",
		"Sheetflow Percent Blockage", "Sheetflow Blockage Type", "Overflow 1 Percent Blockage",
		"Overflow 1 Blockage Type", "Overflow 2 Percent Blockage", "Overflow 2 Blockage Type",
		"Overflow 3 Percent Blockage", "Overflow 3 Blockage Type", "Zone 1 Erosion",
		"Zone 2 Erosion", "Zone 3 Erosion", "Hydrology Concerns", "Zone 1 Length",
		"Section Length", "Sect. 1A Standing Water Depth", "Sect. 1A Siltation Depth",
		"Sect. 1A Liner Present", "Sect. 1A Liner Depth", "Sect. 1A Filter Fabric Present",
		"Sect. 1A Filter Fabric Depth", "Sect. 1A Depth to Native Soils", "Sect. 1A Compacted Surface Soils",
		"Sect. 1A Rain Garden Mix Soil Texture", "Sect. 1A Native Soil Texture",
		"Sect. 1B Standing Water Depth", "Sect. 1B Siltation Depth", "Sect. 1B Liner Present",
		"Sect. 1B Liner Depth", "Sect. 1B Filter Fabric Present", "Sect. 1B Filter Fabric Depth", "Sect. 1B Depth to Native Soils",
		"Sect. 1B Compacted Surface Soils", "Sect. 1B Rain Garden Mix Soil Texture",
		"Sect. 1B Native Soil Texture", "Sect. 1C Standing Water Depth", "Sect. 1C Siltation Depth",
		"Sect. 1C Liner Present", "Sect. 1C Liner Depth", "Sect. 1C Filter Fabric Present",
		"Sect. 1C Filter Fabric Depth", "Sect. 1C Depth to Native Soils", "Sect. 1C Compacted Surface Soils",
		"Sect. 1C Rain Garden Mix Soil Texture", "Sect. 1C Native Soil Texture",
		"Substrate Observations", "1A Type of Mulch", "1B Type of Mulch", "1C Type of Mulch",
		"2 Type of Mulch", "3 Type of Mulch", "1A Mulch Depth", "1B Mulch Depth",
		"1C Mulch Depth", "2 Mulch Depth", "3 Mulch Depth", "Zone 1 Mulch Coverage",
		"Zone 1 Bare Ground Coverage", "Zone 1 Pea Gravel Coverage", "Zone 1 Drain Rock Coverage",
		"Zone 1 2-12\" Rock Coverage", "Zone 1 >12\" Rock/Log Coverage",
		"Zone 1 All Vegetation Coverage", "Zone 1 Target Problem Plants Coverage",
		"Zone 1 Target Problem Plants Vigor", "Zone 1 Non-Target Weeds Coverage",
		"Zone 1 Non-Target Weeds Vigor", "Zone 1 Deciduous Shrubs/Trees Coverage",
		"Zone 1 Deciduous Shrubs/Trees Vigor", "Zone 1 Evergreen Shrubs/Trees Coverage",
		"Zone 1 Evergreen Shrubs/Trees Vigor", "Zone 1 Herbaceous Coverage",
		"Zone 1 Herbaceous Vigor", "Zone 1 Ground Cover Coverage", "Zone 1 Ground Cover Vigor",
		"Zone 2 Mulch Coverage", "Zone 2 Bare Ground Coverage",
		"Zone 2 Pea Gravel Coverage", "Zone 2 Drain Rock Coverage", "Zone 2 2-12\" Rock Coverage",
		"Zone 2 >12\" Rock/Log Coverage", "Zone 2 All Vegetation Coverage",
		"Zone 2 Target Problem Plants Coverage", "Zone 2 Target Problem Plants Vigor",
		"Zone 2 Non-Target Weeds Coverage", "Zone 2 Non-Target Weeds Vigor",
		"Zone 2 Deciduous Shrubs/Trees Coverage", "Zone 2 Deciduous Shrubs/Trees Vigor",
		"Zone 2 Evergreen Shrubs/Trees Coverage", "Zone 2 Evergreen Shrubs/Trees Vigor",
		"Zone 2 Herbaceous Coverage", "Zone 2 Herbaceous Vigor", "Zone 2 Ground Cover Coverage",
		"Zone 2 Ground Cover Vigor", "Zone 3 Mulch Coverage", "Zone 3 Bare Ground Coverage",
		"Zone 3 Pea Gravel Coverage", "Zone 3 Drain Rock Coverage",
		"Zone 3 2-12\" Rock Coverage", "Zone 3 >12\" Rock/Log Coverage",
		"Zone 3 All Vegetation Coverage", "Zone 3 Target Problem Plants Coverage",
		"Zone 3 Target Problem Plants Vigor", "Zone 3 Non-Target Weeds Coverage",
		"Zone 3 Non-Target Weeds Vigor", "Zone 3 Deciduous Shrubs/Trees Coverage",
		"Zone 3 Deciduous Shrubs/Trees Vigor", "Zone 3 Evergreen Shrubs/Trees Coverage",
		"Zone 3 Evergreen Shrubs/Trees Vigor", "Zone 3 Herbaceous Coverage",
		"Zone 3 Herbaceous Vigor", "Zone 3 Ground Cover Coverage", "Zone 3 Ground Cover Vigor",
		"Vegetation Observations", "Site Visibility to Public", "Aesthetically Pleasing",
		"Well Maintained", "Educational Signage", "Other Observations", "End Time"}
}

func createRecord(form formData, water_sources []string) []string {
	return []string{form.site_name, form.address, form.city, form.county,
		form.lat, form.lon, form.sounds_impacts, form.name1, form.numEmail1,
		form.name2, form.numEmail2, form.name3, form.numEmail3, form.name4, form.numEmail4,
		form.survey_date, form.start_time, form.group_code, form.today_rain,
		form.yesterday_rain, form.twodays_rain, siteType[form.site_type], siteAge[form.site_age],
		ageSource[form.age_source_radio], form.age_source_desc, strings.Join(water_sources, " / "), simpleAnswer[form.overflow_1_radio],
		simpleAnswer[form.overflow_2_radio], simpleAnswer[form.overflow_3_radio], percentClass[form.percent_blockage_I1],
		blockageType[form.blockage_type_I1], percentClass[form.percent_blockage_I2], blockageType[form.blockage_type_I2],
		percentClass[form.percent_blockage_I3], blockageType[form.blockage_type_I3], percentClass[form.percent_blockage_SF],
		blockageType[form.blockage_type_SF], percentClass[form.percent_blockage_01], blockageType[form.blockage_type_01],
		percentClass[form.percent_blockage_02], blockageType[form.blockage_type_02], percentClass[form.percent_blockage_03],
		blockageType[form.blockage_type_03], erosionSeverity[form.erosion_Z1], erosionSeverity[form.erosion_Z2], erosionSeverity[form.erosion_Z3],
		form.hydroconc, form.zone1_len, form.section_len, form.section_info_A.water_depth,
		siltationDepth[form.section_info_A.siltation_depth], simpleAnswer[form.section_info_A.liner_present],
		form.section_info_A.liner_depth, simpleAnswer[form.section_info_A.filter_present],
		form.section_info_A.filter_depth, form.section_info_A.native_depth,
		simpleAnswer[form.section_info_A.compacted_soil], soilTexture[form.section_info_A.garden_soil],
		soilTexture[form.section_info_A.native_soil], form.section_info_B.water_depth,
		siltationDepth[form.section_info_B.siltation_depth], simpleAnswer[form.section_info_B.liner_present],
		form.section_info_B.liner_depth, simpleAnswer[form.section_info_B.filter_present],
		form.section_info_B.filter_depth, form.section_info_B.native_depth,
		simpleAnswer[form.section_info_B.compacted_soil], soilTexture[form.section_info_B.garden_soil],
		soilTexture[form.section_info_B.native_soil], form.section_info_C.water_depth,
		siltationDepth[form.section_info_C.siltation_depth], simpleAnswer[form.section_info_C.liner_present],
		form.section_info_C.liner_depth, simpleAnswer[form.section_info_C.filter_present],
		form.section_info_C.filter_depth, form.section_info_C.native_depth,
		simpleAnswer[form.section_info_C.compacted_soil], soilTexture[form.section_info_C.garden_soil],
		soilTexture[form.section_info_C.native_soil], form.substrate_observations,
		mulchType[form.mulch_type_1A], mulchType[form.mulch_type_1B], mulchType[form.mulch_type_1C],
		mulchType[form.mulch_type_2], mulchType[form.mulch_type_3], mulchDepth[form.mulch_depth_1A],
		mulchDepth[form.mulch_depth_1B], mulchDepth[form.mulch_depth_1C], mulchDepth[form.mulch_depth_2], mulchDepth[form.mulch_depth_3],
		percentClass[form.zone_info_1.mulch_coverage], percentClass[form.zone_info_1.bare_ground_coverage],
		percentClass[form.zone_info_1.pea_gravel_coverage], percentClass[form.zone_info_1.drain_rock_coverage],
		percentClass[form.zone_info_1.rock_lower_coverage], percentClass[form.zone_info_1.rock_upper_coverage],
		percentClass[form.zone_info_1.all_vegetation_coverage], percentClass[form.zone_info_1.problem_plants_coverage],
		percentClass[form.zone_info_1.non_target_weed_coverage], percentClass[form.zone_info_1.deciduous_coverage],
		percentClass[form.zone_info_1.evergreen_coverage], percentClass[form.zone_info_1.herbaceous_coverage],
		percentClass[form.zone_info_1.ground_cover_coverage], vigorRank[form.zone_info_1.problem_plants_vigor],
		vigorRank[form.zone_info_1.non_target_weed_vigor], vigorRank[form.zone_info_1.deciduous_vigor],
		vigorRank[form.zone_info_1.evergreen_vigor], vigorRank[form.zone_info_1.herbaceous_vigor],
		vigorRank[form.zone_info_1.ground_cover_vigor], percentClass[form.zone_info_2.mulch_coverage],
		percentClass[form.zone_info_2.bare_ground_coverage], percentClass[form.zone_info_2.pea_gravel_coverage],
		percentClass[form.zone_info_2.drain_rock_coverage], percentClass[form.zone_info_2.rock_lower_coverage],
		percentClass[form.zone_info_2.rock_upper_coverage], percentClass[form.zone_info_2.all_vegetation_coverage],
		percentClass[form.zone_info_2.problem_plants_coverage], percentClass[form.zone_info_2.non_target_weed_coverage],
		percentClass[form.zone_info_2.deciduous_coverage], percentClass[form.zone_info_2.evergreen_coverage],
		percentClass[form.zone_info_2.herbaceous_coverage], percentClass[form.zone_info_2.ground_cover_coverage],
		vigorRank[form.zone_info_2.problem_plants_vigor], vigorRank[form.zone_info_2.non_target_weed_vigor],
		vigorRank[form.zone_info_2.deciduous_vigor], vigorRank[form.zone_info_2.evergreen_vigor],
		vigorRank[form.zone_info_2.herbaceous_vigor], vigorRank[form.zone_info_2.ground_cover_vigor],
		percentClass[form.zone_info_3.mulch_coverage], percentClass[form.zone_info_3.bare_ground_coverage],
		percentClass[form.zone_info_3.pea_gravel_coverage], percentClass[form.zone_info_3.drain_rock_coverage],
		percentClass[form.zone_info_3.rock_lower_coverage], percentClass[form.zone_info_3.rock_upper_coverage],
		percentClass[form.zone_info_3.all_vegetation_coverage], percentClass[form.zone_info_3.problem_plants_coverage],
		percentClass[form.zone_info_3.non_target_weed_coverage], percentClass[form.zone_info_3.deciduous_coverage],
		percentClass[form.zone_info_3.evergreen_coverage], percentClass[form.zone_info_3.herbaceous_coverage],
		percentClass[form.zone_info_3.ground_cover_coverage], vigorRank[form.zone_info_3.problem_plants_vigor],
		vigorRank[form.zone_info_3.non_target_weed_vigor], vigorRank[form.zone_info_3.deciduous_vigor],
		vigorRank[form.zone_info_3.evergreen_vigor], vigorRank[form.zone_info_3.herbaceous_vigor],
		vigorRank[form.zone_info_3.ground_cover_vigor], form.vegetation_observations,
		publicAmenity[form.visible_to_public], publicAmenity[form.aesthetically_pleasing], publicAmenity[form.well_maintained],
		simpleAnswer[form.educational_signage], form.other_observations, form.end_time}
}

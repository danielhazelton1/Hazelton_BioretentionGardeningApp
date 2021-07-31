CREATE TABLE IF NOT EXISTS SimpleAnswer (
  AnswerId smallint PRIMARY KEY,
  AnswerValue varchar(10) UNIQUE
);
INSERT INTO SimpleAnswer
VALUES (0, ""),
  (1, "Yes"),
  (2, "No"),
  (3, "Unknown");

CREATE TABLE IF NOT EXISTS SiteType (
  SiteTypeId smallint PRIMARY KEY,
  SiteTypeName varchar(25) UNIQUE
);
INSERT INTO SiteType
VALUES (0, ""),
  (1, "Rain Garden"),
  (2, "Bioretention"),
  (3, "Unknown");

CREATE TABLE IF NOT EXISTS SiteAge (
  SiteAgeId smallint PRIMARY KEY,
  SiteAgeName varchar(10) UNIQUE
);
INSERT INTO SiteAge
VALUES (0, ""),
  (1, "<1 year"),
  (2, "1-3 years"),
  (3, "3-5 years"),
  (4, ">5 years"),
  (5, "Unknown");

CREATE TABLE IF NOT EXISTS AgeSource (
  AgeSourceId smallint PRIMARY KEY,
  AgeSourceName varchar(20) UNIQUE
);
INSERT INTO AgeSource
VALUES (0, ""),
  (1, "Verifiable source"),
  (2, "Estimate");

CREATE TABLE IF NOT EXISTS WaterSource (
  WaterSourceId smallint PRIMARY KEY,
  WaterSourceName varchar(50) UNIQUE
);
INSERT INTO WaterSource
VALUES
  (0, "Rooftop"),
  (1, "Driveway"),
  (2, "Lawn"),
  (3, "Maintained pasture"),
  (4, "Residential street, low use parking lot"),
  (5, "Moderate use street, high use parking lot"),
  (6, "High use street, livestock confinement area"),
  (7, "Industrial or other high containment area");

CREATE TABLE IF NOT EXISTS PercentClass (
  PercentClassId smallint PRIMARY KEY,
  PercentClassName varchar(10) UNIQUE
);
INSERT INTO PercentClass
VALUES (0, ""),
  (1, "None"),
  (2, "Trace <.5%"),
  (3, ".5-5%"),
  (4, "6-25%"),
  (5, "26-50%"),
  (6, "51-75%"),
  (7, "76-95%"),
  (8, ">96%");

CREATE TABLE IF NOT EXISTS BlockageType (
  BlockageTypeId smallint PRIMARY KEY,
  BlockageTypeName varchar(25) UNIQUE
);
INSERT INTO BlockageType
VALUES (0, ""),
  (1, "None"),
  (2, "Siltation"),
  (3, "Organic"),
  (4, "Rock"),
  (5, "Trash"),
  (6, "Vegetation");

CREATE TABLE IF NOT EXISTS ErosionSeverity (
  ErosionSeverityId smallint PRIMARY KEY,
  ErosionSeverityName varchar(15) UNIQUE
);
INSERT INTO ErosionSeverity
VALUES (0, ""),
  (1, "None"),
  (2, "Minor"),
  (3, "Moderate"),
  (4, "Extensive");

CREATE TABLE IF NOT EXISTS SiltationDepth (
  SiltationDepthId smallint PRIMARY KEY,
  SiltationDepthName varchar(10) UNIQUE
);
INSERT INTO SiltationDepth
VALUES (0, ""),
  (1, "None"),
  (2, "Trace"),
  (3, "<.25\""),
  (4, ">.25\"");

CREATE TABLE IF NOT EXISTS SoilTexture (
  SoilTextureId smallint PRIMARY KEY,
  SoilTextureName varchar(15) UNIQUE
);
INSERT INTO SoilTexture
VALUES (0, ""),
  (1, "Sandy"),
  (2, "Silty"),
  (3, "Clayey"),
  (4, "N/A");

CREATE TABLE IF NOT EXISTS MulchType (
  MulchTypeId smallint PRIMARY KEY,
  MulchTypeName varchar(15) UNIQUE
);
INSERT INTO MulchType
VALUES (0, ""),
  (1, "None"),
  (2, "Shredded Mulch"),
  (3, "Fine Mulch"),
  (4, "Coarse Mulch");

CREATE TABLE IF NOT EXISTS MulchDepth (
  MulchDepthId smallint PRIMARY KEY,
  MulchDepthName varchar(15) UNIQUE
);
INSERT INTO MulchDepth
VALUES (0, ""),
  (1, "None"),
  (2, "Trace <1\""),
  (3, "1-3\""),
  (4, ">3\"");

CREATE TABLE IF NOT EXISTS VigorRank (
  VigorRankId smallint PRIMARY KEY,
  VigorRankName varchar(15) UNIQUE
);
INSERT INTO VigorRank
VALUES (0, ""),
  (1, "Poor"),
  (2, "Moderate"),
  (3, "Robust");

CREATE TABLE IF NOT EXISTS PublicAmenityValue (
  PublicAmenityValueId smallint PRIMARY KEY,
  PublicAmenityValueName varchar(15) UNIQUE
);
INSERT INTO PublicAmenityValue
VALUES (0, ""),
  (1, "Low"),
  (2, "Moderate"),
  (3, "High");

CREATE TABLE IF NOT EXISTS ZoneInfo (
  ZoneInfoId INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  MulchCoverageId SMALLINT,
  BareGroundCoverageId SMALLINT,
  PeaGravelCoverageId SMALLINT,
  DrainRockCoverageId SMALLINT,
  Rock2To12InchCoverageId SMALLINT,
  RockGreaterThan12InchCoverageId SMALLINT,
  AllVegetationCoverageId SMALLINT,
  ProblemPlantsCoverageId SMALLINT,
  NonTargetWeedCoverageId SMALLINT,
  DeciduousCoverageId SMALLINT,
  EvergreenCoverageId SMALLINT,
  HerbaceousCoverageId SMALLINT,
  GroundCoverCoverageId SMALLINT,
  ProblemPlantsVigorId SMALLINT,
  NonTargetWeedVigorId SMALLINT,
  DeciduousVigorId SMALLINT,
  EvergreenVigorId SMALLINT,
  HerbaceousVigorId SMALLINT,
  GroundCoverVigorId SMALLINT,
  CONSTRAINT FK_ZoneInfo_MulchCoverageId FOREIGN KEY (MulchCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_BareGroundCoverageId FOREIGN KEY (BareGroundCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_PeaGravelCoverageId FOREIGN KEY (PeaGravelCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_DrainRockCoverageId FOREIGN KEY (DrainRockCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_Rock2To12InchCoverageId FOREIGN KEY (Rock2To12InchCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_RockGreaterThan12InchCoverageId FOREIGN KEY (RockGreaterThan12InchCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_AllVegetationCoverageId FOREIGN KEY (AllVegetationCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_ProblemPlantsCoverageId FOREIGN KEY (ProblemPlantsCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_NonTargetWeedCoverageId FOREIGN KEY (NonTargetWeedCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_DeciduousCoverageId FOREIGN KEY (DeciduousCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_EvergreenCoverageId FOREIGN KEY (EvergreenCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_HerbaceousCoverageId FOREIGN KEY (HerbaceousCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_GroundCoverCoverageId FOREIGN KEY (GroundCoverCoverageId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_ZoneInfo_ProblemPlantsVigorId FOREIGN KEY (ProblemPlantsVigorId) REFERENCES VigorRank (VigorRankId),
  CONSTRAINT FK_ZoneInfo_NonTargetWeedVigorId FOREIGN KEY (NonTargetWeedVigorId) REFERENCES VigorRank (VigorRankId),
  CONSTRAINT FK_ZoneInfo_DeciduousVigorId FOREIGN KEY (DeciduousVigorId) REFERENCES VigorRank (VigorRankId),
  CONSTRAINT FK_ZoneInfo_EvergreenVigorId FOREIGN KEY (EvergreenVigorId) REFERENCES VigorRank (VigorRankId),
  CONSTRAINT FK_ZoneInfo_HerbaceousVigorId FOREIGN KEY (HerbaceousVigorId) REFERENCES VigorRank (VigorRankId),
  CONSTRAINT FK_ZoneInfo_GroundCoverVigorId FOREIGN KEY (GroundCoverVigorId) REFERENCES VigorRank (VigorRankId)
);

CREATE TABLE IF NOT EXISTS SectionInfo (
  SectionInfoId INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  StandingWaterDepth decimal(6, 3),
  SiltationDepthId smallint,
  LinerPresent SMALLINT,
  LinerDepth DECIMAL(6, 3),
  FilterFabricPresent SMALLINT,
  FilterFabricDepth DECIMAL(6, 3),
  NativeSoilDepth DECIMAL(6, 3),
  CompactedSurfaceSoils SMALLINT,
  RainGardenMixSoilTextureId SMALLINT,
  NativeTextureId SMALLINT,
  CONSTRAINT FK_SectionInfo_SiltationDepthId FOREIGN KEY (SiltationDepthId) REFERENCES SiltationDepth (SiltationDepthId),
  CONSTRAINT FK_SectionInfo_LinerPresent FOREIGN KEY (LinerPresent) REFERENCES SimpleAnswer (AnswerId),
  CONSTRAINT FK_SectionInfo_FilterFabricPresent FOREIGN KEY (FilterFabricPresent) REFERENCES SimpleAnswer (AnswerId),
  CONSTRAINT FK_SectionInfo_CompactedSurfaceSoils FOREIGN KEY (CompactedSurfaceSoils) REFERENCES SimpleAnswer(AnswerId),
  CONSTRAINT FK_SectionInfo_RainGardenMixSoilTextureId FOREIGN KEY (RainGardenMixSoilTextureId) REFERENCES SoilTexture (SoilTextureId),
  CONSTRAINT FK_SectionInfo_NativeTextureId FOREIGN KEY (NativeTextureId) REFERENCES SoilTexture (SoilTextureId)
);

CREATE TABLE IF NOT EXISTS Form (
  FormId INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  Name1 VARCHAR(50),
  Email1 VARCHAR(50),
  Name2 VARCHAR(50),
  Email2 VARCHAR(50),
  Name3 VARCHAR(50),
  Email3 VARCHAR(50),
  Name4 VARCHAR(50),
  Email4 VARCHAR(50),
  SiteName VARCHAR(50),
  SurveyDate VARCHAR(30),
  StartTime VARCHAR(30),
  GroupCode VARCHAR(50),
  Address VARCHAR(100),
  City VARCHAR(50),
  County VARCHAR(50),
  Latitude DECIMAL(10, 6),
  Longitude DECIMAL(10, 6),
  SoundImpactsId VARCHAR(30),
  RainfallToday DECIMAL(6, 3),
  RainfallYesterday DECIMAL(6, 3),
  RainfallTwoDaysAgo DECIMAL(6, 3),
  SiteTypeId SMALLINT,
  SiteAgeId SMALLINT,
  AgeSourceId SMALLINT,
  AgeSourceDescription VARCHAR(250),
  Overflow1Id SMALLINT,
  Overflow2Id SMALLINT,
  Overflow3Id SMALLINT,
  Inflow1BlockagePercentId SMALLINT,
  Inflow2BlockagePercentId SMALLINT,
  Inflow3BlockagePercentId SMALLINT,
  SheetFlowBlockagePercentId SMALLINT,
  Overflow1BlockagePercentId SMALLINT,
  Overflow2BlockagePercentId SMALLINT,
  Overflow3BlockagePercentId SMALLINT,
  Inflow1BlockageTypeId SMALLINT,
  Inflow2BlockageTypeId SMALLINT,
  Inflow3BlockageTypeId SMALLINT,
  SheetFlowBlockageTypeId SMALLINT,
  Overflow1BlockageTypeId SMALLINT,
  Overflow2BlockageTypeId SMALLINT,
  Overflow3BlockageTypeId SMALLINT,
  Zone1ErosionId SMALLINT,
  Zone2ErosionId SMALLINT,
  Zone3ErosionId SMALLINT,
  HydrologyConcerns VARCHAR(250),
  Zone1Length DECIMAL(8, 3),
  Section1AInfoId INT(6) UNSIGNED,
  Section1BInfoId INT(6) UNSIGNED,
  Section1CInfoId INT(6) UNSIGNED,
  SubstrateObservations VARCHAR(250),
  MulchType1AId SMALLINT,
  MulchType1BId SMALLINT,
  MulchType1CId SMALLINT,
  MulchType2Id SMALLINT,
  MulchType3Id SMALLINT,
  MulchDepth1AId SMALLINT,
  MulchDepth1BId SMALLINT,
  MulchDepth1CId SMALLINT,
  MulchDepth2Id SMALLINT,
  MulchDepth3Id SMALLINT,
  Zone1InfoId INT(6) UNSIGNED,
  Zone2InfoId INT(6) UNSIGNED,
  Zone3InfoId INT(6) UNSIGNED,
  VegetationObservations VARCHAR(250),
  VisibleToPublicId SMALLINT,
  AestheticallyPleasingId SMALLINT,
  WellMaintainedId SMALLINT,
  EducationalSignage SMALLINT,
  OtherObservations VARCHAR(250),
  EndTime VARCHAR(30),
  CONSTRAINT FK_Form_SiteAgeId FOREIGN KEY (SiteAgeId) REFERENCES SiteAge (SiteAgeId),
  CONSTRAINT FK_Form_AgeSourceId FOREIGN KEY (AgeSourceId) REFERENCES AgeSource (AgeSourceId),
  CONSTRAINT FK_Form_SiteTypeId FOREIGN KEY (SiteTypeId) REFERENCES SiteType (SiteTypeId),
  CONSTRAINT FK_Form_Overflow1Id FOREIGN KEY (Overflow1Id) REFERENCES SimpleAnswer (AnswerId),
  CONSTRAINT FK_Form_Overflow2Id FOREIGN KEY (Overflow2Id) REFERENCES SimpleAnswer (AnswerId),
  CONSTRAINT FK_Form_Overflow3Id FOREIGN KEY (Overflow3Id) REFERENCES SimpleAnswer (AnswerId),
  CONSTRAINT FK_Form_Inflow1BlockagePercentId FOREIGN KEY (Inflow1BlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_Inflow2BlockagePercentId FOREIGN KEY (Inflow2BlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_Inflow3BlockagePercentId FOREIGN KEY (Inflow3BlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_SheetFlowBlockagePercentId FOREIGN KEY (SheetFlowBlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_Overflow1BlockagePercentId FOREIGN KEY (Overflow1BlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_Overflow2BlockagePercentId FOREIGN KEY (Overflow2BlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_Overflow3BlockagePercentId FOREIGN KEY (Overflow3BlockagePercentId) REFERENCES PercentClass (PercentClassId),
  CONSTRAINT FK_Form_Inflow1BlockageTypeId FOREIGN KEY (Inflow1BlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_Inflow2BlockageTypeId FOREIGN KEY (Inflow2BlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_Inflow3BlockageTypeId FOREIGN KEY (Inflow3BlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_SheetFlowBlockageTypeId FOREIGN KEY (SheetFlowBlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_Overflow1BlockageTypeId FOREIGN KEY (Overflow1BlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_Overflow2BlockageTypeId FOREIGN KEY (Overflow2BlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_Overflow3BlockageTypeId FOREIGN KEY (Overflow3BlockageTypeId) REFERENCES BlockageType (BlockageTypeId),
  CONSTRAINT FK_Form_Zone1ErosionId FOREIGN KEY (Zone1ErosionId) REFERENCES ErosionSeverity (ErosionSeverityId),
  CONSTRAINT FK_Form_Zone2ErosionId FOREIGN KEY (Zone2ErosionId) REFERENCES ErosionSeverity (ErosionSeverityId),
  CONSTRAINT FK_Form_Zone3ErosionId FOREIGN KEY (Zone3ErosionId) REFERENCES ErosionSeverity (ErosionSeverityId),
  CONSTRAINT FK_Form_Section1AInfoId FOREIGN KEY (Section1AInfoId) REFERENCES SectionInfo (SectionInfoId),
  CONSTRAINT FK_Form_Section1BInfoId FOREIGN KEY (Section1BInfoId) REFERENCES SectionInfo (SectionInfoId),
  CONSTRAINT FK_Form_Section1CInfoId FOREIGN KEY (Section1CInfoId) REFERENCES SectionInfo (SectionInfoId),
  CONSTRAINT FK_Form_MulchType1AId FOREIGN KEY (MulchType1AId) REFERENCES MulchType (MulchTypeId),
  CONSTRAINT FK_Form_MulchType1BId FOREIGN KEY (MulchType1BId) REFERENCES MulchType (MulchTypeId),
  CONSTRAINT FK_Form_MulchType1CId FOREIGN KEY (MulchType1CId) REFERENCES MulchType (MulchTypeId),
  CONSTRAINT FK_Form_MulchType2Id FOREIGN KEY (MulchType2Id) REFERENCES MulchType (MulchTypeId),
  CONSTRAINT FK_Form_MulchType3Id FOREIGN KEY (MulchType3Id) REFERENCES MulchType (MulchTypeId),
  CONSTRAINT FK_Form_MulchDepth1AId FOREIGN KEY (MulchDepth1AId) REFERENCES MulchDepth (MulchDepthId),
  CONSTRAINT FK_Form_MulchDepth1BId FOREIGN KEY (MulchDepth1BId) REFERENCES MulchDepth (MulchDepthId),
  CONSTRAINT FK_Form_MulchDepth1CId FOREIGN KEY (MulchDepth1CId) REFERENCES MulchDepth (MulchDepthId),
  CONSTRAINT FK_Form_MulchDepth2Id FOREIGN KEY (MulchDepth2Id) REFERENCES MulchDepth (MulchDepthId),
  CONSTRAINT FK_Form_MulchDepth3Id FOREIGN KEY (MulchDepth3Id) REFERENCES MulchDepth (MulchDepthId),
  CONSTRAINT FK_Form_Zone1InfoId FOREIGN KEY (Zone1InfoId) REFERENCES ZoneInfo (ZoneInfoId),
  CONSTRAINT FK_Form_Zone2InfoId FOREIGN KEY (Zone2InfoId) REFERENCES ZoneInfo (ZoneInfoId),
  CONSTRAINT FK_Form_Zone3InfoId FOREIGN KEY (Zone3InfoId) REFERENCES ZoneInfo (ZoneInfoId),
  CONSTRAINT FK_Form_VisibleToPublicId FOREIGN KEY (VisibleToPublicId) REFERENCES PublicAmenityValue (PublicAmenityValueId),
  CONSTRAINT FK_Form_AestheticallyPleasingId FOREIGN KEY (AestheticallyPleasingId) REFERENCES PublicAmenityValue (PublicAmenityValueId),
  CONSTRAINT FK_Form_WellMaintainedId FOREIGN KEY (WellMaintainedId) REFERENCES PublicAmenityValue (PublicAmenityValueId),
  CONSTRAINT FK_Form_EducationalSignage FOREIGN KEY (EducationalSignage) REFERENCES SimpleAnswer(AnswerId)
);

CREATE TABLE IF NOT EXISTS FormWaterSource
(
  Id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  FormId INT(6) UNSIGNED,
  WaterSourceId SMALLINT,
  CONSTRAINT FK_FormWaterSource_FormId FOREIGN KEY (FormId) REFERENCES Form (FormId),
  CONSTRAINT FK_FormWaterSource_WaterSourceId FOREIGN KEY (WaterSourceId) REFERENCES WaterSource (WaterSourceId)
);


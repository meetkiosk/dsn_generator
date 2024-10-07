package main

// This software generates valid DSNs in a simple format:
// CODE,'value'

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// TODO(vm): S20.G00.07
// TODO(vm): S20.G00.08
// TODO(vm): S21.G00.12
// TODO(vm): S21.G00.13
// TODO(vm): S21.G00.15
// TODO(vm): S21.G00.16
// TODO(vm): S21.G00.20
// TODO(vm): S21.G00.22
// TODO(vm): S21.G00.23
// TODO(vm): S21.G00.31
// TODO(vm): S21.G00.34
// TODO(vm): S21.G00.41
// TODO(vm): S21.G00.44
// TODO(vm): S21.G00.45
// TODO(vm): S21.G00.50

// Transmission represents the transmission data structure
// French: "Envoi"
type Transmission struct {
	SoftwareName     string `dsn:"S10.G00.00.001"` // Nom du logiciel utilisé
	PublisherName    string `dsn:"S10.G00.00.002"` // Nom de l'éditeur
	SoftwareVersion  string `dsn:"S10.G00.00.003"` // Numéro de version du logiciel utilisé
	PreCheckCode     string `dsn:"S10.G00.00.004"` // Code de conformité en pré-contrôle
	FileType         string `dsn:"S10.G00.00.005"` // Code envoi du fichier d'essai ou réel
	StandardVersion  string `dsn:"S10.G00.00.006"` // Numéro de version de la norme utilisée
	SubmissionPoint  string `dsn:"S10.G00.00.007"` // Point de dépôt
	TransmissionType string `dsn:"S10.G00.00.008"` // Type de l'envoi
}

func GenerateTransmission() Transmission {
	fake := gofakeit.New(0)

	return Transmission{
		SoftwareName:     fake.AppName(),
		PublisherName:    fake.Company(),
		SoftwareVersion:  fake.AppVersion(),
		PreCheckCode:     fake.DigitN(1),
		FileType:         fake.RandomString([]string{"01", "02"}),
		StandardVersion:  "P24V01", // This is typically a fixed value for a given period
		SubmissionPoint:  fake.DigitN(2),
		TransmissionType: fake.DigitN(2),
	}
}

// French: Émetteur
type Sender struct {
	SirenNumber         string `dsn:"S10.G00.01.001"` // French: Siren de l'émetteur de l'envoi
	NicNumber           string `dsn:"S10.G00.01.002"` // French: Nic de l'émetteur de l'envoi
	Name                string `dsn:"S10.G00.01.003"` // French: Nom ou raison sociale de l'émetteur
	StreetAddress       string `dsn:"S10.G00.01.004"` // French: Numéro, extension, nature et libellé de la voie
	PostalCode          string `dsn:"S10.G00.01.005"` // French: Code postal
	City                string `dsn:"S10.G00.01.006"` // French: Localité
	CountryCode         string `dsn:"S10.G00.01.007"` // French: Code pays
	ForeignDistribution string `dsn:"S10.G00.01.008"` // French: Code de distribution à l'étranger
	BuildingComplement  string `dsn:"S10.G00.01.009"` // French: Complément de la localisation de la construction
	DeliveryService     string `dsn:"S10.G00.01.010"` // French: Service de distribution, complément de localisation de la voie
}

// GenerateSender creates a new Sender with random data
func GenerateSender() Sender {
	return Sender{
		SirenNumber:         gofakeit.DigitN(9),
		NicNumber:           gofakeit.DigitN(5),
		Name:                gofakeit.Company(),
		StreetAddress:       gofakeit.Street(),
		PostalCode:          gofakeit.Zip(),
		City:                gofakeit.City(),
		CountryCode:         gofakeit.CountryAbr(),
		ForeignDistribution: gofakeit.Word(),
		BuildingComplement:  gofakeit.Name(),
		DeliveryService:     gofakeit.Word(),
	}
}

// SenderContact represents the contact information for the sender in the DSN
// French: Contact Émetteur
type SenderContact struct {
	CivilityCode string `dsn:"S10.G00.02.001"` // French: Code civilité
	FullName     string `dsn:"S10.G00.02.002"` // French: Nom et prénom de la personne à contacter
	Email        string `dsn:"S10.G00.02.004"` // French: Adresse mél du contact émetteur
	PhoneNumber  string `dsn:"S10.G00.02.005"` // French: Adresse téléphonique
	FaxNumber    string `dsn:"S10.G00.02.006"` // French: Adresse fax
}

// GenerateSenderContact creates a new SenderContact with random data
func GenerateSenderContact() SenderContact {
	codes := []string{"01", "02", "03"} // Example codes, adjust as needed
	return SenderContact{
		CivilityCode: sample(codes),
		FullName:     gofakeit.Name(),
		Email:        gofakeit.Email(),
		PhoneNumber:  gofakeit.Phone(),
		FaxNumber:    gofakeit.Phone(),
	}
}

// Declaration represents the declaration information in the DSN
// French: Déclaration
type Declaration struct {
	Nature                 string    `dsn:"S20.G00.05.001"` // French: Nature de la déclaration
	Type                   string    `dsn:"S20.G00.05.002"` // French: Type de la déclaration
	FractionNumber         string    `dsn:"S20.G00.05.003"` // French: Numéro de fraction de déclaration
	OrderNumber            string    `dsn:"S20.G00.05.004"` // French: Numéro d'ordre de la déclaration
	MainDeclarationMonth   time.Time `dsn:"S20.G00.05.005"` // French: Date du mois principal déclaré
	CancelledDeclarationID string    `dsn:"S20.G00.05.006"` // French: Identifiant de la déclaration annulée ou remplacée
	FileCreationDate       time.Time `dsn:"S20.G00.05.007"` // French: Date de constitution du fichier
	DeclarationField       string    `dsn:"S20.G00.05.008"` // French: Champ de la déclaration
	BusinessID             string    `dsn:"S20.G00.05.009"` // French: Identifiant métier
	Currency               string    `dsn:"S20.G00.05.010"` // French: Devise de la déclaration
	TriggerEventNature     string    `dsn:"S20.G00.05.011"` // French: Nature de l'événement déclencheur du signalement
	LastKnownSIRET         string    `dsn:"S20.G00.05.012"` // French: Dernier SIRET connu pour ancien numéro de contrat
	SubstitutionDSNType    string    `dsn:"S20.G00.05.013"` // French: Type de nature de DSN de substitution
}

// GenerateDeclaration creates a new Declaration with random data
func GenerateDeclaration() Declaration {
	return Declaration{
		Nature:                 sample([]string{"01", "02", "03"}),
		Type:                   sample([]string{"01", "02", "03"}),
		FractionNumber:         gofakeit.DigitN(2),
		OrderNumber:            gofakeit.DigitN(3),
		MainDeclarationMonth:   gofakeit.Date(),
		CancelledDeclarationID: gofakeit.UUID(),
		FileCreationDate:       gofakeit.Date(),
		DeclarationField:       sample([]string{"01", "02", "03"}),
		BusinessID:             gofakeit.UUID(),
		Currency:               "EUR", // Assuming Euro is the default currency
		TriggerEventNature:     sample([]string{"01", "02", "03"}),
		LastKnownSIRET:         gofakeit.DigitN(14),
		SubstitutionDSNType:    sample([]string{"01", "02", "03"}),
	}
}

// Company represents the company information in the DSN
// French: Entreprise
type Company struct {
	SIREN                   string `dsn:"S21.G00.06.001"` // French: SIREN
	HeadquartersNIC         string `dsn:"S21.G00.06.002"` // French: NIC du siège
	APENCode                string `dsn:"S21.G00.06.003"` // French: Code APEN
	StreetAddress           string `dsn:"S21.G00.06.004"` // French: Numéro, extension, nature et libellé de la voie
	PostalCode              string `dsn:"S21.G00.06.005"` // French: Code postal
	City                    string `dsn:"S21.G00.06.006"` // French: Localité
	BuildingComplement      string `dsn:"S21.G00.06.007"` // French: Complément de la localisation de la construction
	DeliveryService         string `dsn:"S21.G00.06.008"` // French: Service de distribution, complément de localisation de la voie
	AverageWorkforceOnDec31 int    `dsn:"S21.G00.06.009"` // French: Effectif moyen de l'entreprise au 31 décembre
	CountryCode             string `dsn:"S21.G00.06.010"` // French: Code pays
	ForeignDistribution     string `dsn:"S21.G00.06.011"` // French: Code de distribution à l'étranger
	CompanyLocation         string `dsn:"S21.G00.06.012"` // French: Implantation de l'entreprise
	CollectiveAgreementCode string `dsn:"S21.G00.06.015"` // French: Code convention collective applicable
}

// GenerateCompany creates a new Company with random data
func GenerateCompany() Company {
	return Company{
		SIREN:                   gofakeit.DigitN(9),
		HeadquartersNIC:         gofakeit.DigitN(5),
		APENCode:                generateAPENCode(),
		StreetAddress:           gofakeit.Street(),
		PostalCode:              gofakeit.Zip(),
		City:                    gofakeit.City(),
		BuildingComplement:      gofakeit.Name(),
		DeliveryService:         gofakeit.Word(),
		AverageWorkforceOnDec31: gofakeit.Number(1, 10000),
		CountryCode:             gofakeit.CountryAbr(),
		ForeignDistribution:     gofakeit.Word(),
		CompanyLocation:         sample([]string{"01", "02", "03"}),
		CollectiveAgreementCode: generateCollectiveAgreementCode(),
	}
}

// Helper functions to generate specific types of data
func generateAPENCode() string {
	// APEN code format: 4 digits + 1 letter
	return gofakeit.DigitN(4) + gofakeit.Letter()
}

func generateCollectiveAgreementCode() string {
	// Example format: 4 digits
	return gofakeit.DigitN(4)
}

// Establishment represents the establishment information in the DSN
// French: Établissement
type Establishment struct {
	NIC                         string     `dsn:"S21.G00.11.001"` // French: NIC
	APETCode                    string     `dsn:"S21.G00.11.002"` // French: Code APET
	StreetAddress               string     `dsn:"S21.G00.11.003"` // French: Numéro, extension, nature et libellé de la voie
	PostalCode                  string     `dsn:"S21.G00.11.004"` // French: Code postal
	City                        string     `dsn:"S21.G00.11.005"` // French: Localité
	BuildingComplement          string     `dsn:"S21.G00.11.006"` // French: Complément de la localisation de la construction
	DeliveryService             string     `dsn:"S21.G00.11.007"` // French: Service de distribution, complément de localisation de la voie
	WorkforceAtEndOfPeriod      int        `dsn:"S21.G00.11.008"` // French: Effectif de fin de période déclarée de l'établissement
	ExpatRemunerationType       string     `dsn:"S21.G00.11.009"` // French: Type de rémunération soumise à contributions d'Assurance chômage pour expatriés
	CountryCode                 string     `dsn:"S21.G00.11.015"` // French: Code pays
	ForeignDistribution         string     `dsn:"S21.G00.11.016"` // French: Code de distribution à l'étranger
	EmployerLegalNature         string     `dsn:"S21.G00.11.017"` // French: Nature juridique de l'employeur
	TESECEAJoinDate             *time.Time `dsn:"S21.G00.11.019"` // French: Date d'effet de l'adhésion au dispositif TESE/CEA
	TESECEAExitDate             *time.Time `dsn:"S21.G00.11.020"` // French: Date d'effet de la sortie du dispositif TESE/CEA
	MainCollectiveAgreementCode string     `dsn:"S21.G00.11.022"` // French: Code convention collective principale
	SkillsOperator              string     `dsn:"S21.G00.11.023"` // French: Opérateur de compétences (OPCO)
	DSNExitRequest              string     `dsn:"S21.G00.11.024"` // French: Demande de sortie de la DSN
}

// GenerateEstablishment creates a new Establishment with random data
func GenerateEstablishment() Establishment {
	joinDate := gofakeit.Date()
	exitDate := gofakeit.DateRange(joinDate, time.Now())

	return Establishment{
		NIC:                         gofakeit.DigitN(5),
		APETCode:                    generateAPETCode(),
		StreetAddress:               gofakeit.Street(),
		PostalCode:                  gofakeit.Zip(),
		City:                        gofakeit.City(),
		BuildingComplement:          gofakeit.Name(),
		DeliveryService:             gofakeit.Word(),
		WorkforceAtEndOfPeriod:      gofakeit.Number(1, 1000),
		ExpatRemunerationType:       sample([]string{"01", "02", "03"}),
		CountryCode:                 gofakeit.CountryAbr(),
		ForeignDistribution:         gofakeit.Word(),
		EmployerLegalNature:         sample([]string{"01", "02", "03"}),
		TESECEAJoinDate:             &joinDate,
		TESECEAExitDate:             &exitDate,
		MainCollectiveAgreementCode: generateCollectiveAgreementCode(),
		SkillsOperator:              sample([]string{"OPCO1", "OPCO2", "OPCO3"}),
		DSNExitRequest:              sample([]string{"OPCO1", "OPCO2", "OPCO3"}),
	}
}

// Helper functions to generate specific types of data
func generateAPETCode() string {
	// APET code format: 4 digits + 1 letter
	return gofakeit.DigitN(4) + gofakeit.Letter()
}

// Individual represents the individual information in the DSN
// French: Individu
type Individual struct {
	NIR                            string    `dsn:"S21.G00.30.001"` // French: Numéro d'inscription au répertoire
	LastName                       string    `dsn:"S21.G00.30.002"` // French: Nom de famille
	UsageName                      string    `dsn:"S21.G00.30.003"` // French: Nom d'usage
	FirstNames                     string    `dsn:"S21.G00.30.004"` // French: Prénoms
	Gender                         string    `dsn:"S21.G00.30.005"` // French: Sexe
	BirthDate                      time.Time `dsn:"S21.G00.30.006"` // French: Date de naissance
	BirthPlace                     string    `dsn:"S21.G00.30.007"` // French: Lieu de naissance
	StreetAddress                  string    `dsn:"S21.G00.30.008"` // French: Numéro, extension, nature et libellé de la voie
	PostalCode                     string    `dsn:"S21.G00.30.009"` // French: Code postal
	City                           string    `dsn:"S21.G00.30.010"` // French: Localité
	CountryCode                    string    `dsn:"S21.G00.30.011"` // French: Code pays
	ForeignDistribution            string    `dsn:"S21.G00.30.012"` // French: Code de distribution à l'étranger
	EUCodification                 string    `dsn:"S21.G00.30.013"` // French: Codification UE
	BirthDepartmentCode            string    `dsn:"S21.G00.30.014"` // French: Code département de naissance
	BirthCountryCode               string    `dsn:"S21.G00.30.015"` // French: Code pays de naissance
	BuildingComplement             string    `dsn:"S21.G00.30.016"` // French: Complément de la localisation de la construction
	DeliveryService                string    `dsn:"S21.G00.30.017"` // French: Service de distribution, complément de localisation de la voie
	Email                          string    `dsn:"S21.G00.30.018"` // French: Adresse mél
	CompanyID                      string    `dsn:"S21.G00.30.019"` // French: Matricule de l'individu dans l'entreprise
	TemporaryTechnicalID           string    `dsn:"S21.G00.30.020"` // French: Numéro technique temporaire
	ForeignTaxStatus               string    `dsn:"S21.G00.30.022"` // French: Statut à l'étranger au sens fiscal
	RetirementEmploymentCumulation string    `dsn:"S21.G00.30.023"` // French: Cumul emploi retraite
	HighestEducationLevel          string    `dsn:"S21.G00.30.024"` // French: Niveau de formation le plus élevé obtenu par l'individu
	CurrentDiplomaLevel            string    `dsn:"S21.G00.30.025"` // French: Niveau de diplôme préparé par l'individu
	BirthCountryName               string    `dsn:"S21.G00.30.029"` // French: Libellé du pays de naissance
}

// GenerateIndividual creates a new Individual with random data
func GenerateIndividual() Individual {
	gender := generateGender()
	birthCountry := gofakeit.Country()

	return Individual{
		NIR:                            generateNIR(gender),
		LastName:                       gofakeit.LastName(),
		UsageName:                      gofakeit.LastName(),
		FirstNames:                     gofakeit.FirstName(),
		Gender:                         gender,
		BirthDate:                      gofakeit.DateRange(time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2005, 12, 31, 0, 0, 0, 0, time.UTC)),
		BirthPlace:                     gofakeit.City(),
		StreetAddress:                  gofakeit.Street(),
		PostalCode:                     gofakeit.Zip(),
		City:                           gofakeit.City(),
		CountryCode:                    gofakeit.CountryAbr(),
		ForeignDistribution:            gofakeit.Word(),
		EUCodification:                 sample([]string{"01", "02", "03"}),
		BirthDepartmentCode:            generateDepartmentCode(),
		BirthCountryCode:               gofakeit.CountryAbr(),
		BuildingComplement:             gofakeit.Name(),
		DeliveryService:                gofakeit.Word(),
		Email:                          gofakeit.Email(),
		CompanyID:                      gofakeit.DigitN(8),
		TemporaryTechnicalID:           gofakeit.UUID(),
		ForeignTaxStatus:               sample([]string{"01", "02", "03"}),
		RetirementEmploymentCumulation: sample([]string{"01", "02", "03"}),
		HighestEducationLevel:          sample([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}),
		CurrentDiplomaLevel:            sample([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}),
		BirthCountryName:               birthCountry,
	}
}

const (
	Male   = "01"
	Female = "02"
)

func generateGender() string {
	genderInt := gofakeit.Number(1, 2)
	return fmt.Sprintf("0%d", genderInt)
}

// Helper functions to generate specific types of data
func generateNIR(gender string) string {
	if gender != Male && gender != Female {
		log.Fatalf("cannot generate NIR: invalid gender %v", gender)
	}

	// NIR format: 13 digits + 2 check digits
	base := gofakeit.DigitN(14)

	sexDigit := strings.TrimPrefix(gender, "0")

	return sexDigit + base
}

func generateDepartmentCode() string {
	// French department codes are 2 or 3 characters
	return gofakeit.DigitN(2)
}

type Contrat struct {
	ContractStartDate                 time.Time `dsn:"S21.G00.40.001"` // Date de début du contrat
	EmployeeStatus                    string    `dsn:"S21.G00.40.002"` // Statut du salarié (conventionnel)
	MandatorySupplementaryPensionCode string    `dsn:"S21.G00.40.003"` // Code statut catégoriel Retraite Complémentaire obligatoire
	OccupationCode                    string    `dsn:"S21.G00.40.004"` // Code profession et catégorie socioprofessionnelle (PCS-ESE)
	OccupationCodeExtension           string    `dsn:"S21.G00.40.005"` // Code complément PCS-ESE
	JobTitle                          string    `dsn:"S21.G00.40.006"` // Libellé de l'emploi
	ContractType                      string    `dsn:"S21.G00.40.007"` // Nature du contrat
	PublicPolicyScheme                string    `dsn:"S21.G00.40.008"` // Dispositif de politique publique et conventionnel
	ContractNumber                    string    `dsn:"S21.G00.40.009"` // Numéro du contrat
	ExpectedEndDate                   time.Time `dsn:"S21.G00.40.010"` // Date de fin prévisionnelle du contrat
	WorkTimeUnit                      string    `dsn:"S21.G00.40.011"` // Unité de mesure de la quotité de travail
	CompanyWorkTimeReference          float64   `dsn:"S21.G00.40.012"` // Quotité de travail de référence de l'entreprise pour la catégorie de salarié
	ContractWorkTime                  float64   `dsn:"S21.G00.40.013"` // Quotité de travail du contrat
	WorkTimeArrangement               string    `dsn:"S21.G00.40.014"` // Modalité d'exercice du temps de travail
	MandatorySchemeContribution       string    `dsn:"S21.G00.40.016"` // Complément de base au régime obligatoire
	CollectiveAgreementCode           string    `dsn:"S21.G00.40.017"` // Code convention collective applicable
	HealthInsuranceScheme             string    `dsn:"S21.G00.40.018"` // Code régime de base risque maladie
	WorkplaceID                       string    `dsn:"S21.G00.40.019"` // Identifiant du lieu de travail
	PensionScheme                     string    `dsn:"S21.G00.40.020"` // Code régime de base risque vieillesse
	HiringReason                      string    `dsn:"S21.G00.40.021"` // Motif de recours
	PaidLeaveScheme                   string    `dsn:"S21.G00.40.022"` // Code caisse professionnelle de congés payés
	SpecificDeductionRate             float64   `dsn:"S21.G00.40.023"` // Taux de déduction forfaitaire spécifique pour frais professionnels
	OverseasWorker                    string    `dsn:"S21.G00.40.024"` // Travailleur à l'étranger au sens du code de la Sécurité Sociale
	DSNExclusionReason                string    `dsn:"S21.G00.40.025"` // Motif d'exclusion DSN
	EmploymentStatus                  string    `dsn:"S21.G00.40.026"` // Statut d'emploi du salarié
	UnemploymentInsuranceAssignment   string    `dsn:"S21.G00.40.027"` // Code affectation Assurance chômage
	PublicEmployerInternalNumber      string    `dsn:"S21.G00.40.028"` // Numéro interne employeur public
	UnemploymentInsuranceManagement   string    `dsn:"S21.G00.40.029"` // Type de gestion de l'Assurance chômage
	AdhesionDate                      time.Time `dsn:"S21.G00.40.030"` // Date d'adhésion
	TerminationDate                   time.Time `dsn:"S21.G00.40.031"` // Date de dénonciation
	ManagementAgreementEffectiveDate  time.Time `dsn:"S21.G00.40.032"` // Date d'effet de la convention de gestion"
	ManagementAgreementNumber         string    `dsn:"S21.G00.40.033"` // Numéro de convention de gestion
	HealthInsuranceDelegateCode       string    `dsn:"S21.G00.40.035"` // Code délégataire du risque maladie
	MultipleJobsCode                  string    `dsn:"S21.G00.40.036"` // Code emplois multiples
	MultipleEmployersCode             string    `dsn:"S21.G00.40.037"` // Code employeurs multiples
	WorkAccidentRiskScheme            string    `dsn:"S21.G00.40.039"` // Code régime de base risque accident du travail
	WorkAccidentRiskCode              string    `dsn:"S21.G00.40.040"` // Code risque accident du travail
	CollectiveAgreementPosition       string    `dsn:"S21.G00.40.041"` // Positionnement dans la convention collective
	APECITACategoryCode               string    `dsn:"S21.G00.40.042"` // Code statut catégoriel APECITA
	WorkAccidentContributionRate      float64   `dsn:"S21.G00.40.043"` // Taux de cotisation accident du travail
	PartTimeFullTimeContribution      string    `dsn:"S21.G00.40.044"` // Salarié à temps partiel cotisant à temps plein
	TipBasedRemuneration              string    `dsn:"S21.G00.40.045"` // Rémunération au pourboire
	UserEstablishmentID               string    `dsn:"S21.G00.40.046"` // Identifiant de l'établissement utilisateur
	LivePerformanceServiceProviderID  string    `dsn:"S21.G00.40.048"` // Numéro de label « Prestataire de services du spectacle vivant »
	ShowBusinessLicenseNumber         string    `dsn:"S21.G00.40.049"` // Numéro de licence entrepreneur spectacle
	ShowObjectNumber                  string    `dsn:"S21.G00.40.050"` // Numéro objet spectacle
	ShowOrganizerStatus               string    `dsn:"S21.G00.40.051"` // Statut organisateur spectacle
	StatePublicServicePCSESECode      string    `dsn:"S21.G00.40.052"` // [FP] Code complément PCS-ESE pour la fonction publique d'Etat (NNE)
	PositionNature                    string    `dsn:"S21.G00.40.053"` // Nature du poste
	FullTimeWorkReferenceQuota        float64   `dsn:"S21.G00.40.054"` // [FP] Quotité de travail de référence de l'entreprise pour la catégorie  salarié dans l'hypothèse d'un poste à temps complet"`
	PartTimeWorkRate                  float64   `dsn:"S21.G00.40.055"` // Taux de travail à temps partiel
	ServiceCategoryCode               string    `dsn:"S21.G00.40.056"` // Code catégorie de service
	GrossIndex                        int       `dsn:"S21.G00.40.057"` // [FP] Indice brut
	NetIndex                          int       `dsn:"S21.G00.40.058"` // [FP] Indice majoré
	NewIndexBonus                     int       `dsn:"S21.G00.40.059"` // [FP] Nouvelle bonification indiciaire (NBI)
	OriginalGrossIndex                int       `dsn:"S21.G00.40.060"` // [FP] Indice brut d'origine
	Article15ContributionGrossIndex   int       `dsn:"S21.G00.40.061"` // [FP] Indice brut de cotisation dans un emploi supérieur (article 15)
	FormerPublicEmployer              string    `dsn:"S21.G00.40.062"` // [FP] Ancien employeur public
	FormerPublicEmployeeOriginalIndex int       `dsn:"S21.G00.40.063"` // [FP] Indice brut d'origine ancien salarié employeur public
	FirefighterOriginalIndex          int       `dsn:"S21.G00.40.064"` // [FP] Indice brut d'origine sapeur-pompier professionnel (SPP)
	ContractualOriginalSalary         string    `dsn:"S21.G00.40.065"` // [FP] Maintien du traitement d'origine d'un contractuel titulaire
	SecondmentType                    string    `dsn:"S21.G00.40.066"` // [FP] Type de détachement
	NavigationType                    string    `dsn:"S21.G00.40.067"` // Genre de navigation
	ActiveServiceRate                 float64   `dsn:"S21.G00.40.068"` // Taux de service actif
	RemunerationLevel                 string    `dsn:"S21.G00.40.069"` // Niveau de rémunération
	PayGrade                          string    `dsn:"S21.G00.40.070"` // Echelon
	HierarchicalCoefficient           float64   `dsn:"S21.G00.40.071"` // Coefficient hiérarchique
	DisabledWorkerStatus              string    `dsn:"S21.G00.40.072"` // Statut BOETH
	PublicPolicySchemeComplement      string    `dsn:"S21.G00.40.073"` // Complément de dispositif de politique publique
	ExternalAssignmentCase            string    `dsn:"S21.G00.40.074"` // Cas de mise à disposition externe d'un individu de l'établissement
	FinalClassificationCategory       string    `dsn:"S21.G00.40.075"` // Catégorie de classement finale
	MaritimeEngagementContractID      string    `dsn:"S21.G00.40.076"` // Identifiant du contrat d'engagement maritime
	CNIEGCollege                      string    `dsn:"S21.G00.40.077"` // Collège (CNIEG)
	PartTimeWorkArrangement           string    `dsn:"S21.G00.40.078"` // Forme d'aménagement du temps de travail dans le cadre de l'activité partielle
	Grade                             string    `dsn:"S21.G00.40.079"` // Grade
	IndexSupplementaryTreatment       int       `dsn:"S21.G00.40.080"` // [FP] Indice complément de traitement indiciaire (CTI)
	GeographicFINESS                  string    `dsn:"S21.G00.40.081"` // FINESS géographique
}

func GenerateContract() Contrat {
	return Contrat{
		ContractStartDate:                 gofakeit.Date(),
		EmployeeStatus:                    gofakeit.Letter(),
		MandatorySupplementaryPensionCode: gofakeit.LetterN(2),
		OccupationCode:                    gofakeit.DigitN(4),
		OccupationCodeExtension:           gofakeit.DigitN(2),
		JobTitle:                          gofakeit.JobTitle(),
		ContractType:                      gofakeit.DigitN(2),
		PublicPolicyScheme:                gofakeit.DigitN(2),
		ContractNumber:                    gofakeit.DigitN(5),
		ExpectedEndDate:                   gofakeit.Date(),
		WorkTimeUnit:                      gofakeit.DigitN(2),
		CompanyWorkTimeReference:          gofakeit.Float64Range(0, 100),
		ContractWorkTime:                  gofakeit.Float64Range(0, 100),
		WorkTimeArrangement:               gofakeit.DigitN(2),
		MandatorySchemeContribution:       gofakeit.DigitN(2),
		CollectiveAgreementCode:           gofakeit.DigitN(4),
		HealthInsuranceScheme:             gofakeit.DigitN(3),
		WorkplaceID:                       gofakeit.UUID(),
		PensionScheme:                     gofakeit.DigitN(3),
		HiringReason:                      gofakeit.DigitN(2),
		PaidLeaveScheme:                   gofakeit.DigitN(2),
		SpecificDeductionRate:             gofakeit.Float64Range(0, 100),
		OverseasWorker:                    gofakeit.DigitN(2),
		DSNExclusionReason:                gofakeit.DigitN(2),
		EmploymentStatus:                  gofakeit.DigitN(2),
		UnemploymentInsuranceAssignment:   gofakeit.DigitN(2),
		PublicEmployerInternalNumber:      gofakeit.DigitN(10),
		UnemploymentInsuranceManagement:   gofakeit.DigitN(2),
		AdhesionDate:                      gofakeit.Date(),
		TerminationDate:                   gofakeit.Date(),
		ManagementAgreementEffectiveDate:  gofakeit.Date(),
		ManagementAgreementNumber:         gofakeit.DigitN(10),
		HealthInsuranceDelegateCode:       gofakeit.DigitN(3),
		MultipleJobsCode:                  gofakeit.DigitN(2),
		MultipleEmployersCode:             gofakeit.DigitN(2),
		WorkAccidentRiskScheme:            gofakeit.DigitN(3),
		WorkAccidentRiskCode:              gofakeit.DigitN(6),
		CollectiveAgreementPosition:       gofakeit.DigitN(4),
		APECITACategoryCode:               gofakeit.DigitN(2),
		WorkAccidentContributionRate:      gofakeit.Float64Range(0, 100),
		PartTimeFullTimeContribution:      gofakeit.DigitN(2),
		TipBasedRemuneration:              gofakeit.DigitN(2),
		UserEstablishmentID:               gofakeit.UUID(),
		LivePerformanceServiceProviderID:  gofakeit.DigitN(10),
		ShowBusinessLicenseNumber:         gofakeit.DigitN(10),
		ShowObjectNumber:                  gofakeit.DigitN(10),
		ShowOrganizerStatus:               gofakeit.DigitN(2),
		StatePublicServicePCSESECode:      gofakeit.DigitN(4),
		PositionNature:                    gofakeit.DigitN(2),
		FullTimeWorkReferenceQuota:        gofakeit.Float64Range(0, 100),
		PartTimeWorkRate:                  gofakeit.Float64Range(0, 100),
		ServiceCategoryCode:               gofakeit.DigitN(2),
		GrossIndex:                        gofakeit.IntRange(100, 1000),
		NetIndex:                          gofakeit.IntRange(100, 1000),
		NewIndexBonus:                     gofakeit.IntRange(0, 100),
		OriginalGrossIndex:                gofakeit.IntRange(100, 1000),
		Article15ContributionGrossIndex:   gofakeit.IntRange(100, 1000),
		FormerPublicEmployer:              gofakeit.DigitN(2),
		FormerPublicEmployeeOriginalIndex: gofakeit.IntRange(100, 1000),
		FirefighterOriginalIndex:          gofakeit.IntRange(100, 1000),
		ContractualOriginalSalary:         gofakeit.DigitN(2),
		SecondmentType:                    gofakeit.DigitN(2),
		NavigationType:                    gofakeit.DigitN(2),
		ActiveServiceRate:                 gofakeit.Float64Range(0, 100),
		RemunerationLevel:                 gofakeit.DigitN(2),
		PayGrade:                          gofakeit.DigitN(2),
		HierarchicalCoefficient:           gofakeit.Float64Range(1, 10),
		DisabledWorkerStatus:              gofakeit.DigitN(2),
		PublicPolicySchemeComplement:      gofakeit.DigitN(2),
		ExternalAssignmentCase:            gofakeit.DigitN(2),
		FinalClassificationCategory:       gofakeit.DigitN(2),
		MaritimeEngagementContractID:      gofakeit.UUID(),
		CNIEGCollege:                      gofakeit.DigitN(2),
		PartTimeWorkArrangement:           gofakeit.DigitN(2),
		Grade:                             gofakeit.LetterN(3),
		IndexSupplementaryTreatment:       gofakeit.IntRange(0, 100),
		GeographicFINESS:                  gofakeit.DigitN(9),
	}
}

type Payment struct {
	PaymentDate                   time.Time `dsn:"S21.G00.50.001"` // Date de versement
	TaxableNetRemuneration        float64   `dsn:"S21.G00.50.002"` // Rémunération nette fiscale
	PaymentNumber                 string    `dsn:"S21.G00.50.003"` // Numéro de versement
	NetAmountPaid                 float64   `dsn:"S21.G00.50.004"` // Montant net versé
	WithholdingTaxRate            float64   `dsn:"S21.G00.50.006"` // Taux de prélèvement à la source
	WithholdingTaxRateType        string    `dsn:"S21.G00.50.007"` // Type du taux de prélèvement à la source
	WithholdingTaxRateID          string    `dsn:"S21.G00.50.008"` // Identifiant du taux de prélèvement à la source
	WithholdingTaxAmount          float64   `dsn:"S21.G00.50.009"` // Montant de prélèvement à la source
	NonTaxableIncomeAmount        float64   `dsn:"S21.G00.50.011"` // Montant de la part non imposable du revenu
	TaxBaseDeductionAmount        float64   `dsn:"S21.G00.50.012"` // Montant de l'abattement sur la base fiscale (non déduit de la rémunération nette fiscale)
	AmountSubjectToWithholdingTax float64   `dsn:"S21.G00.50.013"` // Montant soumis au PAS
	MonthlyDSNReferenceMonth      string    `dsn:"S21.G00.50.020"` // Mois de la DSN mensuelle de rattachement des éléments déclarés dans le FCTU
}

func GeneratePayment() Payment {
	paymentDate := gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now())

	return Payment{
		PaymentDate:                   paymentDate,
		TaxableNetRemuneration:        gofakeit.Float64Range(1000, 10000),
		PaymentNumber:                 gofakeit.DigitN(5),
		NetAmountPaid:                 gofakeit.Float64Range(1000, 10000),
		WithholdingTaxRate:            gofakeit.Float64Range(0, 100),
		WithholdingTaxRateType:        gofakeit.LetterN(2),
		WithholdingTaxRateID:          gofakeit.UUID(),
		WithholdingTaxAmount:          gofakeit.Float64Range(0, 1000),
		NonTaxableIncomeAmount:        gofakeit.Float64Range(0, 1000),
		TaxBaseDeductionAmount:        gofakeit.Float64Range(0, 1000),
		AmountSubjectToWithholdingTax: gofakeit.Float64Range(1000, 10000),
		MonthlyDSNReferenceMonth:      gofakeit.Date().Format("2006-01"),
	}
}

type Remuneration struct {
	PayPeriodStartDate             time.Time `dsn:"S21.G00.51.001"` // Date de début de période de paie
	PayPeriodEndDate               time.Time `dsn:"S21.G00.51.002"` // Date de fin de période de paie
	ContractNumber                 string    `dsn:"S21.G00.51.010"` // Numéro du contrat
	Type                           string    `dsn:"S21.G00.51.011"` // Type
	NumberOfHours                  int64     `dsn:"S21.G00.51.012"` // Nombre d'heures
	Amount                         float64   `dsn:"S21.G00.51.013"` // Montant
	AdministrativeStatusPayRate    float64   `dsn:"S21.G00.51.014"` // [FP] Taux de rémunération de la situation administrative
	NuclearPowerPlantOperationRate float64   `dsn:"S21.G00.51.015"` // Taux de conduite centrale nucléaire
	IncreasedRate                  float64   `dsn:"S21.G00.51.016"` // Taux de majoration
	ContributedRemunerationRate    float64   `dsn:"S21.G00.51.019"` // Taux de rémunération cotisée
	FormerApprenticeIncreaseRate   float64   `dsn:"S21.G00.51.020"` // Taux de majoration ex-apprenti/ex-élève
}

func GenerateRemuneration(contractNumber string) Remuneration {
	startDate := gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now())
	endDate := gofakeit.DateRange(startDate, startDate.AddDate(0, 1, 0))

	remunerations := []string{"012", "013", "017", "018"}

	return Remuneration{
		PayPeriodStartDate:             startDate,
		PayPeriodEndDate:               endDate,
		ContractNumber:                 contractNumber,
		Type:                           sample(remunerations),
		NumberOfHours:                  int64(gofakeit.IntRange(0, 200)),
		Amount:                         gofakeit.Float64Range(1000, 10000),
		AdministrativeStatusPayRate:    gofakeit.Float64Range(0, 100),
		NuclearPowerPlantOperationRate: gofakeit.Float64Range(0, 100),
		IncreasedRate:                  gofakeit.Float64Range(0, 100),
		ContributedRemunerationRate:    gofakeit.Float64Range(0, 100),
		FormerApprenticeIncreaseRate:   gofakeit.Float64Range(0, 100),
	}
}

type Activity struct {
	Type            string  `dsn:"S21.G00.53.001"` // Type
	Measure         float64 `dsn:"S21.G00.53.002"` // Mesure
	MeasurementUnit string  `dsn:"S21.G00.53.003"` // Unité de mesure
}

func GenerateActivity() Activity {
	activityTypes := []string{"01"}
	return Activity{
		Type:            sample(activityTypes),
		Measure:         gofakeit.Float64Range(0, 1000),
		MeasurementUnit: gofakeit.DigitN(2),
	}
}

func weirdDateFormat(t *time.Time) string {
	return fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
}

// Serialize converts any struct with dsn tags to a slice of "code,'attribute'" format
func Serialize(v interface{}) ([]string, error) {
	var result []string
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Serialize expects a struct, got %v", rv.Kind())
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		value := rv.Field(i)
		dsnTag := field.Tag.Get("dsn")
		if dsnTag == "" {
			continue
		}

		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Float32, reflect.Float64:
			strValue = fmt.Sprintf("%f", value.Float())
		case reflect.Int32, reflect.Int64, reflect.Int, reflect.Int16, reflect.Int8, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			strValue = fmt.Sprintf("%d", value.Int())
		case reflect.Ptr:
			// Handle pointer types
			if value.IsNil() {
				strValue = "NULL"
			} else if value.Type() == reflect.TypeOf(&time.Time{}) {
				// *time.Time
				t := value.Interface().(*time.Time)
				strValue = weirdDateFormat(t)
			} else {
				strValue = fmt.Sprintf("%v", value.Elem().Interface())
			}
		default:
			// Check if it's a time.Time
			if value.Type() == reflect.TypeOf(time.Time{}) {
				t := value.Interface().(time.Time)
				strValue = weirdDateFormat(&t)
			} else {
				// For other types, use default string conversion
				strValue = fmt.Sprintf("%v", value.Interface())
			}
		}
		result = append(result, fmt.Sprintf("%s,'%s'\n", dsnTag, strValue))
	}

	return result, nil
}

// SerializeToString converts any struct with dsn tags to a single string with "code,'attribute'" format
func SerializeToString(v interface{}) (string, error) {
	serialized, err := Serialize(v)
	if err != nil {
		return "", err
	}
	return strings.Join(serialized, "\n"), nil
}

const nIndividuals = 100

func main() {
	transmission := GenerateTransmission()
	sender := GenerateSender()
	senderContact := GenerateSenderContact()
	declaration := GenerateDeclaration()
	company := GenerateCompany()
	establishment := GenerateEstablishment()

	_, err := os.OpenFile("dsn.txt", os.O_RDONLY, 0644)
	if err == nil {
		err = os.Remove("dsn.txt")
		if err != nil {
			log.Fatalf("could not delete DSN file: %v", err)
		}
	}

	file, err := os.OpenFile("dsn.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString("S10.G00.00,''\n")
	lines, err := Serialize(transmission)
	if err != nil {
		log.Fatalf("cannot serialize transmission: %v", err)
	}
	for _, line := range lines {
		writer.WriteString(line)
	}

	writer.WriteString("S10.G00.01,''\n")
	lines, err = Serialize(sender)
	if err != nil {
		log.Fatalf("cannot serialize sender: %v", err)
	}
	for _, line := range lines {
		writer.WriteString(line)
	}

	writer.WriteString("S10.G00.02,''\n")
	lines, err = Serialize(senderContact)
	if err != nil {
		log.Fatalf("cannot serialize senderContact: %v", err)
	}
	for _, line := range lines {
		writer.WriteString(line)
	}

	writer.WriteString("S20.G00.05,''\n")
	lines, err = Serialize(declaration)
	if err != nil {
		log.Fatalf("cannot serialize declaration: %v", err)
	}
	for _, line := range lines {
		writer.WriteString(line)
	}

	writer.WriteString("S21.G00.06,''\n")
	lines, err = Serialize(company)
	if err != nil {
		log.Fatalf("cannot serialize company: %v", err)
	}
	for _, line := range lines {
		writer.WriteString(line)
	}

	writer.WriteString("S21.G00.11,''\n")
	lines, err = Serialize(establishment)
	if err != nil {
		log.Fatalf("cannot serialize establishment: %v", err)
	}
	for _, line := range lines {
		writer.WriteString(line)
	}

	for range nIndividuals {
		writer.WriteString("S21.G00.30,''\n")
		individual := GenerateIndividual()
		lines, err = Serialize(individual)
		if err != nil {
			log.Fatalf("cannot serialize individual: %v", err)
		}
		for _, line := range lines {
			writer.WriteString(line)
		}

		writer.WriteString("S21.G00.40,''\n")
		contract := GenerateContract()
		lines, err = Serialize(contract)
		if err != nil {
			log.Fatalf("cannot serialize contract: %v", err)
		}
		for _, line := range lines {
			writer.WriteString(line)
		}

		writer.WriteString("S21.G00.50,''\n")
		payment := GeneratePayment()
		lines, err = Serialize(payment)
		if err != nil {
			log.Fatalf("cannot serialize payment: %v", err)
		}
		for _, line := range lines {
			writer.WriteString(line)
		}

		remuneration := GenerateRemuneration(contract.ContractNumber)
		writer.WriteString("S21.G00.51,''\n")
		lines, err = Serialize(remuneration)
		if err != nil {
			log.Fatalf("cannot serialize remuneration: %v", err)
		}
		for _, line := range lines {
			writer.WriteString(line)
		}
	}
}

func sample(s []string) string {
	return s[gofakeit.IntRange(0, len(s)-1)]
}

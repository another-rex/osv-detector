package parsers_test

import (
	"osv-detector/detector/parsers"
	"testing"
)

func TestParseRequirementsTxt_Empty(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/empty.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{})
}

func TestParseRequirementsTxt_CommentsOnly(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/only-comments.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{})
}

func TestParseRequirementsTxt_OneRequirementUnconstrained(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/one-requirement-unconstrained.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{
		{
			Name:      "flask",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
	})
}

func TestParseRequirementsTxt_OneRequirementConstrained(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/one-requirement-constrained.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{
		{
			Name:      "django",
			Version:   "2.2.24",
			Ecosystem: parsers.PipEcosystem,
		},
	})
}

func TestParseRequirementsTxt_MultipleRequirementsConstrained(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/multiple-requirements-constrained.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{
		{
			Name:      "astroid",
			Version:   "2.5.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "beautifulsoup4",
			Version:   "4.9.3",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "boto3",
			Version:   "1.17.19",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "botocore",
			Version:   "1.20.19",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "certifi",
			Version:   "2020.12.5",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "chardet",
			Version:   "4.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "circus",
			Version:   "0.17.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "click",
			Version:   "7.1.2",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "django-debug-toolbar",
			Version:   "3.2.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "django-filter",
			Version:   "2.4.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "django-nose",
			Version:   "1.4.7",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "django-storages",
			Version:   "1.11.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "django",
			Version:   "2.2.24",
			Ecosystem: parsers.PipEcosystem,
		},
	})
}

func TestParseRequirementsTxt_MultipleRequirementsMixed(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/multiple-requirements-mixed.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{
		{
			Name:      "flask",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "flask-cors",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "pandas",
			Version:   "0.23.4",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "numpy",
			Version:   "1.16.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "scikit-learn",
			Version:   "0.20.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "sklearn",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "requests",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "gevent",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
	})
}

func TestParseRequirementsTxt_FileFormatExample(t *testing.T) {
	t.Parallel()

	packages, err := parsers.ParseRequirementsTxt("fixtures/pip/file-format-example.txt")

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	expectPackages(t, packages, []parsers.PackageDetails{
		{
			Name:      "pytest",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "pytest-cov",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "beautifulsoup4",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "docopt",
			Version:   "0.6.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "keyring",
			Version:   "4.1.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "coverage",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "Mopidy-Dirble",
			Version:   "1.1",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "rejected",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
		{
			Name:      "green",
			Version:   "0.0.0",
			Ecosystem: parsers.PipEcosystem,
		},
	})
}

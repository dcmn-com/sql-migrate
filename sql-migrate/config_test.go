package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestReadConfig_YAMLEnvVars(t *testing.T) {
	configTemplate := `
dynamic:
    dialect: ${DIALECT}
    datasource: ${DATASOURCE}
    schema: ${SCHEMA}
    table: ${TABLE}
    dir: ${DIR}`

	ConfigFile := testTempFile(t)
	if err := ioutil.WriteFile(tf, []byte(configTemplate), 600); err != nil {
		t.Fatal(err)
	}

	if err := os.Setenv("DIALECT", "postgres"); err != nil {
		t.Fatal(err)
	}

	if err := os.Setenv("DATASOURCE", "postgres://user:password@localhost:5432/db?sslmode=disable"); err != nil {
		t.Fatal(err)
	}

	if err := os.Setenv("SCHEMA", "myschema"); err != nil {
		t.Fatal(err)
	}

	if err := os.Setenv("TABLE", "migrations"); err != nil {
		t.Fatal(err)
	}

	if err := os.Setenv("DIR", "migrations/sql"); err != nil {
		t.Fatal(err)
	}

	env, err := ReadConfig()
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]*Environment{
		"dynamic": &Environment{
			Dialect:    "postgres",
			DataSource: "postgres://user:password@localhost:5432/db?sslmode=disable",
			TableName:  "migrations",
			Dir:        "migrations/sql",
			SchemaName: "myschema",
		},
	}

	if !reflect.DeepEqual(env, expected) {
		t.Logf("Failed reading config.")
		t.Logf("Got:  %s", env)
		t.Errorf("Want: %s", expected)
	}
}

func testTempFile(t *testing.T) string {
	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	return tf.Name()
}

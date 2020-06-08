package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log"
)


type ExampleSuite struct {
	suite.Suite
	db *gorm.DB
	tx *gorm.DB
}

func (t *ExampleSuite) SetupSuite() {
	db, err := gorm.Open("mysql", "root:my-secret-pw@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	t.db = db
	for _, model := range GetModels() {
		t.db.AutoMigrate(model)
	}
}

func (t *ExampleSuite) TearDownSuite() {
	defer t.db.Close()
	for _, model := range GetModels() {
		t.db.DropTableIfExists(model)
	}
}

func (t *ExampleSuite) SetupTest() {
	t.tx = t.db.Begin()
}

func (t *ExampleSuite) TearDownTest() {
	t.tx.Rollback()
}

func (t *ExampleSuite) TestExampleOne() {
	err := t.tx.Create(Blog{
		Title:   "some_title",
		Content: "some_content",
	}).Error
	require.NoError(t.T(), err)
}

func (t *ExampleSuite) TestExampleTwo() {
	err := t.tx.Create(Blog{
		Title:   "some_title",
		Content: "some_content",
	}).Error
	require.NoError(t.T(), err)
	var count int
	t.tx.Model(Blog{}).Count(&count)
	assert.Equal(t.T(), 1, count)
}



package repo_test

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
	"about-vaccine/src/repo"
	"testing"
)

func TestAdverseEventRepo(t *testing.T) {
	engine, err := dao.NewEngine(dao.DSN)
	if err != nil {
		t.Fatal(err)
	}
	rp := repo.NewAdverseEventRepo(dao.NewDB(engine))
	total, err := rp.Count()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(total)
}

func TestAdverseEventInsert(t *testing.T) {
	engine, err := dao.NewEngine(dao.DSN)
	if err != nil {
		t.Fatal(err)
	}
	rp := repo.NewAdverseEventRepo(dao.NewDB(engine))
	err = rp.CreateOne(&entity.AdverseEvent{
		Description: "dddddddddd",
	})
	if err != nil {
		t.Fatal(err)
	}
}

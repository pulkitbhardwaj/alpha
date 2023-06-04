package main_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pulkitbhardwaj/alpha/social/internal"
	"github.com/pulkitbhardwaj/alpha/social/internal/enttest"
	"github.com/pulkitbhardwaj/alpha/social/internal/migrate"
)

var graph *internal.Client

func TestSocial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Social Suite")
}

var _ = BeforeSuite(func() {
	opts := []enttest.Option{
		enttest.WithOptions(internal.Log(GinkgoWriter.Println)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	graph = enttest.Open(
		GinkgoT(),
		"sqlite3",
		"file:ent?mode=memory&_fk=1",
		opts...,
	)

	Expect(graph).ToNot(BeNil())
})

var _ = AfterSuite(func() {
	Expect(graph.Close()).To(Succeed())
})

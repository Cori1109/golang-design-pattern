package abstractfactory

import "testing"

func getMainAndDetail(factory DAOFactory) (string, string) {
	main := factory.CreateOrderMainDAO()
	detail := factory.CreateOrderDetailDAO()
	return main.SaveOrderMain(), detail.SaveOrderDetail()
}

func TestRdbFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	main, detail := getMainAndDetail(factory)
	if main != "rdb main save" {
		t.Fatal("main save fail")
	}
	if detail != "rdb detail save" {
		t.Fatal("detail save fail")
	}
}

func TestXmlFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	main, detail := getMainAndDetail(factory)
	if main != "xml main save" {
		t.Fatal("main save fail")
	}
	if detail != "xml detail save" {
		t.Fatal("detail save fail")
	}
}

package url

import (
	"net/url"
	"testing"
)

func TestQuery(t *testing.T) {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		t.Fatal(err)
	}

	if u.RawQuery != "q=dotnet" {
		t.Error("RawQuery error")
	}

	u, err = url.Parse("http://bing.com/search?k=v&id=main&id=omit&array[]=first&array[]=second&ids[i]=111&ids[j]=3.14")
	if err != nil {
		t.Fatal(err)
	}

	if u.RawQuery != "k=v&id=main&id=omit&array[]=first&array[]=second&ids[i]=111&ids[j]=3.14" {
		t.Error("RawQuery error")
	}
}

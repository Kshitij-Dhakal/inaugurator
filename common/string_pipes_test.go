package common_test

import "testing"
import "github.com/Kshitij-Dhakal/inaugurator/common"

func TestToCamelCase(t *testing.T)  {
	if(common.ToCamelCase("HelloWorld") != "helloWorld"){
		t.Errorf("Expected %s but got %s", "helloWorld", common.ToCamelCase("HelloWorld"))
	}
	if(common.ToCamelCase("helloWorld") != "helloWorld"){
		t.Errorf("Expected %s but got %s", "helloWorld", common.ToCamelCase("HelloWorld"))
	}
	if(common.ToCamelCase("hello-world") != "helloWorld"){
		t.Errorf("Expected %s but got %s", "helloWorld", common.ToCamelCase("HelloWorld"))
	}
	if(common.ToCamelCase("hello_world") != "helloWorld"){
		t.Errorf("Expected %s but got %s", "helloWorld", common.ToCamelCase("HelloWorld"))
	}
	if(common.ToCamelCase("HELLO_WORLD") != "helloWorld"){
		t.Errorf("Expected %s but got %s", "helloWorld", common.ToCamelCase("HelloWorld"))
	}
	if(common.ToCamelCase("hello world") != "helloWorld"){
		t.Errorf("Expected %s but got %s", "helloWorld", common.ToCamelCase("HelloWorld"))
	}
}

func TestToPascalCase(t *testing.T)  {
	if(common.ToPascalCase("HelloWorld") != "HelloWorld"){
		t.Errorf("Expected %s but got %s", "HelloWorld", common.ToPascalCase("HelloWorld"))
	}
	if(common.ToPascalCase("helloWorld") != "HelloWorld"){
		t.Errorf("Expected %s but got %s", "HelloWorld", common.ToPascalCase("HelloWorld"))
	}
	if(common.ToPascalCase("hello-world") != "HelloWorld"){
		t.Errorf("Expected %s but got %s", "HelloWorld", common.ToPascalCase("HelloWorld"))
	}
	if(common.ToPascalCase("hello_world") != "HelloWorld"){
		t.Errorf("Expected %s but got %s", "HelloWorld", common.ToPascalCase("HelloWorld"))
	}
	if(common.ToPascalCase("HELLO_WORLD") != "HelloWorld"){
		t.Errorf("Expected %s but got %s", "HelloWorld", common.ToPascalCase("HelloWorld"))
	}
	if(common.ToPascalCase("hello world") != "HelloWorld"){
		t.Errorf("Expected %s but got %s", "HelloWorld", common.ToPascalCase("HelloWorld"))
	}
}
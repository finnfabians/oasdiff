package diff_test

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tufin/oasdiff/diff"
	"gopkg.in/yaml.v3"
)

func ExampleGet() {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	s1, err := loader.LoadFromFile("../data/simple1.yaml")
	if err != nil {
		fmt.Printf("failed to load spec with %v", err)
		return
	}

	s2, err := loader.LoadFromFile("../data/simple2.yaml")
	if err != nil {
		fmt.Printf("failed to load spec with %v", err)
		return
	}

	diffReport, err := diff.Get(&diff.Config{}, s1, s2)

	if err != nil {
		fmt.Printf("diff failed with %v", err)
		return
	}

	bytes, err := yaml.Marshal(diffReport)
	if err != nil {
		fmt.Printf("failed to marshal result with %v", err)
		return
	}
	fmt.Printf("%s\n", bytes)

	// Output:
	// paths:
	//     modified:
	//         /api/test:
	//             operations:
	//                 added:
	//                     - POST
	//                 deleted:
	//                     - GET
	// endpoints:
	//     added:
	//         - method: POST
	//           path: /api/test
	//     deleted:
	//         - method: GET
	//           path: /api/test
}

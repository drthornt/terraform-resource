// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package servicecatalogiface provides an interface to enable mocking the AWS Service Catalog service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicecatalogiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
)

// ServiceCatalogAPI provides an interface to enable mocking the
// servicecatalog.ServiceCatalog service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Service Catalog.
//    func myFunc(svc servicecatalogiface.ServiceCatalogAPI) bool {
//        // Make svc.DescribeProduct request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := servicecatalog.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServiceCatalogClient struct {
//        servicecatalogiface.ServiceCatalogAPI
//    }
//    func (m *mockServiceCatalogClient) DescribeProduct(input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServiceCatalogClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ServiceCatalogAPI interface {
	DescribeProductRequest(*servicecatalog.DescribeProductInput) (*request.Request, *servicecatalog.DescribeProductOutput)

	DescribeProduct(*servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error)

	DescribeProductViewRequest(*servicecatalog.DescribeProductViewInput) (*request.Request, *servicecatalog.DescribeProductViewOutput)

	DescribeProductView(*servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error)

	DescribeProvisioningParametersRequest(*servicecatalog.DescribeProvisioningParametersInput) (*request.Request, *servicecatalog.DescribeProvisioningParametersOutput)

	DescribeProvisioningParameters(*servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error)

	DescribeRecordRequest(*servicecatalog.DescribeRecordInput) (*request.Request, *servicecatalog.DescribeRecordOutput)

	DescribeRecord(*servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error)

	ListLaunchPathsRequest(*servicecatalog.ListLaunchPathsInput) (*request.Request, *servicecatalog.ListLaunchPathsOutput)

	ListLaunchPaths(*servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error)

	ListRecordHistoryRequest(*servicecatalog.ListRecordHistoryInput) (*request.Request, *servicecatalog.ListRecordHistoryOutput)

	ListRecordHistory(*servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error)

	ProvisionProductRequest(*servicecatalog.ProvisionProductInput) (*request.Request, *servicecatalog.ProvisionProductOutput)

	ProvisionProduct(*servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error)

	ScanProvisionedProductsRequest(*servicecatalog.ScanProvisionedProductsInput) (*request.Request, *servicecatalog.ScanProvisionedProductsOutput)

	ScanProvisionedProducts(*servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error)

	SearchProductsRequest(*servicecatalog.SearchProductsInput) (*request.Request, *servicecatalog.SearchProductsOutput)

	SearchProducts(*servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error)

	TerminateProvisionedProductRequest(*servicecatalog.TerminateProvisionedProductInput) (*request.Request, *servicecatalog.TerminateProvisionedProductOutput)

	TerminateProvisionedProduct(*servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error)

	UpdateProvisionedProductRequest(*servicecatalog.UpdateProvisionedProductInput) (*request.Request, *servicecatalog.UpdateProvisionedProductOutput)

	UpdateProvisionedProduct(*servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error)
}

var _ ServiceCatalogAPI = (*servicecatalog.ServiceCatalog)(nil)
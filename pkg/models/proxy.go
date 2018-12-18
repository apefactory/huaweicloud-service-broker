package models

import (
	"github.com/pivotal-cf/brokerapi"

	"github.com/huaweicloud/huaweicloud-service-broker/pkg/database"
)

// ServiceBrokerProxy is used to implement details
type ServiceBrokerProxy interface {
	Provision(instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error)

	Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error)

	Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error)

	Unbind(instanceID, bindingID string, details brokerapi.UnbindDetails) error

	Update(instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error)

	LastOperation(instanceID string, operationData database.OperationDetails) (brokerapi.LastOperation, error)

	GetPlanSchemas(serviceID string, planID string, metadata *brokerapi.ServicePlanMetadata) (*brokerapi.PlanSchemas, error)
}

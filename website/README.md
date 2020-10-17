# DOMAINER - Doman name management for the shopr platform

## Overview
Domainer is a service that exposes a set of endpoints that enables authenticated users and shopr-services carry out the following actions.

- Domain name availability search
- Domain name suggestions
- Domain name owner profile creation
- Domain name owner profile update
- Domain name registration
- Domain name renewal (periodic)reminders
- Domain name DNS setting/update
- Domain name renewal

## End Points
- SearchDomains
- RegisterDomain
- RenewDomain
- CreateCustomerProfile
- UpdateCustomerProfile
- DeleteCustomerProfile
- GetDomainSuggestions


### Authentication is handled by the SPA(shopr authentication service)

### Endpoints will be made available in gRPC and JSON

*This is an evolving document*

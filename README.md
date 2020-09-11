# SHOPR - AN ONLINE STORE CREATOR

## Overview
Shopr is an online store creator that enables any merchant with physical products to sell create an online store where they can list products with details and collect payment electronically for purchases with buyer details to fulfill delivery of the order.



## UX - How the user should experience this app

### Merchants and Buyers
1. Merchant User creates a free account through (email of social media one click button).
2. Merchant User confirms their identity and updates their profile.
3. Search for and Choose an available domain name with few suggestions in cases of taken domain.
4. Merchant Picks a user package.
5. Merchant User adds items to store with price, product and delivery details.
6. Merchant User publish items for sale.
7. Items are published on site and on social media (for those that allow).
8. Merchant User manages social interactions independently of the platform.
9. Conversations on social link to item purchase pages on platform.
10. Buyer also access products list through merchants’ home page and domain.
11. Buyer can see items’ details [ reviews, ratings, specification details ].
12. Buyer adds to cart or uses one click checkout.
13. Buyer makes payment for listing.
14. Buyer gets alert with purchase information in their email.
15. Merchant get alert of purchase in their email.
16. Merchants update transaction status as they progress.[admin]
17. Buyer get regular updates on status of order. [admin]
18. Order is complete when item is received by buyer.
19. Buyer gets link to review item.
20. Merchant can view historical orders [admin]
21. Merchant can view, add and edit listings. [admin]
22. Merchant can view a list of previous buyers. [admin]
23. Merchant can view shop activity statistics in dashboard [admin]
24. Merchant can integrate third party tools with one click. [ GoogleAnalytics, PayStack, Stripe, Twitter, Facebook, Instagram ] [admin]
25. Merchant can grant other users shop admin rights based on shop administration roles with a few clicks [admin]
26. Merchant is able to generate coupons with rules [admin]


### Platform Admin
1. View list of users and administer accounts [ Suspend, Re-Activate, View Stats ]
2. See complete list of Items listed


### Features
1. Signup with social media
2. Signup with email address + Verification
3. Profile update for missing information [ Phone, StoreName, PaymentDetails, SubscriptionPackage, LinkSocialProfile, UploadBanner, Verify Govt. Issued ID, ContactDetails  ]
4. Upload Items + Fix Prices + Shipping Cost
5. Payment options for goods sold in stores + For Goods sold in Portal a. Shops with subscription package optional receive direct to bank account b. Commission based merchants payment will be routed to service account.
6. Plans [ Default, Pro, Pro+ ] Portal Admin can add and remove features from plans
7. Limit Product Upload.
8. Social media integration (Instagram)
9. Custom domain.
10. Delivery status updates.
11. Coupon Integration.
12. Analytics.


### Payment Gateways 
- Stripe.
- Paystack.


### Development tools and setup

1. Ubuntu 18.4x [Operating System]
2. Kubernetes[for Orchestration tool]
3. Docker [for containerization]
4. Traefik [for ingress]
5. Linkerd [for Service Mesh]
6. DGraph [for Graph Database]
7. MariaDB [ for Transactional Database]
8. GitLab [for version managment, CI/CD]
9. AWS SES [for Email Delivery]
10. CloudFlare [for DNS Management & DDOS protection]

### Programming Languages
- Go
- PHP
- ReactJS
- JQuery
- SQL
- GraphQL


### Cloud Supplier
1. DigitalOcean
2. AWS

### Domain Vendors with API (Search & Registration)
1. SRSPlus
2. Reseller Club



### MVP

1. Ability for our customers to be able to have their own store in their own domain name.
2. They should be able to integrate their Instagram page to the store such that transactions started on social media can be completed on the website.
3. Payment integration.

## TODO
- Cluster setup [Week1]
- CI/CD Server setup [Week1]
- AWS SES service setup [week1]
- Message Processing service [week2]
- Authentication Service [+UI in React] [week3]
- Signup Service  [+UI in React] [week3]
- User Management [Light with least features] [ Show Users, add & update products, Change order status] [week7+]
- Domain Search & Registration API + Propagation [CloudFlare API + SRS Plus API] [week4]
- Product upload service [Instagram API + Twitter API] [+UI in React] [week5-6]
- Subscription Package Service [ PayStack Integration, Stripe API] [+UI in React]
- Cart Shopping/Checkout Service. [+UI in React] [week6]



## SERVICES

SPA - Handles Authentication, Signup, Profile Management

Subscribe - Handles Subscription

Lister - Handles Product listing management

Ordering - Handled ordering management.

Rating - Handles Rating and Reviews.

Wallet - Handled temporary hold of money value in cases of disputes and refunds

Messager - Handle message routing and delivery within and from the platform to outside the platform.

Domainer - Domain Search & Registration


### Assumptions
*-Client handles their own shipping, so, we don’t really have much to do their.*

*-MVP 1 & 2 will be ready in six weeks assuming we don't have to develop a user interface from scratch*
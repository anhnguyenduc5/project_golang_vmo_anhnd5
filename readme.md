 # Onboarding and Volunteer Service

This repository contains the codebase for the Onboarding and Volunteer Service application. The application is built using the Go programming language with the Gin web framework, GORM ORM, and PostgreSQL database. It aims to streamline the process of onboarding and managing volunteers, providing functionality for managing 2 kinds of user: applicant and volunteer.

## Table of contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [User story](#user-story)
- [API Endpoints "/api/v1"](#api-endpoints-apiv1)
  - [Admin Endpoints: "/admin"](#admin-endpoints-admin)
  - [User Endpoints: "/applicant"](#user-endpoints-applicant)
  - [Application Request Endpoints:"/applicant-request"](#application-request-endpointsapplicant-request)
  - [User Identity Endpoints: "/applicant-identity"](#user-identity-endpoints-applicant-identity)
- [Contributing](#contributing)
- [License](#license)
  
### Project Structure
The project follows a modular structure with clearly defined folders:
├───cmd  
│   ├───migration  
│   └───server  
├───deployment  
├───docs  
├───feature  
│   ├───country    
│   ├───department  
│   ├───middleware  
│   ├───request  
│   ├───role  
│   ├───user  
│   ├───user_identity  
│   └───volunteer  
└───migration  

### Installation
To get started with the Onboarding and Volunteer Service application, follow these steps:

Clone the repository:  
git clone https://github.com/cesc1802/onboarding-and-volunteer-service.git  
cd onboarding-and-volunteer-service

Install dependencies:  
go mod download  
Set up your environment variables. Copy the .env.example to .env and fill in the necessary values.

cp .env.example .env

### Configuration
Make sure to configure the following environment variables in the .env file:

DB.HOST: Database host  
DB.PORT: Database port  
DB.USER: Database user  
DB.PASS: Database password  
DB.NAME: Database name

Database Migration  
Run the database migrations to set up the required tables:  
go run cmd/migration/main.go

### Usage
To start the application, run:  
go run cmd/main.go

The server will start on the port specified in the .env file.

### User story 
I. Guest

As a guest, I want to join the app  so that I can start the registration process.

II. Applicant


As an applicant, I want to log in the system so that I can register as a volunteer. (bỏ qua)  
As an applicant, I would like to register to be a volunteer so that I can participate in volunteer activities.  
As an applicant, I want to change some fields in the application form so that my information is up to date.  
As an applicant, I would like to cancel my registration so that I can withdraw my application.  
As an applicant, I want to view my registration status so that I can track my application progress.  
As an applicant, I would like to activate my email after getting approved so that I can start receiving official communications.  
As an applicant, I would like to know why I’m being rejected so that I can understand the reasons and address them.  
As an applicant, I would like to fix errors according to rejected notes so that I can correct my application and resubmit it.  

III. Volunteer  

As a volunteer, I would like to log in the system so that I can access my volunteer dashboard.  
As a volunteer, I would like to verify that I’m a volunteer so that I can participate in volunteer activities.  
As a volunteer, I want to update my verification form so that my information remains current.  
As a volunteer, I would like to cancel  my verification so that I can withdraw my verification request.  
As a volunteer, I want to view my request status so that I can track the progress of my verification.  
As a volunteer, I would like to activate my email after getting approved so that I can start receiving official communications.   
As a volunteer, I would like to know why I’m being rejected so that I can understand the reasons and address them.   
As a volunteer, I would like to fix errors according to rejected notes so that I can correct my verification and resubmit it.  

IV. Admin  
	
As an admin, I would like to log in the system so that I can access the admin features. (bỏ qua)  
As an admin, I want to see the admin main menu so that I can have an overview of all the features.  
As an admin, I want to redirect to other features from the admin main menu so that I can manage the system effectively.  
As an admin, I would like to search for features in the main menu so that I can quickly access the functionality I need.  
As an admin, I would like to view all verification/application requests so that I can manage pending requests.  
As an admin, I want to search verifications / registrations that have not been approved by email so that I can find specific requests easily.  
As an admin, I want to be able to view detailed forms of verifications / registrations / approved so that I can review all information before making decisions.  
As an admin, I would like to approve requests so that I can accept valid applications.  
As an admin, I would like to reject requests so that I can decline invalid applications.  
As an admin, I want to add reject notes for rejected forms so that applicants know why their request was rejected.    
As an admin, I would like to mark a request that it is viewed so that I can track which requests have been reviewed.  
As an admin, I would like to delete a request so that I can remove invalid or duplicate entries.  
As an admin, I want to send mail to requesters so that I can communicate directly with applicants and volunteers.  
As an admin, I want to view all departments' lists and their locations so that I can manage department information.   
As an admin, I would like to search for a department so that I can find specific departments quickly.  
As an admin, I want to be able to update a department so that I can keep department information current.  
As an admin, I want to be able to add new departments so that I can expand the organization.  
As an admin, I would like to be able to delete a department so that I can remove obsolete departments.  
As an admin, I would like to view the volunteer table so that I can manage volunteer information.  
As an admin, I want to be able to deactivate a volunteer so that I can manage volunteer status.  
As an admin, I would like to search for volunteers by id / name so that I can find specific volunteers quickly.  
As an admin, I would like to filter volunteers by sex so that I can organize volunteer information efficiently.  
As an admin, I would like to add a volunteer manual so that volunteers do not need to make verifications.  
As an admin, I would like to view the volunteer list order in their role so that I can manage volunteers based on their roles (COM: Committee, CVL: Civil volunteer, MNVC: Manager of volunteer coordinators, …)  
As an admin, I want to search for volunteers by role so that I can find volunteers with specific responsibilities.  

### API Endpoints "/api/v1"

#### Admin Endpoints: "/admin" 
Before you get to use the admin api, you must log-in first to get authorize token
GET "/list-request": Get the request list  
GET "/request/:id" : Get a specific request  
POST "/approve-request/:id": Approve a request, change status of request  
POST "/reject-request/:id": Reject a request, change status of request  
POST "/add-reject-notes/:id": Add reject notes to a request  
DELETE "/delete-request/:id": Delete a request  

#### User Endpoints: "/applicant"  
POST "/:" Create a new user  
PUT "/:id" : Update an existing user from register form  
DELETE "/:id" : Delete an user with id  
GET "/:id" : Get information about an user with id

#### Application Request Endpoints:"/applicant-request"  
POST "/" : Create a record request

#### User Identity Endpoints: "/applicant-identity"  
POST "/" : Create a user identity record  
GET "/:id": Find a user identity  
PUT "/:id": Update a user identity record    

### Contributing  

We welcome contributions to enhance the features and functionality of this project. Please follow these steps:

Fork the repository.  
Create a new branch (git checkout -b feature/your-feature).  
Commit your changes (git commit -am 'Add new feature').  
Push to the branch (git push origin feature/your-feature).  
Create a new Pull Request.  

### License  
This project is licensed under the MIT License. See the LICENSE file for details.  


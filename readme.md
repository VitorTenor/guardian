# Guardian
## Authentication Application

Welcome to the comprehensive documentation for the Authentication Application that seamlessly integrates with the MusicStreamer project. This documentation aims to provide an in-depth understanding of the application's architecture, technologies, routes, and functionalities.

## Introduction

The Authentication Application plays a pivotal role in ensuring secure user access to the MusicStreamer platform. By employing modern technologies and best practices, the application guarantees robust authentication mechanisms, protecting user data and enhancing user experience.

## Technologies Utilized

The Authentication Application leverages a combination of cutting-edge technologies to deliver its capabilities:

- **Programming Language: Go (Golang)**
  Go, known for its performance and efficiency, serves as the foundation of the Authentication Application. Its concurrency support and standard libraries make it an ideal choice for building reliable and scalable applications.

- **Password Management: bcrypt**
  The bcrypt library is a critical component in ensuring the security of user credentials. It utilizes a sophisticated hash function with salting and multiple iterations, rendering password decryption attempts significantly more challenging for potential attackers.

- **Routing Framework: Gingonic**
  Gingonic, a lightweight and flexible routing framework, empowers the Authentication Application to define and manage various API endpoints. This facilitates seamless communication between clients and the authentication server.

- **JSON Validation: Validator Playground**
  Validator Playground enhances data integrity by validating incoming JSON payloads against predefined schemas. This step ensures that user inputs are correctly structured and adhere to expected standards.

## Routes and Functionalities

The Authentication Application offers essential routes and functionalities to support secure user authentication:

### Authentication Route

`POST /token`
`POST /token/refresh`

This route enables users to log in to the MusicStreamer platform securely. Users submit their credentials in JSON format.

**Input:**
```json
{
  "email": "email",
  "password": "password"
}

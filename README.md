# Banking API example

> Note this is still work in progress, working on
- Variable naming consistency, review
- Error handling refinement


This code represents a sample banking API for an imaginary bank.

The solution is implemented in Go using a Domain-Driven Design (DDD) approach.

**Why Go?**

Partly because it’s the language I’ve been working with recently, but also because it’s a compiled, fast, and reliable language that is easy to package and deploy. From what I’ve seen, more and more fintech companies are adopting Go—at least for parts of their systems.

**Why DDD (Domain-Driven Design)?**

DDD is also becoming increasingly popular.

This approach might be somewhat overkill for a task of this size. It results in more verbose code and multiple layers of separation. However, it also offers several benefits. Let’s look at the advantages and drawbacks:

**Advantages:**

- Alignment with business goals: Your software reflects the real-world processes and rules of finance, reducing miscommunication between business and tech.
- Scalability & maintainability: Clear boundaries (bounded contexts) make it easier to evolve parts of the system independently.
- Decoupling of concerns: Encourages separation between domain logic and infrastructure, making the system flexible and testable.
- High-quality domain modeling: Explicit models help manage complexity, especially in fintech, where regulations, transactions, and risk rules are intricate.

**Drawbacks:**

- Steep learning curve: Teams need to understand DDD concepts and invest time in modeling the domain.
- Initial complexity: Creating bounded contexts, aggregates, and ubiquitous language requires upfront effort.
- Not a silver bullet: Overkill for simple apps; best for complex, high-domain-value systems.


---

**Layers**
```
┌─────────────────────────────────────────────────────────────┐
│                    Presentation Layer                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   User API  │  │ Account API │  │  Transaction API    │  │
│  │ (handlers)  │  │ (handlers)  │  │    (handlers)       │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                             │
┌─────────────────────────────────────────────────────────────┐
│                   Application Layer                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │User Service │  │Acct Service │  │Transaction Service  │  │
│  │(use cases)  │  │(use cases)  │  │   (use cases)       │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                             │
┌─────────────────────────────────────────────────────────────┐
│                     Domain Layer                            │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │User Entity  │  │Account      │  │Transaction Entity   │  │
│  │+ Value Objs │  │Entity       │  │+ Business Rules     │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                             │
┌─────────────────────────────────────────────────────────────┐
│                Infrastructure Layer                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │User Repo    │  │Account Repo │  │Transaction Repo     │  │
│  │(PostgreSQL) │  │(PostgreSQL) │  │(PostgreSQL)         │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```



**Why it fits fintech**

Fintech is highly regulated, complex, and rapidly changing, with multiple business domains like payments, lending, trading, and compliance. DDD fits perfectly because it:
- Allows isolated bounded contexts for different financial domains (e.g., payments vs. fraud detection), reducing risk of cascading errors.
- Supports auditability and compliance, since domain models make rules explicit and traceable.
- Makes integration with external services (banks, payment networks, KYC providers) easier through decoupled interfaces.

***Key architectural principles in DDD***
- Layered architecture / layer separation:
- Decoupling & dependency inversion:
- Bounded contexts & ubiquitous language:

---

## About the Application

**Command Layer (cmd/)**

This layer is responsible for bootstrapping the project. It sets up dependencies, initializes services, and passes them to the HTTP handlers.

**Presentation Layer (internal/api)**

This layer is not part of the domain model; it only handles presentation concerns.
- ***Middlewares*** manage JSON headers, Bearer authentication, and protection against IDOR (Insecure Direct Object Reference) attacks.
- ***Action handlers*** implement the v1 API. A v2 version is only a placeholder for future expansion.

**Application Layer**

Often referred to as the service layer, this layer orchestrates application behavior and provides a boundary between the presentation layer and the domain.

**Domain Layer**

This layer is completely independent and contains the core business entities. It defines domain objects, entities, and downward-pointing interfaces for repository implementations. In this small example, no aggregates are included.

**Infrastructure Layer**

This layer contains technical implementations such as:
- Persistence repositories (currently using PostgreSQL)
- Generic low-level components, such as configuration handling and database access utilities

---

**Database (PostgreSQL)**

Database migrations are handled using a custom migration tool I wrote.
You can find installation and usage instructions here:
https://github.com/olbrichattila/godbmigrator_cmd


**Database repositories use raw SQL.**

I chose to use raw SQL instead of relying on an ORM. While ORMs are convenient and powerful, in my opinion, for data-critical applications raw SQL (or stored procedures) provides greater control, transparency, and resilience. This approach makes data operations more predictable and easier to optimize, especially when correctness and performance are critical.


**Testing**

Installation, this does not come default like testify:
```bash
go install github.com/onsi/ginkgo/v2/ginkgo@latest
```

Run:
```bash
ginkgo ./...
```

**Mocking**
Used gomock to auto generate mocks: this also needs to be installed:
```bash
go install github.com/golang/mock/mockgen@latest
go get github.com/golang/mock/gomock
```

To re-generate all mocks:
```bash
go generate ./...
```

**A few example BDD (Behavior-Driven Development) test**
***User entity:***

https://github.com/olbrichattila/ddd-bank-api-example/blob/main/internal/domain/user/entity_test.go

***Transaction service:***

https://github.com/olbrichattila/ddd-bank-api-example/blob/main/internal/application/transaction/transaction_test.go


**Makefile**
```bash
# Build docker container
make api-build
# Run build docker container, (not network host works on Linux only, on Mac you need to spin up a postres on the same network, change db settings in .env.migrator)	
make api-run
# Start already built docker container	
make api-start
# Stop background running container	
make api-stop
# Run ginko tests, see setup above	
make run-test
# Auto generate mocks, see setup above	
make gen-mocks
```

## Things to Address: TODO:

- ~~Add transaction support in the API.~~
- ~~Add logger~~
- ~~Use configuration files for database settings instead of hard-coded values.~~
- ~~Implement ACID-compliant transactions using a Unit of Work pattern.~~
- ~~Dockerize the application for easier testing (e.g., multi-stage Docker build).~~
- ~~Implement login flow and JWT token generation for testing (currently commented out).~~
- ~~Apply IDOR protection not only to user resources but to all relevant resource types.~~
- ~~Add at least a few tests id BDD style with gomock and ginko~~
- Value object are not consistent everywhere, account number, currency for example, half way refactored
- Add Linter

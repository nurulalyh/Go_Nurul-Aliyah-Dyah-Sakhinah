# (20) Clean and Hexagonal Architecture

Clean Architecture adalah arsitektur yang bisa terpisahkan secara layer dan modular serta scaleable atau bisa di upgrade dan maintainable. Untuk mengatasi kondisi yang kacau, dapat menggunakan Hexagonal Architecture. Hexagonal Architecture adalah terdiri dari API interface, Domain Service, dan SPI interface. Perkembangan ide arsitektur yang pernah muncul :
- Hexagonal Architecture
- Onion Architecture
- Screaming Architecture
- DCI from Agile Development
- BCE from Object Oriented Software
- Clean Architecture

Hal yang harus dipatuhi sebelum membuat Clean Architecture :
- Independent of Framework
- Testable
- Independent of UI
- Independent of Database
- Independent od any external

Keuntungan menggunakan Clean Architecture :
- Sebuah struktur standar, jadi lebih mudah switching project
- Faster Development
- Mocking dependencies menjadi sesuatu yang selalu diterapkan di seluruh project.
- Mudah switching dari prototypes ke proper solutions.

Layer pada clean arsitektur :
- Entities Layer (optional) => object bisnis yang digunakan
- Use Case - Domain Layer => business logic
- Controller - Presentation layer => menampilkan data ke layar dan menghandle user intecations
- Drivers - Data layer => manages application data.

Pendekatan Clean Architecture :
- Layers in a single module
- One layer per module
- layers inside the feature module
- layers per feature modules

Domain Driven Design dapat membantu untuk mengelompokkan yang dapat digunakan sebagai entity baru. Clean Architecture adalah Software Architecture dan Domain Driven Design afalah software design technique
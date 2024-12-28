# Changelog

## [1.1.0](https://github.com/vp2305/GoSocial/compare/v1.0.0...v1.1.0) (2024-12-28)


### Features

* update api version automatically ([8c2d0cd](https://github.com/vp2305/GoSocial/commit/8c2d0cd5f1a2724944c2ec77225cb118fd4e95d4))

## 1.0.0 (2024-12-28)


### Features

* added automation workflow ([91a0058](https://github.com/vp2305/GoSocial/commit/91a0058a9c8726b95869853fcaaac6d8f5632969))
* **authorization:** Database setup for authorization and role based access control ([2ec2a30](https://github.com/vp2305/GoSocial/commit/2ec2a303fb9f0adbdfe8995844165ed3da6f84b2))
* **cache:** Implemented redis to cache user profile to drastically improve performance. ([6fd0603](https://github.com/vp2305/GoSocial/commit/6fd06032934d590e46dfdc55dcd733183c17a6dd))
* **comments:** Implemented fetch comments by post id ([3b7c01c](https://github.com/vp2305/GoSocial/commit/3b7c01ce7a59c0bd54442ec44b28c2f1fbc9f819))
* **concurrency:** Implemented opitimistic concurrency control in go ([a33075d](https://github.com/vp2305/GoSocial/commit/a33075dd6be21d2db098ebb4485e15bfc6e98555))
* **cors:** Implemented confirm token and cors configuration. ([8651e3a](https://github.com/vp2305/GoSocial/commit/8651e3a7705cee1a843205ca3c1488761643c645))
* **cors:** Implemented confirm token and cors configuration. ([ee44055](https://github.com/vp2305/GoSocial/commit/ee44055925b4357474051302bbf73fd428d0058f))
* **error:** Implemented error handling with logs. ([73cf468](https://github.com/vp2305/GoSocial/commit/73cf4683aafe60e6f8811e2d694ea3382db996a8))
* **follow/unfollow:** Implemented api endpoint to follow and unfollow a user ([4c155fb](https://github.com/vp2305/GoSocial/commit/4c155fb05dddc223095af450ef74951a88e2272e))
* Implemented database config, repository pattern, etc.. ([86be60a](https://github.com/vp2305/GoSocial/commit/86be60ae2dbe57b2b4155a72b3d64fd9e3473cbd))
* **jsonResponse:** Standardising JSON responses ([47066bb](https://github.com/vp2305/GoSocial/commit/47066bb016bf2909aebb417a85dd29884256bc46))
* **jwt:** Implemented basic and stateless token generation for api endpoints ([7f6bf93](https://github.com/vp2305/GoSocial/commit/7f6bf93684bcbd5f49a80dc3904f45fa8019bee0))
* **middleware:** Implemented auth token middleware to authenticate respective api endpoints. ([df7edae](https://github.com/vp2305/GoSocial/commit/df7edae3b2daa06fda5eab0a475016522c2db6e2))
* **pagination:** Implemented pagination on user feed endpoint ([b46b45d](https://github.com/vp2305/GoSocial/commit/b46b45d3f76d9e868a334a39a3a929c8c51dda50))
* **rateLimit:** Implemented rate limiter for all the api endpoints ([0866fc7](https://github.com/vp2305/GoSocial/commit/0866fc703dac8c004cbefaa4ac8af8fba717bbd2))
* **rbac:** Implemented role based actions on api endpoints. ([b735bbd](https://github.com/vp2305/GoSocial/commit/b735bbd5a55ad98e2e7a56cec869f3786c685f10))
* release please script ([bb1fbaa](https://github.com/vp2305/GoSocial/commit/bb1fbaa500a142c8b73e735941b2e9a775be3472))
* **seeding:** Implemented database seeding script along with additional endpoint to create comments ([3703d7b](https://github.com/vp2305/GoSocial/commit/3703d7b9d9906dca66eb7679e1b2fcdc7c82ab32))
* **server:** Implemented graceful shutdown for scaling and good practice ([d2ee132](https://github.com/vp2305/GoSocial/commit/d2ee1322d0849557bf65e351b2eb34331cbaa437))
* **swagger-docs:** Implemented a way to automate api docs ([38ea5d1](https://github.com/vp2305/GoSocial/commit/38ea5d1366d6d078a4d06de96f6255bd15b1fba8))
* **update-delete:** Implemented update/delete for posts api endpoint ([486a3d6](https://github.com/vp2305/GoSocial/commit/486a3d6ce097894c43df3793895e2a99f76eef80))
* **user_activation:** Implemented user activation with the token verification ([be4a27d](https://github.com/vp2305/GoSocial/commit/be4a27dcbb840277632bc1ed9e008c4a69b5bc35))
* **user_creation:** Implemented a user creation flow with end to end invitation through sendgrid ([7bdcd2b](https://github.com/vp2305/GoSocial/commit/7bdcd2bd6e428d01ed83a2a70772f622329be546))
* **userFeed:** Implemented api endpoint to fetch user feed ([09f81c8](https://github.com/vp2305/GoSocial/commit/09f81c8d7dfe03b1048c8d1fb0b2c49246a0b3ea))
* **users:** Implemented api endpoint to fetch user by id ([223fff0](https://github.com/vp2305/GoSocial/commit/223fff0ce01f79a6e8a1bbf44cb0cee363a3c6a7))
* **validator:** Implemented validator for the payload. ([163aa60](https://github.com/vp2305/GoSocial/commit/163aa607565fe8b8da2e5f4e2b3226713866b34e))


### Bug Fixes

* staticcheck installation step in audit workflow ([aa2c90c](https://github.com/vp2305/GoSocial/commit/aa2c90ca9e4a5d31f52da4bb1ff466a7f9716f91))

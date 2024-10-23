# Evaluation of test harnesses for LEAF

This repository was used in an effort to overhaul automated API testing in LEAF. Go was ultimately selected, and the tests reside in https://github.com/department-of-veterans-affairs/LEAF-Automated-Tests

Go and TypeScript (Deno) are interesting because they include first-party support for test runners, are well documented, and have small runtimes. Tests were implemented in both languages to explore the respective ecosystems and determine viability. Although Go is more verbose, there is significantly more potential to reuse code.

The existing PHPUnit test suite was also evaluated, as well as another solution using TestNG. Both solutions were determined to be less viable primarily due to variability and increased overhead.

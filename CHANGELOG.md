# Changelog

## Release v1.0.0: Sprout Genesis 🌱 (2024-10-xx)

> 🌱 A New Era of Functionality, Flexibility, and Performance!

We’re thrilled to announce the **v1.0.0-rc1** release of **Sprout**, marking a significant step forward in our mission to create the most powerful and flexible templating library for Go developers. This release introduces major features, critical fixes, and exciting new tools to help you build more efficient and secure templates.

_This section's comparisons are based on Sprig v3.2.3. If you're totally new, welcome! Go ahead with [Getting started](https://docs.atom.codes/sprout/introduction/getting-started)_

### 🚀 **New Core Features**
- **Migration to Function Handler**: All functions have been migrated under a unified function handler to streamline function management. [Documentation](https://docs.atom.codes/sprout/features/loader-system-registry).
- **Registry System (Loader)**: Introduced a new registry system for modular function management, supporting easier extension and organization of functions. [Documentation](https://docs.atom.codes/sprout/features/loader-system-registry).
- **Safe Functions**: New safe versions of functions that follow Go's template standards, providing flexible error handling options. [Documentation](https://docs.atom.codes/sprout/features/safe-functions).
- **Function Notices**: Added real-time notices when specific functions are called to warn or inform users of critical behavior. [Documentation](https://docs.atom.codes/sprout/features/function-notices).
- **Function Aliases**: Added function aliases to ensure smooth transition and backward compatibility. [Documentation](https://docs.atom.codes/sprout/features/function-aliases).

### 🔄 **Backward Compatibility**
- **Reimport Functions from Sprig**: Maintained backward compatibility by reimporting core functions from Sprig. [Documentation](https://docs.atom.codes/sprout/migration-from-sprig).
- **Backward Compatibility Documentation**: Updated docs to ensure seamless migration and backward compatibility. [Documentation](https://docs.atom.codes/sprout/migration-from-sprig).

### 🛠 **Enhancements and Fixes**
- **Optimized Memory Footprint**: Performance improvements were made to reduce memory usage across the board. [Benchmarks](https://github.com/go-sprout/sprout/tree/main/benchmarks).
- **Fixed CamelCase Logic**: Updated CamelCase and PascalCase transformation logic to handle edge cases better. [Documentation](https://docs.atom.codes/sprout/migration-from-sprig#tocamelcase-topascalcase)
- **Never more panics**: Rework functions how cause panics on template engines to ensure a better stability. [Documentation](https://docs.atom.codes/sprout/migration-from-sprig#panicking-functions).

### 🌐 **New Utilities**
- **Batch of New Functions**: Introduced a wide range of functions for slices, regex manipulations, and conversions, expanding Sprout's toolkit significantly. See [PR 70](https://github.com/go-sprout/sprout/pull/70).
- **Network Registry**: New functions for handling IP, CIDR, and MAC address manipulations in templates. See [PR 71](https://github.com/go-sprout/sprout/pull/71).
- **SHA512 Checksums**: Added `sha512sum` to the checksum registry with useful notices for end-users. See [PR 59](https://github.com/go-sprout/sprout/pull/59).
- **New Struct Method `hasField`**: Added a method for checking struct fields dynamically. See [PR 61](https://github.com/go-sprout/sprout/pull/61).
- **String Capitalization Functions**: New string capitalization functions with full Unicode and Latin rune support. See [PR 62](https://github.com/go-sprout/sprout/pull/62).
- **`toDuration` Conversion Function**: A new utility to simplify time conversions across templates. See [PR 27](https://github.com/go-sprout/sprout/pull/27).

### 📚 **Documentation**
- **Fancy and complete documentation**: Create a fancy and complete documentation, ensuring they are up-to-date and aligned with Sprout’s growth. [Documentation](https://sprout.atom.codes)

---

### 🏆 **A Special Thanks to Our Contributors**

A heartfelt thank you to everyone who contributed to this v1.0.0 journey, particularly [@42atomys](https://github.com/42atomys), whose tireless work and commitment have made this release possible. Special thanks to [@mbezhanov](https://github.com/mbezhanov), [@andig](https://github.com/andig), [@ccoVeille](https://github.com/ccoVeille) for their valuable contributions and to [@caarlos0](https://github.com/caarlos0) for the support in making decisions and for being the second maintainer of the Sprout organization.

---

### 🔮 **Looking Ahead**

This release candidate is a crucial step towards the official v1.0.0 release. We encourage you to test the new features, provide feedback, and help us fine-tune the final version. We’re incredibly excited for what’s to come and can’t wait to see how Sprout will evolve with your help!

Let’s continue growing Sprout together and make this library the best tool for Go developers everywhere!

## Release v0.6.0: Sprout Evolution 🌱 (2024-09-16)

> 💡 Cultivating Precision, One Function at a Time!

### 🚀 New Features
- **Function Call Notices**: Added notifications to inform or warn end-users when functions are called. See [PR 58](https://github.com/go-sprout/sprout/pull/58).
- **Safe Functions Compliance**: Aligned safe functions with Go template standards and Sprout conventions. See [PR 65](https://github.com/go-sprout/sprout/pull/65).
- **SHA512 Checksum in Registry**: Introduced `sha512sum` to the checksum registry, complete with informative notices. See [PR 59](https://github.com/go-sprout/sprout/pull/59).
- **`hasField` for Structs**: A new `hasField` method to check fields in structs is now available. See [PR 61](https://github.com/go-sprout/sprout/pull/61).
- **String Capitalization Functions**: Added new functions to capitalize strings, fully supporting Unicode and Latin runes. See [PR 62](https://github.com/go-sprout/sprout/pull/62) and [PR 63](https://github.com/go-sprout/sprout/pull/63).

### 🛠 Fixes & Improvements
- **Dropped v0.1 Error Handling**: Removed legacy error handling until a safe/must decision is finalized in the RFC. See [PR 52](https://github.com/go-sprout/sprout/pull/52).
- **Documentation Updates**: Added function signatures to the conventions, making it easier to understand their usage. See [PR 64](https://github.com/go-sprout/sprout/pull/64).

### 🐛 Bug Fixes
- **Unicode Capitalization Fix**: Resolved issues with string capitalization involving Unicode and Latin runes. See [PR 63](https://github.com/go-sprout/sprout/pull/63).
- **Release Candidate Fixes**: Addressed problems with v0.6.0-rc.1 to ensure stability and performance. See [PR 68](https://github.com/go-sprout/sprout/pull/68).

*Read more about notices on [official documentation](https://sprout.atom.codes/features/function-notices).*

*Read more about safe functions on [official documentation](https://sprout.atom.codes/features/safe-functions).*


## Release v0.5.1: Sprout Growth 🌿 (2024-09-15)

> 💡 Cultivating Code, Growing Solutions!

### ⚡ Performance Improvements
- **Memory Footprint Reduction**: Reduced overall memory footprint for better performance. See (@42atomys) [PR 56](https://github.com/go-sprout/sprout/pull/56).

### 🐛 Bug Fixes
- **Default Logger Initialization**: Fixed an issue where the default logger had a bad initialization. See (@42atomys) [PR 48](https://github.com/go-sprout/sprout/pull/48).
- **Logger Accessibility**: Resolved a problem where loggers were not accessible due to duplicated pointers. See (@42atomys) [PR 50](https://github.com/go-sprout/sprout/pull/50).
- **Sprigin CamelCase Consistency**: Ensured that `sprigin` camelcase returns are consistent. See (@42atomys) [PR 55](https://github.com/go-sprout/sprout/pull/55).

### 🛠️ Chores
- **Go Task Integration**: Replaced Makefile with Go Task for better task management. See (@42atomys) [PR 54](https://github.com/go-sprout/sprout/pull/54).

## Release v0.5.0: Sprout Growth 🌿 (2024-08-15)

> 💡 Nurturing Ideas, Harvesting Innovation!

### 🌟 Major Feature: Registry System Unleashed!
- **Revamped Architecture**: Introducing the powerful registry system (aka loader). This refactor modularizes all methods into separate registries. See (@42atomys) [PR 46](https://github.com/go-sprout/sprout/pull/46)
- **Handler & Registry Interfaces**: New interfaces with clear rules to streamline function management.
- **Seamless Migration**: All functions are now in registries, with backward compatibility via `FuncsMap` in `springin` [See Transitioning from Sprig](https://github.com/go-sprout/sprout?tab=readme-ov-file#transitioning-from-sprig).

*Read more about the registry system in the [official documentation](https://sprout.atom.codes/features/loader-system-registry).*

### 📚 Fully Documented
- **In-Depth Docs**: Detailed documentation and a handy glossary are now available. Explore more [sprout.atom.codes](https://sprout.atom.codes).
- **README**: Updated the documentation and README to reflect all recent changes. Check out the latest [README.md](https://raw.githubusercontent.com/go-sprout/sprout/main/README.md).

### 🐛 Bug Fixes
- **`toDuration` Doc Update**: Added a practical example showing how to convert durations to seconds using `toDuration`. This is based on real test cases to make time formatting easier. See (@cbandy) [PR 44](https://github.com/go-sprout/sprout/pull/44).


## Release v0.4.1: Sprout Blossom 🌸 (2024-06-17)

> 💡 Cultivating Innovation, One Sprig at a Time!

### 🚀 Features
- **Reducing YAML Dependencies Footprint**: Improved project efficiency by reducing dependencies. See (@andig) [PR 38](https://github.com/go-sprout/sprout/pull/38).
  
### 🐛 Bugs fixes
- **Timezone Leak Fix in `toDate` Method**: Resolved an issue affecting date calculations. See (@42atomys) [PR 42](https://github.com/go-sprout/sprout/pull/42).
- **Backward Compatibility Documentation**: Added documentation for ensuring seamless upgrades. See (@42atomys) [PR 43](https://github.com/go-sprout/sprout/pull/43).

## Release v0.4.0: Sprout Blossom 🌸 (2024-05-16)

> 💡 Cultivating code is something beautiful.

### 🚀 Features
- **Enhanced Conversions Group**: New functions (`toBool`, `toUint`, `toUint64`) and comprehensive documentation have been added to the conversions group, broadening our library's functionality and making it more user-friendly. See (@42atomys) [PR 33](https://github.com/go-sprout/sprout/pull/33).
- **YAML Functions Unleashed**: Implementing YAML functions (`fromYaml`, `toYaml`. `mustFromYaml`, `mustToYaml`) inspired by Helm's robust toolset, we've extended our configuration management capabilities. See (@42atomys) [PR 36](https://github.com/go-sprout/sprout/pull/36).

### 🛠 Fixes from sprig issues
- **Merge Function Improvement**: The merge function has been tweaked to preserve the zero value in destination structs, ensuring more predictable and accurate data handling. See (@42atomys) [PR 34](https://github.com/go-sprout/sprout/pull/34).
- **String Transformation Logic Update**: Corrected the logic for transforming strings to CamelCase and PascalCase to avoid previous inconsistencies and errors. See (@42atomys) [PR 35](https://github.com/go-sprout/sprout/pull/35).


## Release v0.3.0: Moved Farm 🌾 (2024-05-09)

> 💡 Sprouting New Possibilities in Every Release!

> [!IMPORTANT]
> The project has moved to a new GitHub home [**github.com/go-sprout/sprout**](https://github.com/go-sprout/sprout) !

### 🚀 Features
- **Unified Function Management**: All functions are now neatly organized under the new FunctionHandler, streamlining how functionalities are handled within the library. This consolidation is crucial for enhancing library operations and future development. See (@42atomys) [PR 14](https://github.com/go-sprout/sprout/pull/14).
- **Introducing `toDuration` Conversion**: A new utility function, `toDuration`, has been added to simplify time conversions across various formats, enhancing our toolkit's versatility. See (@42atomys) [PR 27](https://github.com/go-sprout/sprout/pull/27).

### 🛠 Documentation and Community
- **Project's New Home**: The project has moved to a new GitHub home, centralizing where updates and community interactions will take place. Visit us at: [Sprout on GitHub](https://github.com/go-sprout/sprout).
- **Community Files Update**: All community-related files have been refreshed to better support our growing community of developers and contributors. See (@42atomys) [PR 12](https://github.com/go-sprout/sprout/pull/12).

## Release v0.2.0: Garden Genesis 🌱 (2024-04-03)

> 💡 Cultivating Innovation, One Sprig at a Time!

### 🚀 Features
- **Creating the Root of the Sprout**: Sprouts are now an evolution of Sprig with a standalone function handler. See (@42atomys) [PR 2](https://github.com/go-sprout/sprout/pull/2).
- **Allowing Function Aliasing**: Enables developers to use aliases for their templates. In Sprout, this feature is used for backward compatibility with Sprig. See (@42atomys) [PR 3](https://github.com/go-sprout/sprout/pull/3). 
  - Full documentation available here: https://docs.atom.codes/sprout/function-aliases

### 🛠 Chore
- **Documentation Available**: Documentation can be found at https://docs.atom.codes/sprout.
- **README Refactor**: The README has been updated to reflect the project's vision and its future. See (@42atomys) [PR 4](https://github.com/go-sprout/sprout/pull/4).


## Release v0.1.0: New Seed 🌱 (2024-03-29)

We are excited to announce the release of Sprout v0.1, a modern and evolved variant of the [Masterminds/sprig](https://github.com/Masterminds/sprig) library, specifically reimagined and redesigned for contemporary Go development environments. Our mission with Sprout is to reignite the innovation that made Sprig an indispensable tool for Go developers, providing a robust set of functions and helpers that enhance productivity and code clarity.

### Vision and Goals
Our vision for Sprout is to not only match but exceed the functionality and reliability that made Sprig a cornerstone in many Go projects. We aim to bring Sprout into the modern Go ecosystem, ensuring compatibility with the latest versions of Go and introducing a stream of new features and improvements that reflect the needs and requests of the community. We recognize the importance of maintaining a vibrant and up-to-date toolset for developers and commit to an active development cycle for Sprout.

### What's New in v0.1
Sprout v0.1 is designed to align seamlessly with Sprig v3.2.3, providing a familiar yet enhanced experience for developers transitioning from Sprig. 

Key features and enhancements include:
**Enhanced Compatibility**: Sprout is fully compatible with modern Go versions, starting with Go 1.19 and above, addressing the compatibility issues faced by Sprig users in newer Go environments.
**New Functions and Improvements**: We will introduce additional functions and enhancements to existing ones, carefully designed to increase productivity and simplify common coding tasks in Go.
**Performance Optimizations**: Sprout includes significant performance improvements, making your applications faster and more efficient.
**Community-Driven Development**: Sprout is committed to being a community-focused project, welcoming contributions, and suggestions from developers to shape the future of the library.

### Future Directions
Looking ahead, Sprout will continue to evolve with the Go ecosystem. Our roadmap includes the integration of more features and utilities, drawing from the feedback and needs of our growing community of users. We aim to foster a vibrant ecosystem around Sprout, encouraging contributions, and collaboration to ensure that Sprout remains at the forefront of Go development tools.

### Getting Started with Sprout
To start using Sprout in your Go projects, please visit our GitHub repository at [Sprout's GitHub Page](https://github.com/go-sprout/sprout). You'll find comprehensive documentation, installation instructions, and examples to help you get started.

We are thrilled to embark on this journey with you, the Go developer community, and look forward to seeing the incredible applications you will build with Sprout. Thank you for your support, and welcome to Sprout v0.1!

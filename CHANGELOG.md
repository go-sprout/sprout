# Changelog

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
